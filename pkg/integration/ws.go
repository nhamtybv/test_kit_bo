package integration

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
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
	logText := ">> calling web service at %s\n"
	logText += "================================================================"
	logText += "\n%s\n"
	logText += "================================================================"

	log.Printf(logText, service_url, req.String())

	resp, err := http.NewRequest(http.MethodPost, service_url, req)
	if err != nil {
		return nil, fmt.Errorf(">> preparing request\n%w", err)
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
		return nil, fmt.Errorf(">> calling http %w", err)
	}

	httpBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf(">> service: parsing response %w", err)
	}
	defer httpResp.Body.Close()

	soapResp := &entity.SoapReponse{}

	err = xml.Unmarshal(httpBody, soapResp)
	if err != nil {
		return nil, fmt.Errorf(">> service: parsing xml %w", err)
	}

	f := soapResp.Body.Fault
	if f != nil {
		return nil, fmt.Errorf(">> service: parsing response %+v", f)
	}

	return soapResp, nil

}
