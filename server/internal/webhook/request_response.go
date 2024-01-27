package webhook

type StarWebhookRequest struct {
	Action     string                       `json:"action"`
	Repository StarWebhookRepositoryRequest `json:"repository"`
	Sender     StarWebhookSenderRequest     `json:"sender"`
}

type StarWebhookRepositoryRequest struct {
	Name          string                  `json:"name"`
	Id            int64                   `json:"id"`
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
