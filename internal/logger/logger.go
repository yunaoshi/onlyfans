package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
	once   sync.Once
)

func Init() error {
	var err error
	once.Do(func() {
		today := fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02"))
		file, ferr := os.OpenFile(today, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if ferr != nil {
			err = ferr
			return
		}
		fileEncoderCfg := zap.NewProductionEncoderConfig()
		fileCore := zapcore.NewCore(
			zapcore.NewJSONEncoder(fileEncoderCfg),
			zapcore.AddSync(file),
			zap.InfoLevel,
		)
		consoleEncoderCfg := zap.NewProductionEncoderConfig()
		consoleEncoderCfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		}
		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(consoleEncoderCfg),
			zapcore.AddSync(os.Stdout),
			zap.InfoLevel,
		)
		logger = zap.New(zapcore.NewTee(fileCore, consoleCore))
	})
	return err
}

// Logger returns the global zap.Logger instance.
func Logger() *zap.Logger {
	if logger == nil {
		// No-op logger if not initialized
		return zap.NewNop()
	}
	return logger
}
