package postgres

import (
	"database/sql"
	pbp "exam/post_service/genproto/post"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type postRepo struct {
	db *sqlx.DB
}

// NewPostRepo ...
func NewProductRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (p *postRepo) CreatePost(req *pbp.PostRequest) (*pbp.PostWithoutReview, error) {
	newPost := &pbp.PostWithoutReview{}
	err := p.db.QueryRow(`INSERT INTO posts(
		name, 
		description, 
		customer_id) VALUES($1, $2, $3) RETURNING 
	id, 
	name, 
	description`, req.Name, req.Description, req.CustomerId).Scan(
		&newPost.Id,
		&newPost.Name,
		&newPost.Description,
	)
	if err != nil {
		fmt.Println("error while inserting into posts", err)
		return &pbp.PostWithoutReview{}, err
	}
	medias := []*pbp.Media{}
	for _, media := range req.Medias {
		media_ := &pbp.Media{}
		err := p.db.QueryRow(`INSERT INTO medias(
			name, 
			link, 
			type, 
			post_id)
		 VALUES($1, $2, $3, $4) RETURNING 
		id, 
		name, 
		link, 
		type`, media.Name, media.Link, media.Type, newPost.Id).Scan(
			&media_.Id,
			&media_.Name,
			&media_.Link,
			&media_.Type,
		)
		if err != nil {
			fmt.Println("error while inserting into medias", err)
			return &pbp.PostWithoutReview{}, err
		}
		medias = append(medias, media_)
	}
	newPost.Medias = medias
	return newPost, nil
}

func (p *postRepo) GetPostWithCustomerInfo(req *pbp.Id) (*pbp.PostWithCustomerInfo, error) {
	post := &pbp.PostWithCustomerInfo{}
	err := p.db.QueryRow(`SELECT 
	id, 
	name, 
	description, 
	customer_id FROM posts WHERE id = $1`, req.Id).Scan(
		&post.Id,
		&post.Name,
		&post.Description,
		&post.CustomerId,
	)

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("error while getting from posts by id", err)
		return &pbp.PostWithCustomerInfo{}, err
	}

	post_medias, err := p.db.Query(`SELECT 
	id, 
	name, 
	link, 
	type 
	FROM medias WHERE post_id = $1`, post.Id)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("error while getting from medias", err)
		return &pbp.PostWithCustomerInfo{}, err
	}
	medias := []*pbp.Media{}
	for post_medias.Next() {
		media := &pbp.Media{}
		err = post_medias.Scan(
			&media.Id,
			&media.Name,
			&media.Link,
			&media.Type,
		)
		if err != nil {
			fmt.Println("Error while selecting from medias", err)
			return &pbp.PostWithCustomerInfo{}, err
		}
		medias = append(medias, media)
	}
	post.Medias = medias
	return post, nil
}

func (p *postRepo) UpdatePost(req *pbp.PostWithoutReview) (*pbp.PostWithoutReview, error) {
	_, err := p.db.Exec(`UPDATE posts SET name = $1, description = $2, customer_id = $3 WHERE id = $4 and customer_id = $5`, req.Name, req.Description, req.CustomerId, req.Id, req.CustomerId)
	if err != nil {
		fmt.Println("error while updating posts", err)
		return &pbp.PostWithoutReview{}, err
	}

	
	for _, media := range req.Medias {
		_, err := p.db.Exec(`UPDATE medias SET name = $1, link = $2, type = $3 WHERE id = $4 and post_id = $5`, media.Name, media.Link, media.Type, media.Id, req.Id)
		if err != nil {
			fmt.Println("error while updating medias", err)
			return &pbp.PostWithoutReview{}, err
		}
	}
	
	post_medias, err := p.db.Query(`SELECT 
	id, 
	name, 
	link, 
	type 
	FROM medias WHERE post_id = $1`, req.Id)
	if err != nil {
		fmt.Println("error while getting from medias", err)
		return &pbp.PostWithoutReview{}, err
	}
	medias := []*pbp.Media{}
	for post_medias.Next() {
		media := &pbp.Media{}
		err = post_medias.Scan(
			&media.Id,
			&media.Name,
			&media.Link,
			&media.Type,
		)
		if err != nil {
			fmt.Println("Error while selecting from medias", err)
			return &pbp.PostWithoutReview{}, err
		}
		medias = append(medias, media)
	}
	req.Medias = medias
	return req, nil
}
func (p *postRepo) DoesPostExist(req int32) (*pbp.Exist, error) {
	var yes int
	err := p.db.QueryRow(`SELECT 1 FROM posts WHERE id = $1`, req).Scan(&yes)
	if err == sql.ErrNoRows {
		return &pbp.Exist{
			Exists: false,
		}, nil
	}
	if err != nil {
		fmt.Println("error while selecting 1 from posts storage/postgres/post.go", err)
		return &pbp.Exist{}, err
	}

	if yes == 0 {
		return &pbp.Exist{
			Exists: false,
		}, nil
	}

	return &pbp.Exist{
		Exists: true,
	}, nil
}
func (p *postRepo) DeletePost(req int32) (*pbp.IsDeleted, error) {
	row, err := p.db.Exec(`UPDATE posts SET deleted_at = NOW() WHERE id = $1`, req)
	if err != nil {
		fmt.Println("Error while deleting from posts", err)
		return &pbp.IsDeleted{
			PostDeleted: false,
		}, err
	}
	_, err = row.RowsAffected()
	if err != nil {
		return &pbp.IsDeleted{
			PostDeleted: false,
		}, err
	}
	return &pbp.IsDeleted{
		PostDeleted: true,
	}, nil
}

