package features

import (
	"context"
	"errors"
	"strconv"
	"strings"
	_ "time"

	"github.com/arangodb/go-driver"
)

// всего получено BTC на адрес
func TotalGetAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	query := "RETURN SUM(FOR doc IN btcOut FILTER doc._to == @addr RETURN doc.spentBtc)"
	bindVars := map[string]interface{}{
		"addr": addr,
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		return 0, err
	}
	defer cursor.Close()

	var income int64
	_, err = cursor.ReadDocument(ctx, &income)
	if err != nil {
		return 0, err
	}

	if income == 0 {
		err = errors.New("There is no such address")
		return 0, err
	}
	return income, nil
}

// баланс BTC
func BalanceAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	query := "RETURN SUM(FOR doc IN btcIn FILTER doc._from == @addr RETURN doc.spentBtc)"
	bindVars := map[string]interface{}{
		"addr": addr,
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		return 0, err
	}
	defer cursor.Close()

	var spend int64
	_, err = cursor.ReadDocument(ctx, &spend)
	if err != nil {
		return 0, err
	}

	var income int64
	income, err = TotalGetAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	if income == 0 && spend == 0 {
		err = errors.New("There is no such address")
		return 0, err
	}
	return income - spend, nil
}

// время первого появления адреса
func FirstTimeAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	query := "RETURN MIN(FOR doc IN btcOut FILTER doc._to == @addr RETURN doc.time)"
	bindVars := map[string]interface{}{
		"addr": addr,
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		return 0, err
	}
	defer cursor.Close()

	var time int64
	_, err = cursor.ReadDocument(ctx, &time)
	if err != nil {
		return 0, err
	}

	if time == 0 {
		err = errors.New("There is no such address")
		return 0, err
	}
	return time, nil
}

// время последнего появления адреса
func LastTimeAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	query := "RETURN MAX(FOR doc IN btcIn FILTER doc._from == @addr RETURN doc.time)"
	bindVars := map[string]interface{}{
		"addr": addr,
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		return 0, err
	}
	defer cursor.Close()

	var time int64
	_, err = cursor.ReadDocument(ctx, &time)
	if err != nil {
		return 0, err
	}

	if time == 0 {
		err = errors.New("There is no such address")
		return 0, err
	}
	return time, nil
}

// кол-во входящих транзакций
func CountOutTx(ctx context.Context, db driver.Database, addr string) (int64, error) {
	query := "RETURN SUM(FOR doc IN btcOut FILTER doc._to == @addr RETURN 1)"
	bindVars := map[string]interface{}{
		"addr": addr,
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		return 0, err
	}
	defer cursor.Close()

	var count int64
	_, err = cursor.ReadDocument(ctx, &count)
	if err != nil {
		return 0, err
	}

	if count == 0 {
		err = errors.New("There is no such address")
		return 0, err
	}
	return count, nil
}

// кол-во исходящих транзакций
func CountInTx(ctx context.Context, db driver.Database, addr string) (int64, error) {
	query := "RETURN SUM(FOR doc IN btcIn FILTER doc._from == @addr RETURN 1)"
	bindVars := map[string]interface{}{
		"addr": addr,
	}
	cursor, err := db.Query(ctx, query, bindVars)
	if err != nil {
		return 0, err
	}
	defer cursor.Close()

	var count int64
	_, err = cursor.ReadDocument(ctx, &count)
	if err != nil {
		return 0, err
	}

	if count == 0 {
		err = errors.New("There is no such address")
		return 0, err
	}
	return count, nil
}

// кол-во адресов на которые уходили средства + список этих адресов
func countInAddr(ctx context.Context, db driver.Database, addr string) (int64, map[string]bool, error) {
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

	inAddr := make(map[string]bool)
	var countAddr int64
	for {
		var doc string
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return 0, nil, err
		}
		_, found := inAddr[doc]
		if !found && doc != addr {
			inAddr[doc] = false
			countAddr++
		}
	}

	if countAddr == 0 {
		err = errors.New("There is no such address")
		return 0, nil, err
	}
	return countAddr, inAddr, nil
}

// кол-во адресов на которые уходили средства
func CountInAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	c, _, err := countInAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	return c, nil
}

// кол-во адресов с которых приходили средства + список этих адресов
func countOutAddr(ctx context.Context, db driver.Database, addr string) (int64, map[string]bool, error) {
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

	outAddr := make(map[string]bool)
	var countAddr int64
	for {
		var doc string
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return 0, nil, err
		}

		_, found := outAddr[doc]
		if !found && doc != addr {
			outAddr[doc] = true
			countAddr++
		}
	}
	if countAddr == 0 {
		err = errors.New("There is no such address")
		return 0, nil, err
	}
	return countAddr, outAddr, nil
}

// кол-во адресов с которых приходили средства
func CountOutAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	c, _, err := countOutAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	return c, nil
}

