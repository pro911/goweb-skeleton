package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var (
	logger *zap.Logger
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func NewZapLogger() {

	//同时将日志输入系统输入输出屏幕上
	consoleSyncer := zapcore.AddSync(os.Stdout)

	writeSyncer := getLogWriter()

	encoder := getEncoder()

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, zap.DebugLevel),
		zapcore.NewCore(encoder, consoleSyncer, zap.DebugLevel),
	)

	// zap打印时将整个调用的stack链路会存放到内存中，默认打印调用处的caller信息。所以需要再初始化zap时额外增加AddCallerSkip跳过指定层级的caller
	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	defer logger.Sync()

	Logger = logger

	Sugar = Logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = CustomTimeEncoder

	//if conf.Conf.Log.Format == "json" {
	//	return zapcore.NewJSONEncoder(encoderConfig)
	//}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/golang/task-info.logger", viper.GetString("LOG_DIR")), // 日志文件的位置
		MaxSize:    100,                                                                   // 在进行切割之前，日志文件的最大大小(MB为单位)
		MaxBackups: 5,                                                                     // 保留旧文件的最大个数
		MaxAge:     7,                                                                     // 保留旧文件的最大天数
		Compress:   false,                                                                 // 是否压缩/归档旧文件
	}

	return zapcore.AddSync(lumberJackLogger)
}

// CustomTimeEncoder 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
