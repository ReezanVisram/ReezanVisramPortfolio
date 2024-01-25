package database

import (
	"context"
	domain "reezanvisramportfolio/domain/project"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepository interface {
	InsertProject(project domain.Project) error
	RemoveProjectByName(name string) error
}

type projectRepository struct {
	projectCollection *mongo.Collection
}

func NewProjectRepository(projectCollection *mongo.Collection) ProjectRepository {
	return &projectRepository{
		projectCollection: projectCollection,
	}
}

func (pr *projectRepository) InsertProject(project domain.Project) error {
	_, err := pr.projectCollection.InsertOne(context.TODO(), project)

	return err
}

func (pr *projectRepository) RemoveProjectByName(name string) error {
	filter := bson.D{{Key: "name", Value: name}}

	_, err := pr.projectCollection.DeleteOne(context.TODO(), filter)

	return err
}
