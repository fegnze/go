package db

import "fmt"

//VersionData 数据库映射
type VersionData struct {
	Release string `json:"release"`
	Test    string `json:"test"`
}

//GetGameVersion -获取版本号
func GetGameVersion(ret *VersionData) error {
	sql := fmt.Sprint("SELECT * FROM version")
	return DB.QueryRow(sql).Scan(&ret.Release, &ret.Test)
}
