package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp,err:=http.Get("https://book.douban.com/")

	if err !=nil{
		panic(err)
	}
	defer resp.Body.Close()      //延迟关闭

	if resp.StatusCode !=http.StatusOK {//statusOK 值为两百，像404等等为服务器异常
		fmt.Printf("Error status code : %d",resp.StatusCode)

	}
	result,err:=ioutil.ReadAll(resp.Body)
	if err !=nil{
		panic(err)
	}
	fmt.Printf("%s",result)
}
