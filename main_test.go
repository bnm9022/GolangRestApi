package main

import (
    "testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"os"
	"bytes"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func TestGetApplicationMD(t *testing.T) {
	req, err := http.NewRequest("GET", "/applicationMDs", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnAllApplicationMDs)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[]`
	if strings.TrimSuffix(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strings.TrimSuffix(rr.Body.String(), "\n"), expected)
	}
}

func TestPostAndDeleteApplicationMD(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/applicationMD", bytes.NewReader([]byte(path + "/testdata/valid2.yaml")))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createNewApplicationMD)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	req, err = http.NewRequest("POST", "/applicationMD", bytes.NewReader([]byte(path + "/testdata/valid1.yaml")))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(createNewApplicationMD)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	req, err = http.NewRequest("DELETE", "/applicationMD?title=Valid%20App%202", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(deleteApplicationMD)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	req, err = http.NewRequest("GET", "/applicationMDs", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(returnAllApplicationMDs)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "App 1") {
		t.Errorf("Handler returned unexpected body: not got App 1 within %v",
		rr.Body.String())
	}
	if strings.Contains(rr.Body.String(), "App 2") {
		t.Errorf("Handler returned unexpected body: got App 2 within %v",
		rr.Body.String())
	}
}

func TestPostAndGetApplicationMD(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/applicationMD", bytes.NewReader([]byte(path + "/testdata/valid1.yaml")))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createNewApplicationMD)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var expectedApplicationMD ApplicationMD
	var responseApplicationMD ApplicationMD
	yamlFile, err := ioutil.ReadFile(path + "/testdata/valid1.yaml")
	err = yaml.Unmarshal(yamlFile, &expectedApplicationMD)
	if err != nil {
		t.Fatal(err)
	}
	err = yaml.Unmarshal([]byte(rr.Body.String()), &responseApplicationMD)
	if err != nil {
		t.Fatal(err)
	}

	if expectedApplicationMD.Title != responseApplicationMD.Title {
		t.Errorf("handler returned unexpected body: got %v want %v",
		expectedApplicationMD.Title, responseApplicationMD.Title)
	}

	req, err = http.NewRequest("GET", "/applicationMDs", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(returnAllApplicationMDs)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), expectedApplicationMD.Title) {
		t.Errorf("Handler returned unexpected body: not got %v within %v",
		rr.Body.String(), responseApplicationMD.Title)
	}
}

func TestPostAndGetByParametersApplicationMD(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/applicationMD", bytes.NewReader([]byte(path + "/testdata/valid1.yaml")))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createNewApplicationMD)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var expectedApplicationMD ApplicationMD
	var responseApplicationMD ApplicationMD
	yamlFile, err := ioutil.ReadFile(path + "/testdata/valid1.yaml")
	err = yaml.Unmarshal(yamlFile, &expectedApplicationMD)
	if err != nil {
		t.Fatal(err)
	}
	err = yaml.Unmarshal([]byte(rr.Body.String()), &responseApplicationMD)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("GET", "/applicationMD?title=Valid%20App%201", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(returnApplicationMDLists)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), expectedApplicationMD.Title) {
		t.Errorf("Handler returned unexpected body: not got %v within %v",
		rr.Body.String(), responseApplicationMD.Title)
	}
}

func TestPutAndGetApplicationMD(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/applicationMD", bytes.NewReader([]byte(path + "/testdata/valid2.yaml")))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createNewApplicationMD)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var expectedApplicationMD ApplicationMD
	var responseApplicationMD ApplicationMD
	yamlFile, err := ioutil.ReadFile(path + "/testdata/valid2.yaml")
	err = yaml.Unmarshal(yamlFile, &expectedApplicationMD)
	if err != nil {
		t.Fatal(err)
	}
	err = yaml.Unmarshal([]byte(rr.Body.String()), &responseApplicationMD)
	if err != nil {
		t.Fatal(err)
	}

	if expectedApplicationMD.Title != responseApplicationMD.Title {
		t.Errorf("handler returned unexpected body: got %v want %v",
		expectedApplicationMD.Title, responseApplicationMD.Title)
	}

	req, err = http.NewRequest("PUT", "/applicationMD?title=Valid%20App%202", bytes.NewReader([]byte(path + "/testdata/valid3.yaml")))
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(putApplicationMD)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), "app 3") {
		t.Errorf("Handler returned unexpected body: not got app 3 within %v",
		rr.Body.String())
	}
}


