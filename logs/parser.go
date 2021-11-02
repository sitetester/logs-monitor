package logs

import (
	"regexp"
)

type Parser struct{}

type ParseResult struct {
	TotalUniqueIps int
	TotalRequests  int
}

func (p Parser) Parse(logs []string) ParseResult {
	return ParseResult{
		TotalUniqueIps: len(p.uniqueIps(logs)),
		TotalRequests:  len(logs),
	}
}

func (p Parser) uniqueIps(logs []string) map[string]bool {
	var ips = make(map[string]bool)
	var r = regexp.MustCompile(`([:0-9])+`)
	for _, log := range logs {
		ip := r.FindString(log)
		_, ok := ips[ip]
		if !ok {
			ips[ip] = true
		}
	}

	return ips
}

func getCountByStatusCode(logs []string, statusCode int) int {
	countByStatusCode := 0
	var r = regexp.MustCompile(`\s(\d{3})\s(\d{3})`)
	for _, log := range logs {
		matches := r.FindStringSubmatch(log)
		if len(matches) > 0 {
			statusCodeInt := strToInt(matches[1])
			if statusCodeInt == statusCode {
				countByStatusCode += 1
			}
		}
	}

	return countByStatusCode
}
