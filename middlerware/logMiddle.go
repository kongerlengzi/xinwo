package middlerware
import (
	"encoding/json"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func init() {
	InitLogger()
	defer sugarLogger.Sync()
}

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
	defer sugarLogger.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}


func getLogWriter() zapcore.WriteSyncer {
	timeStr := time.Now().Local().Format("2006-01-02")
	filename := timeStr + ".log"
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./" + filename,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     1,
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

//func simpleHttpGet(url string) {
//	sugarLogger.Debugf("Trying to hit GET request for %s", url)
//	resp, err := http.Get(url)
//	if err != nil {
//		sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
//	} else {
//		sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
//		resp.Body.Close()
//	}
//}

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求头
		reqContentType := c.Request.Header.Get("Content-Type")

		//请求参数
		c.Request.ParseMultipartForm(1024)
		reqFormTmp := c.Request.Form
		reqForm,_ := json.Marshal(reqFormTmp)
		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		sugarLogger.Infof("| %3d | %13v | %15s | %s | %s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
			reqContentType,
			reqForm,
		)
	}
}
