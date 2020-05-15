package engine

type ParseResult struct {
	Requesrts []Request //返回数组
	Items     []Item    //可能会获取的内容
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult //标识如何解析这个Url地址
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
