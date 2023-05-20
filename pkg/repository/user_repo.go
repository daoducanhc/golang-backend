package repository

import (
	"context"

	"std/pkg/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*entity.UserEntity, error)
	// CheckLogin(ctx context.Context, username, password string) (*entity.UserEntity, error)
	Update(user *entity.UserEntity) (*entity.UserEntity, error)
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db.Table("usr"),
	}
}

func (u *userRepository) Update(user *entity.UserEntity) (*entity.UserEntity, error) {
	return user, u.DB.Clauses(clause.Insert{Modifier: "IGNORE"}).Save(&user).Error
}

func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (*entity.UserEntity, error) {
	user := &entity.UserEntity{}
	err := u.DB.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// func (u *userRepository) CheckLogin(ctx context.Context, username, password string) (*entity.UserEntity, error) {
// 	user := &entity.UserEntity{}
// 	err := u.DB.First(user).Where("username = ?", username).Error
// 	if err != nil {
// 		return nil, err
// 		// #! TODO: return log user not found
// 	}
// 	if user.Password != password {
// 		return nil, err
// 		// #! TODO: return log wrong password
// 	}
// 	return user, nil
// }
