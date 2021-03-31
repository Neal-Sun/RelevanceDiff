package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ProductIn struct {
	SearchKeyWord string `json:"searchKeyword"`
}

type TestSearchResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

var url = "http://localhost:9998/product/search"

func main() {
	var productIn []ProductIn

	// get search query
	byteValue, err := ioutil.ReadFile("./query3.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 去掉utf-BOM的编码格式的前缀
	byteValue = bytes.TrimPrefix(byteValue, []byte("\xef\xbb\xbf"))
	err = json.Unmarshal([]byte(byteValue), &productIn)
	if err != nil {
		fmt.Println(err)
		fmt.Println("get search query failed")
	}

	//fmt.Println("-------query text--------")
	//for _,product:= range productIn {
	//      fmt.Println(product)
	//}

	httpClient := &http.Client{}
	for _, product := range productIn {
		request := fmt.Sprintf(`{"query": "%s","size": %d,"from": %d, "business_type": "%s"}`, product.SearchKeyWord, 10, 0, "query_diff")
		// request := fmt.Sprintf(`{"query": "%s","size": %d,"from": %d}`, product.SearchKeyWord, 10, 0)

		var jsonStr = []byte(request)
		var reqBuffer = bytes.NewBuffer(jsonStr)
		//httpResp, err := httpClient.Post(url, "application/json", reqBuffer)
		_, err := httpClient.Post(url, "application/json", reqBuffer)
		if err != nil {
			fmt.Println(err)
		}

		//fmt.Println("-------response--------")
		//bytes, err := ioutil.ReadAll(httpResp.Body)
		//fmt.Println(string(bytes))
		//response := &TestSearchResponse{}
		//err = json.Unmarshal(bytes, &response)
		//fmt.Println(response.Message)
	}
}
