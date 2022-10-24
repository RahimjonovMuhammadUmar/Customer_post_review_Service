package repo

import (
	pbp "exam/post_service/genproto/post"
)

// PostStorage
type PostStorage interface {
	CreatePost(*pbp.PostRequest) (*pbp.Post, error)
	GetPostWithCustomerInfo(*pbp.Id) (*pbp.PostWithCustomerInfo, error)
	UpdatePost(*pbp.Post) (*pbp.Post, error)
	DoesPostExist(id int32) (*pbp.Exist, error)
	DeletePost(id int32) (*pbp.IsDeleted, error)
	DeletePostByCustomerId(id int32) (*pbp.IsDeleted, []int32, error)
	GetAllPostsWithCustomer(*pbp.Empty) (*pbp.AllPosts, error)
	GetPostsOfCustomer(*pbp.Id) (*pbp.Posts, error)
}
