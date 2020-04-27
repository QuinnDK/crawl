package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//resp,err:=http.Get("https://book.douban.com/top250?start=0")  豆瓣已经加入的发爬虫机制，这种方法不能用了
	res, err := http.NewRequest("GET", "https://book.douban.com", nil)

	if err != nil {
		panic(err)
	}
	res.Header.Add("User-Agent", "test")

	client := &http.Client{}
	resp, err2 := client.Do(res)

	if err2 != nil {
		err = err2
		return
	}
	defer resp.Body.Close() //延迟关闭

	if resp.StatusCode != http.StatusOK { //statusOK 值为两百，像404等等为服务器异常
		fmt.Printf("Error status code : %d", resp.StatusCode)

	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", result)
}
