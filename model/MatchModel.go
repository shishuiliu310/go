package model

//修改歌曲赞数
func AddPraiseNum(sql string) bool {
	rows, err := db.Query(sql)
	defer rows.Close()
	if err == nil {
		return true
	} else {
		return false
	}

}
