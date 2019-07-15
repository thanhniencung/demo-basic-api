package repo_impl

import (
	"basic-api/repository"
	"basic-api/db"
	"basic-api/model"
	"fmt"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepository {
	return &UserRepoImpl{
		sql: sql,
	}
}

func (u *UserRepoImpl) Insert(user model.User) error {
	fmt.Println(user)
	// logic insert to database
	insertCommand := `INSERT INTO users(name, email) VALUES (:name, :email)` // RAW QUERY

	_, err := u.sql.Db.NamedExec(insertCommand, user)
	return err

}



