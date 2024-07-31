package http

import (
	"bytes"
	"context"
	"encoding/json"
	l "service-hf-orch-p5/external/logger"
	ps "service-hf-orch-p5/external/strings"
	"service-hf-orch-p5/internal/core/domain/entity/dto"
	httpi "service-hf-orch-p5/internal/core/domain/http"
	"io"
	"net/http"
	"strings"
	"time"
)

var _ httpi.PaymentAPI = (*mercadoPagoAPI)(nil)

type clientHTTPRequest interface {
	Do(req *http.Request) (*http.Response, error)
}

var httpClient clientHTTPRequest = &http.Client{}

type mercadoPagoAPI struct {
	URL     string
	Headers map[string]string
	Timeout time.Duration
}

func NewMercadoPagoAPI(url string, headers map[string]string, timeout time.Duration) *mercadoPagoAPI {
	return &mercadoPagoAPI{URL: url, Headers: headers, Timeout: timeout}
}

func (m *mercadoPagoAPI) DoPayment(ctx context.Context, input dto.InputPaymentAPI) (*dto.OutputPaymentAPI, error) {

	ctxReq, cancel := context.WithTimeout(ctx, m.Timeout)
	defer cancel()

	l.Infof("DoPayment received input: ", " | ", ps.MarshalString(input))

	var buff bytes.Buffer

	if _, err := buff.ReadFrom(strings.NewReader(ps.MarshalString(input))); err != nil {
		l.Errorf("DoPayment error: ", " | ", err)
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctxReq, http.MethodPost, m.URL, &buff)

	if err != nil {
		l.Errorf("DoPayment error: ", " | ", err)
		return nil, err
	}

	for k, v := range m.Headers {
		req.Header.Add(k, v)
	}

	resp, err := httpClient.Do(req)

	if err != nil {
		l.Errorf("DoPayment error: ", " | ", err)
		return nil, err
	}

	defer resp.Body.Close()

	rBody, err := io.ReadAll(resp.Body)

	if err != nil {
		l.Errorf("DoPayment error: ", " | ", err)
		return nil, err
	}

	cleanSpaces(&rBody)

	l.Infof("received httpAdapter response: ", " | ", string(rBody))

	var out = new(dto.OutputPaymentAPI)

	out.HTTPStatus = resp.StatusCode

	if err := json.Unmarshal(rBody, out); err != nil {
		l.Errorf("DoPayment error: ", " | ", err)
		return nil, err
	}

	l.Infof("DoPayment output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func cleanSpaces(b *[]byte) {
	*b = bytes.ReplaceAll(*b, []byte("  "), []byte(" "))
	*b = bytes.ReplaceAll(*b, []byte("\t"), []byte(""))
	*b = bytes.ReplaceAll(*b, []byte("\n"), []byte(""))
}
