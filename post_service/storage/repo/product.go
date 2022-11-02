package repo

import (
	pbp "exam/post_service/genproto/post"
)

// PostStorage
type PostStorage interface {
	CreatePost(*pbp.PostRequest) (*pbp.PostWithoutReview, error)
	GetPostWithCustomerInfo(post_id int32) (*pbp.PostWithCustomerInfo, error)
	UpdatePost(*pbp.PostWithoutReview) (*pbp.PostWithoutReview, error)
	DoesPostExist(id int32) (*pbp.Exist, error)
	DeletePost(id int32) (*pbp.IsDeleted, error)
	DeletePostByCustomerId(id int32) (*pbp.IsDeleted, []int32, error)
	GetAllPostsWithCustomer(*pbp.Empty) (*pbp.AllPosts, error)
	GetPostsOfCustomer(*pbp.Id) (*pbp.Posts, error)
	GetPostsByPage(page, limit int32) (*pbp.PostsByPage, error)
	GetPostInfoOnly(id int32) (*pbp.PostInfoOnly, error)
}
