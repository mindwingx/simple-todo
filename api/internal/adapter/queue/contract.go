package queue

import "microservice/internal/core/domain"

//go:generate mockgen -source=./contract.go -destination=./mocks/sqs_mock.go -package=sqs_mocks
type IQueue interface {
	Init()
	Send(todo *domain.Todo) error
}
