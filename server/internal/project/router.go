package project

import (
	"log/slog"
	"net/http"
	"reezanvisramportfolio/internal/custom_logging"
)

type ProjectRouter interface {
	GetProjects(w http.ResponseWriter, r *http.Request)
}

type projectRouter struct {
	logger         *slog.Logger
	projectService ProjectService
}

func NewProjectRouter(logger *slog.Logger, projectService ProjectService) ProjectRouter {
	return &projectRouter{
		logger:         logger,
		projectService: projectService,
	}
}

func (pr *projectRouter) GetProjects(w http.ResponseWriter, r *http.Request) {
	pr.logger.Info("projectRouter.GetProject", "path", "/projects/", "method", "GET", "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))

	projects, err := pr.projectService.GetAllProjects(r.Context())
	if err != nil {
		pr.logger.Error("projectRouter.GetProject", "err", err.Error(), "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
		encodeError(w, err)
		return
	}

	encodeResponse(w, projects)
}
