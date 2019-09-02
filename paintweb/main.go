package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

<<<<<<< HEAD
	"github.com/MDGSF/go-paint/paintdom"
=======
	"github.com/qiniu/qpaint/paintdom"
>>>>>>> 4089f4170c5ffb5be6ad5b55f4d7494391d6a47c
)

func newReverseProxy(baseURL string) *httputil.ReverseProxy {
	rpURL, _ := url.Parse(baseURL)
	return httputil.NewSingleHostReverseProxy(rpURL)
}

func handleDefault(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
<<<<<<< HEAD
		http.ServeFile(w, req, "www/index.html")
=======
		http.ServeFile(w, req, "www/index.htm")
>>>>>>> 4089f4170c5ffb5be6ad5b55f4d7494391d6a47c
		return
	}
	req.URL.RawQuery = "" // skip "?params"
	wwwServer.ServeHTTP(w, req)
}

var (
	apiReverseProxy = newReverseProxy("http://localhost:9999")
	wwwServer       = http.FileServer(http.Dir("www"))
)

func main() {
	go paintdom.Main()
	http.Handle("/api/", http.StripPrefix("/api/", apiReverseProxy))
	http.HandleFunc("/", handleDefault)
	http.ListenAndServe(":8888", nil)
}
