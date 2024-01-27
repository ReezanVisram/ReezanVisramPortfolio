package webhook

import (
	"context"
	"fmt"
	"log/slog"
	domain "reezanvisramportfolio/domain/project"
	"reezanvisramportfolio/internal/custom_logging"
	"reezanvisramportfolio/internal/database"
)

type WebhookService interface {
	HandleStarWebhookCreated(
		ctx context.Context,
		repoName string,
		repoDescription string,
		repoLink string,
		repoReleaseLink string,
		repoDefaultBranch string,
		repoTags []string) error
	HandleStarWebhookDeleted(ctx context.Context, repoName string) error
}

type webhookService struct {
	logger            *slog.Logger
	projectRepository database.ProjectRepository
}

func NewWebhookService(logger *slog.Logger, projectRepository database.ProjectRepository) WebhookService {
	return &webhookService{
		logger:            logger,
		projectRepository: projectRepository,
	}
}

func (ws *webhookService) HandleStarWebhookCreated(
	ctx context.Context,
	repoName string,
	repoDescription string,
	repoLink string,
	repoReleaseLink string,
	repoDefaultBranch string,
	repoTags []string) error {
	ws.logger.Info("webhookService.HandleStarWebhookCreated", "repo_name", repoName, "repo_description", repoDescription, "repo_link", repoLink, "repo_release_link", repoReleaseLink, "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
	project := domain.Project{
		Name:        repoName,
		Description: repoDescription,
		RepoLink:    repoLink,
		ReleaseLink: repoReleaseLink,
		ImageLink:   fmt.Sprintf("%s/blob/%s/featured_screenshot.png", repoLink, repoDefaultBranch),
		IsHardware:  false,
	}

	for _, tag := range repoTags {
		if tag == "hardware" {
			ws.logger.Info("webhookService.HandleStarWebhookCreated", "repo_is_hardware", "true", "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
			project.IsHardware = true
		} else if tag != "software" {
			project.Technologies = append(project.Technologies, tag)
		}
	}

	err := ws.projectRepository.InsertProject(ctx, project)
	if err != nil {
		ws.logger.Error("webhookService.HandleStarWebhookCreated", "err", err.Error(), "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
		return err
	}

	return nil
}

func (ws *webhookService) HandleStarWebhookDeleted(ctx context.Context, repoName string) error {
	ws.logger.Info("webhookService.HandleStarWebhookDeleted", "repo_name", repoName, "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))

	err := ws.projectRepository.RemoveProjectByName(ctx, repoName)
	if err != nil {
		ws.logger.Error("webhookService.HandleStarWebhookDeleted", "err", err.Error(), "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
		return err
	}

	return nil
}
