mocks:
	mockgen -source=./internal/webhook/service.go -destination=./internal/mocks/webhook/service.go
	mockgen -source=./internal/database/project_repository.go -destination=./internal/mocks/database/project_repository.go
	mockgen -source=./internal/project/service.go -destination=./internal/mocks/project/service.go
	mockgen -source=./internal/database/experience_repository.go -destination=./internal/mocks/database/experience_repository.go
	mockgen -source=./internal/experience/service.go -destination=./internal/mocks/experience/service.go

tests:
	go test -cover ./...