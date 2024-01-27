mocks:
	mockgen -source=./internal/webhook/service.go -destination=./internal/mocks/webhook/service.go
	mockgen -source=./internal/database/project_repository.go -destination=./internal/mocks/database/project_repository.go

tests:
	go test -cover ./...