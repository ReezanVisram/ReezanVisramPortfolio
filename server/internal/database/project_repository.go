package database

import (
	"context"
	domain "reezanvisramportfolio/domain/project"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepository interface {
	GetProjectById(ctx context.Context, id int64) (*domain.Project, error)
	InsertProject(ctx context.Context, project domain.Project) error
	RemoveProjectById(ctx context.Context, id int64) error
	GetAllProjects(ctx context.Context) ([]domain.Project, error)
}

type projectRepository struct {
	projectCollection *mongo.Collection
}

func NewProjectRepository(projectCollection *mongo.Collection) ProjectRepository {
	return &projectRepository{
		projectCollection: projectCollection,
	}
}

func (pr *projectRepository) GetProjectById(ctx context.Context, id int64) (*domain.Project, error) {
	filter := bson.D{{Key: "id", Value: id}}

	var result domain.Project
	err := pr.projectCollection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (pr *projectRepository) InsertProject(ctx context.Context, project domain.Project) error {
	_, err := pr.projectCollection.InsertOne(ctx, project)

	return err
}

func (pr *projectRepository) RemoveProjectById(ctx context.Context, id int64) error {
	filter := bson.D{{Key: "id", Value: id}}

	_, err := pr.projectCollection.DeleteOne(ctx, filter)

	return err
}

func (pr *projectRepository) GetAllProjects(ctx context.Context) ([]domain.Project, error) {
	filter := bson.D{{}}
	cursor, err := pr.projectCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []domain.Project
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
