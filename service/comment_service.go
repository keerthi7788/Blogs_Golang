package service

import (
	"Blogs/modals"
	"Blogs/repositories"
	"context"
	"fmt"
)

type CommentService interface {
	CreateComment(ctx context.Context, CommentModal modals.Comment) (modals.Comment, error)
	GetCommentByID(ctx context.Context, id string) (modals.Comment, error)
	GetAllComments(ctx context.Context) ([]modals.Comment, error)
	DeleteAllComments(ctx context.Context) ([]modals.Comment, error)
	DeleteCommentById(ctx context.Context, id string) (modals.Comment, error)
	UpdateCommentById(ctx context.Context, id, Content string) (modals.Comment, error)
	UpdateAllComments(ctx context.Context, id string, UpdateAllComments modals.Comment) (modals.Comment, error)
}

type CommentSer struct {
	repo repositories.CommentRep
}

func NewCommentService(repo repositories.CommentRep) *CommentSer {

	return &CommentSer{repo: repo}
}
func (s *CommentSer) CreateComment(ctx context.Context, CommentModal modals.Comment) (modals.Comment, error) {
	results, err := s.repo.CreateComment(ctx, CommentModal)
	if err != nil {
		fmt.Println("eror while creating user in service", err)

	}
	return results, nil
}

func (s *CommentSer) GetAllComments(ctx context.Context) ([]modals.Comment, error) {
	result, err := s.repo.GetAllComments(ctx)
	if err != nil {
		return []modals.Comment{}, fmt.Errorf("uanble to get the usrs in service layer %d", err)
	}
	return result, nil
}

func (s *CommentSer) GetCommentByID(ctx context.Context, id string) (modals.Comment, error) {

	users, err := s.repo.GetCommentByID(ctx, id)
	if err != nil {
		return modals.Comment{}, fmt.Errorf("unable to find the user %d", err)

	}
	return users, nil
}
func (s *CommentSer) UpdateCommentById(ctx context.Context, id, Content string) (modals.Comment, error) {

	updateduser, err := s.repo.UpdateCommentById(ctx, id, Content)
	if err != nil {
		return modals.Comment{}, err
	}
	return updateduser, nil
}
func (s *CommentSer) UpdateAllComments(ctx context.Context, id string, UpdateAllComments modals.Comment) (modals.Comment, error) {
	updateduserdetails, err := s.repo.UpdateAllComments(ctx, id, UpdateAllComments)
	if err != nil {
		return modals.Comment{}, fmt.Errorf("user details updated %d", err)

	}
	return updateduserdetails, nil

}
func (s *CommentSer) DeleteCommentById(ctx context.Context, id string) (modals.Comment, error) {
	results, err := s.repo.DeleteCommentById(ctx, id)

	if err != nil {
		return modals.Comment{}, fmt.Errorf("unable to delete the user %d", err)
	}
	return results, nil

}
func (s *CommentSer) DeleteAllComments(ctx context.Context) ([]modals.Comment, error) {
	deletedcomment, err := s.repo.DeleteAllComments(ctx)
	if err != nil {
		return []modals.Comment{}, fmt.Errorf("unable to deleteAll the user %d", err)
	}
	return deletedcomment, nil
}
