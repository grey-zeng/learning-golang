package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println(check(r.Form["username"], "string"))
	}
}

func check(val []string, t string) bool {
	switch t {
	// 检查是否数字
	case "number":
		_, err := strconv.Atoi(val[0])
		if err != nil {
			return false
		}
		return true
	// 检查是否字符串
	case "string":
		return true
	// 检查邮箱
	case "email":
		m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, val[0])
		return m
	// 备选项
	case "select":
		slice := []string{"apple", "pear", "banana"}
		for _, item := range slice {
			if val[0] == item {
				return true
			}
		}
		return false
	// 默认都是返回false
	default:
		return false
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) // 设置访问路由
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
