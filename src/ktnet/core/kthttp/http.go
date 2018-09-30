package kthttp

import (
	"encoding/json"
	"gamecore/core/Ktlog"
	"net/http"
	"net/url"
)

//ProtoInterface 路由
type ProtoInterface interface {
	excute(query *url.Values) ProtoError
	response(w *http.ResponseWriter)
}

//ProtoError error
type ProtoError struct {
	Code rune
	Msg  string
}

//Server KtHttp server实例
type Server struct {
	protoAcount int
	httpServer  *http.Server
	mux         *http.ServeMux
}

//RegistRout 注册路由表
func (server *Server) RegistRout(path string, proto ProtoInterface, method string) {
	//mux.Handle("/", http.HandlerFunc(testHandler))
	//mux.HandleFunc("/name", testHandler)
	server.protoAcount++
	server.mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if method == "get" {
			queryData := r.URL.Query()
			if err := proto.excute(&queryData); err.Code != 0 {
				proto.response(&w)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Content-Type", "apllication/json")
				data, jsonerr := json.Marshal(err)
				if jsonerr != nil {
					data, _ = json.Marshal(ProtoError{
						Code: 1000,
						Msg:  "protol excute return failed,and parse to jsonString failed!",
					})
				}
				if _, error := w.Write(data); error != nil {
					Ktlog.Error("http get [%s] write error:%s", path, error.Error())
				}

			}
		} else if method == "post" {

		}

	})
}

//RegistServer 注册一个http server并监听一个端口
func RegistServer(hs *http.Server) *Server {
	// hs := {
	// 	//Addr: "localhost:3000",
	// 	//Handler: mux,
	// 	// TLSConfig:    *tls.Config,
	// 	// ReadTimeout:    time.Duration,
	// 	// ReadHeaderTimeout:    time.Duration,
	// 	//WriteTimeout: 4 * time.Second,
	// 	// IdleTimeout:    time.Duration,
	// 	// MaxHeaderBytes:    int,
	// 	// TLSNextProto:    map[string]func(*Server, *tls.Conn, Handler),
	// 	// ConnState:    func(net.Conn, ConnState),
	// 	// ErrorLog:    *log.Logger,
	// 	// disableKeepAlives:    int32,
	// 	// inShutdown:        int32,
	// 	// nextProtoOnce:        sync.Once,
	// 	// nextProtoErr:        error,
	// 	// mu:        sync.Mutex,
	// 	// listeners:    map[net.Listener]struct{},
	// 	// activeConn:    map[*conn]struct{},
	// 	// doneChan:    chan struct{},
	// 	// onShutdown:    []func(),
	// }
	mux := http.NewServeMux()

	return &Server{
		httpServer:  hs,
		mux:         mux,
		protoAcount: 0,
	}
}

//StartService 开启http服务
func (server *Server) StartService() {
	if server.httpServer == nil || server.httpServer.Addr == "" {
		Ktlog.Panic("http server Start failed.")
		return
	}
	if server.protoAcount <= 0 {
		Ktlog.Error("http server has no proto registed.")
	}
	Ktlog.Info("Listening on %s...", server.httpServer.Addr)
	//会一直阻塞执行
	if err := server.httpServer.ListenAndServe(); err != nil {
		Ktlog.Info("Listening error", err.Error())
	}
}

// func testHandler(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query()
// 	r.ParseForm()
// 	defer r.Body.Close()
// 	con, _ := ioutil.ReadAll(r.Body)
// 	fmt.Println(r.Method, query["pid"], query["role"], r.PostForm, string(con))
// 	type testRet struct {
// 		Name string `json:"name"`
// 		Age  rune   `json:"age"`
// 		Sex  string `json:"sex"`
// 	}
// 	var test testRet
// 	test.Name = "小明"
// 	test.Age = 20
// 	test.Sex = "男"
// 	data, _ := json.Marshal(test)

// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "apllication/json")
// 	ret, error := w.Write(data)
// 	Ktlog.Info(string(ret), error)
// }
