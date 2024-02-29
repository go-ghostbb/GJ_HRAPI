package status

import gbcode "ghostbb.io/gb/errors/gb_code"

var (
	CodeNil                       = gbcode.New(-1, "", nil)                             // No error code specified.
	CodeOK                        = gbcode.New(0, "OK", nil)                            // It is OK.
	CodeFailed                    = gbcode.New(10000, "Failed", nil)                    // It is failed.
	CodeInternalError             = gbcode.New(50, "Internal Error", nil)               // An error occurred internally.
	CodeValidationFailed          = gbcode.New(51, "Validation Failed", nil)            // Data validation failed.
	CodeDbOperationError          = gbcode.New(52, "Database Operation Error", nil)     // Database operation error.
	CodeInvalidParameter          = gbcode.New(53, "Invalid Parameter", nil)            // The given parameter for current operation is invalid.
	CodeMissingParameter          = gbcode.New(54, "Missing Parameter", nil)            // Parameter for current operation is missing.
	CodeInvalidOperation          = gbcode.New(55, "Invalid Operation", nil)            // The function cannot be used like this.
	CodeInvalidConfiguration      = gbcode.New(56, "Invalid Configuration", nil)        // The configuration is invalid for current operation.
	CodeMissingConfiguration      = gbcode.New(57, "Missing Configuration", nil)        // The configuration is missing for current operation.
	CodeNotImplemented            = gbcode.New(58, "Not Implemented", nil)              // The operation is not implemented yet.
	CodeNotSupported              = gbcode.New(59, "Not Supported", nil)                // The operation is not supported yet.
	CodeOperationFailed           = gbcode.New(60, "Operation Failed", nil)             // I tried, but I cannot give you what you want.
	CodeNotAuthorized             = gbcode.New(61, "Not Authorized", nil)               // Not Authorized.
	CodeSecurityReason            = gbcode.New(62, "Security Reason", nil)              // Security Reason.
	CodeServerBusy                = gbcode.New(63, "Server Is Busy", nil)               // Server is busy, please try again later.
	CodeUnknown                   = gbcode.New(64, "Unknown Error", nil)                // Unknown error.
	CodeNotFound                  = gbcode.New(65, "Not Found", nil)                    // Resource does not exist.
	CodeInvalidRequest            = gbcode.New(66, "Invalid Request", nil)              // Invalid request.
	CodeNecessaryPackageNotImport = gbcode.New(67, "Necessary Package Not Import", nil) // It needs necessary package import.
	CodeInternalPanic             = gbcode.New(68, "Internal Panic", nil)               // A panic occurred internally.
	CodeBusinessValidationFailed  = gbcode.New(300, "Business Validation Failed", nil)  // Business validation failed.
)

/*
錯誤碼
*/
var (
	CodeRequestInvalidMethod = gbcode.New(10001, "invalid HTTP method", nil)
	CodeRequestInvalidParam  = gbcode.New(10002, "invalid dynamic route parameter", nil)
	CodeRequestInvalidQuery  = gbcode.New(10003, "invalid request parameter", nil)
	CodeRequestInvalidBody   = gbcode.New(10004, "invalid request body", nil)
	CodeRequestInvalidForm   = gbcode.New(10005, "invalid form data", nil)
	CodeRequestInvalidHeader = gbcode.New(10006, "invalid request header", nil)
)

var (
	CodeAuthTokenIllegal = gbcode.New(20001, "invalid authentication token", nil)
	CodeAuthTokenExpired = gbcode.New(20002, "expired authentication token", nil)
)

var (
	CodeLoginFailed = gbcode.New(30001, "login failed", nil)
)
