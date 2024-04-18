package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestArticleListHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "aaa", resultCode: http.StatusBadRequest},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			router.GET("/article/list", aCon.ArticleListHandler)

			url := fmt.Sprintf("/article/list?page=%s", tt.query)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode for %s: want %d but %d\n", tt.name, tt.resultCode, w.Code)
			}
		})
	}
}

func TestArticleDetailHandler(t *testing.T) {
	var tests = []struct {
		name       string
		articleID  string
		resultCode int
	}{
		{name: "number pathparam", articleID: "1", resultCode: http.StatusOK},
		{name: "alphabet pathparam", articleID: "aaa", resultCode: http.StatusBadRequest},
	}

	gin.SetMode(gin.TestMode)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := gin.Default()
			router.GET("/article/:id", aCon.ArticleDetailHandler)

			url := fmt.Sprintf("/article/%s", tt.articleID)
			req := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode for %s: want %d but %d\n", tt.name, tt.resultCode, w.Code)
			}
		})
	}
}
