package v1

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"go-gql/gql"
	"go-gql/resolver"
	"net/http"
)

// Params graphql请求体的标准格式
type Params struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func GraphQL(c *gin.Context) {
	params := new(Params)
	err := c.ShouldBindJSON(&params)
	if err != nil {
		return
	}
	schema := graphql.MustParseSchema(gql.SchemaString, &resolver.RootResolver{})
	data := schema.Exec(c.Request.Context(), params.Query, params.OperationName, params.Variables)
	c.JSON(http.StatusOK, data)
}
