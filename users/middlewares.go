package users

import (
	"errors"
	"net/http"
	"simple_crud/common"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func stripBearerPrefixFromTokenString(tok string) (string, error) {
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func AuthMiddleware(auto401 bool, isAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)
		token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.NBSecretPassword))
			return b, nil
		})
		if err != nil {
			if auto401 {
				// c.AbortWithError(http.StatusUnauthorized, err)
				c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewError("error", errors.New("Unauthorized")))
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userId := int(claims["id"].(float64))
			if isAdmin {
				role := claims["role"]
				if role != "admin" {
					c.AbortWithStatusJSON(http.StatusUnauthorized, common.NewError("error", errors.New("Unauthorized")))
					return
				}
			}
			UpdateContextUserModel(c, userId)
		}
	}
}

func UpdateContextUserModel(c *gin.Context, userId int) {
	var user User
	if userId != 0 {
		db := common.GetDB()
		db.First(&user, userId)
	}
	c.Set("user_id", userId)
	c.Set("user", user)
}
