package testing

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RAMESSESII2/go-ledger/server/services"
)

//testing handler services.SayHello()
func TestSayHello(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:9000/hello", nil)
	if err != nil {
		t.Fatal("error in 'http.NewRequest'" + err.Error())
	}
	//NewRecorder returns an initialized ResposeRecorder i.e., the w http.ResponseWriter object which is an interface
	rr := httptest.NewRecorder()
	services.SayHello(rr, req)
	res := rr.Result()
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("couldn't read the respose: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 'OK'; got %v", res.StatusCode)
	}
	if string(b) != "hi!" {
		t.Fatalf("Expected respose 'hi!'; found %v", string(b))
	}
}

func TestGetLedger(t *testing.T) {
}
