package main

import (
	fairy_admin "github.com/zhongsdjn/fairy-admin"
)

func main() {
	r := fairy_admin.Routers()
	err := r.Run("127.0.0.1:8080")
	if err != nil {
		return
	}
}
