package main

import (
	"github.com/gin-gonic/gin"
	"ibsen-back-go/authors"
	"ibsen-back-go/common"
)

func main() {
	common.InitDB("bolt://39.107.25.37:27687", "neo4j", "password")
	r := gin.Default()
	v1 := r.Group("api")
	authors.AuthorsRegister(v1)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
