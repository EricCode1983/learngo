package main

import (
	"learngo/retriever/realPkg"
	"fmt"
	"time"
)

type Retriever interface {
	Get(url string) string
}


func download(r Retriever) string{
	return r.Get("http://www.imooc.com")
}

func main() {

  var r Retriever
  r= &realPkg.Retriever{UserAgent:"Mozilla/5.0",TimeOut:time.Minute}
  fmt.Printf("%T %v\n",r,r)
  //fmt.Printf(download(r))
}