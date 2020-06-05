package repositories

import (
	"blog-for-go/datamodels"
	"database/sql"
	"log"
)

type PostRepo interface {
	Query(index int, size int) []datamodels.Post
	Insert(post datamodels.Post) bool
	Update(post datamodels.Post) bool
}

func NewPostRepo(source *sql.DB) PostRepo {
	return &postRepo{source: source}
}

type postRepo struct {
	source *sql.DB
}

func (r *postRepo) Query(index int, size int) []datamodels.Post {
	// 通过切片存储
	posts := make([]datamodels.Post, 0)

	idx := 0
	if index > 0 {
		idx = (index - 1) * size
	}

	rows, _ := r.source.Query("SELECT * FROM `post` limit ?,?", idx, size)

	// 遍历
	var post datamodels.Post
	for rows.Next() {
		rows.Scan(&post.Id, &post.Title, &post.PostUser, &post.Description, &post.Content,
			&post.CreateTime, &post.UpdateTime, &post.CreateUser, &post.UpdateUser)
		posts = append(posts, post)
	}

	return posts
}

func (r *postRepo) Insert(post datamodels.Post) bool {
	ret, _ := r.source.Exec("INSERT INTO `post`(title, description, post_user, content, create_user, update_user) values(?,?,?,?,?,?)",
		&post.Title, &post.Description, &post.PostUser, &post.Content, &post.CreateUser, &post.UpdateUser)

	//插入数据的主键id
	lastInsertID, _ := ret.LastInsertId()
	log.Println("LastInsertID:", lastInsertID)

	//影响行数
	affectedRows, _ := ret.RowsAffected()
	log.Println("RowsAffected:", affectedRows)

	return affectedRows > 0
}

func (r *postRepo) Update(post datamodels.Post) bool {
	ret, _ := r.source.Exec("UPDATE `post` set title=?, description=?, post_user=?, content=?, create_user=?, update_user=?"+
		" where id=?", &post.Title, &post.Description, &post.PostUser, &post.Content, &post.CreateUser, &post.UpdateUser, &post.Id)
	affectedRows, _ := ret.RowsAffected()

	log.Println("RowsAffected:", affectedRows)

	return affectedRows > 0
}
