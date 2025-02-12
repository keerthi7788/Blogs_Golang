package repositories

import (
	"Blogs/modals"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user modals.Users) (primitive.ObjectID, error)
	GetAllUsers(ctx context.Context) ([]modals.Users, error)
	GetUserByID(ctx context.Context, id string) (modals.Users, error)
	UpdateUserByID(ctx context.Context, id, Name, Email string) (modals.Users, error)
	UpdateUserDetails(ctx context.Context, id string, UpdatedDetails modals.Users) (modals.Users, error)
	DeleteUserByID(ctx context.Context, id string) (modals.Users, error)
	DeleteAllUsers(ctx context.Context) ([]modals.Users, error)
	GetUserByEmail(ctx context.Context, email string) (modals.Users, error)
}

type UserRepo struct {
	Client         *mongo.Client
	DatabaseName   string
	CollectionName string
}

// GetAllUsers implements UserRepository.

/*Creating Db and collection to stire data*/

func NewUserRepository(client *mongo.Client, databaseName string, collectionName string) *UserRepo {
	return &UserRepo{
		Client:         client,
		DatabaseName:   databaseName,
		CollectionName: collectionName,
	}

}
func (r *UserRepo) CreateUser(ctx context.Context, user modals.Users) (primitive.ObjectID, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("unable to create the user %w", err)
	}
	return result.InsertedID.(primitive.ObjectID), nil

}
func (r *UserRepo) GetAllUsers(ctx context.Context) ([]modals.Users, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
	var users []modals.Users
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return []modals.Users{}, err

	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user modals.Users
		if err := cursor.Decode(&user); err != nil {
			fmt.Printf("UNABLE TO DECODE: %v\n", err)
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepo) GetUserByID(ctx context.Context, id string) (modals.Users, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Users{}, fmt.Errorf("invalid ID format: %w", err)
	}
	var user modals.Users
	err = collection.FindOne(ctx, bson.M{"_id": ObjectID}).Decode(&user)
	if err != nil {
		return modals.Users{}, fmt.Errorf("unable to find the user %d", err)

	}
	return user, nil
}
func (r *UserRepo) UpdateUserByID(ctx context.Context, id, UpdatedName, UpdatedEmail string) (modals.Users, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Users{}, fmt.Errorf("invalid ID format: %w", err)
	}
	filter := bson.M{"_id": ObjectID}
	update := bson.M{"$set": bson.M{"Name": UpdatedName, "Email": UpdatedEmail}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return modals.Users{}, nil
	}
	var updateduser modals.Users
	err = collection.FindOne(ctx, filter).Decode(&updateduser)
	if err != nil {
		return modals.Users{}, err
	}
	return updateduser, nil
}
func (r *UserRepo) UpdateUserDetails(ctx context.Context, id string, UpdatedDetails modals.Users) (modals.Users, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Users{}, fmt.Errorf("invalid ID format: %w", err)
	}
	filter := bson.M{"_id": ObjectID}

	update := bson.M{"$set": UpdatedDetails}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return modals.Users{}, fmt.Errorf("user details updated %d", err)
	}
	var updateduserDetails modals.Users
	err = collection.FindOne(ctx, filter).Decode(&updateduserDetails)
	if err != nil {
		return modals.Users{}, fmt.Errorf("unable to display the updated user document %d", err)
	}
	return updateduserDetails, nil

}
func (r *UserRepo) DeleteUserByID(ctx context.Context, id string) (modals.Users, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
	var deletedUser modals.Users
	ObjectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return modals.Users{}, fmt.Errorf("invalid ID format: %w", err)
	}
	err = collection.FindOne(ctx, bson.M{"_id": ObjectID}).Decode(&deletedUser)
	if err != nil {
		return modals.Users{}, err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		return modals.Users{}, fmt.Errorf("unable to delete the user %d", err)
	}

	return deletedUser, nil

}
func (r *UserRepo) DeleteAllUsers(ctx context.Context) ([]modals.Users, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
	var deletedUser []modals.Users
	// err := r.collection.Find(ctx, bson.D{}).Decode(&deletedUser)
	_, err := collection.DeleteMany(ctx, bson.D{})
	if err != nil {
		return []modals.Users{}, fmt.Errorf("unable to deleteAll the user %d", err)
	}
	return deletedUser, nil
}
func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (modals.Users, error) {
	collection := r.Client.Database(r.DatabaseName).Collection(r.CollectionName)
	var user modals.Users
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return modals.Users{}, err
	}
	return user, nil
}
