package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/RAMESSESII2/go-ledger/server/models"
)

// reprents the HTTP client which communicates with the server APIs
type HTTPClient struct {
	client     *http.Client
	BackendURI string
}

//to create a new instance of  HTTPClient
func NewHTTPClient(uri string) HTTPClient {
	return HTTPClient{
		client:     &http.Client{},
		BackendURI: uri,
	}
}

//create() call
func (c HTTPClient) Create(fname, lname string, debit, credit int64) ([]byte, error) {
	requestBody := models.Transaction{
		FirstName: fname,
		LastName:  lname,
		DebitAmt:  debit,
		CreditAmt: credit,
	}
	return c.apiCall(http.MethodPost, "/ledger", &requestBody, http.StatusOK)
}

//edit() call
func (c HTTPClient) Edit(id int, fname, lname string, debit, credit int64) ([]byte, error) {
	requestBody := models.Transaction{
		FirstName: fname,
		LastName:  lname,
		DebitAmt:  debit,
		CreditAmt: credit,
	}
	return c.apiCall(http.MethodPatch, "/ledger/"+fmt.Sprint(id), &requestBody, http.StatusOK)
}

//fetch() call
func (c HTTPClient) Fetch(id int) ([]byte, error) {
	return c.apiCall(http.MethodGet, "/ledger/"+fmt.Sprint(id), nil, http.StatusOK)

}

//delete() call
func (c HTTPClient) Delete(id int) error {
	_, err := c.apiCall(http.MethodDelete, "/ledger/"+fmt.Sprint(id), nil, http.StatusOK)
	return err
}

// generic calling structure
func (c HTTPClient) apiCall(method, path string, body interface{}, resCode int) ([]byte, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		e := wrapError("Couldn't marshall request body", err)
		return nil, e
	}

	req, err := http.NewRequest(method, c.BackendURI+path, bytes.NewReader(bs))
	if err != nil {
		e := wrapError("Couldn't create request", err)
		return []byte{}, e
	}
	res, err := c.client.Do(req)
	if err != nil {
		e := wrapError("Couldn't make http call", err)
		return []byte{}, e
	}
	resBody, err := c.readResposeBody(res.Body)
	if err != nil {
		return []byte{}, err
	}
	if res.StatusCode != resCode {
		if len(resBody) > 0 {
			fmt.Printf("got this response body: \n%s\n", resBody)
		}
		return []byte{}, fmt.Errorf("exptected respose code: %d, got: %d", resCode, res.StatusCode)
	}
	return []byte(resBody), err
}

//for reading response body
func (c HTTPClient) readResposeBody(b io.Reader) (string, error) {
	bs, err := ioutil.ReadAll(b)
	if err != nil {
		return "", wrapError("Couldn't read respose body", err)
	}
	if len(bs) == 0 {
		return "", nil
	}
	var buff bytes.Buffer
	if err := json.Indent(&buff, bs, "", "\t"); err != nil {
		return "", wrapError("Couldn't indent json", err)
	}
	return buff.String(), nil
}
