// Package constants message code
// @author Valentino <daud.darianus@kudo.co.id>
package constants

// B2B response code & status definition
const (
	// B2bAPIGeneralSuccess response code
	B2bAPIGeneralSuccess = 9000

	// B2bAPIInvalidAuthentication response code
	B2bAPIInvalidAuthentication = 9101

	// B2bAPIInvalidSignature response code
	B2bAPIInvalidSignature = 9102

	// B2bAPIInvalidValidation response code
	B2bAPIInvalidValidation = 9103

	// B2bAPIDataNotFound response code
	B2bAPIDataNotFound = 9201

	// B2bAPITransactionProcessing response code
	B2bAPITransactionProcessing = 9301

	// B2bAPITransactionPending response code
	B2bAPITransactionPending = 9302

	// B2bAPITransactionFailed response code
	B2bAPITransactionFailed = 9303

	// B2bAPITransactionReachLimit response code
	B2bAPITransactionReachLimit = 9304

	// APIUnknownError response code
	APIUnknownError = 2000

	// TransactionSuccess status success
	TransactionSuccess = "success"

	// TransactionPending status pending
	TransactionPending = "pending"

	// TransactionProcess status processing
	TransactionProcess = "processing"

	// TransactionFailed status failed
	TransactionFailed = "failed"

	// TransactionOnQueue status on queue
	TransactionOnQueue = "queue"

	// AirtimeServiceName is the identification name application
	AppServiceName = "biller_connector_ptpos"
)
