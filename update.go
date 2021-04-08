package sqlxz

import "github.com/jmoiron/sqlx"

// UpdateAndSaveX ...
func UpdateAndSaveX(sqlx *sqlx.DB, srs interface{}, tbl2 string, id int64, fid func(int64)) error {
	sic := TableIdxColumn{Table: TBL + tbl2, IDVal: id}
	if sqr, sps, err := CreateUpdateSQLByNamedAndSkipNilAndSet(sic, srs); err != nil {
		return err
	} else if res, err := sqlx.NamedExec(sqr, sps); err != nil {
		return err
	} else if id == 0 && fid != nil {
		if lii, err := res.LastInsertId(); err == nil {
			fid(lii)
		}
	}
	return nil
}

// UpdateAndSaveX2 ...
func UpdateAndSaveTIC(sqlx *sqlx.DB, srs interface{}, sic *TableIdxColumn, fres func(interface{})) error {
	if sqr, sps, err := CreateUpdateSQLByNamedAndSkipNilAndSet(*sic, srs); err != nil {
		return err
	} else if res, err := sqlx.NamedExec(sqr, sps); err != nil {
		return err
	} else if fres != nil {
		fres(res)
	}
	return nil
}
