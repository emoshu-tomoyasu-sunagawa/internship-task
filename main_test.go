package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var memberJSON = `{"profile_img": "https://google.com", "full_name": "テストマン"}`

func aTestCreateMember(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/member", strings.NewReader(memberJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"profile_img": "https://google.com", "full_name": "Testing User"}`, rec.Body.String())
}

func TestHelloApi(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "http://localhost:3030/hello", nil)
	fmt.Println(req)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello, API!", rec.Body.String())
}
