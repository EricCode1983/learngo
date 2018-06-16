package parser

import (
	"learngo/crawler/engine"
	"regexp"
)

const citylistRe=`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte)(engine.ParserResult) {

	re := regexp.MustCompile(citylistRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	//limit:=3
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Request = append(result.Request, engine.Request{Url: string(m[1]), ParserFunc:ParseCity})
		//limit--
		//if limit==0{
		//	break
		//}
	}

	return result
}