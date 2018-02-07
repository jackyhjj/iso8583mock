// Package controllers of health check
// @author Valentino <daud.darianus@kudo.co.id>
package controllers

import (
	"net/http"

	"github.com/nkristianto/tcp_server/constants"
	"github.com/nkristianto/tcp_server/helpers"
	"github.com/nkristianto/tcp_server/viewmodels"
)

// HealthCheckController representation health check controller
type HealthCheckController struct {
}

// HeartBeat is a server health response
func (c *HealthCheckController) HeartBeat(w http.ResponseWriter, r *http.Request) {

	body := &viewmodels.ServiceResponse{
		Code:    constants.B2bAPIGeneralSuccess,
		Message: "OK",
	}
	helpers.APIResponse(w, 200, body)
	return
}
