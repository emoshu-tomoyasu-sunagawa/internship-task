package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// var memberJSON = `{"profile_img": "https://google.com", "full_name": "テストマン"}`

// func TestCreateMember(t *testing.T) {
// 	router := NewRouter()
// 	req := httptest.NewRequest(echo.POST, "/member", strings.NewReader(memberJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()

// 	router.ServeHTTP(rec, req)

// 	assert.Equal(t, http.StatusCreated, rec.Code)
// 	assert.JSONEq(t, `{"profile_img": "https://google.com", "full_name": "Testing User"}`, rec.Body.String())
// }

func TestGetAllMembers(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(echo.GET, "/members", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"profile_img": "https://google.com", "full_name": "Testing User"}`, rec.Body.String())
}

func TestHelloApi(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, API!", rec.Body.String())
}
