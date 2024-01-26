package database

import (
	"context"
	domain "reezanvisramportfolio/domain/project"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepository interface {
	InsertProject(ctx context.Context, project domain.Project) error
	RemoveProjectByName(ctx context.Context, name string) error
}

type projectRepository struct {
	projectCollection *mongo.Collection
}

func NewProjectRepository(projectCollection *mongo.Collection) ProjectRepository {
	return &projectRepository{
		projectCollection: projectCollection,
	}
}

func (pr *projectRepository) InsertProject(ctx context.Context, project domain.Project) error {
	_, err := pr.projectCollection.InsertOne(ctx, project)

	return err
}

func (pr *projectRepository) RemoveProjectByName(ctx context.Context, name string) error {
	filter := bson.D{{Key: "name", Value: name}}

	_, err := pr.projectCollection.DeleteOne(ctx, filter)

	return err
}
