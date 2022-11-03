package v1

import (
	"context"
	pbp "exam/api_gateway/genproto/post"
	l "exam/api_gateway/pkg/logger"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
)

// CreatePost creates post
// @Summary create post api
// @Description this api creates new post
// @Tags post
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param customer body post.PostRequest true "Post"
// @Success 201 {json} customer.PostWithoutReview
// @Router /v1/post [post]
func (h *handlerV1) CreatePost(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
		body        pbp.PostRequest
	)

	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json CreatePost", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	created_post, err := h.serviceManager.PostService().CreatePost(ctx, &body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error while created_post, err := h.serviceManager.PostService().CreatePost(ctx, &body)", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, created_post)
}

// GetPostWithCustomerInfo
// @Summary      Get post with customer information
// @Description  Get Post infos with id
// @Tags         post
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param        id path int true "post_id"
// @Success      200  {object}  post.PostWithCustomerInfo
// @Router       /v1/post/{id} [get]
func (h *handlerV1) GetPostWithCustomerInfo(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetPostWithCustomerInfo(
		ctx, &pbp.Ids{
			Id: int32(id),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get post")
		return
	}
	c.JSON(http.StatusOK, response)

}

// GetPostsOfCustomer
// @Summary      Gets post by customers id
// @Description  Get posts of customer
// @Tags         post
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param        id path int true  "customer_id"
// @Success      200  {object}  post.Posts
// @Router       /v1/post/customers_posts/{id} [get]
func (h *handlerV1) GetPostsOfCustomer(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().GetPostsOfCustomer(
		ctx, &pbp.Id{
			Id: int32(id),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get posts of customer")
		return
	}
	c.JSON(http.StatusOK, response)

}

// UpdatePost
// @Summary      Updates post by id
// @Description  update post api
// @Tags         post
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param        post body post.PostWithoutReview true "Update post by id"
// @Success      200  {object}  post.PostWithoutReview
// @Router       /v1/post [put]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	var (
		jspbMarshal protojson.MarshalOptions
		body        pbp.PostWithoutReview
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	claims, _ := h.GetSub(c)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	_, err = h.serviceManager.PostService().GetPostInfoOnly(ctx, &pbp.Id{
		Id: body.Id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("error getting post by id to compare ids of customer for updating", l.Error(err))
		return
	}
	idFromToken := cast.ToInt32(claims["sub"])
	if idFromToken != 500 && idFromToken != 999 {
		c.JSON(http.StatusForbidden, "You are not the owner of this post!!")
		return
	}

	response, err := h.serviceManager.PostService().UpdatePost(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update customer")
		return
	}

	c.JSON(http.StatusOK, response)

}

// DeletePost
// @Summary      Delete post from database
// @Description  Delete Post and it's reviews by Id
// @Tags         post
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param        id path int true "post_id"
// @Success      200
// @Router       /v1/post/{id} [delete]
func (h *handlerV1) DeletePost(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeletePost(
		ctx, &pbp.Id{
			Id: int32(id),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get product")
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeletePostByCustomerId
// @Summary      Delete customers posts
// @Description  Delete Post by Customer Id
// @Tags         post
// @Security BearerAuth
// @Accept       json
// @Produce      json
// @Param        id path int true "customer_id"
// @Success      200
// @Router       /v1/post/delete_customers_posts/{id} [delete]
func (h *handlerV1) DeletePostByCustomerId(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")
	id, err := strconv.ParseInt(guid, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed parse string to int", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeletePostByCustomerId(
		ctx, &pbp.Id{
			Id: int32(id),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete post by customer id")
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *handlerV1) GetSub(c *gin.Context) (jwt.MapClaims, error) {
	authorization := c.GetHeader("Authorization")
	if c.Request.Header.Get("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, "not authorized")
		h.log.Error("not authorized")
	}
	h.jwthandler.Token = authorization
	claims, err := h.jwthandler.ExtractClaims()
	if err != nil {
		c.JSON(http.StatusUnauthorized, "267 post.go")
		h.log.Error("268 post.go")
	}
	return claims, nil
}
