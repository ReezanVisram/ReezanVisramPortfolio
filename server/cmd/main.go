package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HomeResponse struct {
	API     string `json:"api"`
	Version string `json:"version"`
}

type StarWebhookRequest struct {
	Action     string                       `json:"action"`
	Repository StarWebhookRepositoryRequest `json:"repository"`
	Sender     StarWebhookSenderRequest     `json:"sender"`
}

type StarWebhookRepositoryRequest struct {
	Name          string                  `json:"name"`
	IsPrivate     bool                    `json:"private"`
	Owner         StarWebhookOwnerRequest `json:"owner"`
	Description   string                  `json:"description"`
	RepoLink      string                  `json:"html_url"`
	ReleaseLink   string                  `json:"homepage"`
	Tags          []string                `json:"topics"`
	NumStars      int                     `json:"num_stargazers"`
	IsFork        bool                    `json:"fork"`
	DefaultBranch string                  `json:"default_branch"`
}

type StarWebhookOwnerRequest struct {
	Username string `json:"login"`
}

type StarWebhookSenderRequest struct {
	Username string `json:"login"`
}

type Project struct {
	Name         string   `bson:"name"`
	Description  string   `bson:"description"`
	RepoLink     string   `bson:"repo_link"`
	ReleaseLink  string   `bson:"release_link"`
	ImageLink    string   `bson:"image_link"`
	IsHardware   bool     `bson:"is_hardware"`
	Technologies []string `bson:"technologies"`
}

func main() {
	r := chi.NewRouter()

	mongoURI := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("reezanvisramportfolio").Collection("projects")

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

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

	r.Post("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Github-Event") != "star" {
			w.WriteHeader(400)
			return
		}

		webhookSecret := os.Getenv("WEBHOOK_SECRET")

		var signature string
		if signature = r.Header.Get("X-Hub-Signature-256"); signature == "" {
			w.WriteHeader(403)
			fmt.Printf("Signature is missing!\n")
			return
		}

		h := hmac.New(sha256.New, []byte(webhookSecret))

		payload, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(500)
			fmt.Printf("Unable to read request body\n")
			return
		}

		h.Write(payload)

		expectedSignature := fmt.Sprintf("sha256=%s", hex.EncodeToString(h.Sum(nil)))

		if expectedSignature != signature {
			w.WriteHeader(403)
			fmt.Printf("Signature received does not match expected\n")
			return
		}

		fmt.Printf("Signature does match what was expected\n")
		starWebhookRequest := &StarWebhookRequest{}
		if err := json.Unmarshal(payload, starWebhookRequest); err != nil {
			w.WriteHeader(400)
			fmt.Printf("Unable to unmarshal request")
			return
		}
		fmt.Printf("%+v\n", starWebhookRequest)
		if starWebhookRequest.Sender.Username != "ReezanVisram" {
			w.WriteHeader(412)
			fmt.Printf("A user other than ReezanVisram is not permitted to add projects")
			return
		}

		if starWebhookRequest.Repository.Owner.Username != "ReezanVisram" {
			w.WriteHeader(412)
			fmt.Printf("Only ReezanVisram's projects can be featured")
			return
		}

		if starWebhookRequest.Repository.IsPrivate {
			w.WriteHeader(412)
			fmt.Printf("Only public repositories can be featured")
			return
		}

		if starWebhookRequest.Repository.IsFork {
			w.WriteHeader(412)
			fmt.Printf("Forks cannot be featured")
			return
		}

		fmt.Printf("Webhook Action: %s\n", starWebhookRequest.Action)
		if starWebhookRequest.Action == "created" {
			newProject := Project{
				Name:        starWebhookRequest.Repository.Name,
				Description: starWebhookRequest.Repository.Description,
				RepoLink:    starWebhookRequest.Repository.RepoLink,
				ReleaseLink: starWebhookRequest.Repository.ReleaseLink,
				ImageLink:   fmt.Sprintf("%s/blob/%s/featured_screenshot.png", starWebhookRequest.Repository.RepoLink, starWebhookRequest.Repository.DefaultBranch),
			}
			newProject.IsHardware = false
			for _, tag := range starWebhookRequest.Repository.Tags {
				if tag == "hardware" {
					newProject.IsHardware = true
				} else if tag != "software" {
					newProject.Technologies = append(newProject.Technologies, tag)
				}
			}

			_, err := coll.InsertOne(context.TODO(), newProject)
			if err != nil {
				w.WriteHeader(500)
				fmt.Printf("Error occurred inserting element into mongo: %s\n", err.Error())
				return
			}
		} else if starWebhookRequest.Action == "deleted" {
			filter := bson.D{{"name", starWebhookRequest.Repository.Name}}

			_, err := coll.DeleteOne(context.TODO(), filter)

			if err != nil {
				w.WriteHeader(500)
				fmt.Printf("Could not delete element from mongo: %s\n", err.Error())
				return
			}
		}
	})

	http.ListenAndServe(":3000", r)
}
