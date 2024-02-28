package domain

type Message struct {
	Name       string `json:"name" bson:"name"`
	Email      string `json:"email" bson:"email"`
	Subject    string `json:"subject" bson:"subject"`
	Message    string `json:"message" bson:"message"`
	HasAlerted bool   `json:"has_alerted" bson:"has_alerted"`
}
