package service

import (
	"Blogs/modals"
	"Blogs/repositories"
	"context"
	"fmt"
	"time"
)

type PostService interface {
	CreatePost(ctx context.Context, postModal modals.Post) (modals.Post, error)
	GetPostByID(ctx context.Context, id string) (modals.Post, error)
	GetAllPosts(ctx context.Context) ([]modals.Post, error)
	DeleteAllPosts(ctx context.Context) ([]modals.Post, error)
	DeletePostById(ctx context.Context, id string) (modals.Post, error)
	UpdatePostById(ctx context.Context, id, PostTitle, Content string) (modals.Post, error)
	UpdateAllPosts(ctx context.Context, id string, updateAllPosts modals.Post) (modals.Post, error)
}
type PostServ struct {
	repo repositories.Postrepositor
}

// GetAllPosts implements PostService.

/*Creating Db and collection to stire data*/

func NewpostService(repo repositories.Postrepositor) *PostServ {

	return &PostServ{repo: repo}

}
func (s *PostServ) CreatePost(ctx context.Context, postModal modals.Post) (modals.Post, error) {
	created_at := time.Now().UTC()
	postModal.CreatedAt = &created_at
	results, err := s.repo.CreatePost(ctx, postModal)
	if err != nil {
		fmt.Println("eror while creating user in service", err)

	}
	return results, nil
}

func (s *PostServ) GetAllPosts(ctx context.Context) ([]modals.Post, error) {
	result, err := s.repo.GetAllPosts(ctx)
	if err != nil {
		return []modals.Post{}, fmt.Errorf("uanble to get the usrs in service layer %d", err)
	}
	return result, nil
}

func (s *PostServ) GetPostByID(ctx context.Context, id string) (modals.Post, error) {

	users, err := s.repo.GetPostByID(ctx, id)
	if err != nil {
		return modals.Post{}, fmt.Errorf("unable to find the user %d", err)

	}
	return users, nil
}
func (s *PostServ) UpdatePostById(ctx context.Context, id, PostTitle, Content string) (modals.Post, error) {

	updateduser, err := s.repo.UpdatePostById(ctx, id, PostTitle, Content)
	if err != nil {
		return modals.Post{}, err
	}
	return updateduser, nil
}
func (s *PostServ) UpdateAllPosts(ctx context.Context, id string, updateAllPosts modals.Post) (modals.Post, error) {
	updateduserdetails, err := s.repo.UpdateAllPosts(ctx, id, updateAllPosts)
	if err != nil {
		return modals.Post{}, fmt.Errorf("user details updated %d", err)

	}
	return updateduserdetails, nil

}
func (s *PostServ) DeletePostById(ctx context.Context, id string) (modals.Post, error) {
	results, err := s.repo.DeletePostById(ctx, id)

	if err != nil {
		return modals.Post{}, fmt.Errorf("unable to delete the user %d", err)
	}
	return results, nil

}
func (s *PostServ) DeleteAllPosts(ctx context.Context) ([]modals.Post, error) {
	deletedUser, err := s.repo.DeleteAllPosts(ctx)
	if err != nil {
		return []modals.Post{}, fmt.Errorf("unable to deleteAll the user %d", err)
	}
	return deletedUser, nil
}
