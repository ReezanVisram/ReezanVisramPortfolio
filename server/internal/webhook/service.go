package webhook

import (
	"context"
	"fmt"
	domain "reezanvisramportfolio/domain/project"
	"reezanvisramportfolio/internal/database"
)

type WebhookService interface {
	HandleStarWebhookCreated(ctx context.Context, repoName string, repoDescription string, repoLink string, repoReleaseLink string, repoDefaultBranch string, repoTags []string) error
	HandleStarWebhookDeleted(ctx context.Context, repoName string) error
}

type webhookService struct {
	projectRepository database.ProjectRepository
}

func NewWebhookService(projectRepository database.ProjectRepository) WebhookService {
	return &webhookService{
		projectRepository: projectRepository,
	}
}

func (ws *webhookService) HandleStarWebhookCreated(ctx context.Context, repoName string, repoDescription string, repoLink string, repoReleaseLink string, repoDefaultBranch string, repoTags []string) error {
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
			project.IsHardware = true
		} else if tag != "software" {
			project.Technologies = append(project.Technologies, tag)
		}
	}

	err := ws.projectRepository.InsertProject(ctx, project)

	return err
}

func (ws *webhookService) HandleStarWebhookDeleted(ctx context.Context, repoName string) error {
	err := ws.projectRepository.RemoveProjectByName(ctx, repoName)

	return err
}
