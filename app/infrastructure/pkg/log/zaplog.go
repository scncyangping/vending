package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	logger                         *zap.SugaredLogger
	sp                             = string(filepath.Separator)
	errWS, warnWS, infoWS, debugWS zapcore.WriteSyncer       // IO输出
	debugConsoleWS                 = zapcore.Lock(os.Stdout) // 控制台标准输出
	errorConsoleWS                 = zapcore.Lock(os.Stderr)
)

// Config logger
type Config struct {
	LogFileDir    string `yaml:"logFileDir"` //文件保存地方
	AppName       string `yaml:"appName"`    //日志文件前缀
	ErrorFileName string `yaml:"errorFileName"`
	WarnFileName  string `yaml:"warnFileName"`
	InfoFileName  string `yaml:"infoFileName"`
	DebugFileName string `yaml:"debugFileName"`
	MaxSize       int    `yaml:"maxSize"`     //日志文件小大（M）
	MaxBackups    int    `yaml:"maxBackups"`  // 最多存在多少个切片文件
	MaxAge        int    `yaml:"maxAge"`      //保存的最大天数
	Development   bool   `yaml:"development"` //是否是开发模式
	Level         int8   `yaml:"level"`       //日志等级
}

type Options struct {
	SelfConfig *Config
	zap.Config
}

type ZapLogger struct {
	*zap.Logger
	sync.RWMutex
	Opts        *Options `json:"opts"`
	zapConfig   zap.Config
	initialized bool
}

func New(c *Config) {
	if logger == nil {
		logger = c.newZapLogger().Sugar()
	}
}

func Logger() *zap.SugaredLogger {
	return logger
}

func (c *Config) newZapLogger() *zap.Logger {
	l := &ZapLogger{}
	l.Lock()
	defer l.Unlock()
	if l.initialized {
		l.Info("[NewZapLogger] logger init")
		return nil
	}

	l.Opts = &Options{
		SelfConfig: c,
	}

	if l.Opts.SelfConfig.LogFileDir == "" {
		l.Opts.SelfConfig.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
		l.Opts.SelfConfig.LogFileDir += sp + "logs" + sp
	}

	if l.Opts.Development {
		l.zapConfig = zap.NewDevelopmentConfig()
		l.zapConfig.EncoderConfig.EncodeTime = timeEncoder
	} else {
		l.zapConfig = zap.NewProductionConfig()
		l.zapConfig.EncoderConfig.EncodeTime = timeUnixNano
	}

	if l.Opts.OutputPaths == nil || len(l.Opts.OutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stdout"}
	}

	if l.Opts.ErrorOutputPaths == nil || len(l.Opts.ErrorOutputPaths) == 0 {
		l.zapConfig.OutputPaths = []string{"stderr"}
	}

	l.zapConfig.Level.SetLevel(zapcore.Level(l.Opts.SelfConfig.Level))

	l.build()

	l.initialized = true

	l.Info("[NewLogger] success")
	return l.Logger
}

func (l *ZapLogger) build() {
	l.setSyncers()
	var err error
	l.Logger, err = l.zapConfig.Build(l.cores())
	if err != nil {
		panic(err)
	}
	defer l.Logger.Sync()
}

func (l *ZapLogger) setSyncers() {
	f := func(fN string) zapcore.WriteSyncer {
		if fN == "" {
			fN = "log.log"
		}
		if l.Opts.SelfConfig.MaxSize == 0 {
			l.Opts.SelfConfig.MaxSize = 100
		}
		if l.Opts.SelfConfig.MaxBackups == 0 {
			l.Opts.SelfConfig.MaxBackups = 60
		}
		if l.Opts.SelfConfig.MaxAge == 0 {
			l.Opts.SelfConfig.MaxAge = 30
		}
		return zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.Opts.SelfConfig.LogFileDir + sp + l.Opts.SelfConfig.AppName + "-" + fN,
			MaxSize:    l.Opts.SelfConfig.MaxSize,
			MaxBackups: l.Opts.SelfConfig.MaxBackups,
			MaxAge:     l.Opts.SelfConfig.MaxAge,
			Compress:   true,
			LocalTime:  true,
		})
	}
	errWS = f(l.Opts.SelfConfig.ErrorFileName)
	warnWS = f(l.Opts.SelfConfig.WarnFileName)
	infoWS = f(l.Opts.SelfConfig.InfoFileName)
	debugWS = f(l.Opts.SelfConfig.DebugFileName)
	return
}

func (l *ZapLogger) cores() zap.Option {
	fileEncoder := zapcore.NewJSONEncoder(l.zapConfig.EncoderConfig)
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	errPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel && zapcore.ErrorLevel-l.zapConfig.Level.Level() > -1
	})
	warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel && zapcore.WarnLevel-l.zapConfig.Level.Level() > -1
	})
	infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel && zapcore.InfoLevel-l.zapConfig.Level.Level() > -1
	})
	debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel && zapcore.DebugLevel-l.zapConfig.Level.Level() > -1
	})
	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, errWS, errPriority),
		zapcore.NewCore(fileEncoder, warnWS, warnPriority),
		zapcore.NewCore(fileEncoder, infoWS, infoPriority),
		zapcore.NewCore(fileEncoder, debugWS, debugPriority),
	}
	if l.Opts.Development {
		cores = append(cores, []zapcore.Core{
			zapcore.NewCore(consoleEncoder, errorConsoleWS, errPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWS, warnPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWS, infoPriority),
			zapcore.NewCore(consoleEncoder, debugConsoleWS, debugPriority),
		}...)
	}
	return zap.WrapCore(func(c zapcore.Core) zapcore.Core {
		return zapcore.NewTee(cores...)
	})
}
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
func timeUnixNano(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.UnixNano() / 1e6)
}
