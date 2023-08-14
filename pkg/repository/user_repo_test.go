package repository

import (
	"context"
	"database/sql"
	"regexp"
	"std/pkg/entity"
	"testing"

	"gorm.io/driver/mysql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	u    UserRepository
	user *entity.UserEntity
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	assert.NoError(s.T(), err)

	// dsn := "root:123456@tcp(localhost)/users?charset=utf8mb4&parseTime=True&loc=Local"
	// s.DB, err = gorm.Open("mysql", db)
	s.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))
	if err != nil {
		panic(err)
	}
	assert.NoError(s.T(), err)

	s.u = NewUserRepository(s.DB)
}

func (s *Suite) Test_repo_GetUsername() {
	var (
		ctx         context.Context
		username    = "10"
		password    = "10"
		nickname    = ""
		picture_url = ""
	)

	s.user = &entity.UserEntity{Username: username, Nickname: nickname, Password: password, Picture_url: picture_url}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		// `SELECT * FROM "usr" WHERE username = ?`)).
		"SELECT * FROM `usr` WHERE username = ?")).
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows([]string{"username", "nickname", "password", "picture_url"}).
			AddRow(username, nickname, password, picture_url))

	res, err := s.u.GetUserByUsername(ctx, username)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.user, res)
}

func (s *Suite) Test_repo_Update() {
	var (
		ctx         context.Context
		username    = "10"
		password    = "10"
		picture_url = "abc"
	)

	s.user = &entity.UserEntity{Username: username, Nickname: "changed and tested", Password: password, Picture_url: picture_url}

	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(
		// `SELECT * FROM "usr" WHERE username = ?`)).
		// "UPDATE `usr` SET username=?,password=?,nickname=?,picture_url=?")).
		"UPDATE `usr` SET `username`=?,`password`=?,`nickname`=?,`picture_url`=? WHERE username = ?")).
		WithArgs(username, password, "changed and tested", picture_url, username).
		WillReturnResult(sqlmock.NewResult(int64(0), 1))
	s.mock.ExpectCommit()
	res, err := s.u.Update(ctx, s.user)

	assert.NoError(s.T(), err)
	assert.Equal(s.T(), s.user, res)
}

func TestMain(t *testing.T) {
	suite.Run(t, new(Suite))
}

// func add(a, b int) int {
// 	return a + b
// }

// func TestAdd(t *testing.T) {
// 	tests := []struct {
// 		a int
// 		b int
// 		r int
// 	}{
// 		{
// 			a: 1,
// 			b: 2,
// 			r: 3,
// 		},
// 		{
// 			a: 1,
// 			b: 2,
// 			r: 3,
// 		},
// 	}

// 	for i, test := range tests {
// 		t.Run(strconv.Itoa(i), func(t *testing.T) {
// 			result := add(test.a, test.b)
// 			assert.Equal(t, test.r, result)
// 		})
// 	}
// }
