package main

import (
	"ibsen-back-go/authors"
	"ibsen-back-go/common"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB("bolt://127.0.0.1:11003", "neo4j", "beilin123")
	r := gin.Default()
	v1 := r.Group("api")
	authors.AuthorsRegister(v1)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
