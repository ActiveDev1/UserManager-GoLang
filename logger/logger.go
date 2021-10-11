package logger

import (
	"fmt"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ZapConfig zap.Config        `json:"zap_config" yaml:"zap_config"`
	LogRotate lumberjack.Logger `json:"log_rotate" yaml:"log_rotate"`
}

type Logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger(env string) *Logger {
	configYaml, err := ioutil.ReadFile("./zaplogger." + env + ".yml")
	if err != nil {
		fmt.Printf("Failed to read logger configuration: %s", err)
		os.Exit(2)
	}
	var myConfig *Config
	if err = yaml.Unmarshal(configYaml, &myConfig); err != nil {
		fmt.Printf("Failed to read zap logger configuration: %s", err)
		os.Exit(2)
	}
	var zap *zap.Logger
	zap, err = build(myConfig)
	if err != nil {
		fmt.Printf("Failed to compose zap logger : %s", err)
		os.Exit(2)
	}
	sugar := zap.Sugar()
	// set package varriable logger.
	logger := &Logger{Zap: sugar}
	logger.Zap.Infof("Success to read zap logger configuration: zaplogger." + env + ".yml")
	_ = zap.Sync()
	return logger
}

// GetZapLogger returns zapSugaredLogger
func (log *Logger) GetZapLogger() *zap.SugaredLogger {
	return log.Zap
}
