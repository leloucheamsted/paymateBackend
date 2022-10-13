package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

func (fam *FirebaseAuthMiddleware) VerifyToken(c *gin.Context) {

	token := c.Param("token")
	fmt.Println(token)

	_, err := fam.cli.VerifyIDToken(context.Background(), token)

	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("Token is valid")
	c.JSON(http.StatusOK, gin.H{"error": "", "message": "token valid"})

}

type FirebaseAuthMiddleware struct {
	cli *auth.Client
}

// New is constructor of the middleware
func NewMiddleware(app *firebase.App) (*FirebaseAuthMiddleware, error) {

	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return &FirebaseAuthMiddleware{
		cli: auth,
	}, nil
}

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": code, "message": message})
}
func (fam *FirebaseAuthMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		token := strings.Replace(authHeader, "Bearer ", "", 1)
		idToken, err := fam.cli.VerifyIDToken(context.Background(), token)

		if err != nil {
			RespondWithError(c, http.StatusUnauthorized, "Invalid API token")
			return
		}

		log.Println("User ID is " + idToken.UID)
		//c.Set(valName, idToken.UID)
		c.Next()
	}
}
