package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JsonObject(key string, value string) []byte {
	mapobj := map[string]string{
		key: value,
	}
	res, _ := json.Marshal(mapobj)
	return res
}

// Todo:= add handler dir
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	res := JsonObject("status", "healthy")
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	res := JsonObject("message", "404 Page Not Found")
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

// Todo:= add fiber
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", HandleHealth)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/health" {
			NotFoundHandler(w, r)
			return
		}
	})
	fmt.Println("Server Listen : 8080")
	http.ListenAndServe(":8080", mux)
}
