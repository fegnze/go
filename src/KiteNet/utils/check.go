package utils

import (
	"KiteNet/log"
	"errors"
)

//CheckErrA 简单检查err
func CheckErrA(param interface{}, err error, msg ...interface{}) interface{} {
	if err != nil {
		glog.Deepin(3)
		glog.Error(err)
	}
	return param
}

//CheckErr 简单断言err
func CheckErr(err error, msg ...interface{}) bool {
	if err != nil {
		glog.Deepin(3)
		if msg == nil {
			glog.Error(err, "undefined error!")
		} else {
			glog.Error(err, msg...)
		}
		return true
	}
	return false
}

//CheckNil 断言空值
func CheckNil(v interface{}, msg ...interface{}) bool {
	if v == nil {
		glog.Deepin(3)
		glog.Error(errors.New("Unexpected nil value!"))
		return true
	}
	return false
}

//CheckNilAndErr 断言空值
func CheckNilAndErr(v interface{}, err error, msg ...interface{}) bool {
	glog.Deepin(4)
	if CheckErr(err,msg) {
		return true
	}
	glog.Deepin(4)
	return CheckNil(v,msg)
}

//CheckNilMsg 断言空字符串
func CheckNilMsg(s string, msg ...interface{}) bool {
	if s == "" {
		glog.Deepin(3)
		glog.Error(errors.New("Unexpected empty string!"))
		return true
	}
	return false
}
