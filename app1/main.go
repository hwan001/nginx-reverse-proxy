package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 새로운 Gin 엔진 생성
	router := gin.Default()
	BasePath := "/app1"

	// 기본 GET 엔드포인트
	router.GET(BasePath+"/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello, this is " + BasePath,
		})
	})

	// 기본 GET 엔드포인트
	router.GET(BasePath+"/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// POST 엔드포인트 예제
	router.POST(BasePath+"/echo", func(c *gin.Context) {
		var json map[string]interface{}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"received": json,
		})
	})

	// 서버 실행
	router.Run(":8080") // 기본 포트 8080에서 서버 실행
}
