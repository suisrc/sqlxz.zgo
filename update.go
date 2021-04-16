package sqlxz

// UpdateAndSaveX ...
func UpdateAndSaveX(sqlx ExecDB, srs interface{}, tbl string, sid int64, fid func(int64)) (rsa int64, err error) {
	sic := TableIdxColumn{Table: TBL + tbl, IDVal: sid}
	if sqr, sps, err := CreateUpdateSQLByNamedAndSkipNilAndSet(sic, srs); err != nil {
		return 0, err
	} else if res, err := sqlx.NamedExec(sqr, sps); err != nil {
		return 0, err
	} else if cii, err := res.RowsAffected(); err != nil {
		return 0, err
	} else if cii == 0 { // 操作失败, 没有更新
		return 0, nil
	} else if sid == 0 { // install id
		if nid, err := res.LastInsertId(); err != nil {
			return cii, err
		} else {
			fid(nid)
		}
		return cii, nil
	} else {
		return cii, nil
	}
}

// UpdateAndSaveX2 ...
func UpdateAndSaveTIC(sqlx ExecDB, srs interface{}, sic *TableIdxColumn, fres func(interface{})) (rsa int64, err error) {
	if sqr, sps, err := CreateUpdateSQLByNamedAndSkipNilAndSet(*sic, srs); err != nil {
		return 0, err
	} else if res, err := sqlx.NamedExec(sqr, sps); err != nil {
		return 0, err
	} else if cii, err := res.RowsAffected(); err != nil {
		return 0, err
	} else if cii == 0 { // 操作失败, 没有更新
		return 0, nil
	} else if fres != nil {
		fres(res)
		return cii, nil
	} else {
		return cii, nil
	}
}
