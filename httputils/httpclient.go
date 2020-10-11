package httputils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func SimpleGet(httpUrl string) (interface{}, error) {
	resp, err := http.Get(httpUrl)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return string(body), err
}

func CustomGet(httpUrl string) (interface{}, error) {
	req, err := http.NewRequest("GET", httpUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Golang-test")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return string(body), err
}

func SimpePost(httpUrl string, jsonData string) (string, error) {
	var jsonStr = []byte(jsonData)

	req, err := http.NewRequest("POST", httpUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

func PostForm(httpUrl string) (interface{}, error) {
	formData := url.Values{
		"name": {"masnun"},
	}
	resp, err := http.PostForm(httpUrl, formData)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	log.Println(result["form"])
	return result, nil
}
