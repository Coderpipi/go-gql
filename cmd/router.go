package cmd

import (
	"github.com/gin-gonic/gin"
	"go-gql/actions/api"
	v1 "go-gql/actions/api/v1"
	"go-gql/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("ping", api.Ping)
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("hello", v1.Hello)
		apiV1.POST("graphql", v1.GraphQL)
	}
	return r
}
