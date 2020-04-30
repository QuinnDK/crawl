package engine

type ParseResult struct {
	Requesrts []Request     //返回数组
	Items     []interface{} //可能会获取的内容
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult //标识如何解析这个Url地址
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
