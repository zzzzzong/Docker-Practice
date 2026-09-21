package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func newProxy(targetURL string) *httputil.ReverseProxy {
	target, _ := url.Parse(targetURL)
	return httputil.NewSingleHostReverseProxy(target)
}

// corsAndLogMiddleware 處理 CORS 並印出 redirect 資訊
func corsAndLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 根據路徑印出 log
		if strings.HasPrefix(r.URL.Path, "/api/fastapi") {
			log.Println("redirect to python")
		} else if strings.HasPrefix(r.URL.Path, "/api/go") {
			log.Println("redirect to go")
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/api/fastapi", newProxy("http://127.0.0.1:8001"))
	mux.Handle("/api/go", newProxy("http://127.0.0.1:8002"))

	log.Println("Gateway listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsAndLogMiddleware(mux)))
}
