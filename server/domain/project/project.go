package domain

type Project struct {
	Name         string   `json:"name" bson:"name"`
	Id           int64    `json:"id" bson:"id"`
	Description  string   `json:"description" bson:"description"`
	RepoLink     string   `json:"repo_link" bson:"repo_link"`
	ReleaseLink  string   `json:"release_link" bson:"release_link"`
	ImageLink    string   `json:"image_link" bson:"image_link"`
	IsHardware   bool     `json:"is_hardware" bson:"is_hardware"`
	Technologies []string `json:"technologies" bson:"technologies"`
}
