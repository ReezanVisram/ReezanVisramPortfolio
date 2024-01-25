package webhook

import "errors"

var (
	ErrMissingSignature              = errors.New("webhook signature is missing")
	ErrInvalidSignature              = errors.New("webhook signature is incorrect")
	ErrCouldNotReadBody              = errors.New("request body is malformed")
	ErrInvalidStarWebhookRequestBody = errors.New("request body is not a correct star webhook request")
	ErrInvalidSender                 = errors.New("that user is not permitted to add a project")
	ErrInvalidOwner                  = errors.New("that user's repos are not permitted to be featured as projects")
	ErrRepoPrivate                   = errors.New("private repos are not permitted to be featured as projects")
	ErrIsFork                        = errors.New("forks are not permitted to be featured as projects")
)
