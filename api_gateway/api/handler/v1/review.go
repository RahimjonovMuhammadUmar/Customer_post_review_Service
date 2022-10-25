package v1

import (
	"context"
	pbr "exam/api_gateway/genproto/review"
	l "exam/api_gateway/pkg/logger"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
)

// @Summary create review api
// @Description this api creates new review
// @Tags review
// @Accept json
// @Produce json
// @Param review body review.ReviewRequest true "Review"
// @Success 201 {json} review.Review
// @Router /v1/review [post]
func (h *handlerV1) CreateReview(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
		body        pbr.ReviewRequest
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json CreateReview", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	createdReview, err := h.serviceManager.ReviewService().CreateReview(ctx, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed createdReview, err := h.serviceManager.ReviewService().CreateReview(ctx, &body)", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, createdReview)
}

// DeleteReview deletes customer
// @Summary Delete review api
// @Description this api deletes review from database
// @Tags review
// @Accept json
// @Product json
// @Param id path int true "id"
// @Succes 200 {json} review.Empty
// @Router /v1/review/{id} [delete]
func (h *handlerV1) DeleteReview(c *gin.Context) {
	review_idStr := c.Param("id")
	review_id, err := strconv.ParseInt(review_idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert id to int64 DeleteReview", l.Error(err))
		return
	}
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	is_deleted, err := h.serviceManager.ReviewService().DeleteReview(ctx, &pbr.ReviewId{
		Id: int32(review_id),
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

// GetReview gets review
// @Summary Get review api
// @Description this api gets review from database
// @Tags review
// @Accept json
// @Product json
// @Param id path int true "id"
// @Succes 200 {json} review.Review
// @Router /v1/review/{id} [get]
func (h *handlerV1) GetReview(c *gin.Context) {
	review_idStr := c.Param("id")
	review_id, err := strconv.ParseInt(review_idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to convert id to int64 GetReview", l.Error(err))
		return
	}
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	review, err := h.serviceManager.ReviewService().GetReview(ctx, &pbr.ReviewId{
		Id: int32(review_id),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to send id to GetCustomer", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, review)
}
