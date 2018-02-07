// Package controllers of health check
// @author Valentino <daud.darianus@kudo.co.id>
package controllers

import (
	"net/http"

	"github.com/nkristianto/tcp_server/helpers"
)

// HealthCheckController representation health check controller
type HealthCheckController struct {
}

// HeartBeat is a server health response
func (c *HealthCheckController) HeartBeat(w http.ResponseWriter, r *http.Request) {

	body := &viewmodels.ServiceResponse{
		Code:    "1000",
		Message: "OK",
	}
	helpers.APIResponse(w, 200, body)
	return
}
