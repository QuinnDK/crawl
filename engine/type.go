package engine

type ParseResult struct {
	requesrts []Request
	Items     []interface{}
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}
