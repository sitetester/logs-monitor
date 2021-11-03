package logs

type Parser struct{}

type ParseResult struct {
	TotalRequests int
}

func (p Parser) Parse(logs []string) ParseResult {
	return ParseResult{
		TotalRequests: len(logs),
	}
}
