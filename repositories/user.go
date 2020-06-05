package repositories

import (
	"blog-for-go/datamodels"
	"database/sql"
	"log"
)

type UserRepo interface {
	Query(index int, size int) []datamodels.User
	QueryBy(name string, password string) []datamodels.User
	Insert(post datamodels.User) bool
	Update(post datamodels.User) bool
}

func NewUserRepo(source *sql.DB) UserRepo {
	return &userRepo{source: source}
}

type userRepo struct {
	source *sql.DB
}

func (r *userRepo) Query(index int, size int) []datamodels.User {
	// 通过切片存储
	users := make([]datamodels.User, 0)

	idx := 0
	if index > 0 {
		idx = (index - 1) * size
	}

	rows, _ := r.source.Query("SELECT * FROM `user` limit ?,?", idx, size)

	// 遍历
	var user datamodels.User
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Phone,
			&user.CreateTime, &user.UpdateTime, &user.CreateUser, &user.UpdateUser)
		users = append(users, user)
	}

	return users
}

func (r *userRepo) QueryBy(name string, password string) []datamodels.User {
	// 通过切片存储
	users := make([]datamodels.User, 0)

	rows, _ := r.source.Query("SELECT * FROM `user` WHERE name = ? and password = ?", name, password)

	// 遍历
	var user datamodels.User
	for rows.Next() {
		rows.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Phone,
			&user.CreateTime, &user.UpdateTime, &user.CreateUser, &user.UpdateUser)
		users = append(users, user)
	}

	return users
}

func (r *userRepo) Insert(user datamodels.User) bool {
	ret, _ := r.source.Exec("INSERT INTO `user`(name, password, email, phone, create_user, update_user) values(?,?,?,?,?,?)",
		&user.Name, &user.Password, &user.Email, &user.Phone, &user.CreateUser, &user.UpdateUser)

	//插入数据的主键id
	lastInsertID, _ := ret.LastInsertId()
	log.Println("LastInsertID:", lastInsertID)

	//影响行数
	affectedRows, _ := ret.RowsAffected()
	log.Println("RowsAffected:", affectedRows)

	return affectedRows > 0
}

func (r *userRepo) Update(user datamodels.User) bool {
	ret, _ := r.source.Exec("UPDATE `user` set name=?, password=?, email=?, phone=?, create_user=?, update_user=?"+
		" where id=?", &user.Name, &user.Password, &user.Email, &user.Phone, &user.CreateUser, &user.UpdateUser, &user.Id)
	affectedRows, _ := ret.RowsAffected()

	log.Println("RowsAffected:", affectedRows)

	return affectedRows > 0
}
