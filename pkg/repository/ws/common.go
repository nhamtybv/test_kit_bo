package ws

import (
	"bytes"
	"encoding/xml"
	"fmt"

	"github.com/nhamtybv/test_kit_bo/pkg/entity"
)

func Otob(msg entity.SoapEnvelope) (*bytes.Buffer, error) {
	out, err := xml.MarshalIndent(&msg, " ", "  ")
	if err != nil {
		return nil, fmt.Errorf("service: marshal xml error >> %w", err)
	}

	buf := string(out)

	return bytes.NewBufferString(buf), nil
}
