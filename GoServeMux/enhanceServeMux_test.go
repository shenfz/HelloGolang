package GoServeMux__test

import (
	"fmt"
	"net/http"
	"testing"
)

/**
* @Author shenfz
* @Date 2024/5/16 15:46
* @Email 1328919715@qq.com
* @Description: go1.22 新版路由
            不同路由设置之间存在交集，这就需要路由匹配优先级规则
            增强后的ServeMux可能会影响向后兼容性，使用GODEBUG=httpmuxgo121=1可以保留原先的ServeMux行为。
**/
/*

 $curl localhost:8080/index.html
match /index.html

$curl example.com:8080/static/abc
match "example.com/"

$curl localhost:8080/static/abc
match "GET /static/"

$curl example.com:8080/
match "example.com/{$}"

$curl example.com:8080/b/mybucket/o/myobject/tonybai
match "example.com/"

$curl localhost:8080/b/mybucket/o/myobject/tonybai
match /b/{bucket}/o/{objectname...}:bucket=mybucket,objectname=myobject/tonybai

*/
func Test_ServeMux(t *testing.T) {
	mux := http.NewServeMux()
	// "/index.html"路由将匹配任何主机和方法的路径"/index.html"；
	mux.HandleFunc("/index.html", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match /index.html`)
	})
	// "GET /static/"将匹配路径以"/static/"开头的GET请求；
	mux.HandleFunc("GET /static/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match "GET /static/"`)
	})
	// "example.com/"可以与任何指向主机为"example.com"的请求匹配；
	mux.HandleFunc("example.com/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match "example.com/"`)
	})
	// "example.com/{$}"会匹配主机为"example.com"、路径为"/"的请求，即"example.com/"；
	mux.HandleFunc("example.com/{$}", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, `match "example.com/{$}"`)
	})

	// "/b/{bucket}/o/{objectname...}"匹配第一段为"b"、第三段为"o"的路径。名称"bucket"表示第二段，"objectname"表示路径的其余部分。
	mux.HandleFunc("/b/{bucket}/o/{objectname...}", func(w http.ResponseWriter, req *http.Request) {
		bucket := req.PathValue("bucket")
		objectname := req.PathValue("objectname")
		fmt.Fprintln(w, `match /b/{bucket}/o/{objectname...}`+":"+"bucket="+bucket+",objectname="+objectname)
	})

	http.ListenAndServe(":8080", mux)
}
