package utils

import (
	"github.com/rs/xid"
	"math/rand"
	"time"
)

var r *rand.Rand

func init(){
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

//CreateUniqueCode 随机数
func CreateUniqueCode() xid.ID{
	return xid.New()
}

//CreateRandA 随机字符串
func CreateRandString(len int) string {
	buf := make([]byte,len)
	for i := 0; i < len; i++{
		n := r.Intn(26)+65
		buf[i] = byte(n)
	}
	return string(buf)
}