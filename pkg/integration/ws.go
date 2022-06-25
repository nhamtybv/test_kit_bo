package integration

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

type WSHandler interface {
	Call(ctx context.Context, service_url string, req *bytes.Buffer) (*entity.SoapReponse, error)
}

type wsHandler struct {
}

func NewWSHandler() WSHandler {
	return &wsHandler{}
}

// Call implements WSHandler
func (w *wsHandler) Call(ctx context.Context, service_url string, req *bytes.Buffer) (*entity.SoapReponse, error) {
	resp, err := http.NewRequest(http.MethodPost, service_url, req)
	if err != nil {
		return nil, fmt.Errorf("preparing request error: %w", err)
	}

	transport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   3600 * time.Second,
			KeepAlive: 3600 * time.Second,
		}).Dial,
		// We use ABSURDLY large keys, and should probably not.
		TLSHandshakeTimeout: 3600 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{
		Transport: transport,
	}

	httpResp, err := httpClient.Do(resp)
	if err != nil {
		return nil, fmt.Errorf("calling http error: %w", err)
	}

	httpBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing response error: %w", err)
	}
	defer httpResp.Body.Close()

	soapResp := &entity.SoapReponse{}

	err = xml.Unmarshal(httpBody, soapResp)
	if err != nil {
		return nil, fmt.Errorf("parsing xml error: %w", err)
	}

	return soapResp, nil

}
