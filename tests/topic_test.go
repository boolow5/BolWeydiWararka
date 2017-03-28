package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/boolow5/BolWeydi/controllers"

	. "github.com/franela/goblin"
	"gopkg.in/gin-gonic/gin.v1"
)

// TestUserAdd tests user creation
func TestTopic(t *testing.T) {
	gin.SetMode(gin.TestMode)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/topic", controllers.AddTopic)
	r.PUT("/topic", controllers.UpdateTopic)
	r.DELETE("/topic", controllers.DeleteTopic)

	g := Goblin(t)

	g.Describe("AddTopic", func() {
		// create request
		body := strings.NewReader(`{"text":"Mowduuca Koowaad"}`)
		req, err := http.NewRequest(http.MethodPost, "/topic", body)

		g.It("Should not have request error", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Should return 200 success code", func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		expected_data := "Saved successfully"
		g.It("Should not return error", func() {
			g.Assert(responseJson["error"] == nil).Equal(true)
		})
		g.It("Should return success", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should return the expected message", func() {
			g.Assert(responseJson["success"]).Equal(expected_data)
		})
	})

	g.Describe("UpdateTopic", func() {
		// create request
		body := strings.NewReader(`{"text":"Waa mowduuc cusub"}`)
		req, err := http.NewRequest(http.MethodPut, "/topic?topic_id=1", body)
		g.It("Should not have request error", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Should return 200 success code", func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		expected_data := "updated successfully"
		g.It("Should not return error", func() {
			g.Assert(responseJson["error"] == nil).Equal(true)
		})
		g.It("Should return success", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should return the expected message", func() {
			g.Assert(responseJson["success"]).Equal(expected_data)
		})
	})

	g.Describe("DeleteTopic", func() {
		// create request
		body := strings.NewReader(``)
		req, err := http.NewRequest(http.MethodDelete, "/topic?topic_id=1", body)
		g.It("Should not have request error", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It("Should return 200 success code", func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		expected_data := "deleted successfully"
		g.It("Should not return error", func() {
			g.Assert(responseJson["error"] == nil).Equal(true)
		})
		g.It("Should return success", func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should return the expected message", func() {
			g.Assert(responseJson["success"]).Equal(expected_data)
		})
	})
}
