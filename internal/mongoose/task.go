package mongoose

import (
	"context"

	task "solid-task-management/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	TaskRepository interface {
		Create(ctx context.Context, payload task.CreateReq) (*task.Task, error)
		FindOne(ctx context.Context, id string) (*task.Task, error)
		FindAll(ctx context.Context, payload task.FetchTasksReq) ([]*task.Task, error)
		Update(ctx context.Context, payload task.UpdateReq) (*task.Task, error)
	}

	TaskRepositoryImpl struct {
		tasksCollection *mongo.Collection
	}
)

func TaskRepositoryProvider(db *mongo.Database) TaskRepository {
	return TaskRepositoryImpl{
		tasksCollection: db.Collection("tasks"),
	}
}

func (r TaskRepositoryImpl) FindOne(ctx context.Context, id string) (*task.Task, error) {
	print("this is id", id)

	taskId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var result *task.Task
	err = r.tasksCollection.FindOne(context.TODO(), primitive.M{"_id": taskId}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r TaskRepositoryImpl) Create(ctx context.Context, spec task.CreateReq) (*task.Task, error) {
	res, err := r.tasksCollection.InsertOne(context.TODO(), spec)
	if err != nil {
		return nil, err
	}

	insertedID := res.InsertedID.(primitive.ObjectID).Hex()
	newTask, err := r.FindOne(ctx, insertedID)

	if err != nil {
		return nil, err
	}

	return newTask, nil
}

func (r TaskRepositoryImpl) FindAll(ctx context.Context, payload task.FetchTasksReq) ([]*task.Task, error) {
	query := bson.M{}
	if payload.Description != "" {
		query["description"] = payload.Description
	}
	if payload.Title != "" {
		query["title"] = payload.Title
	}

	skip := (payload.PerPage - 1) * payload.Page

	qOption := options.Find()
	qOption.SetLimit(payload.PerPage)
	qOption.SetSkip(skip)

	cursor, err := r.tasksCollection.Find(ctx, query, qOption)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*task.Task
	for cursor.Next(ctx) {
		var t task.Task
		if err := cursor.Decode(&t); err != nil {
			return nil, err
		}

		tasks = append(tasks, &t)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r TaskRepositoryImpl) Update(ctx context.Context, payload task.UpdateReq) (*task.Task, error) {
	// Convert the ID to a MongoDB ObjectID
	taskId, err := primitive.ObjectIDFromHex(payload.ID)
	if err != nil {
		return nil, err
	}

	// Create the update document
	update := bson.M{
		"$set": bson.M{
			"title":       payload.Title,
			"description": payload.Description,
			"status":      payload.Status,
		},
	}

	// Perform the update operation
	_, err = r.tasksCollection.UpdateOne(ctx, bson.M{"_id": taskId}, update)
	if err != nil {
		return nil, err
	}

	// Fetch and return the updated task
	updatedTask, err := r.FindOne(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}
