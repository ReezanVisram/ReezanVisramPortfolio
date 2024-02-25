package domain

type Experience struct {
	Name            string   `json:"name" bson:"name"`
	JobTitle        string   `json:"job_title" bson:"job_title"`
	StartAndEndDate string   `json:"start_and_end_date" bson:"start_and_end_date"`
	BulletPoints    []string `json:"bullet_points" bson:"bullet_points"`
	Tools           []string `json:"tools" bson:"tools"`
}
