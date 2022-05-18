/**
 learn to use net/http
 */
//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//)
//
//func main() {
//	http.HandleFunc("/", indexHandler)
//	http.HandleFunc("/hello", helloHandler)
//	log.Fatal(http.ListenAndServe(":9999", nil))
//}
//
//func helloHandler(writer http.ResponseWriter, request *http.Request) {
//	for k, v := range request.Header {
//		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
//	}
//}
//
//func indexHandler(writer http.ResponseWriter, request *http.Request) {
//	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
//}

/**
  custom define http handle
*/
//package main

//type Engine struct{}
//
//func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	switch req.URL.Path {
//	case "/":
//		fmt.Fprintf(w, "URL.PATH = %q, HOST = %q, REMOTE_ADDR = %q\n", req.URL.Path, req.Host, req.RemoteAddr)
//	case "/hello":
//		for k, v := range req.Header {
//			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
//		}
//	default:
//		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
//
//	}
//}
//
//func main() {
//	engine := new(Engine)
//	log.Fatal(http.ListenAndServe(":9999", engine))
//}

package main

import (
	"fmt"
	"gan/HttpBase"
	"net/http"
)

func main() {
	r := HttpBase.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.GET("/hello",func (writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")
}