package config

type Queue struct {
	Conn                 string `mapstructure:"QUEUE_CONNECTION"`
	Prefix               string `mapstructure:"QUEUE_PREFIX"`
	SqsKey               string `mapstructure:"QUEUE_SQS_KEY"`
	SqsSecret            string `mapstructure:"QUEUE_SQS_SECRET"`
	SqsRegion            string `mapstructure:"QUEUE_SQS_REGION"`
	SqsQueue             string `mapstructure:"QUEUE_SQS_QUEUE"`
	SqsMaxNumMessages    int    `mapstructure:"QUEUE_SQS_MAX_NUM_MESSAGES"`
	SqsVisibilityTimout  int    `mapstructure:"QUEUE_SQS_VISIBILITY_TIMOUT"`
	AwsHttpClientTimeout int    `mapstructure:"AWS_HTTP_CLIENT_TIMEOUT"`
}
