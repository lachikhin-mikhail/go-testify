package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	city := "moscow"
	target := fmt.Sprintf("/cafe?count=%d&city=%s", totalCount+1, city)
	req := httptest.NewRequest("GET", target, nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	status := responseRecorder.Code
	require.Equal(t, status, http.StatusOK)

	cafes := responseRecorder.Body.String()
	cafeList := strings.Split(cafes, ",")

	require.Equal(t, totalCount, len(cafeList))
}

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	count := 2
	city := "moscow"
	target := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)
	req := httptest.NewRequest("GET", target, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	answer := responseRecorder.Body
	require.Equal(t, http.StatusOK, status)
	require.NotEmpty(t, answer)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	count := 2
	city := "mossycow"
	target := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)
	req := httptest.NewRequest("GET", target, nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	errReturned := responseRecorder.Body.String()
	require.Equal(t, http.StatusBadRequest, status)
	require.Equal(t, "wrong city value", errReturned)
}
