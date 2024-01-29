package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"reezanvisramportfolio/internal/custom_middleware"
	"reezanvisramportfolio/internal/database"
	"reezanvisramportfolio/internal/project"
	"reezanvisramportfolio/internal/webhook"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HomeResponse struct {
	API     string `json:"api"`
	Version string `json:"version"`
}

func main() {
	WEBHOOK_SECRET := os.Getenv("WEBHOOK_SECRET")
	MONGODB_CONNECTION_METHOD := os.Getenv("MONGODB_CONNECTION_METHOD")
	MONGODB_USERNAME := os.Getenv("MONGODB_USERNAME")
	MONGODB_PASSWORD := os.Getenv("MONGODB_PASSWORD")
	MONGODB_HOST := os.Getenv("MONGODB_HOST")
	MONGODB_CONNECTION_OPTIONS := os.Getenv("MONGODB_CONNECTION_OPTIONS")

	r := chi.NewRouter()

	client, err := connectToMongo(MONGODB_CONNECTION_METHOD, MONGODB_USERNAME, MONGODB_PASSWORD, MONGODB_HOST, MONGODB_CONNECTION_OPTIONS)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	projectCollection := client.Database("reezanvisramportfolio").Collection("projects")
	projectRepo := database.NewProjectRepository(projectCollection)

	webhookService := webhook.NewWebhookService(logger, projectRepo)
	webhookRouter := webhook.NewWebhookRouter(logger, WEBHOOK_SECRET, webhookService)

	projectService := project.NewProjectService(logger, projectRepo)
	projectRouter := project.NewProjectRouter(logger, projectService)

	r.Use(custom_middleware.ContentTypeMiddleware)
	r.Use(custom_middleware.CorrelationIdMiddleware)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := json.Marshal(HomeResponse{
			API:     "reezanvisramportfolio",
			Version: "0.0.1",
		})

		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Write(response)
	})

	r.Route("/webhooks", func(r chi.Router) {
		r.Post("/", webhookRouter.PostWebhookHandler)
	})

	r.Route("/projects", func(r chi.Router) {
		r.Get("/", projectRouter.GetProjects)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r)
}

func connectToMongo(mongoConnMethod string, mongoUsername string, mongoPassword string, mongoHost string, mongoConnOptions string) (*mongo.Client, error) {
	mongoURI := fmt.Sprintf("%s://%s:%s@%s/%s", mongoConnMethod, mongoUsername, mongoPassword, mongoHost, mongoConnOptions)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	return client, nil
}
