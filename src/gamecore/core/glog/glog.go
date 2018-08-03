package glog

// //Debug debug
// func Debug() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			log.Println("|E|", r)
// 		}
// 	}()

// 	log.Print("This is a log to print", 1, "\n")
// 	log.Panicf("This is a log to  panicf,%d", 2) //抛出异常，有recover的话，会被捕获
// 	log.Fatalln("This is a log to println", 3)   //直接中断
// }

//Error for log
func Error(s string) {

}

//Debug for log
func Debug(s string) {

}

//Info for log
func Info(s string) {

}

//Verbose for log
func Verbose(s string) {

}
