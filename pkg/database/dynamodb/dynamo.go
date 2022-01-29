package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/devlibx/gox-base/util"
	dynamoOrm "github.com/guregu/dynamo"
	"go.uber.org/fx"
	"net/http"
	"time"
)

// DynamoConfig DB setup configuration
type DynamoConfig struct {
	AWSAccessKey      string `json:"aws_access_key" yaml:"aws_access_key"`
	AWSSecretKey      string `json:"aws_secret_key" yaml:"aws_secret_key"`
	SessionToken      string `json:"session_token" yaml:"session_token"`
	URL               string `json:"url" yaml:"url"`
	DaxUrl            string `json:"dax_url" yaml:"dax_url"`
	Timeout           int64  `json:"timeout" yaml:"timeout"`
	Region            string `json:"region" yaml:"region"`
	DaxGetTimeoutInMs int    `json:"dax_get_timeout_ms" yaml:"dax_get_timeout_ms"`
}

var DatabaseModule = fx.Options(
	fx.Provide(func(dynamoConfig *DynamoConfig) (*Dynamo, error) { return buildDynamo(dynamoConfig) }),
	fx.Provide(newUserDao),
)

func buildDynamo(dynamoConfig *DynamoConfig) (*Dynamo, error) {
	var err error
	d := &Dynamo{}

	awsConfig := &aws.Config{
		Region:   &dynamoConfig.Region,
		Endpoint: &dynamoConfig.URL,
	}

	if !util.IsStringEmpty(dynamoConfig.AWSAccessKey) {
		awsConfig.Credentials = credentials.NewStaticCredentials(
			dynamoConfig.AWSAccessKey,
			dynamoConfig.AWSSecretKey,
			dynamoConfig.SessionToken,
		)
	}

	if dynamoConfig.Timeout != 0 {
		awsConfig.HTTPClient = &http.Client{
			Timeout: time.Duration(dynamoConfig.Timeout) * time.Second,
		}
	}

	if d.Session, err = session.NewSession(awsConfig); err == nil {
		d.DynamoDb = dynamoOrm.New(d.Session, awsConfig)
	}

	return d, err
}

type Dynamo struct {
	Session  *session.Session
	DynamoDb *dynamoOrm.DB
}
