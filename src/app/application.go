package app

import (
	"github.com/jnre/oauth/src/clients/cassandra"
	"github.com/jnre/oauth/src/domain/access_token"
	"github.com/jnre/oauth/src/http"
	"github.com/jnre/oauth/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()	
	atService := access_token.NewService(db.NewRepository())
	atHandler := http.NewAccessTokenHandler(atService)

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token",atHandler.Create)

	router.Run(":8080")
}
