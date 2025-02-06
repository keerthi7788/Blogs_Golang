package repositories

import (
	"Blogs/modals"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Commentepositor interface {
	CreateComment(ctx context.Context, CommentModal modals.Comment) (modals.Comment, error)
	GetCommentByID(ctx context.Context, id string) (modals.Comment, error)
	GetAllComments(ctx context.Context) ([]modals.Comment, error)
	DeleteAllComments(ctx context.Context) ([]modals.Comment, error)
	DeleteCommentById(ctx context.Context, id string) (modals.Comment, error)
	UpdateCommentById(ctx context.Context, id, Content string) (modals.Comment, error)
	UpdateAllComments(ctx context.Context, id string, UpdateAllComments modals.Comment) (modals.Comment, error)
}
type CommentRep struct {
	Client         *mongo.Client
	DatabaseName   string
	CollectionName string
}

func NewCommentRepository(client *mongo.Client, databaseName, collectionName string) *CommentRep {

	return &CommentRep{
		Client:         client,
		DatabaseName:   databaseName,
		CollectionName: collectionName,
	}
}
func (r *CommentRep) CreateComment(ctx context.Context, CommentModal modals.Comment) (modals.Comment, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	result, err := collection.InsertOne(ctx, CommentModal)
	if err != nil {
		return modals.Comment{}, fmt.Errorf("unable to create post %w", err)
	}
	CommentModal.ID = result.InsertedID.(primitive.ObjectID)
	return CommentModal, err
}
func (r *CommentRep) GetCommentByID(ctx context.Context, id string) (modals.Comment, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Comment{}, fmt.Errorf("unable to decode id %w", err)
	}
	var Comments modals.Comment
	err = collection.FindOne(ctx, bson.M{"_id": objectid}).Decode(&Comments)
	if err != nil {
		return modals.Comment{}, fmt.Errorf("unable to find te postby id%w", err)
	}
	return Comments, nil
}
func (r *CommentRep) GetAllComments(ctx context.Context) ([]modals.Comment, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	var Comments []modals.Comment
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []modals.Comment{}, err
	}
	cursor.Close(ctx)
	for cursor.Next(ctx) {
		var comment modals.Comment
		err = cursor.Decode(&comment)
		if err != nil {
			return []modals.Comment{}, err
		}
		Comments = append(Comments, comment)
	}
	return Comments, err
}

func (r *CommentRep) DeleteAllComments(ctx context.Context) ([]modals.Comment, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	var deletedpos []modals.Comment
	_, err := collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return []modals.Comment{}, fmt.Errorf("unable to deleteAll the user %d", err)
	}
	return deletedpos, nil
}
func (r *CommentRep) DeleteCommentById(ctx context.Context, id string) (modals.Comment, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Comment{}, fmt.Errorf("unable to decode id %w", err)
	}
	var deletedcomments modals.Comment

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectid})
	if err != nil {
		return modals.Comment{}, fmt.Errorf("unable to delete the post %w", err)

	}
	return deletedcomments, nil
}
func (r *CommentRep) UpdateCommentById(ctx context.Context, id, Content string) (modals.Comment, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Comment{}, fmt.Errorf("unable to decode id %w", err)
	}

	filter := bson.M{"_id": objectid}
	updated := bson.M{
		"$set": bson.M{"Content": Content}}
	_, err = collection.UpdateOne(ctx, filter, updated)
	if err != nil {
		return modals.Comment{}, err
	}
	var updatedcommentDetails modals.Comment
	_, err = collection.Find(ctx, filter)
	if err != nil {
		return modals.Comment{}, fmt.Errorf("unable to find the updated posts %w", err)
	}
	return updatedcommentDetails, nil
}

func (r *CommentRep) UpdateAllComments(ctx context.Context, id string, UpdateAllComments modals.Comment) (modals.Comment, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Comment{}, fmt.Errorf("unable to decode id %w", err)
	}
	filter := bson.M{"_id": objectid}
	updated := bson.M{"$set": UpdateAllComments}

	_, err = collection.UpdateMany(ctx, filter, updated)
	if err != nil {
		return modals.Comment{}, fmt.Errorf("uanable to update posts %w", err)
	}
	var comments modals.Comment
	_, err = collection.Find(ctx, filter)
	if err != nil {
		return modals.Comment{}, err
	}
	return comments, nil
}
