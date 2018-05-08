package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ngo275/gae-blueprint/src/handler"
)

var _ = Describe("Ping", func() {
	var pong handler.PingRes

	BeforeEach(func() {
		req, _ := http.NewRequest("GET", "/ping", nil)
		res := httptest.NewRecorder()
		handler.PingHandler(res, req)

		json.Unmarshal([]byte(res.Body.String()), &pong)
	})

	Describe("GET /ping", func() {
		It("returns pong", func() {
			Ω(pong.Ping).To(Equal("pong"))
		})
	})
})
