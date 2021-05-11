package gee

import "net/http"

type H map[string]interface{}

type Context struct{
	//origin objects
	Writer http.ResponseWriter
	Req *http.Request
	//request info
	Path string
	Method string
	//response info
	StatusCode int
}

func newContext(w http.ResponseWriter,req *http.Request){

}