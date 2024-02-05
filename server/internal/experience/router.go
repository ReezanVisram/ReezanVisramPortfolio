package experience

import (
	"log/slog"
	"net/http"
	"reezanvisramportfolio/internal/custom_logging"
)

type ExperienceRouter interface {
	GetExperience(w http.ResponseWriter, r *http.Request)
}

type experienceRouter struct {
	logger            *slog.Logger
	experienceService ExperienceService
}

func NewExperienceRouter(logger *slog.Logger, experienceService ExperienceService) ExperienceRouter {
	return &experienceRouter{
		logger:            logger,
		experienceService: experienceService,
	}
}

func (er *experienceRouter) GetExperience(w http.ResponseWriter, r *http.Request) {
	er.logger.Info("experienceRouter.GetExperience", "path", "/experience/", "method", "GET", "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))

	experience, err := er.experienceService.GetExperience(r.Context())
	if err != nil {
		er.logger.Error("experienceRouter.GetExperience", "err", err.Error(), "correlation_id", r.Context().Value(custom_logging.KeyCorrelationId))
		encodeError(w, err)
		return
	}

	encodeResponse(w, experience)
}
