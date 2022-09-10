package middlewares

import (
	"log"
	"reflect"

	"github.com/Leonardo-lucas-pereira/tcc-api/services"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Beare_schema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Beare_schema):]
		claim, _ := services.NewJWTServices().ExtractClaims(token)
		log.Print(reflect.TypeOf(claim))
		log.Print(claim)
		log.Print(claim["sum"])
		log.Print(claim["is_adm"])

		if !services.NewJWTServices().ValidateToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
