package logs

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Reader struct{}

func (r *Reader) Read(filePath string, previousLineNum int) ([]string, int) {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)
	currentLineNum := 0

	var re = regexp.MustCompile(`([0-9]{1,2})/([a-zA-Z]{3})/(\d{4}):(\d{2}):(\d{2}):(\d{2})`)

	for scanner.Scan() {
		currentLineNum += 1
		line := scanner.Text()

		if currentLineNum > previousLineNum {
			if isWithin10Sec(parseDateTime(line, re)) {
				lines = append(lines, line)
			}
		}
	}

	return lines, currentLineNum
}

func getLocalTime() time.Time {
	localTimeNow := time.Now()
	localTime := time.Date(
		localTimeNow.Year(), localTimeNow.Month(), localTimeNow.Day(), localTimeNow.Hour(), localTimeNow.Minute(), localTimeNow.Second(), 0, time.UTC,
	)

	return localTime
}

func isWithin10Sec(oldTime time.Time) bool {
	return int(getLocalTime().Sub(oldTime).Seconds()) <= 10
}

func parseDateTime(line string, re *regexp.Regexp) time.Time {
	matches := re.FindStringSubmatch(line)
	return time.Date(
		strToInt(matches[3]), time.Month(monthNameToNum(matches[2])), strToInt(matches[1]),
		strToInt(matches[4]), strToInt(matches[5]), strToInt(matches[6]), 0, time.UTC,
	)
}

func monthNameToNum(name string) int {
	m := make(map[string]int)
	m["Jan"] = 1
	m["Feb"] = 2
	m["Mar"] = 3
	m["Apr"] = 4
	m["May"] = 5
	m["June"] = 6
	m["Jul"] = 7
	m["Aug"] = 8
	m["Sep"] = 9
	m["Oct"] = 10
	m["Nov"] = 11
	m["Dec"] = 12

	return m[name]
}

func strToInt(str string) int {
	intVal, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}

	return intVal
}
