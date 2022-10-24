package v1

import (
	"context"
	pbc "exam/api_gateway/genproto/customer"
	l "exam/api_gateway/pkg/logger"
	"net/http"
	"strconv"
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

func (h *handlerV1) GetCustomer(c *gin.Context) {
	customer_idStr := c.Param("id")
	customer_id, err := strconv.ParseInt(customer_idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert id to int64", l.Error(err))
		return
	}
	var jspbMarshal protojson.MarshalOptions

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	_, err = h.serviceManager.CustomerService().GetCustomer(ctx, &pbc.CustomerId{
		Id: int32(customer_id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get from customer from customer service", l.Error(err))
		return
	}
		

}
