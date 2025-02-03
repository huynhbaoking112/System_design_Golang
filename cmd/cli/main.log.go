package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()

	// sugar.Infof("Hello name %s, age: %d", "KingHuynh", 21)

	// //logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "KingHuynh"), zap.Int("age", 21))

	// logger := zap.NewExample()

	// logger.Info("Hello")

	//Development
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello from dev")

	//Production

	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// 3.

	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Info log", zap.Int("line", 2))
}

// Format log
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1738572361.1037612 -> 2025-02-03T15:46:01.102+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	//ts -> Time
	encodeConfig.TimeKey = "time"

	//info -> INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {

	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)

}
