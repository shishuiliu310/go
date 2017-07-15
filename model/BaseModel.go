package model

import (
	"conf"
	//"fmt"
)

var (
	db    = conf.Open()
	count int
)

//将从MySQL中读取到的数量
func Count(sql string) int {
	rows, _ := db.Query(sql)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&count)
	}
	return count
}

//将从MySQL中读取到的数据生成切片
func DataToSlice(sql string) []map[string]string {
	rows, _ := db.Query(sql)
	defer rows.Close()
	s := make([]map[string]string, 0, 0) //声明切片用来存储用户数据
	cols, _ := rows.Columns()
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	scanArgs := make([]interface{}, len(cols))
	values := make([]interface{}, len(cols))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		records := make(map[string]string)
		//将行数据保存到records字典,records不能放在for外部不然会引用同一个地址将其records中原来的值覆盖掉
		rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				records[cols[i]] = string(col.([]byte))
			}
		}
		s = append(s, records)
	}
	return s
}

//使用copy向切片中追加元素
func AppendString(slice []map[string]string, data ...map[string]string) []map[string]string {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]map[string]string, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
