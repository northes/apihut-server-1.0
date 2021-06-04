package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	ErrRequest = errors.New("请求失败")
)

var client = http.Client{
	Timeout: 10 * time.Second,
}

func HttpGetRequest(url string, result interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)

	return err
}
