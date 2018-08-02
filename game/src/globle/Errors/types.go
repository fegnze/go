package Errors

/*Errors 导出*/
var Errors = map[string]string{"test1": "error1"}

func ts() {
	Errors["test1"] = "错误1"
}

func main() {
	ts()
}
