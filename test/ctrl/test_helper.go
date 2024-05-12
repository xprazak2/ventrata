package ctrl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/xprazak2/ventrata/internal/controller"
)

func Request(method string, endpoint string, body interface{}, respBody interface{}, priced bool) error {
	var reqBody *bytes.Buffer = &bytes.Buffer{}

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		reqBody = bytes.NewBuffer(bodyBytes)
	}

	reqUrl := "http://localhost:3000" + endpoint

	req, err := http.NewRequest(method, reqUrl, reqBody)
	if err != nil {
		return err
	}

	if priced {
		req.Header.Set(controller.CapbHeader, controller.Pricing)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		rBody, _ := io.ReadAll(resp.Body)
		strRespBody := string(rBody)

		return fmt.Errorf("failed to decode response, got: %s", strRespBody)
	}

	return nil
}
