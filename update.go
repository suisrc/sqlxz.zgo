package sqlxz

import "errors"

// UpdateAndSaveX ...
func UpdateAndSaveX(sqlx ExecDB, srs interface{}, tbl string, sid int64, fid func(int64)) error {
	sic := TableIdxColumn{Table: TBL + tbl, IDVal: sid}
	if sqr, sps, err := CreateUpdateSQLByNamedAndSkipNilAndSet(sic, srs); err != nil {
		return err
	} else if res, err := sqlx.NamedExec(sqr, sps); err != nil {
		return err
	} else if cii, err := res.RowsAffected(); err != nil {
		return err
	} else if cii == 0 { // 操作失败
		return errors.New("sql: no rows in result set")
	} else if sid == 0 { // install id
		if nid, err := res.LastInsertId(); err != nil {
			return err
		} else {
			fid(nid)
		}
	}
	return nil
}

// UpdateAndSaveX2 ...
func UpdateAndSaveTIC(sqlx ExecDB, srs interface{}, sic *TableIdxColumn, fres func(interface{})) error {
	if sqr, sps, err := CreateUpdateSQLByNamedAndSkipNilAndSet(*sic, srs); err != nil {
		return err
	} else if res, err := sqlx.NamedExec(sqr, sps); err != nil {
		return err
	} else if cii, err := res.RowsAffected(); err != nil {
		return err
	} else if cii == 0 { // 操作失败
		return errors.New("sql: no rows in result set")
	} else if fres != nil {
		fres(res)
	}
	return nil
}
