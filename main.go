package main

import (
	cache2 "github.com/huayun321/cache/cache"
	"github.com/huayun321/cache/http"
)

func main() {
	cache := cache2.New("inmemory")
	http.NewServer(cache).Listen()
}
