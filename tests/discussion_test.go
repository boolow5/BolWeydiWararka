package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/boolow5/BolWeydi/controllers"

	. "github.com/franela/goblin"
	"gopkg.in/gin-gonic/gin.v1"
)

// TestUserAdd tests user creation
func TestDiscussion(t *testing.T) {
	gin.SetMode(gin.TestMode)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/discussion", controllers.AddDiscussion)
	r.PUT("/discussion", controllers.UpdateDiscussion)
	r.DELETE("/discussion", controllers.DeleteDiscussion)

	g := Goblin(t)

	g.Describe("AddDiscussion", func() {
		// create request
		body := strings.NewReader(`{"guests":[{"user_id":1}]}`)
		req, err := http.NewRequest(http.MethodPost, "/discussion", body)

		g.It("Should not have request error", func() {
			g.Assert(err).Equal(nil)
		})
		// create response recorder
		rec := httptest.NewRecorder()
		// make the call
		r.ServeHTTP(rec, req)
		// check status
		g.It(fmt.Sprintf("Should return 200 success code = %d", rec.Code), func() {
			g.Assert(rec.Code).Equal(200)
		})
		// Unmarshal the response body
		var responseJson map[string]interface{}
		json.Unmarshal(rec.Body.Bytes(), &responseJson)
		// check response body
		expected_data := "Saved successfully"
		g.It(fmt.Sprintf("Should not return error, error = %v", responseJson["error"]), func() {
			g.Assert(responseJson["error"]).Equal(nil)
		})
		g.It(fmt.Sprintf("Should return success. Response = %v", responseJson), func() {
			g.Assert(responseJson["success"] == nil).Equal(false)
		})
		g.It("Should return the expected message", func() {
			g.Assert(responseJson["success"]).Equal(expected_data)
		})
	})

	g.Describe("UpdateDiscussion", func() {
		// create request
		body := strings.NewReader(`{"guests":[{"user_id":1}]}`)
		req, err := http.NewRequest(http.MethodPut, "/discussion?discussion_id=1", body)
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

	g.Describe("DeleteDiscussion", func() {
		// create request
		body := strings.NewReader(``)
		req, err := http.NewRequest(http.MethodDelete, "/discussion?discussion_id=1", body)
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
