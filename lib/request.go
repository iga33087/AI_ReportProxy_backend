package lib

import (
    "fmt"
	"encoding/json"
	"bytes"
	"time"
	"io"
	//"os"
	"net/http"
)

var timeoutSec int = 120

func createRequest(url string,method string,header map[string]string, payload map[string]any) (*http.Request, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("JSON 序列化失敗: %v\n", err)
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("建立請求失敗: %v\n", err)
		return nil, err
	}
	for i,val := range header {
	  req.Header.Set(i, val)
	}
	return req,nil
}

func createClient() *http.Client {
	client := &http.Client{
		Timeout: time.Duration(timeoutSec) * time.Second,
	}
	return client
}

func Call(url string,method string,header map[string]string, payload map[string]any) (any,error) {
	req,_ := createRequest(url,method,header,payload)
	client := createClient()
    resp, err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("讀取回應失敗: %v\n", err)
		return nil,err
	}
    if resp.StatusCode != http.StatusOK { // 或者 resp.StatusCode >= 400
        fmt.Printf("API 發生錯誤，狀態碼：%d\n", resp.StatusCode)
        return nil,fmt.Errorf("API 請求失敗，狀態碼：%d / 錯誤訊息：%d ", resp.StatusCode,string(body))
    }
	return string(body),nil
}