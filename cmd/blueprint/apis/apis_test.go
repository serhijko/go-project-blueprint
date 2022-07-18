package apis

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/serhijko/go-project-blueprint/cmd/blueprint/config"
	"github.com/serhijko/go-project-blueprint/cmd/blueprint/test_data"
	"github.com/stretchr/testify/assert"
)

type apiTestCase struct {
	tag              string
	method           string
	urlToServe       string
	urlToHit         string
	body             string
	function         gin.HandlerFunc
	status           int
	responseFilePath string
}

func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	config.Config.DB = test_data.ResetDB()

	return router
}

func testAPI(router *gin.Engine, method string, urlToServe string, urlToHit string, function gin.HandlerFunc, body string) *httptest.ResponseRecorder {
	router.Handle(method, urlToServe, function)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(method, urlToHit, bytes.NewBufferString(body))
	router.ServeHTTP(res, req)
	return res
}

func runAPITests(t *testing.T, tests []apiTestCase) {
	for _, test := range tests {
		router := newRouter()
		res := testAPI(router, test.method, test.urlToServe, test.urlToHit, test.function, test.body)
		assert.Equal(t, test.status, res.Code, test.tag)
		if test.responseFilePath != "" {
			response, _ := ioutil.ReadFile(test.responseFilePath)
			assert.JSONEq(t, string(response), res.Body.String(), test.tag)
		}
	}
}
