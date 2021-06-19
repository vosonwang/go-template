package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	path   string
	level  = zapcore.InfoLevel
	sugar  *zap.SugaredLogger
	logger *zap.Logger
)

// SetPath 开启打印到日志文件
func SetPath(logPath string) {
	path = logPath
}

// SetLevel 设置日志级别
// 默认level为info，默认不会打印比info（0）级别低的日志，即不会打印debug(-1)
// 级别参考:zapcore.Level
func SetLevel(l zapcore.Level) {
	level = l
}

func Init(project string, hooks ...func(zapcore.Entry) error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		CallerKey:      "file",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	// 打印到控制台
	ws := []zapcore.WriteSyncer{zapcore.AddSync(os.Stdout)}

	// 如果写入文件地址不为空，则打印到文件
	if path != "" {
		hook := lumberjack.Logger{
			Filename:   path, // 日志文件路径
			MaxSize:    128,  // 每个日志文件保存的最大尺寸 单位：M
			MaxBackups: 30,   // 日志文件最多保存多少个备份
			MaxAge:     7,    // 文件最多保存多少天
			Compress:   true, // 是否压缩
		}
		// 增加打印到文件
		ws = append(ws, zapcore.AddSync(&hook))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		zapcore.NewMultiWriteSyncer(ws...),    // 输出目标
		atomicLevel,                           // 日志级别
	)

	// 设置初始化字段
	filed := zap.Fields(zap.String("project", project))

	// 设置调用方层级
	skip := zap.AddCallerSkip(1)

	logger = zap.New(core, zap.AddCaller(), filed, skip, zap.Hooks(hooks...))

	sugar = logger.Sugar()
}

func Debug(args ...interface{}) {
	sugar.Debug(args...)
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

func Warn(args ...interface{}) {
	sugar.Warn(args...)
}

func Error(args ...interface{}) {
	sugar.Error(args...)
}
