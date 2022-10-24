package v1

import (
	"context"
	pbc "exam/api_gateway/genproto/customer"
	l "exam/api_gateway/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)




// Create creates customer
// @Summary create customer api
// @Description this api creates new customer
// @Tags product
// @Accept json
// @Produce json
// @Param customer body customer.CustomerRequest true "Customer"
// @Success 201 {json} customer.Customer
// @Router /v1/customer [post]
func (h *handlerV1) CreateCustomer(c *gin.Context) {
	var (
		body        pbc.CustomerRequest
		jspbMarshal protojson.MarshalOptions
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.CustomerService().CreateCustomer(ctx, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while declaring reponse api/hanlder/v1/customer.go", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)

}
