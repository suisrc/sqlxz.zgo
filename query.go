package sqlxz

import (
	"fmt"
	"reflect"
	"strings"
)

// QueryGet 查询单个
func QueryGet(sqlx EntityDB, srs interface{}, table, after string, sps ...interface{}) error {
	slc := SelectColumns(srs)
	sqr := fmt.Sprintf(`select %s from %s%s %s`, slc, TBL, table, after)
	err := sqlx.Get(srs, sqr, sps...)
	return err
}

// QuerySelect 查询集合
func QuerySelect(sqlx ArrayDB, srs interface{}, table, after string, sps ...interface{}) error {
	val := reflect.ValueOf(srs)
	vty := val.Type().Elem().Elem()
	vpc := reflect.New(vty)
	bas := vpc.Interface()
	slc := SelectColumns(bas)
	sqr := fmt.Sprintf(`select %s from %s%s %s`, slc, TBL, table, after)
	err := sqlx.Select(srs, sqr, sps...)
	return err
}

// QueryGet 查询单个
func QueryGetTBL(sqlx EntityDB, srs interface{}, sql string, sps ...interface{}) error {
	slc := SelectColumns(srs)
	sqr := fmt.Sprintf(`select %s from %s`, slc, strings.ReplaceAll(sql, "[TBL]", TBL))
	err := sqlx.Get(srs, sqr, sps...)
	return err
}

// QuerySelect 查询集合
func QuerySelectTBL(sqlx ArrayDB, srs interface{}, sql string, sps ...interface{}) error {
	val := reflect.ValueOf(srs)
	vty := val.Type().Elem().Elem()
	vpc := reflect.New(vty)
	bas := vpc.Interface()
	slc := SelectColumns(bas)
	sqr := fmt.Sprintf(`select %s from %s`, slc, strings.ReplaceAll(sql, "[TBL]", TBL))
	err := sqlx.Select(srs, sqr, sps...)
	return err
}

// QuerySelect 查询集合
func QuerySelectDistinctTBL(sqlx ArrayDB, srs interface{}, sql string, sps ...interface{}) error {
	val := reflect.ValueOf(srs)
	vty := val.Type().Elem().Elem()
	vpc := reflect.New(vty)
	bas := vpc.Interface()
	slc := SelectColumns(bas)
	sqr := fmt.Sprintf(`select distinct %s from %s`, slc, strings.ReplaceAll(sql, "[TBL]", TBL))
	err := sqlx.Select(srs, sqr, sps...)
	return err
}
