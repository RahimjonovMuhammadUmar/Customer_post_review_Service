package v1

import (
	"context"
	"encoding/json"
	"exam/api_gateway/api/handler/models"
	pbc "exam/api_gateway/genproto/customer"
	"fmt"
	"strconv"

	l "exam/api_gateway/pkg/logger"
	"exam/api_gateway/pkg/utils"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/encoding/protojson"
)

// Register customer
// @Summary Register for authentication
// @Tags Auth
// @Accept json
// @Produce json
// @Param userData body models.CustomerRegister true "user data"
// @Success 200 "Message sended to your email succesfully"
// @Failure 400 {object} models.Error
// @Router /v1/register [post]
func (h *handlerV1) RegisterCustomer(c *gin.Context) {
	var (
		newUser     models.CustomerRegister
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	email, err := utils.IsValidMail(newUser.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email address",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	exist, err := h.serviceManager.CustomerService().CheckField(ctx, &pbc.FieldCheck{
		Field:           "email",
		EmailOrUsername: newUser.Email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while checking if email exists", l.Error(err))
		return
	}

	if exist.Exists {
		c.JSON(http.StatusBadRequest, "Registration failed, account with such email already exists")
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		h.log.Error("error while hashing password", l.Error(err))
		fmt.Println("error -> register.go 80 ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	newUser.Password = string(hashPass)
	code := utils.RandomNum(6)
	customerData := models.CustomerDataToSave{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Bio:       newUser.Bio,
		Password:  newUser.Password,
		Email:     newUser.Email,
		Addresses: newUser.Addresses,
		Code:      code,
	}
	_, err = h.InMemoryStorage.Get(fmt.Sprint(customerData.Email))
	if err != nil {
		h.log.Error("error while checking if key with such email exists in redis", l.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while creating user",
		})
		return
	}
	jsNewUser, err := json.Marshal(customerData)
	if err != nil {
		h.log.Error("error while marshaling new user, inorder to insert it to redis", l.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while creating user",
		})
		return
	}
	fmt.Println(code)
	if err = h.InMemoryStorage.SetWithTTl(fmt.Sprint(customerData.Email), string(jsNewUser), 86000); err != nil {
		h.log.Error("error while inserting new user into redis")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong, please try again",
		})
		return
	}

	newUser.Email = email
	_, cancel = context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	fmt.Println(customerData.Code)
	defer cancel()
	// res, err := EmailVerification("Verification code", fmt.Sprint(customerData.Code), email)
	// if err != nil {
	// 	fmt.Println("error is here ->", err)
	// 	h.log.Error("error while sending verification code to new user", l.Error(err))
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": "something went wrong, please try again",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, "Done")

}
func EmailVerification(subject, code, email string) (string, error) {

	// Sender data.
	from := "asliddinvstalim@gmail.com"
	password := "gnradbxvloedrkti"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(fmt.Sprintf("%s %s", subject, code))

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return "Error with sending message", err
	}
	return "Message sended to your email succesfully", nil
}

// Verify for customer
// @Summary Verify for authentication
// @Tags Auth
// @Accept json
// @Produce json
// @Param code path int true "code"
// @Param email path string true "email"
// @Success 200 {json} models.CustomerRegister
// @Failure 400 {object} models.Error
// @Router /v1/register/{code}/{email} [get]
func (h *handlerV1) VerifyRegistration(c *gin.Context) {
	input_code := c.Param("code")
	code, err := strconv.ParseInt(input_code, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert id to int64 Verify", l.Error(err))
		return
	}
	customer_email := c.Param("email")
	storedData, err := h.InMemoryStorage.Get(fmt.Sprint(customer_email))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to send code to redis", l.Error(err))
		return
	}
	data := cast.ToString(storedData)
	userInfo := models.CustomerDataToSave{}

	err = json.Unmarshal([]byte(data), &userInfo)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to send code to redis", l.Error(err))
		return
	}
	fmt.Println(userInfo.Code, code)
	if !(userInfo.Code == int(code)) {
		c.JSON(http.StatusConflict, gin.H{
			"Incorrect": "Code does not match",
		})
		h.log.Info("Incorrect input for code")
		return
	}
	customerRequest := &pbc.CustomerRequest{
		FirstName:   userInfo.FirstName,
		LastName:    userInfo.LastName,
		Bio:         userInfo.Bio,
		Email:       userInfo.Email,
		PhoneNumber: userInfo.Password,
	}
	for _, address := range userInfo.Addresses {
		customerRequest.Addresses = append(customerRequest.Addresses, &pbc.AddressRequest{
			HouseNumber: address.House_number,
			Street:      address.Street,
		})
	}
	// Generating refresh and jwt tokens
	h.jwthandler.Iss = "user"
	h.jwthandler.Sub = customerRequest.Bio
	h.jwthandler.Role = "authorized"
	h.jwthandler.Aud = []string{"exam-app"}
	h.jwthandler.SignInKey = "UmarSecret"
	h.jwthandler.Log = h.log
	accessToken, refreshToken, err := h.jwthandler.GenerateAuthJWT()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	if err != nil {
		h.log.Error("error occured while generating tokens", l.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong,please try again",
		})
		return
	}
	customerRequest.Token = refreshToken
	response, err := h.serviceManager.CustomerService().CreateCustomer(ctx, customerRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while declaring reponse api/hanlder/v1/customer.go", l.Error(err))
		return
	}
	response.RefreshToken = refreshToken
	response.AccessToken = accessToken

	c.JSON(http.StatusCreated, response)
}
