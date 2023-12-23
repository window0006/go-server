package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	client *http.Client
	// ctx    context.Context
}

func NewHttpClient() *HttpClient {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return &HttpClient{
		client: client,
	}
}

func (hc *HttpClient) SendPostRequest(url string, data []byte, headers map[string]interface{}) (respJson map[string]interface{}, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	hc.addHeaders(req, headers)
	return hc.sendRequest(req)
}

func (hc *HttpClient) SendGetRequest(url string, params map[string]interface{}, headers map[string]interface{}) (respJson map[string]interface{}, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v.(string))
	}
	// 将参数编码到 url 中
	req.URL.RawQuery = q.Encode()
	hc.addHeaders(req, headers)
	return hc.sendRequest(req)
}

func (hc *HttpClient) addHeaders(req *http.Request, headers map[string]interface{}) {
	for k, v := range headers {
		req.Header.Set(k, v.(string))
	}
}

func (hc *HttpClient) sendRequest(req *http.Request) (respJson map[string]interface{}, err error) {
	resp, err := hc.client.Do(req)
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respBytes, &respJson)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
