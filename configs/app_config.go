package configs

import (
	"fmt"
	"github.com/eviccari/simple-sns-producer/adapters"
	"github.com/spf13/viper"
)

const (
	DefaultAppPort                  = 8080
	DefaultHTTPRequestTimeoutInSecs = 10
	DefaultAppLogName               = "SimpleSNSProducer"
	ConfigPath                      = "./cmd/.env"
	ConfigFile                      = "configs.json"
)

var (
	appPort                  int
	requestHTTPTimeoutInSecs int32
	appName                  string
	awsEndpointURL           string
)

// GetAppPort - Get application server port
func GetAppPort() int {
	return appPort
}

// GetRequestHTTPTimeoutInSecs - Get HTTP Request timeout in seconds
func GetRequestHTTPTimeoutInSecs() int32 {
	return requestHTTPTimeoutInSecs
}

// GetAppName - Get application name to provide a custom log identification
func GetAppName() string {
	return appName
}

// init - Load app configuration file
func init() {
	logger := adapters.NewBasicLogger()
	logger.Info(fmt.Sprintf("Loading configurations from path......: %s", ConfigPath))

	viper.SetConfigType("json")
	viper.SetConfigName("configs")
	viper.AddConfigPath(ConfigPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	appName = viper.GetString("appName")
	if appName == "" {
		appName = DefaultAppLogName
	}
	logger.Info(fmt.Sprintf("Starting with appName.................: %s", appName))

	requestHTTPTimeoutInSecs = viper.GetInt32("requestHTTPTimeoutInSecs")
	if requestHTTPTimeoutInSecs == 0 {
		requestHTTPTimeoutInSecs = DefaultHTTPRequestTimeoutInSecs
	}
	logger.Info(fmt.Sprintf("Starting with requestHTTPTimeoutInSecs: %d", requestHTTPTimeoutInSecs))

	appPort = viper.GetInt("appPort")
	if appPort == 0 {
		appPort = DefaultAppPort
	}
	logger.Info(fmt.Sprintf("Starting with appPort.................: %d", appPort))
}
