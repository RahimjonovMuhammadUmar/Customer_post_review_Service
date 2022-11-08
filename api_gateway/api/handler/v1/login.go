package v1

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"exam/api_gateway/api/handler/models"
	pbc "exam/api_gateway/genproto/customer"
	"exam/api_gateway/pkg/logger"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/encoding/protojson"
)

// Login customer
// @Summary      Login customer
// @Description  Logins customer
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        email  path string true "email"
// @Param        password   path string true "password"
// @Success         200                   {object}  customer.CustomerWithoutPost
// @Failure         500                   {object}  models.Error
// @Failure         400                   {object}  models.Error
// @Failure         404                   {object}  models.Error
// @Router      /v1/login/{email}/{password} [get]
func (h *handlerV1) Login(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	email := c.Param("email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	res, err := h.serviceManager.CustomerService().GetCustomerForLogin(ctx, &pbc.Email{
		Email: email,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.Error{
			Error:       err,
			Description: "Couln't find matching information, Have you registered before?",
		})
		h.log.Error("Error while getting customer by email", logger.Any("post", err))
		return
	}
	password := c.Param("password")
	err = bcrypt.CompareHashAndPassword([]byte(res.PhoneNumber), []byte(password))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Error{
			Error:       err,
			Description: "Error comparing passwords",
		})
		h.log.Error("Error while comparing hashed passwords", logger.Any("login", err))
		return
	}
	h.jwthandler.Iss = "user"
	h.jwthandler.Sub = strconv.Itoa(int(res.Id))
	h.jwthandler.Role = "authorized"
	h.jwthandler.Aud = []string{"exam-app"}
	h.jwthandler.SignInKey = h.cfg.SignInKey
	h.jwthandler.Log = h.log
	AccesToken, RefreshToken, err := h.jwthandler.GenerateAuthJWT()
	accessToken := AccesToken
	refreshToken := RefreshToken
	
	if err != nil {
		h.log.Error("error occured while generating tokens")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong,please try again",
		})
		return
	}
	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	res.PhoneNumber = ""
	c.JSON(http.StatusOK, res)
}
