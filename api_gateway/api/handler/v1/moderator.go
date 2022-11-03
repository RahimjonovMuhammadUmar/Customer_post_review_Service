package v1

import (
	"context"
	"exam/api_gateway/api/handler/models"
	pbc "exam/api_gateway/genproto/customer"
	l "exam/api_gateway/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login moderator
// @Summary      Login moderator
// @Description  Logins moderator
// @Tags         Moderato
// @Accept       json
// @Produce      json
// @Param        username  path string true "username"
// @Param        password   path string true "password"
// @Success         200                   {object}  models.adminResponse
// @Failure         500                   {object}  models.Error
// @Failure         400                   {object}  models.Error
// @Failure         404                   {object}  models.Error
// @Failure         409                   {object}  models.Error
// @Router      /v1/moderator/login/{username}/{password} [get]
func (h *handlerV1) ModeratorLogin(c *gin.Context) {
	username := c.Param("username")
	moderator, err := h.serviceManager.CustomerService().IsModerator(context.Background(), &pbc.Admin{
		Username: username,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error moderator not found", l.Error(err))
		return
	}
	password := c.Param("password")
	if password != moderator.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect password",
		})
		h.log.Error("error wrong password", l.Error(err))
		return

	}
	// a, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	// fmt.Println(string(password))
	// err = bcrypt.CompareHashAndPassword([]byte(password), []byte(admin.Password))
	responseModerator := models.AdminResponse{
		Username: username,
		Password: "",
	}
	// Generating refresh and jwt tokens
	h.jwthandler.Iss = "moderator"
	h.jwthandler.Sub = "500"
	h.jwthandler.Role = "moderator"
	h.jwthandler.Aud = []string{"exam-app"}
	h.jwthandler.SignInKey = "UmarSecret"
	h.jwthandler.Log = h.log
	accessToken, _, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		h.log.Error("error occured while generating tokens to admin")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong,please try again",
		})
		return
	}

	responseModerator.AccessToken = accessToken
	c.JSON(http.StatusOK, responseModerator)
}
