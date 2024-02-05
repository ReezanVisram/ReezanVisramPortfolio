package database

import (
	"context"
	domain "reezanvisramportfolio/domain/experience"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExperienceRepository interface {
	GetExperience(ctx context.Context) ([]domain.Experience, error)
}

type experienceRepository struct {
	experienceCollection *mongo.Collection
}

func NewExperienceRepository(experienceCollection *mongo.Collection) ExperienceRepository {
	return &experienceRepository{
		experienceCollection: experienceCollection,
	}
}

func (er *experienceRepository) GetExperience(ctx context.Context) ([]domain.Experience, error) {
	filter := bson.D{{}}
	cursor, err := er.experienceCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []domain.Experience
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
