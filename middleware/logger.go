package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        fmt.Printf("[%s] %s %s\n", r.Method, r.URL.Path, start.Format("2006-01-02 15:04:05"))
        next.ServeHTTP(w, r)
    })
}
