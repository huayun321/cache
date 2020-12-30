package main

import (
	"github.com/huayun321/cache/cache"
	"github.com/huayun321/cache/http"
	"github.com/huayun321/cache/tcp"
)

func main() {
	ca := cache.New("inmemory")
	go tcp.New(ca).Listen()
	http.NewServer(ca).Listen()
}
