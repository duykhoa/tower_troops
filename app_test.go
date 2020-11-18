package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"tower_troops/controllers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	w := proccessRequest(http.MethodGet, "/ping", nil, controllers.SetupRoute)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "{\"msg\":\"OK\"")
}

func TestCreateTowerAPI(t *testing.T) {
	w := proccessRequest(http.MethodPost, "/towers", nil, controllers.SetupRoute)

	fmt.Println(w.Body.String())

	assert.Equal(t, http.StatusCreated, w.Code)
}

func proccessRequest(httpMethod string, path string, body io.ReadCloser, routerFn func(*gin.Engine) *gin.Engine) *httptest.ResponseRecorder {
	gin.SetMode(gin.ReleaseMode)
	req := httptest.NewRequest(httpMethod, path, nil)

	if body != nil {
		req.Body = body
	}

	w := httptest.NewRecorder()
	r := gin.New()

	routerFn(r).ServeHTTP(w, req)

	return w
}
