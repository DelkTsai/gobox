package http

import (
	"andals/gobox/log"
	logWriter "andals/gobox/log/writer"
	"fmt"
	"testing"
	"time"
)

func TestClientGet(t *testing.T) {
	path := "/tmp/test_http_client.log"
	w, _ := logWriter.NewFileWriter(path)
	logger, _ := log.NewSimpleLogger(w, log.LEVEL_INFO)

	client := NewClient(logger).SetTimeout(time.Second * 3)
	extHeaders := map[string]string{
		"GO-CLIENT-1": "gobox-httpclient-1",
		"GO-CLIENT-2": "gobox-httpclient-2",
	}
	req, _ := NewRequestForGet("http://www.vdocker.com/test.php", "127.0.0.1", extHeaders)

	cb, err := client.Do(req, 1)
	fmt.Println(string(cb), err)
}
