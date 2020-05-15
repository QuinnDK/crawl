package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//原生方式
//func Fetch(url string) ([]byte, error) {
//	//resp,err:=http.Get("https://book.douban.com/top250?start=0")  豆瓣已经加入的发爬虫机制，这种方法不能用了
//	res, err := http.NewRequest("GET", "https://book.douban.com", nil)
//
//	if err != nil {
//		panic(err)
//	}
//	res.Header.Add("User-Agent", "test")
//
//	client := &http.Client{}
//	resp, err2 := client.Do(res)
//
//	if err2 != nil {
//		err = err2
//		return nil, nil
//	}
//	defer resp.Body.Close() //延迟关闭
//
//	if resp.StatusCode != http.StatusOK { //statusOK 值为两百，像404等等为服务器异常
//		fmt.Printf("Error status code : %d", resp.StatusCode)
//
//	}
//	return ioutil.ReadAll(resp.Body)
//
//}

var ratelimit = time.Tick(20 * time.Millisecond)

//模拟浏览器模式
func Fetch(url string) ([]byte, error) {
	//resp,err:=http.Get("https://book.douban.com/top250?start=0")  豆瓣已经加入的发爬虫机制，这种方法不能用了

	//client := &http.Client{}
	//req,err:=http.NewRequest("GET",url,nil)
	//if err != nil{
	//	return nil,fmt.Errorf("ERROR: get url:%s",url)
	//}
	//req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36")
	//
	//resp,err:= client.Do(req)

	//似乎我之前用的就是浏览器模式，但是豆瓣的翻爬虫机制似乎没有解决
	<-ratelimit
	//
	//proxy:= func(_*http.Request) (*url.URL,error){
	//	return url.Parse("http://127.0.0.1:1087")
	//}
	//Transport:=&http.Transport{Proxy:proxy}
	//client:=&http.Client{Transport:Transport}

	res, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}
	//res.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	res.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36")

	client := &http.Client{}
	resp, err2 := client.Do(res)

	if err2 != nil {
		err = err2
		return nil, nil
	}
	defer resp.Body.Close() //延迟关闭

	if resp.StatusCode != http.StatusOK { //statusOK 值为两百，像404等等为服务器异常
		fmt.Printf("Error status code : %d", resp.StatusCode)

	}
	return ioutil.ReadAll(resp.Body)

}

//func ParseContent(content []byte) engine.ParseResult{
//	//<a href="/tag/哲学" class="tag">哲学</a>
//	//re:=regexp.MustCompile(`<a href="([^"]+)">([^</a>]+)</a>`)
//	re := regexp.MustCompile(`<a href="([^"]+)" class="tag">([^"]+)</a>`) //外边的小点 ` 并非是引号，我靠，花了好长时间
//	matches := re.FindAllSubmatch(content, -1)
//
//	for _, m := range matches {
//		fmt.Printf("url:%s\n", "https://book.douban.com"+string(m[1]))
//
//	}
