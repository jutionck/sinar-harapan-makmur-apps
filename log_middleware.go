package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	// Open file
	filePath := "/Users/jutioncandrakirana/Documents/GitHub/enigma/GOLANG/go-sinar-harapan-makmur-api/LOG.txt"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logger := log.New(file, "", 0)

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now().Sub(startTime)

		entryLog := EntryLog{
			StartTime:    startTime,
			EndTime:      endTime,
			StatusCode:   c.Writer.Status(),
			ClientAIP:    c.ClientIP(),
			Method:       c.Request.Method,
			RelativePath: c.Request.URL.Path,
			UserAgent:    c.Request.UserAgent(),
		}

		entryLogString := fmt.Sprintf("[LOG] : %v\n", entryLog)
		logger.Println(entryLogString)
	}
}
