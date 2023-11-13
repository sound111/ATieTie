package logger

import (
	"TieTie/settings"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init() (err error) {
	writeSyncer := getLogWriter(
		settings.Conf.LogConfig.FileName,
		settings.Conf.LogConfig.MaxSize,
		settings.Conf.LogConfig.MaxAge,
		settings.Conf.LogConfig.MaxBackups,
		settings.Conf.LogConfig.Compress)
	if err != nil {
		return
	}
	encoder := getEncoder()

	var core zapcore.Core
	if settings.Conf.AppConfig.Mode == "debug" {
		consoleEncoding := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(consoleEncoding, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
			zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	}

	lg := zap.New(core, zap.AddCaller())
	// 替换zap库中全局的logger zap.L()
	zap.ReplaceGlobals(lg)
	return
}

func getLogWriter(filename string, maxsize, maxage, maxbackups int, compress bool) zapcore.WriteSyncer {
	//库不支持按时间切割日志文件
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxsize,    //单位MB
		MaxBackups: maxbackups, //最多保存的文件数量
		MaxAge:     maxage,     //文件最多保存30天
		Compress:   compress,   //是否压缩
	}

	return zapcore.AddSync(lumberjackLogger)
}

func getEncoder() zapcore.Encoder {
	conf := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(conf)
	//企业大多采用json格式
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("myError", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
// stack参数表示是否记录堆栈信息
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

//func HTTPGet(url string) (err error) {
//	resp, err := http.Get(url)
//
//	if err != nil {
//		Logger.Error(
//			"error fetching url ",
//			zap.Error(err),
//			zap.String("url", url))
//	} else {
//		Logger.Info(
//			"success ",
//			zap.String("statusCode", resp.Status),
//			zap.String("url", url))
//
//		err = resp.Body.Close()
//	}
//
//	return
//}
