package main
import (
    "github.com/elazarl/goproxy"
    "log"
    "net/http"
)

func main() {
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true
    proxy.OnRequest().HandleConnect(goproxy.AlwaysReject)
    log.Fatal(http.ListenAndServe(":8080", proxy))

}