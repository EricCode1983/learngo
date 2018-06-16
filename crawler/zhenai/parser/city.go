package parser

import (
	"regexp"
	"learngo/crawler/engine"
)
const cityRe=`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParserResult{
	//fmt.Printf("%s", contents)
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, m := range matches {
		name:=string(m[2])
		result.Items = append(result.Items, "User "+name)
		result.Request = append(result.Request, engine.Request{Url: string(m[1]),
		ParserFunc: func(c []byte)engine.ParserResult{
			return ParseProfile(c,name)
		},
		})
	}
	return result
}
