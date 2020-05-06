package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var tmpArticleList []article

// Used for setup before executing the test function
func TestMain(m *testing.M) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// run the other tests
	os.Exit(m.Run())
}

func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
	}
	return r
}

func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	// create a response recorder
	w := httptest.NewRecorder()

	// create the service and process the above request
	r.ServeHTTP(w, req)
	if !f(w) {
		t.Fail()
	}
}

func saveLists() {
	tmpArticleList = articleList
}

func restoreLists() {
	articleList = tmpArticleList
}


