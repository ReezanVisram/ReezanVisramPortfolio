package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"reezanvisramportfolio/internal/custom_middleware"
	"reezanvisramportfolio/internal/database"
	"reezanvisramportfolio/internal/experience"
	"reezanvisramportfolio/internal/project"
	"reezanvisramportfolio/internal/webhook"

	"cloud.google.com/go/storage"
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
	CLOUDSTORAGE_BUCKET_NAME := os.Getenv("CLOUDSTORAGE_BUCKET_NAME")
	CLOUDSTORAGE_FILENAME_TO_FETCH := os.Getenv("CLOUDSTORAGE_FILENAME_TO_FETCH")

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

	storageClient, err := storage.NewClient(context.TODO())
	if err != nil {
		panic(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	projectCollection := client.Database("reezanvisramportfolio").Collection("projects")
	experienceCollection := client.Database("reezanvisramportfolio").Collection("experience")
	projectRepo := database.NewProjectRepository(projectCollection)
	experienceRepo := database.NewExperienceRepository(experienceCollection)

	webhookService := webhook.NewWebhookService(logger, projectRepo)
	webhookRouter := webhook.NewWebhookRouter(logger, WEBHOOK_SECRET, webhookService)

	projectService := project.NewProjectService(logger, projectRepo)
	projectRouter := project.NewProjectRouter(logger, projectService)

	experienceService := experience.NewExperienceService(logger, experienceRepo)
	experienceRouter := experience.NewExperienceRouter(logger, experienceService)

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

	r.Get("/resume", func(w http.ResponseWriter, r *http.Request) {
		rc, err := storageClient.Bucket(CLOUDSTORAGE_BUCKET_NAME).Object(CLOUDSTORAGE_FILENAME_TO_FETCH).NewReader(r.Context())
		if err != nil {
			w.WriteHeader(500)
			return
		}

		defer rc.Close()
		body, err := io.ReadAll(rc)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Header().Add("Content-Type", "application/pdf")
		w.Write(body)
	})

	r.Group(func(r chi.Router) {
		r.Use(custom_middleware.ContentTypeMiddleware)
		r.Route("/webhooks", func(r chi.Router) {
			r.Post("/", webhookRouter.PostWebhookHandler)
		})

		r.Route("/projects", func(r chi.Router) {
			r.Get("/", projectRouter.GetProjects)
		})

		r.Route("/experience", func(r chi.Router) {
			r.Get("/", experienceRouter.GetExperience)
		})
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
