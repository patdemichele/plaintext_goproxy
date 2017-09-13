package main
import (
    "github.com/elazarl/goproxy"
    "log"
    "net/http"
    "fmt"
)

func main() {
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true
    proxy.OnRequest().DoFunc(
    func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
    	fmt.Printf("Yeah I do this")
    	r.Header.Set("X-GoProxy","yxorPoG-X")
        return r,nil
	})
    log.Fatal(http.ListenAndServe(":8080", proxy))

}