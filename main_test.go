package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct {
	mock.Mock
}

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
	// routerを設定
	router := NewRouter()
	router.GET("/members", getAllMembers)

	// リクエストを作成
	req, err := http.NewRequest("GET", "/members", nil)
	if err != nil {
		fmt.Println("bbbbbbb")
		panic("bbbbb")
	}

	// テスト用サーバーを作成
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var members []Member
	json.Unmarshal(w.Body.Bytes(), &members)

	// assertする
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 24, len(members))
	// router := NewRouter()
	// req := httptest.NewRequest(echo.GET, "/members", nil)
	// rec := httptest.NewRecorder()

	// router.ServeHTTP(rec, req)

	// assert.Equal(t, http.StatusOK, rec.Code)
	// assert.JSONEq(t, `{"profile_img": "https://google.com", "full_name": "Testing User"}`, rec.Body.String())
}

func aTestAddMember(t *testing.T) {
	// newMember := Member{
	// 	ProfileImg: "https://emoshu.co.jp",
	// 	FullName:   "エモッシュ　タロウ",
	// }
	var newMember = `{"profile_img": "https://google.com", "full_name": "テストマン"}`
	// requestBody, _ := json.Marshal(newMember)
	// request, _ := http.NewRequest("POST", "/member", bytes.NewBuffer(requestBody))
	request, _ := http.NewRequest("POST", "/member", bytes.NewBuffer([]byte(newMember)))
	writer := httptest.NewRecorder()
	NewRouter().ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)

	// writer := makeRequest("POST", "/auth/register", newUser, false)
	// assert.Equal(t, http.StatusCreated, writer.Code)
}

// func TestHelloApi(t *testing.T) {
// 	router := NewRouter()
// 	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
// 	rec := httptest.NewRecorder()
// 	router.ServeHTTP(rec, req)

// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, "Hello, API!", rec.Body.String())
// }

// func aTestCeate(t *testing.T) {
// 	member := &Member{
// 		ProfileImg: "https://emoshu.co.jp",
// 		FullName:   "エモッシュ　タロウ",
// 	}
// }
