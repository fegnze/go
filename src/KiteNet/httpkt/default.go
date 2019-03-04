package httpkt

import (
	"KiteNet/log"
	"KiteNet/utils"
	"net/http"
	"path/filepath"
)

//FaviconHandler 默认处理函数
func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	glog.Debug("FaviconHandler Exec ...")
	dir, err := filepath.Abs("./")
	utils.CheckErr(err)
	http.ServeFile(w,r, dir+"/favicon.ico")
}