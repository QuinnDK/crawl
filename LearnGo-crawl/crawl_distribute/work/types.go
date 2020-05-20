package work

import (
	"crawl/LearnGo-crawl/engine"
	"crawl/LearnGo-crawl/parse/zhenai"
	"errors"
	"fmt"
	"log"
)

//type SerializeParer struct {
//	Name string
//	Args interface{}
//}
//
//type Request struct {
//	Url   string
//	Parse SerializeParer
//}
//
//type ParseResult struct {
//	items    []engine.Item
//	Requests []Request
//}
//
//func SerializeResult(r engine.ParseResult) ParseResult {
//	result := ParseResult{items: r.Items}
//
//	for _, req := range r.Requesrts {
//		result.Requests = append(result.Requests, SerializeRequest(req))
//
//	}
//	return result
//}
//
//func SerializeRequest(r engine.Request) Request {
//
//	name, args := r.Parse.Serialize()
//
//	return Request{
//		Url: r.Url,
//		Parse: SerializeParer{
//			Name: name,
//			Args: args,
//		},
//	}
//
//}
//
//func DeserializeResult(r ParseResult) engine.ParseResult {
//
//	result := engine.ParseResult{
//		Items: r.items,
//	}
//	for _, req := range r.Requests {
//		engineReq, err := DeserializeRequest(req)
//		if err != nil {
//			log.Printf("error deserializeing:%V", err)
//			continue
//		}
//		result.Requesrts = append(result.Requesrts, engineReq)
//	}
//	return result
//}
//
//func DeserializeRequest(r Request) (engine.Parser, error) {
//
//	parse, err := deserializeParse(r.Parse)
//	if err != nil {
//		return engine.Request{}, err
//	}
//
//	return engine.Request{
//		Url:   r.Url,
//		Parse: parse,
//	}, nil
//}
//
//func deserializeParse(p SerializeParer) (engine.Parser, error) {
//	switch p.Name {
//	case "ParseCityList":
//		return engine.NewFuncparse(zhenai.ParseCityList, "ParseCityList")
//	case "ParseCity":
//		return engine.NewFuncparse(zhenai.ParseCity, "ParseCity")
//	case "ProfileParse":
//		if useName, ok := p.Args.(string); ok {
//			return zhenai.NewprofileParse(useName), nil
//		} else {
//			return nil, fmt.Errorf("invilid args:%v", p.Args)
//		}
//	case "Nilparse":
//		return engine.Nilparse{}, nil
//	default:
//		return nil, errors.New("unkown parse name")
//
//	}
//}

type SerializeParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url   string
	Parse SerializeParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{Items: r.Items}

	for _, req := range r.Requesrts {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}
func SerializeRequest(r engine.Request) Request {

	name, args := r.Parse.Serialize()

	return Request{
		Url: r.Url,
		Parse: SerializeParser{
			Name: name,
			Args: args,
		},
	}
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			log.Printf("error deserializeing:%v", err)
			continue
		}
		result.Requesrts = append(result.Requesrts, engineReq)
	}

	return result
}
func DeserializeRequest(r Request) (engine.Request, error) {
	parse, err := deserializeParse(r.Parse)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:   r.Url,
		Parse: parse,
	}, nil

}
func deserializeParse(p SerializeParser) (engine.Parser, error) {

	switch p.Name {
	case "ParseCityList":
		return engine.NewFuncparse(zhenai.ParseCityList, "ParseCityList"), nil
	case "Parsecity":
		return engine.NewFuncparse(zhenai.ParseCity, "Parsecity"), nil

	case "ProfileParse":
		if useName, ok := p.Args.(string); ok {
			return zhenai.NewprofileParse(useName), nil
		} else {
			return nil, fmt.Errorf("invilid args:%v", p.Args)
		}
	case "Nilparse":
		return engine.Nilparse{}, nil
	default:
		return nil, errors.New("unkown parse name")

	}
}
