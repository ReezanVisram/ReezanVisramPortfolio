package experience

import (
	"context"
	"log/slog"
	domain "reezanvisramportfolio/domain/experience"
	"reezanvisramportfolio/internal/custom_logging"
	"reezanvisramportfolio/internal/database"
)

type ExperienceService interface {
	GetExperience(ctx context.Context) ([]domain.Experience, error)
}

type experienceService struct {
	logger               *slog.Logger
	experienceRepository database.ExperienceRepository
}

func NewExperienceService(logger *slog.Logger, experienceRepository database.ExperienceRepository) ExperienceService {
	return &experienceService{
		logger:               logger,
		experienceRepository: experienceRepository,
	}
}

func (es *experienceService) GetExperience(ctx context.Context) ([]domain.Experience, error) {
	es.logger.Info("experienceService.GetExperience", "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))

	experience, err := es.experienceRepository.GetExperience(ctx)
	if err != nil {
		es.logger.Error("experienceService.GetExperience", "err", err.Error(), "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
		return nil, err
	}

	return experience, nil
}
