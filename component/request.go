package component

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"proxy-collect/component/logger"
	"time"
)

func WebRequest(req *http.Request) string {
	client := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Error("http get error", err)
		return ""
	}
	if resp.StatusCode != 200 {
		resp.Body.Close()
		logger.Error("http status error ", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error", err)
		return ""
	}
	return string(body)
}
