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
	Call(ctx context.Context, action string, req *bytes.Buffer) (*entity.SoapReponse, error)
}

type wsHandler struct {
	wsAddr string
}

func NewWSHandler(address string) WSHandler {
	return &wsHandler{
		wsAddr: address,
	}
}

// Call implements WSHandler
func (w *wsHandler) Call(ctx context.Context, action string, req *bytes.Buffer) (*entity.SoapReponse, error) {
	resp, err := http.NewRequest(http.MethodPost, w.wsAddr, req)
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
		Timeout:   3600,
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