// кол-во общих адресов среди countInAddr и countOutAddr
func CountSharedAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	_, outAddr, err := countOutAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	_, inAddr, err := countInAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	var count int64
	for key := range inAddr {
		_, found := outAddr[key]
		if found {
			count++
		}
	}
	return count, nil
}

// общее кол-во адресов среди countInAddr и countOutAddr
func TotalCountAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	count, outAddr, err := countOutAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	_, inAddr, err := countInAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	if err != nil {
		return 0, err
	}
	for key := range inAddr {
		_, found := outAddr[key]
		if !found {
			count++
		}
	}
	return count, nil
}

// кол-во уникальных адресов среди countInAddr и countOutAddr
// inAddr = 5 общих + 3 уникальных
// outAddr = 5 общих + 2 уникальных
// count = 7
// 7 + 3 - 5 = 5  (3 раза зайдет в if !found и 5 раз в else)
func CountUniqueAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	count, outAddr, err := countOutAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	_, inAddr, err := countInAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	for key := range inAddr {
		_, found := outAddr[key]
		if !found {
			count++
		} else {
			count--
		}
	}
	return count, nil
}

// среднее кол-во адресов во входных транзакциях
func AverageCountOutAddr(ctx context.Context, db driver.Database, addr string) (float64, error) {
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

	outAddr := make(map[string]int64)
	var countAddr float64
	var countTx float64
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
	if countAddr == 0 && countTx == 0 {
		err = errors.New("There is no such address")
		return 0, err
	}
	return countAddr / countTx, nil
}

// среднее кол-во адресов в выходящих транзакциях
func AverageCountInAddr(ctx context.Context, db driver.Database, addr string) (float64, error) {
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

	inAddr := make(map[string]int64)
	var countAddr float64
	var countTx float64
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

	if countAddr == 0 && countTx == 0 {
		err = errors.New("There is no such address")
		return 0, err
	}
	return countAddr / countTx, nil
}

func NmotifAddr(ctx context.Context, db driver.Database, addr1 string, addr2 string, n int) ([][]string, error) {
	var (
		path     [][]string
		i        int
		addr1Key string
		addr2Key string
	)
	if 0 < n && n < 4 {
		parts := strings.Split(addr1, "/")
		if len(parts) > 1 {
			addr1Key = parts[1]
		} else {
			err := errors.New("Адрес не содержит символа /")
			return nil, err
		}
		parts = strings.Split(addr2, "/")
		if len(parts) > 1 {
			addr2Key = parts[1]
		} else {
			err := errors.New("Адрес не содержит символа /")
			return nil, err
		}

		query := `FOR bn IN btcNext FILTER bn.address == @addr RETURN bn._to`
		bindVars := map[string]interface{}{
			"addr": addr1Key,
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

			var btx bTx
			_, err = cursor2.ReadDocument(ctx, &btx)
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
				"startVertex": btx,
				"n":           n,
			}
			cursor3, err := db.Query(ctx, query, bindVars)
			if err != nil {
				return nil, err
			}
			defer cursor3.Close()

			var p map[string]btcNext
			for {
				var one_path []string
				_, err := cursor3.ReadDocument(ctx, &p)
				if driver.IsNoMoreDocuments(err) {
					break
				} else if err != nil {
					return nil, err
				}
				one_path = append(one_path, addr1Key)
				if len(p) == n {
					for j := 0; j < n; j++ {
						val, _ := p[strconv.Itoa(j)]
						id := val.From
						one_path = append(one_path, id)
						address := val.Address
						one_path = append(one_path, address)
					}
				}
				if one_path[n*2] == addr2Key {
					path = append(path, one_path)
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
				"startVertex": btx,
				"n":           n,
			}
			cursor4, err := db.Query(ctx, query, bindVars)
			if err != nil {
				return nil, err
			}
			defer cursor4.Close()

			for {
				var one_path []string
				_, err := cursor4.ReadDocument(ctx, &p)
				if driver.IsNoMoreDocuments(err) {
					break
				} else if err != nil {
					return nil, err
				}
				if len(p) == n+1 {
					for j := 0; j < n+1; j++ {
						val, _ := p[strconv.Itoa(j)]
						address := val.Address
						one_path = append(one_path, address)
						if j != n {
							id := val.From
							one_path = append(one_path, id)
						}
					}
				}
				if one_path[n*2] == addr2Key {
					path = append(path, one_path)
				}
			}
			i++
		}
	} else {
		err := errors.New("Invalid n, 1 <= n <= 3")
		return nil, err
	}
	if i == 0 {
		err := errors.New("invalid addr1")
		return nil, err
	}
	if len(path) == 0 {
		query := `FOR baddr IN btcAddress FILTER baddr._id == @addr RETURN 1`
		bindVars := map[string]interface{}{
			"addr": addr2,
		}
		cursor, err := db.Query(ctx, query, bindVars)
		if err != nil {
			return nil, err
		}
		defer cursor.Close()

		var doc int
		_, err = cursor.ReadDocument(ctx, &doc)

		if doc != 1 {
			err := errors.New("invalid addr2")
			return nil, err
		}
	}
	return path, nil
}
