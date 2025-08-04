package queue

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.uber.org/zap"
	"log"
	"microservice/config"
	"microservice/internal/adapter/locale"
	"microservice/internal/adapter/logger"
	"microservice/internal/adapter/registry"
	"microservice/internal/core/domain"
	"microservice/internal/server/http/status"
	"microservice/pkg/meta"
	"net/http"
	"time"
)

type queue struct {
	lgr     logger.ILogger
	l       locale.ILocale
	service *config.Service
	config  *config.Queue
	sqs     *sqs.SQS
}

func New(registry registry.IRegistry, locale locale.ILocale) IQueue {
	q := new(queue)
	registry.Parse(&q.service)
	registry.Parse(&q.config)
	q.l = locale
	return q
}

func (q *queue) Init() {
	var err error
	awsConf := aws.Config{
		Region:     aws.String(q.config.SqsRegion),
		HTTPClient: &http.Client{Timeout: time.Duration(q.config.AwsHttpClientTimeout) * time.Second},
	}

	if q.service.Env == "development" {
		// secret for specific queue
		awsConf.Credentials = credentials.NewStaticCredentials(q.config.SqsKey, q.config.SqsSecret, "")
	}

	awsSession, err := session.NewSession(&awsConf)

	if err != nil {
		q.lgr.Error("aws.init", zap.Error(err))
		log.Fatalf("[queue][sqs] init err: %s", err)
		return
	}

	q.sqs = sqs.New(awsSession)
	if err != nil {
		return
	}

	q.lgr.Info("aws.init.success", zap.String("env", q.service.Env))
	return
}

func (q *queue) Send(todo *domain.Todo) (err error) {
	queueUrl := aws.String(fmt.Sprintf("%s_%s", q.config.Prefix, "todo"))

	message, err := json.Marshal(todo)
	if err != nil {
		q.lgr.Error("queue.sqs.json.maral", zap.String("todo.id", todo.UUID().String()), zap.Error(err))
		err = meta.ServiceErr(status.Failed)
		return
	}

	_, err = q.sqs.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(message)),
		QueueUrl:    queueUrl,
	})

	if err != nil {
		q.lgr.Error("queue.sqs.send", zap.String("todo.id", todo.UUID().String()), zap.Error(err))
		err = meta.ServiceErr(status.Failed)
		return
	}

	q.lgr.Info("queue.sqs.sent.success", zap.String("todo.id", todo.UUID().String()))
	return
}
