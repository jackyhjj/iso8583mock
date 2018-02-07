package helpers

import (
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/nkristianto/tcp_server/constants"
)

func TestAPIResponse(t *testing.T) {

	//req := httptest.NewRequest("GET", "test", nil)
	t.Run("Test helper api response", func(t *testing.T) {
		rw := httptest.NewRecorder()
		APIResponse(rw, 200, nil)
	})
}

func TestGetMessage(t *testing.T) {
	t.Run("Test Get Message", func(t *testing.T) {
		r := GetMessage(9000)

		assert.Equal(t, r, apiMessage[9000])
	})
}

func TestGetTransactionMessage(t *testing.T) {
	t.Run("Test get transaction message", func(t *testing.T) {
		r := GetTransactionMessage(constants.TransactionPending)
		assert.Equal(t, r, transactionMessages[constants.TransactionPending])
	})
}

func TestTransformStatusToCode(t *testing.T) {
	t.Run("want got status failed code", func(t *testing.T) {
		r := TransformStatusToCode(constants.TransactionFailed)
		assert.Equal(t, r, uint(constants.B2bAPITransactionFailed))
	})

	t.Run("want got status success code", func(t *testing.T) {
		r := TransformStatusToCode(constants.TransactionSuccess)
		assert.Equal(t, r, uint(constants.B2bAPIGeneralSuccess))
	})

	t.Run("want got status pending code", func(t *testing.T) {
		r := TransformStatusToCode(constants.TransactionPending)
		assert.Equal(t, r, uint(constants.B2bAPITransactionPending))
	})

	t.Run("want got status processing code", func(t *testing.T) {
		r := TransformStatusToCode(constants.TransactionProcess)
		assert.Equal(t, r, uint(constants.B2bAPITransactionProcessing))
	})

}

func TestTransformStatus(t *testing.T) {
	t.Run("want got status failed", func(t *testing.T) {
		r := TransformStatus(constants.TransactionFailed)
		assert.Equal(t, r, constants.TransactionFailed)
	})

	t.Run("want got status success", func(t *testing.T) {
		r := TransformStatus(constants.TransactionSuccess)
		assert.Equal(t, r, constants.TransactionSuccess)
	})

	t.Run("want got status pending", func(t *testing.T) {
		r := TransformStatus(constants.TransactionPending)
		assert.Equal(t, r, constants.TransactionPending)
	})

	t.Run("want got status processing", func(t *testing.T) {
		r := TransformStatus(constants.TransactionProcess)
		assert.Equal(t, r, constants.TransactionProcess)
	})
}
