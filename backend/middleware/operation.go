package middleware

import (
	"bytes"
	"encoding/json"
	"gin-pro/global"
	"gin-pro/model/system"
	"gin-pro/service"
	"gin-pro/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

var bufferSize = 1024

func OperationRecordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var body []byte
		var userId int
		if c.Request.Method != http.MethodGet {
			var err error
			//读取 HTTP 请求的 body 内容
			body, err = io.ReadAll(c.Request.Body)
			if err != nil {
				global.GVA_LOG.Error("read body from request error:", zap.Error(err))
			} else {
				//将读取的请求体内容重新赋值给 c.Request.Body，以便后续的中间件或处理函数可以再次读取这个请求体
				c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
			}
		} else {
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			arr := strings.Split(query, "&")
			m := make(map[string]string)
			for _, v := range arr {
				item := strings.Split(v, "=")
				m[item[0]] = item[1]
			}
			//数据结构（如结构体、切片、映射等）编码为 JSON 格式的字节切片
			body, _ = json.Marshal(&m)
		}
		claims, _ := utils.GetClaims(c)
		if claims == nil && claims.BaseClaims.Id != 0 {
			userId = int(claims.BaseClaims.Id)
		} else {
			id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
			if err != nil {
				userId = 0
			}
			userId = id
		}

		record := system.SysOperationRecord{
			Ip:     c.ClientIP(),
			Method: c.Request.Method,
			Path:   c.Request.URL.Path,
			Agent:  c.Request.UserAgent(),
			Body:   "",
			UserID: userId,
		}
		if strings.Contains(c.GetHeader("Content-Type"), "multipart/form-data") {
			record.Body = "[文件]"
		} else {
			if len(body) > bufferSize {
				record.Body = "[超出记录]"
			} else {
				record.Body = string(body)
			}
		}
		//重写writer
		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()
		c.Next()
		latency := time.Since(now)
		record.Latency = latency
		record.Status = c.Writer.Status()
		record.Resp = writer.body.String()
		//私有错误类型信息
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		//当前的响应可能是一个较大的文件下载或不应被缓存的内容。
		if strings.Contains(c.Writer.Header().Get("Pragma"), "public") ||
			strings.Contains(c.Writer.Header().Get("Expires"), "0") ||
			strings.Contains(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/force-download") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
			strings.Contains(c.Writer.Header().Get("Content-Type"), "application/download") ||
			strings.Contains(c.Writer.Header().Get("Content-Disposition"), "attachment") ||
			strings.Contains(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
			if len(record.Resp) > bufferSize {
				record.Body = "私有错误超出记录长度"
			}
		}
		var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
		if err := operationRecordService.CreateOperationRecord(record); err != nil {
			global.GVA_LOG.Error("create operation record error:", zap.Error(err))
		}

	}
}
