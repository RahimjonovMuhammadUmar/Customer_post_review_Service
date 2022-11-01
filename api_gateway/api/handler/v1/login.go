package v1

// import (
// 	"context"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"exam/api_gateway/api/handler/models"
// 	"exam/api_gateway/genproto/customer"
// 	"exam/api_gateway/pkg/etc"
// 	"exam/api_gateway/pkg/logger"
// 	"google.golang.org/protobuf/encoding/protojson"
// )

// // Login customer
// // @Summary      Login customer
// // @Description  Logins customer
// // @Tags         Customer
// // @Accept       json
// // @Produce      json
// // @Param        email  path string true "email"
// // @Param        password   path string true "password"
// // @Success         200                   {object}  customer.LoginRes
// // @Failure         500                   {object}  models.Error
// // @Failure         400                   {object}  models.Error 
// // @Failure         404                   {object}  models.Error 
// // @Router      /login/{email}/{password} [get]
// func (h *handlerV1) Login(c *gin.Context) {
// 	var jspbMarshal protojson.MarshalOptions
// 	jspbMarshal.UseProtoNames = true

// 	var (
// 		password = c.Param("password")
// 		email    = c.Param("email")
// 	)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()

// 	res, err := h.serviceManager.CustomerService().SearchCustomer(ctx, &customer.InfoForSearch{
		
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, models.Error{
// 			Error:       err,
// 			Description: "Couln't find matching information, Have you registered before?",
// 		})
// 		h.log.Error("Error while getting customer by email", logger.Any("post", err))
// 		return
// 	}

// 	if !etc.CheckPasswordHash(password, res.Password) {
// 		c.JSON(http.StatusNotFound, models.Error{
// 			Description: "Password or Email error",
// 			Code:        http.StatusBadRequest,
// 		})
// 		return
// 	}

// 	h.jwthandler.Iss = "user"
// 	h.jwthandler.Sub = res.Id
// 	h.jwthandler.Role = "authorized"
// 	h.jwthandler.Aud = []string{"exam-app"}
// 	h.jwthandler.SignInKey = h.cfg.SignInKey
// 	h.jwthandler.Log = h.log
// 	tokens, err := h.jwthandler.GenerateAuthJWT()
// 	accessToken := tokens[0]
// 	refreshToken := tokens[1]

// 	if err != nil {
// 		h.log.Error("error occured while generating tokens")
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "something went wrong,please try again",
// 		})
// 		return
// 	}
// 	res.AccessToken = accessToken
// 	res.Refreshtoken = refreshToken
// 	res.Password = ""
// 	c.JSON(http.StatusOK, res)
// }
