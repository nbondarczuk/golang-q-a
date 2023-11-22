package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerJSON2XML(t *testing.T) {
	str := `{"id":1,"name":"xxx","email":"yyy"}`
	reader := strings.NewReader(str)
	req := httptest.NewRequest(http.MethodGet, "/json2xml", reader)
	res := httptest.NewRecorder()

	json2xml(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusOK)
	}
}

func TestHandlerXML2JSON(t *testing.T) {
	str := `<User><ID>1</ID><Name>xxx</Name><Email>yyy</Email></User>`
	reader := strings.NewReader(str)
	req := httptest.NewRequest(http.MethodGet, "/xml2json", reader)
	res := httptest.NewRecorder()

	xml2json(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusOK)
	}

}
