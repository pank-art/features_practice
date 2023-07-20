package features

import (
	"context"
	"github.com/arangodb/go-driver"
	"log"
	"strconv"
)

type Book struct {
	Title   string `json:"title"`
	NoPages int    `json:"no_pages"`
}

func GetCluster(ctx context.Context, db driver.Database, walletId string) (chan []string, error) {
	ch := make(chan []string)

	go func() {
		defer close(ch)

		query := `FOR addr IN btcAddress FILTER addr._walletId == @walletId RETURN addr._id`
		bindVars := map[string]interface{}{
			"walletId": walletId,
		}
		cursor, err := db.Query(ctx, query, bindVars)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close()

		var cluster []string
		for {
			var doc string
			_, err := cursor.ReadDocument(ctx, &doc)
			if driver.IsNoMoreDocuments(err) {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			cluster = append(cluster, doc)

			if len(cluster) >= 10 {
				ch <- cluster
				cluster = []string{}
			}
		}

		if len(cluster) > 0 {
			ch <- cluster
		}
	}()

	return ch, nil
}

func GetCluster_key(ctx context.Context, db driver.Database, walletId string) (chan []string, error) {
	ch := make(chan []string)

	go func() {
		defer close(ch)

		query := `FOR addr IN btcAddress FILTER addr._walletId == @walletId RETURN addr._key`
		bindVars := map[string]interface{}{
			"walletId": walletId,
		}
		cursor, err := db.Query(ctx, query, bindVars)
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close()

		var cluster []string
		for {
			var doc string
			_, err := cursor.ReadDocument(ctx, &doc)
			if driver.IsNoMoreDocuments(err) {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			cluster = append(cluster, doc)

			if len(cluster) >= 10 {
				ch <- cluster
				cluster = []string{}
			}
		}

		if len(cluster) > 0 {
			ch <- cluster
		}
	}()

	return ch, nil
}

func AddrKeyInCluster(ctx context.Context, db driver.Database, addrKey string, walletId string) (bool, error) {
	query := `FOR addr IN btcAddress FILTER addr._key == @addr RETURN addr._walletId`
	bindVars := map[string]interface{}{
		"addr": addrKey,
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		return false, err
	}
	defer cursor.Close()

	var doc string
	_, err = cursor.ReadDocument(ctx, &doc)

	return doc == walletId, nil
}

func TotalGetClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	clusterCh, err := GetCluster(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	var total int64
	for cluster := range clusterCh {
		for _, addr := range cluster {
			tmp, err := TotalGetAddr(ctx, db, addr)
			if err != nil {
				return 0, err
			}
			total += tmp
		}
	}
	return total, nil
}

func BalanceClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	clusterCh, err := GetCluster(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	var balance int64
	for cluster := range clusterCh {
		for _, addr := range cluster {
			tmp, err := BalanceAddr(ctx, db, addr)
			if err != nil {
				return 0, err
			}
			balance += tmp
		}
	}
	return balance, nil
}

func FirstTimeClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	clusterCh, err := GetCluster(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	var firstTime int64
	for cluster := range clusterCh {
		for i, addr := range cluster {
			if i != 0 {
				firstAddr, err := FirstTimeAddr(ctx, db, addr)
				if err != nil {
					return 0, err
				}
				if firstAddr < firstTime {
					firstTime = firstAddr
				}
			} else {
				firstTime, err = FirstTimeAddr(ctx, db, cluster[0])
				if err != nil {
					return 0, err
				}
			}
		}
	}
	return firstTime, nil
}

func LastTimeClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	clusterCh, err := GetCluster(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	var lastTime int64
	for cluster := range clusterCh {
		for i, addr := range cluster {
			if i != 0 {
				lastAddr, err := FirstTimeAddr(ctx, db, addr)
				if err != nil {
					return 0, err
				}
				if err != nil {
					return 0, err
				}
				if lastAddr > lastTime {
					lastTime = lastAddr
				}
			} else {
				lastTime, err = FirstTimeAddr(ctx, db, cluster[0])
				if err != nil {
					return 0, err
				}
			}
		}
	}
	return lastTime, nil
}

// кол-во входящих транзакций
func CountInTxClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	clusterCh, err := GetCluster(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	var countIn int64
	for cluster := range clusterCh {
		for _, addr := range cluster {
			tmp, err := CountInAddr(ctx, db, addr)
			if err != nil {
				return 0, err
			}
			countIn += tmp
		}
	}
	return countIn, nil
}

// кол-во исходящих транзакций
func CountOutTxClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	clusterCh, err := GetCluster(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	var countOut int64
	for cluster := range clusterCh {
		for _, addr := range cluster {
			tmp, err := CountOutAddr(ctx, db, addr)
			if err != nil {
				return 0, err
			}
			countOut += tmp
		}
	}
	return countOut, nil
}

// кол-во адресов на которые уходили средства
func countOutClust(ctx context.Context, db driver.Database, walletId string) (int64, map[string]bool, error) {
	clusterCh, err := GetCluster(ctx, db, walletId)
	if err != nil {
		return 0, nil, err
	}
	var count int64
	var outClust map[string]bool
	for cluster := range clusterCh {
		for _, addr := range cluster {
			query := `LET tx = (FOR bin IN btcIn FILTER bin._from == @addr RETURN bin._to)
				FOR bout IN btcOut 
					FILTER bout._from IN tx
						RETURN bout._to`
			bindVars := map[string]interface{}{
				"addr": addr,
			}
			cursor, err := db.Query(ctx, query, bindVars)
			if err != nil {
				return 0, nil, err
			}
			defer cursor.Close()

			for {
				var doc string
				_, err := cursor.ReadDocument(ctx, &doc)
				if driver.IsNoMoreDocuments(err) {
					break
				} else if err != nil {
					return 0, nil, err
				}
				_, found := outClust[doc]
				if !found && doc != addr {
					outClust[doc] = false
					count++
				}
			}
		}
	}
	return count, outClust, nil
}

func CountOutClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	c, _, err := countOutClust(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	return c, nil
}

// кол-во адресов с которых приходили средства
func countInClust(ctx context.Context, db driver.Database, walletId string) (int64, map[string]bool, error) {
	clusterCh, err := GetCluster(ctx, db, walletId)
	if err != nil {
		return 0, nil, err
	}
	var count int64
	var inClust map[string]bool
	for cluster := range clusterCh {
		for _, addr := range cluster {
			query := `LET tx = (FOR bout IN btcOut FILTER bout._to == @addr RETURN bout._from)
				FOR bin IN btcIn 	
					FILTER bin._to IN tx 
						RETURN bin._from`
			bindVars := map[string]interface{}{
				"addr": addr,
			}
			cursor, err := db.Query(ctx, query, bindVars)
			if err != nil {
				return 0, nil, err
			}
			defer cursor.Close()

			for {
				var doc string
				_, err := cursor.ReadDocument(ctx, &doc)
				if driver.IsNoMoreDocuments(err) {
					break
				} else if err != nil {
					return 0, nil, err
				}
				_, found := inClust[doc]
				if !found && doc != addr {
					inClust[doc] = true
					count++
				}
			}
		}
	}
	return count, inClust, nil
}

func CountInClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	c, _, err := countInClust(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	return c, nil
}

func CountSharedClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	_, inClust, err := countInClust(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	_, outClust, err := countOutClust(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	var count int64
	for key := range inClust {
		_, found := outClust[key]
		if found {
			count++
		}
	}
	return count, nil
}

func TotalCountClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	_, inClust, err := countInClust(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	count, outClust, err := countOutClust(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	for key := range inClust {
		_, found := outClust[key]
		if !found {
			count++
		}
	}
	return count, nil
}

// inAddr = 5 общих + 3 уникальных
// outAddr = 5 общих + 2 уникальных
// count = 7
// 7 + 3 - 5 = 5  (3 раза зайдет в if !found и 5 раз в else)
func CountUniqueClust(ctx context.Context, db driver.Database, walletId string) (int64, error) {
	_, inClust, err := countInClust(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	count, outClust, err := countOutClust(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	for key := range inClust {
		_, found := outClust[key]
		if !found {
			count++
		} else {
			count--
		}
	}
	return count, nil
}

// среднее кол-во адресов во входных транзакциях
func AverageCountInClust(ctx context.Context, db driver.Database, walletId string) (float64, error) {
	clusterCh, err := GetCluster(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	inAddr := make(map[string]int64)
	var countAddr float64
	var countTx float64
	for cluster := range clusterCh {
		for _, addr := range cluster {
			query := `LET tx = (FOR bout IN btcOut FILTER bout._to == @addr RETURN bout._from)
				FOR bin IN btcIn 	
					FILTER bin._to IN tx 
						RETURN bin._to`
			bindVars := map[string]interface{}{
				"addr": addr,
			}
			cursor, err := db.Query(ctx, query, bindVars)
			if err != nil {
				return 0, err
			}
			defer cursor.Close()

			for {
				var doc string
				_, err := cursor.ReadDocument(ctx, &doc)
				if driver.IsNoMoreDocuments(err) {
					break
				} else if err != nil {
					return 0, err
				}

				_, found := inAddr[doc]
				if found {
					inAddr[doc]++
				} else {
					inAddr[doc] = 1
					countTx++
				}
				countAddr++
			}
		}
	}
	return countAddr / countTx, nil
}

// среднее кол-во адресов в выходящих транзакциях
func AverageCountOutClust(ctx context.Context, db driver.Database, walletId string) (float64, error) {
	clusterCh, err := GetCluster(ctx, db, walletId)
	if err != nil {
		return 0, err
	}
	outAddr := make(map[string]int64)
	var countAddr float64
	var countTx float64
	for cluster := range clusterCh {
		for _, addr := range cluster {
			query := `LET tx = (FOR bin IN btcIn FILTER bin._from == @addr RETURN bin._to)
				FOR bout IN btcOut 
					FILTER bout._from IN tx 
						RETURN bout._from`
			bindVars := map[string]interface{}{
				"addr": addr,
			}
			cursor, err := db.Query(ctx, query, bindVars)
			if err != nil {
				return 0, err
			}
			defer cursor.Close()

			for {
				var doc string
				_, err := cursor.ReadDocument(ctx, &doc)
				if driver.IsNoMoreDocuments(err) {
					break
				} else if err != nil {
					return 0, err
				}

				_, found := outAddr[doc]
				if found {
					outAddr[doc]++
				} else {
					outAddr[doc] = 1
					countTx++
				}
				countAddr++
			}
		}
	}
	return countAddr / countTx, nil
}

func Nmotif(ctx context.Context, db driver.Database, walletId1 string, walletId2 string, n int) ([][]interface{}, error) {
	var path [][]interface{}
	if 0 < n && n < 4 {
		clusterCh, err := GetCluster_key(ctx, db, walletId1)
		if err != nil {
			return nil, err
		}
		for cluster := range clusterCh {
			for _, addr := range cluster {
				query := `FOR bn IN btcNext FILTER bn.address == @addr RETURN bn._to`
				bindVars := map[string]interface{}{
					"addr": addr,
				}
				cursor, err := db.Query(ctx, query, bindVars)
				if err != nil {
					return nil, err
				}
				defer cursor.Close()

				for {
					var doc string
					_, err := cursor.ReadDocument(ctx, &doc)
					if driver.IsNoMoreDocuments(err) {
						break
					} else if err != nil {
						return nil, err
					}
					query = `FOR bTx IN btcTx FILTER bTx._id == @bT RETURN bTx`
					bindVars = map[string]interface{}{
						"bT": doc,
					}
					cursor2, err := db.Query(ctx, query, bindVars)
					if err != nil {
						return nil, err
					}
					defer cursor2.Close()

					var bTx map[string]interface{}
					_, err = cursor2.ReadDocument(ctx, &bTx)
					if driver.IsNoMoreDocuments(err) {
						break
					} else if err != nil {
						return nil, err
					}
					if n == 1 {
						query = `FOR v, e, p
							IN @n..@n
								ANY @startVertex
							GRAPH "graphNext"
							FILTER p.edges[0]._to != @startVertex._id
								Return {"0": p.edges[0]}`
					} else if n == 2 {
						query = `FOR v, e, p
							IN @n..@n
								ANY @startVertex
							GRAPH "graphNext"
							FILTER p.edges[0]._to != @startVertex._id
								Return {"0": p.edges[0], "1": p.edges[1]}`
					} else if n == 3 {
						query = `FOR v, e, p
							IN @n..@n
								ANY @startVertex
							GRAPH "graphNext"
							FILTER p.edges[0]._to != @startVertex._id
								Return {"0": p.edges[0], "1": p.edges[1], "2": p.edges[2]}`
					}
					bindVars = map[string]interface{}{
						"startVertex": bTx,
						"n":           n,
					}
					cursor3, err := db.Query(ctx, query, bindVars)
					if err != nil {
						return nil, err
					}
					defer cursor3.Close()

					var p map[string]map[string]interface{}
					for {
						var abc []interface{}
						_, err := cursor3.ReadDocument(ctx, &p)
						if driver.IsNoMoreDocuments(err) {
							break
						} else if err != nil {
							return nil, err
						}
						abc = append(abc, addr)
						if len(p) == n {
							for j := 0; j < n; j++ {
								val, _ := p[strconv.Itoa(j)]
								id, _ := val["_from"]
								abc = append(abc, id)
								address := val["address"]
								abc = append(abc, address)
							}
						}
						f, err := AddrKeyInCluster(ctx, db, abc[n*2].(string), walletId2)
						if err != nil {
							return nil, err
						}
						if f {
							path = append(path, abc)
						}

					}
					if n == 1 {
						query = `FOR v, e, p
							IN @n+1..@n+1
								ANY @startVertex
							GRAPH "graphNext"
							FILTER p.edges[0]._to == @startVertex._id
								Return {"0": p.edges[0], "1": p.edges[1]}`
					} else if n == 2 {
						query = `FOR v, e, p
							IN @n+1..@n+1
								ANY @startVertex
							GRAPH "graphNext"
							FILTER p.edges[0]._to == @startVertex._id
								Return {"0": p.edges[0], "1": p.edges[1], "2": p.edges[2]}`
					} else if n == 3 {
						query = `FOR v, e, p
							IN @n+1..@n+1
								ANY @startVertex
							GRAPH "graphNext"
							FILTER p.edges[0]._to == @startVertex._id
								Return {"0": p.edges[0], "1": p.edges[1], "2": p.edges[2], "3": p.edges[3]}`
					}

					bindVars = map[string]interface{}{
						"startVertex": bTx,
						"n":           n,
					}
					cursor4, err := db.Query(ctx, query, bindVars)
					if err != nil {
						return nil, err
					}
					defer cursor4.Close()

					for {
						var abc []interface{}
						_, err := cursor4.ReadDocument(ctx, &p)
						if driver.IsNoMoreDocuments(err) {
							break
						} else if err != nil {
							return nil, err
						}
						if len(p) == n+1 {
							for j := 0; j < n+1; j++ {
								val, _ := p[strconv.Itoa(j)]
								address := val["address"]
								abc = append(abc, address)
								if j != n {
									id, _ := val["_from"]
									abc = append(abc, id)
								}
							}
						}
						f, err := AddrKeyInCluster(ctx, db, abc[n*2].(string), walletId2)
						if err != nil {
							return nil, err
						}
						if f {
							path = append(path, abc)
						}
					}
				}
			}
		}
	}
	return path, nil
}
