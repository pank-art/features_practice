package features

import (
	"context"
	"errors"
	_ "time"

	"github.com/arangodb/go-driver"
)

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

// кол-во адресов на которые уходили средства
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

func CountInAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	c, _, err := countInAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	return c, nil
}

// кол-во адресов с которых приходили средства
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

func CountOutAddr(ctx context.Context, db driver.Database, addr string) (int64, error) {
	c, _, err := countOutAddr(ctx, db, addr)
	if err != nil {
		return 0, err
	}
	return c, nil
}

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
