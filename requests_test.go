package requests

import (
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://httpbin.org/get"
	headers := Headers{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.146 Safari/537.36",
	}
	params := Params{
		"key": "value",
	}
	response, err := Get(url, params, headers)
	if err != nil{
		t.Log(err)
	}else{
		t.Log(response.Text)
	}
}

func TestPostForm(t *testing.T) {
	url := "https://httpbin.org/Post"
	headers := Headers{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.146 Safari/537.36",
	}
	params := Data{
		"key": "value",
	}
	response, err := Get(url, params, headers)
	if err != nil{
		t.Log(err)
	}else{
		t.Log(response.Text)
	}
}

func TestPostJson(t *testing.T) {
	url := "https://httpbin.org/Post"
	headers := Headers{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.146 Safari/537.36",
	}
	params := Json{
		"key": "value",
	}
	response, err := Get(url, params, headers)
	if err != nil{
		t.Log(err)
	}else{
		t.Log(response.Text)
	}
}