func (p *postRepo) GetAllPostsWithCustomer(*pbp.Empty) (*pbp.AllPosts, error) {
	posts, err := p.db.Query(`SELECT id, name, description, customer_id FROM posts WHERE deleted_at is NULL`)
	if err != nil {
		fmt.Println("error while selecting all posts", err)
		return &pbp.AllPosts{}, err
	}
	AllPosts := &pbp.AllPosts{}
	for posts.Next() {
		post := &pbp.PostWithOnlyCustomerInfo{}
		err = posts.Scan(&post.Id, &post.Name, &post.Description, &post.CustomerId)
		if err != nil {
			fmt.Println("Error while scanning to PostWithOnlyCustomerInfo", err)
			return &pbp.AllPosts{}, err
		}
		post_medias, err := p.db.Query(`SELECT 
			id, 
			name, 
			link, 
			type 
			FROM medias WHERE post_id = $1`, post.Id)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println("error while getting from medias", err)
			return &pbp.AllPosts{}, err
		}
		medias := []*pbp.Media{}
		for post_medias.Next() {
			media := &pbp.Media{}
			err = post_medias.Scan(
				&media.Id,
				&media.Name,
				&media.Link,
				&media.Type,
			)
			if err != nil {
				fmt.Println("Error while selecting from medias", err)
				return &pbp.AllPosts{}, err
			}
			medias = append(medias, media)
		}
		post.Medias = medias
		AllPosts.Posts = append(AllPosts.Posts, post)
	}
	return AllPosts, nil
}

func (p *postRepo) GetPostsOfCustomer(req *pbp.Id) (*pbp.Posts, error) {
	PostsOfCustomer := &pbp.Posts{}
	posts, err := p.db.Query(`SELECT id, name, description FROM posts WHERE customer_id = $1`, req.Id)
	if err != nil {
		fmt.Println("error while selecting from posts with customer_id", err)
		return &pbp.Posts{}, err
	}
	for posts.Next() {
		post := &pbp.Post{}
		err = posts.Scan(
			&post.Id,
			&post.Name,
			&post.Description,
		)
		if err != nil {
			fmt.Println("Error while scanning to post GetPostsOfCustomer", err)
			return &pbp.Posts{}, err
		}
		PostsOfCustomer.Posts = append(PostsOfCustomer.Posts, post)
	}
	for _, post := range PostsOfCustomer.Posts {
		medias, err := p.db.Query(`SELECT id, name, link, type FROM medias WHERE post_id = $1`, post.Id)
		if err != nil {
			fmt.Println("error while selecting from medias by post id", err)
			return &pbp.Posts{}, err
		}

		for medias.Next() {
			media := &pbp.Media{}
			medias.Scan(
				&media.Id,
				&media.Name,
				&media.Link,
				&media.Type,
			)
			post.Medias = append(post.Medias, media)
		}
	}
	return PostsOfCustomer, nil
}

func (p *postRepo) DeletePostByCustomerId(req int32) (*pbp.IsDeleted, []int32, error) {
	delets, err := p.db.Query(`SELECT id FROM posts WHERE customer_id = $1`, req)
	if err != nil {
		fmt.Println("Error while selecting from posts by customer_id for deleting", err)
		return &pbp.IsDeleted{}, nil, err
	}
	ids := []int32{}
	for delets.Next() {
		var id int32
		err = delets.Scan(&id)
		if err != nil {
			fmt.Println("error while scanning to id ", err)
			return &pbp.IsDeleted{}, nil, err
		}
		ids = append(ids, id)
	}

	row, err := p.db.Exec(`UPDATE posts SET deleted_at = NOW() WHERE customer_id = $1`, req)
	if err != nil {
		fmt.Println("Error while deleting from posts by customer_id", err)
		return &pbp.IsDeleted{
			PostDeleted: false,
		}, nil, err
	}
	_, err = row.RowsAffected()
	if err != nil {
		return &pbp.IsDeleted{
			PostDeleted: false,
		}, nil, err
	}
	return &pbp.IsDeleted{
		PostDeleted: true,
	}, ids, nil
}
