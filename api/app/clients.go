package app

import (
	"fmt"
	"go.uber.org/mock/gomock"
	"log"
	src "microservice"
	"microservice/config"
	"microservice/internal/adapter/locale"
	"microservice/internal/adapter/logger"
	"microservice/internal/adapter/orm"
	"microservice/internal/adapter/queue"
	sqs "microservice/internal/adapter/queue/mocks"
	"microservice/internal/adapter/registry"
	"microservice/pkg/mock"
	"time"
)

func (c *App) InitClients() {
	c.initRegistry()
	c.initService()
	c.initLogger()
	c.initLocale()
	c.initDatabase()
	c.initQueue()
}

// Clients

func (c *App) initRegistry() {
	c.registry = registry.New()
	c.registry.Init()
}

func (c *App) Registry() registry.IRegistry {
	return c.registry
}

func (c *App) initService() {
	c.config = new(config.Service)
	c.registry.Parse(&c.config)

	location, err := time.LoadLocation(c.config.TimeZone)
	if err != nil {
		log.Fatal("[usecase] error loading timezone")
		return
	}

	time.Local = location
}

// Config Service configs
func (c *App) Config() *config.Service {
	return c.config
}

func (c *App) initLocale() {
	c.locale = locale.New(c.registry)
	c.locale.Init()
}

func (c *App) Locale() locale.ILocale {
	return c.locale
}

func (c *App) initLogger() {
	c.logger = logger.New(c.registry)
	c.logger.Init()
}

func (c *App) Logger() logger.ILogger {
	return c.logger
}

func (c *App) initDatabase() {
	c.database = orm.New(c.Config(), c.registry, c.locale)
	c.database.Init()
	c.database.Migrate(fmt.Sprintf("%s/schema/psql", src.Root()))
	c.database.Seed() //NOTE: it is recommended to handle the seeder with CMD or with LIQUIBASE
}

func (c *App) DB() orm.ISql {
	return c.database
}

func (c *App) initQueue() {
	//c.queue = queue.New(c.registry, c.locale) // the main implemented sqs conn.

	ctrl := gomock.NewController(mock.NewController())
	defer ctrl.Finish()

	mockQueue := sqs.NewMockIQueue(ctrl)
	mockQueue.EXPECT().Init().Times(1)
	mockQueue.EXPECT().Send(gomock.Any()).Return(nil).AnyTimes()

	c.queue = mockQueue
	c.queue.Init()
}

func (c *App) Queue() queue.IQueue {
	return c.queue
}
