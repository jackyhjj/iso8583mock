package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/nkristianto/tcp_server/viewmodels"
)

func TestHeartBeat(t *testing.T) {
	var result, rsp viewmodels.ServiceResponse
	var req *http.Request
	var err error
	result.Code = 9000
	result.Message = "OK"

	ctrl := new(HealthCheckController)

	router := mux.NewRouter()
	router.HandleFunc("/healthcheck", ctrl.HeartBeat).Methods("GET")
	rw := httptest.NewRecorder()

	t.Run("health check status", func(t *testing.T) {
		req = httptest.NewRequest("GET", "/healthcheck", nil)
		router.ServeHTTP(rw, req)
		err = json.NewDecoder(rw.Body).Decode(&rsp)
		assert.NoError(t, err)
		assert.Equal(t, rsp.Code, result.Code)
	})
}
