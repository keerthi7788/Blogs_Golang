package repositories

import (
	"Blogs/modals"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Postrepositor interface {
	CreatePost(ctx context.Context, postModal modals.Post) (modals.Post, error)
	GetPostByID(ctx context.Context, id string) (modals.Post, error)
	GetAllPosts(ctx context.Context) ([]modals.Post, error)
	DeleteAllPosts(ctx context.Context) ([]modals.Post, error)
	DeletePostById(ctx context.Context, id string) (modals.Post, error)
	UpdatePostById(ctx context.Context, id, PostTitle, Content string) (modals.Post, error)
	UpdateAllPosts(ctx context.Context, id string, updateAllPosts modals.Post) (modals.Post, error)
}
type PostRep struct {
	Client         *mongo.Client
	DatabaseName   string
	CollectionName string
}

func NewPostRepository(client *mongo.Client, databaseName string, collectionName string) *PostRep {

	return &PostRep{
		Client:         client,
		DatabaseName:   databaseName,
		CollectionName: collectionName,
	}
}
func (r *PostRep) CreatePost(ctx context.Context, postModal modals.Post) (modals.Post, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	result, err := collection.InsertOne(ctx, postModal)
	if err != nil {
		return modals.Post{}, fmt.Errorf("unable to create post %w", err)
	}
	postModal.ID = result.InsertedID.(primitive.ObjectID)
	return postModal, err
}
func (r *PostRep) GetPostByID(ctx context.Context, id string) (modals.Post, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Post{}, fmt.Errorf("unable to decode id %w", err)
	}
	var posts modals.Post
	err = collection.FindOne(ctx, bson.M{"_id": objectid}).Decode(&posts)
	if err != nil {
		return modals.Post{}, fmt.Errorf("unable to find te postby id%w", err)
	}
	return posts, nil
}
func (r *PostRep) GetAllPosts(ctx context.Context) ([]modals.Post, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
	var posts []modals.Post
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []modals.Post{}, err
	}
	cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post modals.Post
		err = cursor.Decode(&post)
		if err != nil {
			fmt.Printf("UNABLE TO DECODE: %v\n", err)
			continue
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRep) DeleteAllPosts(ctx context.Context) ([]modals.Post, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	var deletedpos []modals.Post
	_, err := collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return []modals.Post{}, fmt.Errorf("unable to deleteAll the user %d", err)
	}
	return deletedpos, nil
}
func (r *PostRep) DeletePostById(ctx context.Context, id string) (modals.Post, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
	var deletedPosts modals.Post
	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Post{}, fmt.Errorf("unable to decode id %w", err)
	}
	err = collection.FindOne(ctx, bson.M{"_id": objectid}).Decode(&deletedPosts)
	if err != nil {
		return modals.Post{}, err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectid})
	if err != nil {
		return modals.Post{}, fmt.Errorf("unable to delete the post %w", err)

	}
	return deletedPosts, nil
}
func (r *PostRep) UpdatePostById(ctx context.Context, id, PostTitle, Content string) (modals.Post, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Post{}, fmt.Errorf("unable to decode id %w", err)
	}

	filter := bson.M{"_id": objectid}
	updated := bson.M{
		"$set": bson.M{"Content": Content, "Title": PostTitle}}
	_, err = collection.UpdateOne(ctx, filter, updated)
	if err != nil {
		return modals.Post{}, err
	}
	var updatedPostDetails modals.Post
	_, err = collection.Find(ctx, filter)
	if err != nil {
		return modals.Post{}, fmt.Errorf("unable to find the updated posts %w", err)
	}
	return updatedPostDetails, nil
}

func (r *PostRep) UpdateAllPosts(ctx context.Context, id string, updateAllPosts modals.Post) (modals.Post, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)

	objectid, err := primitive.ObjectIDFromHex("id")
	if err != nil {
		return modals.Post{}, fmt.Errorf("unable to decode id %w", err)
	}
	filter := bson.M{"_id": objectid}
	updated := bson.M{"$set": updateAllPosts}

	_, err = collection.UpdateMany(ctx, filter, updated)
	if err != nil {
		return modals.Post{}, fmt.Errorf("uanable to update posts %w", err)
	}
	var posts modals.Post
	_, err = collection.Find(ctx, filter)
	if err != nil {
		return modals.Post{}, err
	}
	return posts, nil
}
