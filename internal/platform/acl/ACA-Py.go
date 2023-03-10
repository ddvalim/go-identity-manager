package acl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ACAPyResponse struct {
}

type ACAPyACL struct {
	RequestURL string
	Method     string
}

func NewACAPyACL(requestURL string, method string) *ACAPyACL {
	return &ACAPyACL{
		RequestURL: requestURL,
		Method:     method,
	}
}

func (a ACAPyACL) Create() (*ACAPyResponse, error) {
	requestURL := a.RequestURL

	req, err := http.NewRequest(a.Method, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}

	var ACA ACAPyResponse

	err = json.Unmarshal(resBody, &ACA)
	if err != nil {
		fmt.Printf("client: could not unmarshall response body: %s\n", err)
		return nil, err
	}

	fmt.Printf("client: response body: %s\n", resBody)

	return &ACA, err
}
