// Package helpers Api response
// @author Novian Kristianto <novian.kristianto@kudo.co.id>
// based on helper from Valentino <daud.darianus@kudo.co.id>
package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/nkristianto/tcp_server/constants"
)

var (
	transactionMessages = map[string]string{
		constants.TransactionFailed:  "the transaction is failed",
		constants.TransactionSuccess: "the transaction is success",
		constants.TransactionPending: "the transaction is pending",
		constants.TransactionProcess: "the transaction is processing",
		constants.TransactionOnQueue: "the transaction is processing",
	}

	apiMessage = map[int]string{
		constants.B2bAPIGeneralSuccess:        "success",
		constants.B2bAPIInvalidAuthentication: "invalid authentication",
		constants.B2bAPIInvalidSignature:      "invalid signature",
		constants.B2bAPIInvalidValidation:     "invalid validation",
		constants.B2bAPIDataNotFound:          "the data not found in our system",
		constants.B2bAPITransactionProcessing: "the transaction is processing",
		constants.B2bAPITransactionPending:    "the transaction is pending",
		constants.B2bAPITransactionFailed:     "the transaction is fail",
		constants.B2bAPITransactionReachLimit: "the transaction has reached limit",
		constants.APIUnknownError:             "unknown error, please contact administrator",
	}
)

// APIResponse  - render response json output
func APIResponse(w http.ResponseWriter, httpStatus int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	json.NewEncoder(w).Encode(body)
}

// GetMessage by message code
func GetMessage(code int) string {
	return apiMessage[code]
}

// GetTransactionMessage by status transaction
func GetTransactionMessage(status string) string {
	return transactionMessages[status]
}

// TransformStatus airtime to grab
func TransformStatus(status string) string {
	switch status {
	case constants.TransactionFailed:
		return constants.TransactionFailed
	case constants.TransactionPending:
		return constants.TransactionPending
	case constants.TransactionSuccess:
		return constants.TransactionSuccess
	default:
		return constants.TransactionProcess
	}
}

// TransformStatusToCode status the transaction
func TransformStatusToCode(status string) uint {
	switch status {
	case constants.TransactionPending:
		return constants.B2bAPITransactionPending
	case constants.TransactionFailed:
		return constants.B2bAPITransactionFailed
	case constants.TransactionSuccess:
		return constants.B2bAPIGeneralSuccess
	default:
		return constants.B2bAPITransactionProcessing
	}
}
