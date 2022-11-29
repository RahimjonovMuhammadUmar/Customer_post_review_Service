package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"exam/api_gateway/api/handler/models"
	token "exam/api_gateway/api/token"
	"exam/api_gateway/config"

	"github.com/casbin/casbin/v2"
	jwtg "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtRoleAuth struct {
	enforcer   *casbin.Enforcer
	cnf        config.Config
	jwtHandler token.JWTHandler
}

func NewAuth(enforce *casbin.Enforcer, jwtHandler token.JWTHandler, cfg config.Config) gin.HandlerFunc {
	a := &JwtRoleAuth{
		enforcer:   enforce,
		cnf:        cfg,
		jwtHandler: jwtHandler,
	}

	return func(c *gin.Context) {
		allow, err := a.CheckPermission(c.Request)
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n")

		fmt.Println(allow)
		if err != nil {
			v, _ := err.(*jwtg.ValidationError)
			if v.Errors == jwtg.ValidationErrorExpired {
				a.RequireRefresh(c)
			} else {
				a.RequirePermission(c)
			}
		} else if !allow {
			a.RequirePermission(c)
		}
	}
}

// GetRole gets role from Authorization header if there is a token then it is
// parsed and in role got from role claim. If there is no token then role is
// unauthorized
func (a *JwtRoleAuth) GetRole(r *http.Request) (string, error) {
	var (
		role   string
		claims jwtg.MapClaims
		err    error
	)

	jwtToken := r.Header.Get("Authorization")
	if jwtToken == "" {
		return "unauthorized", nil
	} else if strings.Contains(jwtToken, "Basic") {
		return "unauthorized", nil
	}

	a.jwtHandler.Token = jwtToken
	claims, err = a.jwtHandler.ExtractClaims()
	if err != nil {
		return "", err
	}
	if claims["role"].(string) == "authorized" {
		role = "authorized"
	} else if claims["role"].(string) == "admin" {
		role = "admin"
	} else if claims["role"].(string) == "moderator" {
		role = "moderator"
	} else {
		role = "unknown"
	}
	fmt.Println(role)
	return role, nil
}

// CheckPermission checks whether user is allowed to use certain endpoint
func (a *JwtRoleAuth) CheckPermission(r *http.Request) (bool, error) {
	user, err := a.GetRole(r)
	if err != nil {
		return false, err
	}
	method := r.Method
	path := r.URL.Path
	fmt.Println(user, path, method)
	fmt.Printf("\n\n\n\n\n\n\n\n\n\n")
	allowed, err := a.enforcer.Enforce(user, path, method)
	if err != nil {
		panic(err)
	}

	return allowed, nil
}

// RequirePermission aborts request with 403 status
func (a *JwtRoleAuth) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(403)
}

// RequireRefresh aborts request with 401 status
func (a *JwtRoleAuth) RequireRefresh(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, models.ResponseError{
		Error: models.ServerError{
			Status:  "UNAUTHORIZED",
			Message: "Token is expired",
		},
	})
	c.AbortWithStatus(401)
}
