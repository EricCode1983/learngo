package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/net/html/charset"
	"regexp"
)

func main() {
	resp,err:=http.Get("http://www.zhenai.com/zhenghun")
	if err!=nil{
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK {
		fmt.Println("Error status code", resp.StatusCode)
		return
	}
    //获取网页语言编码
	e:=determineEncoding(resp.Body)
	utf8Reader:=transform.NewReader(resp.Body,e.NewDecoder())
	all,err:=	ioutil.ReadAll(utf8Reader)
	if err!=nil {
		println(err)
	}
	printCityList(all)
	//fmt.Printf("%s\n",all)
}




func determineEncoding(r io.Reader) encoding.Encoding{
    bytes,err:=	bufio.NewReader(r).Peek(1024)
    if err!=nil{
    	panic(err)
	}
	e,_,_:=charset.DetermineEncoding(bytes,"")
	return e

}

func printCityList(contents []byte){
	re:=regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	matchs:=re.FindAllSubmatch(contents,-1)
	for _,m:=range matchs{
		fmt.Printf("City:%s Url:%s ", m[2],m[1])
		fmt.Println()
	}
	fmt.Printf("Matchs found: %d\n",len(matchs))
}