package service

import (
	"context"
	"errors"
	"os"

	"std/pkg/entity"
	"std/pkg/repository"
)

type UserService interface {
	GetUserByUsername(ctx context.Context, username string) (*entity.UserEntity, error)
	CheckLogin(ctx context.Context, username, password string) (bool, error)
	ChangeNickname(ctx context.Context, username, nickname string) (bool, error)
	ChangePicture(ctx context.Context, username, filename string) (bool, error)
}

type userService struct {
	repository *repository.Repository
}

func NewUserService(repo *repository.Repository) UserService {
	return &userService{
		repository: repo,
	}
}

func (s *userService) ChangePicture(ctx context.Context, username, filename string) (bool, error) {
	user := &entity.UserEntity{}
	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	old_url := user.Picture_url
	user.Picture_url = filename

	_, err = s.repository.Update(ctx, user)
	if err != nil {
		return false, err
	}

	if len(old_url) != 0 {
		os.Remove(old_url)
	}
	return true, nil
}

func (s *userService) ChangeNickname(ctx context.Context, username, nickname string) (bool, error) {
	user := &entity.UserEntity{}
	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	user.Nickname = nickname

	_, err = s.repository.Update(ctx, user)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*entity.UserEntity, error) {
	return s.repository.GetUserByUsername(ctx, username)
}

func (s *userService) CheckLogin(ctx context.Context, username, password string) (bool, error) {
	user := &entity.UserEntity{}
	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		return false, err
	}

	if user.Password != password {
		return false, errors.New("incorrect password")
	}

	return true, nil

}
