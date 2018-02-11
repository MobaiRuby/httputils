package httputils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpConPool struct {
	Conn *http.Client
}

// Get a http pool for http client
func GetHttpPool(max_conn int, duration int64) *HttpConPool {
	hpool := new(HttpConPool)
	hpool.Conn = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: max_conn,
		},
		Timeout: time.Duration(duration) * time.Millisecond,
	}

	return hpool
}

// send a http request of post or get
func (h *HttpConPool) Request(url string, method string, data string, header map[string]string) (interface{}, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	for h, v := range header {
		req.Header.Set(h, v)
	}

	response, err := h.Conn.Do(req)

	if err != nil {
		return nil, err
	} else if response != nil {
		defer response.Body.Close()

		r_body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		} else {
			return string(r_body), nil
		}
	} else {
		return nil, nil
	}
}
