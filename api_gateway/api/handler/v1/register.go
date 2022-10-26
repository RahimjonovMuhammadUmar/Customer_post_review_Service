package v1

import (
	"context"
	"encoding/json"
	"exam/api_gateway/api/handler/models"
	pbc "exam/api_gateway/genproto/customer"
	"fmt"

	l "exam/api_gateway/pkg/logger"
	"exam/api_gateway/pkg/utils"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/encoding/protojson"
)

// Register godoc
// @Summary Register for authentication
// @Tags Auth
// @Accept json
// @Produce json
// @Param userData body models.UserRegister true "user data"
// @Success 200 "Message sended to your email succesfully"
// @Failure 400 {object} models.Error
// @Router /register [post]
func (h *handlerV1) RegisterCustomer(c *gin.Context) {
	var (
		newUser     models.UserRegister
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

	exist, err = h.serviceManager.CustomerService().CheckField(ctx, &pbc.FieldCheck{
		Field:           "username",
		EmailOrUsername: newUser.Username,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while checking if username exists", l.Error(err))
		return
	}

	if exist.Exists {
		c.JSON(http.StatusBadRequest, "Registration failed, account with such username already exists")
		return
	}
	hashPass, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		h.log.Error("error while hashing password", l.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong",
		})
		return
	}
	newUser.Password = string(hashPass)

	jsNewUser, err := json.Marshal(newUser)
	if err != nil {
		h.log.Error("error while marshaling new user, inorder to insert it to redis", l.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while creating user",
		})
		return
	}

	code := utils.RandomNum()

	_, err = h.InMemoryStorage.Get(fmt.Sprint(code))
	if err == nil {
		code = utils.RandomNum()
	}

	if err = h.InMemoryStorage.SetWithTTl(fmt.Sprint(code), string(jsNewUser), 86000); err != nil {
		fmt.Println(err)
		h.log.Error("error while inserting new user into redis")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong, please try again",
		})
		return
	}

	newUser.Email = email

	_, cancel = context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	res, err := EmailVerification("Verigication", fmt.Sprint(code), email)
	if err != nil {

		h.log.Error("error while sending verification code to new user", l.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went wrong, please try again",
		})
		return
	}

	c.JSON(http.StatusOK, res)

}

func EmailVerification(subject, code, email string) (string, error) {

	// Sender data.
	from := "asliddinvstalim@gmail.com"
	password := "gnradbxvloedrkti"

	// Receiver email address.
	to := []string{
		"rahimzanovmuhammadumar@gmail.com",
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
