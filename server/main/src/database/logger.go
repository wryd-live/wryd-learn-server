package database

import "github.com/wryd-live/wryd-learn-server/server/main/src/logging"

var DebugMode bool
var logger *logging.SeelogWrapper

func init() {
	var err error
	logger, err = logging.New()
	if err != nil {
		panic(err)
	}
	Debug(false)
}

func Debug(debugMode bool) {
	DebugMode = debugMode
	if debugMode {
		logger.SetLevel("debug")
	} else {
		logger.SetLevel("info")
	}
}
