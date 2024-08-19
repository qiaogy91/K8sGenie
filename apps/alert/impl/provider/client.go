package provider

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

var _ Service = &HttpClient{}

type HttpClient struct {
}

func (h *HttpClient) SendCard(ctx context.Context, url string, data []byte) []byte {
	// req
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")

	// send
	rsp, err := http.DefaultClient.Do(req)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		fmt.Printf("http default client send req failed: %v", err)
	}

	// rsp
	content, _ := io.ReadAll(rsp.Body)
	return content
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}
