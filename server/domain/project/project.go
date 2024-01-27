package domain

type Project struct {
	Name         string   `bson:"name"`
	Id           int64    `bson:"id"`
	Description  string   `bson:"description"`
	RepoLink     string   `bson:"repo_link"`
	ReleaseLink  string   `bson:"release_link"`
	ImageLink    string   `bson:"image_link"`
	IsHardware   bool     `bson:"is_hardware"`
	Technologies []string `bson:"technologies"`
}
