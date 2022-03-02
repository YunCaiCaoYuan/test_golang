package main

import (
	"go.uber.org/zap"
	"time"
)

type gameOfData struct {
	GameId int32
	Val    int32
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "test"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	// 字段大写，才会打印
	list := make([]*gameOfData, 0)
	list = append(list, &gameOfData{GameId: 1, Val: 1})
	logger.Info("test", zap.Any("list", list))
}
