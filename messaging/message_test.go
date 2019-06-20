package messaging

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var apiKey = os.Getenv("MANDRILL_API_KEY")

var _ = Describe("Send email", func() {

	email := Email{From: "abc@example.com", To: "xyz@example.com", Subject: "Testing microservice", Message: "Any message to test", TemplateName: "Template"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(email)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, request)

	Describe("Send email message", func() {
		Context("send", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Send email", func() {

	os.Setenv("API_KEY", apiKey)
	email := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(email)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, request)

	Describe("Send email message", func() {
		Context("send", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Send email", func() {

	os.Setenv("API_KEY", apiKey)
	email := Email{From: "rohits@heaptrace.com", To: "demot636@gmail.com", Subject: "Testing microservice", Message: "Any message to test", TemplateName: "Welcome Text"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(email)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, request)

	Describe("Send email message", func() {
		Context("send", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
