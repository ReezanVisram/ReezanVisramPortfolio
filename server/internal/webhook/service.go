package webhook

import (
	"context"
	"fmt"
	"log/slog"
	domain "reezanvisramportfolio/domain/project"
	"reezanvisramportfolio/internal/custom_logging"
	"reezanvisramportfolio/internal/database"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type WebhookService interface {
	HandleStarWebhookCreated(
		ctx context.Context,
		repoName string,
		repoId int64,
		repoDescription string,
		repoLink string,
		repoReleaseLink string,
		repoDefaultBranch string,
		repoTags []string) error
	HandleStarWebhookDeleted(ctx context.Context, repoId int64) error
}

var specialTechnologyCases = map[string]string{"opengl": "OpenGL", "cpp": "C++", "nodejs": "Node.js", "mysql": "MySQL", "fpga": "FPGA", "html": "HTML", "css": "CSS", "axi4-stream": "AXI4-Stream", "i2s": "I2S", "uart": "UART", "vga": "VGA", "mongodb": "MonogDB"}

type webhookService struct {
	logger            *slog.Logger
	projectRepository database.ProjectRepository
	caser             cases.Caser
}

func NewWebhookService(logger *slog.Logger, projectRepository database.ProjectRepository) WebhookService {
	return &webhookService{
		logger:            logger,
		projectRepository: projectRepository,
		caser:             cases.Title(language.BritishEnglish),
	}
}

func (ws *webhookService) HandleStarWebhookCreated(
	ctx context.Context,
	repoName string,
	repoId int64,
	repoDescription string,
	repoLink string,
	repoReleaseLink string,
	repoDefaultBranch string,
	repoTags []string) error {
	ws.logger.Info("webhookService.HandleStarWebhookCreated", "repo_name", repoName, "repo_description", "repo_id", strconv.FormatInt(repoId, 10), repoDescription, "repo_link", repoLink, "repo_release_link", repoReleaseLink, "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))

	if ws.projectExists(ctx, repoId) {
		return ErrProjectExists
	}

	project := domain.Project{
		Name:        strings.Replace(repoName, "-", " ", -1),
		Id:          repoId,
		Description: repoDescription,
		RepoLink:    repoLink,
		ReleaseLink: repoReleaseLink,
		ImageLink:   fmt.Sprintf("%s/%s/featured_screenshot.png", strings.Replace(repoLink, "github.com", "raw.githubusercontent.com", 1), repoDefaultBranch),
		IsHardware:  false,
	}

	for _, tag := range repoTags {
		if tag == "hardware" {
			ws.logger.Info("webhookService.HandleStarWebhookCreated", "repo_is_hardware", "true", "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
			project.IsHardware = true
		} else if tag != "software" {
			project.Technologies = append(project.Technologies, ws.convertTag(tag))
		}
	}

	err := ws.projectRepository.InsertProject(ctx, project)
	if err != nil {
		ws.logger.Error("webhookService.HandleStarWebhookCreated", "err", err.Error(), "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
		return err
	}

	return nil
}

func (ws *webhookService) HandleStarWebhookDeleted(ctx context.Context, repoId int64) error {
	ws.logger.Info("webhookService.HandleStarWebhookDeleted", "repo_id", strconv.FormatInt(repoId, 10), "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))

	if !ws.projectExists(ctx, repoId) {
		return ErrProjectDoesNotExist
	}

	err := ws.projectRepository.RemoveProjectById(ctx, repoId)
	if err != nil {
		ws.logger.Error("webhookService.HandleStarWebhookDeleted", "err", err.Error(), "correlation_id", ctx.Value(custom_logging.KeyCorrelationId))
		return err
	}

	return nil
}

func (ws *webhookService) projectExists(ctx context.Context, repoId int64) bool {
	_, err := ws.projectRepository.GetProjectById(ctx, repoId)

	return !(err == mongo.ErrNoDocuments)
}

func (ws *webhookService) convertTag(tag string) string {
	convertedString, ok := specialTechnologyCases[tag]
	if ok {
		return convertedString
	}

	return ws.caser.String(tag)
}
