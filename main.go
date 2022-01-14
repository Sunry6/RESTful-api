package main

import (
	"github.com/gin-gonic/gin"
	"go/found/restful/router"
	"log"
	"net/http"
)

func main() {
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	// 路由
	router.Load(g, middlewares...)

	go func() {

		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}
