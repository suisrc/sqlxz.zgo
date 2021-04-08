package sqlxz

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
)

var (
	TBL = "zgo_"
)

// QueryGet 查询单个
func QueryGet(sqlx *sqlx.DB, srs interface{}, table, after string, sps ...interface{}) error {
	slc := SelectColumns(srs)
	sqr := fmt.Sprintf(`select %s from %s%s %s`, slc, TBL, table, after)
	err := sqlx.Get(srs, sqr, sps...)
	return err
}

// QuerySelect 查询集合
func QuerySelect(sqlx *sqlx.DB, srs interface{}, table, after string, sps ...interface{}) error {
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
func QueryGetTBL(sqlx *sqlx.DB, srs interface{}, sql string, sps ...interface{}) error {
	slc := SelectColumns(srs)
	sqr := fmt.Sprintf(`select %s from %s`, slc, strings.ReplaceAll(sql, "[TBL]", TBL))
	err := sqlx.Get(srs, sqr, sps...)
	return err
}

// QuerySelect 查询集合
func QuerySelectTBL(sqlx *sqlx.DB, srs interface{}, sql string, sps ...interface{}) error {
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
func QuerySelectDistinctTBL(sqlx *sqlx.DB, srs interface{}, sql string, sps ...interface{}) error {
	val := reflect.ValueOf(srs)
	vty := val.Type().Elem().Elem()
	vpc := reflect.New(vty)
	bas := vpc.Interface()
	slc := SelectColumns(bas)
	sqr := fmt.Sprintf(`select distinct %s from %s`, slc, strings.ReplaceAll(sql, "[TBL]", TBL))
	err := sqlx.Select(srs, sqr, sps...)
	return err
}
