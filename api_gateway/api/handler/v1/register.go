package v1

import (
	pbc "exam/api_gateway/genproto/customer"
	l "exam/api_gateway/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

func (h *handlerV1) RegisterCustomer(c *gin.Context) {
	var (
		body        pbc.FieldCheck
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

}
