package basic

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* ------------------以 JSON 格式输出所有请求头 ------------------ */
func handleHeader(w http.ResponseWriter, req *http.Request) {
	retMap := make(map[string]interface{})
	for k, _ := range req.Header {
		retMap[k] = req.Header.Get(k)
	}
	retMap["Method"] = req.Method     //方法
	retMap["Host"] = req.Host         // 主机
	retMap["Referer"] = req.Referer() // 引用头

	ret, err := json.Marshal(retMap)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

/* ---------------------以 JSON 返回body参数---------------------------- */
func handleBody(w http.ResponseWriter, req *http.Request) {

	retMap := make(map[string]interface{})
	// 判断表单编码
	contentType := req.Header.Get("Content-Type")
	fmt.Println(req.Header.Get("Content-Type"))
	if req.Method == "POST" {
		if contentType == "application/x-www-form-urlencoded" {
			req.ParseForm()
			for k, _ := range req.Form {
				retMap[k] = req.Form.Get(k)
			}

		} else {
			req.ParseMultipartForm(32 << 20)
			if req.MultipartForm.Value != nil {
				for k, _ := range req.MultipartForm.Value {
					retMap[k] = req.MultipartForm.Value[k][0]
				}
			}
		}
	}

	ret, err := json.Marshal(retMap)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

/* ---------------------以 JSON 返回请求Query参数----------------------- */
func handleQuery(w http.ResponseWriter, req *http.Request) {
	retMap := make(map[string]interface{})
	for k, _ := range req.URL.Query() {
		retMap[k] = req.URL.Query().Get(k)
	}

	ret, err := json.Marshal(retMap)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(ret)
}

func Start() {
	http.HandleFunc("/header", handleHeader)
	http.HandleFunc("/body", handleBody)
	http.HandleFunc("/query", handleQuery)

	http.ListenAndServe(":8080", nil)
}
