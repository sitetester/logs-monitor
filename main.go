package main

import (
	"fmt"
	"log-monitor/logs"
	"time"
)

func main() {

	filePath := "/usr/local/var/log/httpd/access_log"

	var logsReader logs.Reader
	var logsParser logs.Parser
	var previousLineNum = 0
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case _ = <-ticker.C:
			parsedLogs, lastReadNum := logsReader.Read(filePath, previousLineNum)
			previousLineNum = lastReadNum

			t := time.Now().Format("2006-01-02 15:04:05")
			if len(parsedLogs) > 0 {
				r := logsParser.Parse(parsedLogs)
				fmt.Printf("%s: %+v\n", t, r)
			} else {
				fmt.Printf("%s: %+v\n", t, logs.ParseResult{
					TotalRequests: 0,
				})
			}
		}
	}
}
