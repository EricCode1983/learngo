package main

import (
	"learngo/retriever/realPkg"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}


func download(r Retriever) string{
	return r.Get("http://www.imooc.com")
}

func main() {

  var r Retriever
  r= realPkg.Retriever{}
  fmt.Printf(download(r))
}