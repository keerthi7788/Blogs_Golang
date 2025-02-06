package service

import (
	"Blogs/modals"
	"Blogs/repositories"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	CreateUser(ctx context.Context, user modals.Users) (primitive.ObjectID, error)
	GetAllUsers(ctx context.Context) ([]modals.Users, error)
	GetUserByID(ctx context.Context, id string) (modals.Users, error)
	UpdateUserByID(ctx context.Context, id, Name, Email string) (modals.Users, error)
	UpdateUserDetails(ctx context.Context, id string, UpdatedDetails modals.Users) (modals.Users, error)
	DeleteUserByID(ctx context.Context, id string) (modals.Users, error)
	DeleteAllUsers(ctx context.Context) ([]modals.Users, error)
}
type UserServ struct {
	repo repositories.UserRepository
}

/*Creating Db and collection to stire data*/

func NewUserService(repo repositories.UserRepository) *UserServ {
	return &UserServ{repo: repo}

}
func (s *UserServ) CreateUser(ctx context.Context, user modals.Users) (primitive.ObjectID, error) {
	results, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		fmt.Println("eror while creating user in service", err)

	}
	return results, nil
}

func (s *UserServ) GetAllUsers(ctx context.Context) ([]modals.Users, error) {
	result, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return []modals.Users{}, fmt.Errorf("uanble to get the usrs in service layer %d", err)
	}
	return result, nil
}

func (s *UserServ) GetUserByID(ctx context.Context, id string) (modals.Users, error) {

	users, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return modals.Users{}, fmt.Errorf("unable to find the user %d", err)

	}
	return users, nil
}
func (s *UserServ) UpdateUserByID(ctx context.Context, id, UpdatedName, UpdatedEmail string) (modals.Users, error) {

	updateduser, err := s.repo.UpdateUserByID(ctx, id, UpdatedName, UpdatedEmail)
	if err != nil {
		return modals.Users{}, err
	}
	return updateduser, nil
}
func (s *UserServ) UpdateUserDetails(ctx context.Context, id string, UpdatedDetails modals.Users) (modals.Users, error) {
	updateduserdetails, err := s.repo.UpdateUserDetails(ctx, id, UpdatedDetails)
	if err != nil {
		return modals.Users{}, fmt.Errorf("user details updated %d", err)

	}
	return updateduserdetails, nil

}
func (s *UserServ) DeleteUserByID(ctx context.Context, id string) (modals.Users, error) {
	results, err := s.repo.DeleteUserByID(ctx, id)

	if err != nil {
		return modals.Users{}, fmt.Errorf("unable to delete the user %d", err)
	}
	return results, nil

}
func (s *UserServ) DeleteAllUsers(ctx context.Context) ([]modals.Users, error) {

	deletedUser, err := s.repo.DeleteAllUsers(ctx)
	if err != nil {
		return []modals.Users{}, fmt.Errorf("unable to deleteAll the user %d", err)
	}
	return deletedUser, nil
}
