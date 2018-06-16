package main

import (
	"regexp"
	"fmt"
)

const text=`
my email is ccmouse@gmail.com  email is 
eric@abc.com
email2 is kkk@qq.com
email3 is ddd@abc.com.cn
`
func main() {

	//.+  .*
	//` 不会转换所给的字符串 (）子匹配 [] 特殊字符不需要\ 外面需要\.
	re:=regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)([.a-zA-Z0-9.]+)`)
	match:=re.FindAllStringSubmatch(text,-1)
	for _,m:=range match{
		fmt.Println(m)
	}

}
