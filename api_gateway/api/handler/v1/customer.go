package v1

import (
	"context"
	pbc "exam/api_gateway/genproto/customer"
	l "exam/api_gateway/pkg/logger"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// // Create creates customer
// // @Summary create customer api
// // @Description this api creates new customer
// // @Tags customer
// // @Accept json
// // @Produce json
// // @Param customer body customer.CustomerRequest true "Customer"
// // @Success 201 {json} customer.CustomerWithoutPost
// // @Router /v1/customer [post]
// func (h *handlerV1) CreateCustomer(c *gin.Context) {
// 	var (
// 		body        pbc.CustomerRequest
// 		jspbMarshal protojson.MarshalOptions
// 	)
// 	jspbMarshal.UseProtoNames = true
// 	err := c.ShouldBindJSON(&body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("failed to bind json", l.Error(err))
// 		return
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
// 	defer cancel()
// 	response, err := h.serviceManager.CustomerService().CreateCustomer(ctx, &body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		h.log.Error("error while declaring reponse api/hanlder/v1/customer.go", l.Error(err))
// 		return
// 	}
// 	c.JSON(http.StatusCreated, response)
// }

// Get finds customer
// @Summary get customer api
// @Description this api finds existing customer
// @Tags customer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 201 {json} customer.Customer
// @Router /v1/customer/{id} [get]
func (h *handlerV1) GetCustomer(c *gin.Context) {
	customer_idStr := c.Param("id")
	customer_id, err := strconv.ParseInt(customer_idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert id to int64 GetCustomer", l.Error(err))
		return
	}
	var jspbMarshal protojson.MarshalOptions

	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	customer, err := h.serviceManager.CustomerService().GetCustomer(ctx, &pbc.CustomerId{
		Id: int32(customer_id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get from customer from customer service", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, customer)
}

// UpdateCustomer Update updates customer
// @Summary Update customer api
// @Description this api updates customer by id in database
// @Tags customer
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param customer body customer.CustomerWithoutPost true "Customer"
// @Success 200 {json} customer.CustomerWithoutPost
// @Router /v1/customer [put]
func (h *handlerV1) UpdateCustomer(c *gin.Context) {
	var (
		body        pbc.CustomerWithoutPost
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error to bind json UpdateCustomer", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	updated_product, err := h.serviceManager.CustomerService().UpdateCustomer(ctx, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update product inside UpdateProduct", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, updated_product)
}

// DeleteCustomer deletes customer
// @Summary Delete customer api
// @Description this api deletes customer from database
// @Tags customer
// @Security        BearerAuth
// @Accept json
// @Product json
// @Param id path int true "id"
// @Succes 200 {json} customer.CustomerDeleted
// @Router /v1/customer/{id} [delete]
func (h *handlerV1) DeleteCustomer(c *gin.Context) {
	customer_idStr := c.Param("id")
	customer_id, err := strconv.ParseInt(customer_idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert id to int64 DeleteCustomer", l.Error(err))
		return
	}
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	is_deleted, err := h.serviceManager.CustomerService().DeleteCustomer(ctx, &pbc.CustomerId{
		Id: int32(customer_id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to send id to deleteCustomer", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, is_deleted)
}

// // Search searches alike customers
// @Summary search customer api
// @Description this api searches customer
// @Tags customer
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param limit  query int true "Limit"
// @Param page   query int true "Page"
// @Param orderBy query  string true "Order:DescOrAsc" example = "last_name:desc"
// @Param fieldValue  query string true "Field:Value" example = "first_name:asl"
// @Success 200 {json} customer.PossibleCustomers
// @Failure 400
// @Failure 500
// @Router /v1/customer/search [get]
func (h *handlerV1) SearchCustomer(c *gin.Context) {
	params := c.Request.URL.Query()
	fieldWithValues := strings.Split(params["fieldValue"][0], ":")
	orderType := strings.Split(params["orderBy"][0], ":")

	limit, err := strconv.Atoi(params["limit"][0])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error converting limit to int", l.Error(err))
		return
	}

	page, err := strconv.Atoi(params["page"][0])
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error converting page to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	possible_customers, err := h.serviceManager.CustomerService().SearchCustomer(ctx, &pbc.InfoForSearch{
		Field:     fieldWithValues[0],
		Value:     fieldWithValues[1],
		Limit:     int32(limit),
		Page:      int32(page),
		OrderBy:   orderType[1],
		AscOrDesc: orderType[1],
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while sending to search service to customerService", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, possible_customers)
}
