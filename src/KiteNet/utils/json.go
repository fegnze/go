package utils

import (
	"strings"
)

//JsonNormalize Json字符串规范化
func JsonNormalize(buf []byte)[]byte{
	js := string(buf)
	begin := strings.Index(js,"//")
	if begin < 0 {
		return []byte(js)
	}
	end := begin + strings.IndexAny(js[begin:],"\n")
	old := js[begin:end]
	js = strings.Replace(js,old,"",-1)
	return JsonNormalize([]byte(js))
}
