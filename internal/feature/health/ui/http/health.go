package http

import "net/http"

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("OK"))
}
