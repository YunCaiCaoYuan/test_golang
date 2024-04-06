package main

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	lumberjacklogger := &lumberjack.Logger{
		Filename:   "./log-rotate-test.log",
		MaxSize:    1,    // megabytes,M为单位,达到这个设置数后就进行日志切割
		MaxBackups: 3,    // 保留旧文件最大份数
		MaxAge:     28,   // days,旧文件最大保存天数
		Compress:   true, // disabled by default，是否压缩日志归档，默认不压缩
	}
	defer lumberjacklogger.Close()
	config := zap.NewProductionEncoderConfig()
	// 设置时间格式
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	core := zapcore.NewCore(
		fileEncoder,                       //编码设置
		zapcore.AddSync(lumberjacklogger), //输出到文件
		zap.InfoLevel,                     //日志等级
	)
	logger := zap.New(core)
	defer logger.Sync()

	// 测试分割日志
	for i := 0; i < 8000; i++ {
		logger.With(
			zap.String("url", fmt.Sprintf("www.test%d.com", i)),
			zap.String("name", "jimmmyr"),
			zap.Int("age", 23),
			zap.String("agradege", "no111-000222"),
		).Info("test info ")
	}
}
