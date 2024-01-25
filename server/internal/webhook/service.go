package webhook

import (
	"fmt"
	domain "reezanvisramportfolio/domain/project"
	"reezanvisramportfolio/internal/database"
)

type WebhookService interface {
	HandleStarWebhookCreated(repoName string, repoDescription string, repoLink string, repoReleaseLink string, repoDefaultBranch string, repoTags []string) error
	HandleStarWebhookDeleted(repoName string) error
}

type webhookService struct {
	projectRepository database.ProjectRepository
}

func NewWebhookService(projectRepository database.ProjectRepository) WebhookService {
	return &webhookService{
		projectRepository: projectRepository,
	}
}

func (ws *webhookService) HandleStarWebhookCreated(repoName string, repoDescription string, repoLink string, repoReleaseLink string, repoDefaultBranch string, repoTags []string) error {
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

	err := ws.projectRepository.InsertProject(project)

	return err
}

func (ws *webhookService) HandleStarWebhookDeleted(repoName string) error {
	err := ws.projectRepository.RemoveProjectByName(repoName)

	return err
}
