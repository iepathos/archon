package archon_test

import (
	. "github.com/iepathos/archon"

	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Archon", func() {
	Describe("GET GetHandler", func() {
		Context("Simple GET request", func() {
			req, _ := http.NewRequest("GET", "/health-get", nil)

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetHandler)

			handler.ServeHTTP(rr, req)
			It("should return 200 status code", func() {
				Expect(rr.Code).To(Equal(http.StatusOK))
			})

			It("Should have the expected body", func() {
				expected := `{"alive": true}`
				Expect(rr.Body.String()).To(Equal(expected))
			})
		})
	})

	Describe("POST PostHandler", func() {
		Context("Simple POST request", func() {
			jsonStr := []byte(`{"service":"test"}`)
			req, _ := http.NewRequest("POST", "/health-post", bytes.NewBuffer(jsonStr))

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(PostHandler)

			handler.ServeHTTP(rr, req)
			It("should return 200 status code", func() {
				Expect(rr.Code).To(Equal(http.StatusOK))
			})

			It("Should have the expected body", func() {
				expected := `{"service":"test"}`
				Expect(rr.Body.String()).To(Equal(expected))
			})
		})
	})
})
