package project

import (
	"context"
	"log/slog"
	domain "reezanvisramportfolio/domain/project"
	"reezanvisramportfolio/internal/custom_logging"
	"reezanvisramportfolio/internal/database"
)

type ProjectService interface {
	GetAllProjects(ctx context.Context) ([]domain.Project, error)
}

type projectService struct {
	logger            *slog.Logger
	projectRepository database.ProjectRepository
}

func NewProjectService(logger *slog.Logger, projectRepository database.ProjectRepository) ProjectService {
	return &projectService{
		logger:            logger,
		projectRepository: projectRepository,
	}
}

func (ps *projectService) GetAllProjects(ctx context.Context) ([]domain.Project, error) {
	ps.logger.Info("projectService.GetAllProjects", "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))

	projects, err := ps.projectRepository.GetAllProjects(ctx)
	if err != nil {
		ps.logger.Error("projectService.GetAllProjects", "err", err.Error(), "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
		return nil, err
	}

	return projects, nil
}
