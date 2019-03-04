package db

import (
	"KiteNet/log"
	"fmt"
	"time"
)

//RegionInfo -服务器列表,数据库game_region表映射
type RegionInfo struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Hot      int       `json:"hot"`
	State    int       `json:"state"`
	InitTime time.Time `json:"initTime"`
}

//GetRegionList -获取服务器列表
func GetRegionList() []RegionInfo {
	sql := fmt.Sprint("SELECT * FROM game_region")
	rows, err := DB.Query(sql)
	defer rows.Close()
	if err != nil {
		glog.Error(err,"查询服务器列表失败")
	}
	list := make([]RegionInfo, 0)
	for rows.Next() {
		region := RegionInfo{}
		if err := rows.Scan(&region.ID, &region.Name, &region.Hot, &region.State, &region.InitTime); err != nil {
			glog.Error(err,"读取服务器列表失败")
		}
		list = append(list, region)
	}

	return list
}
