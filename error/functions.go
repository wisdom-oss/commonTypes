package error

import (
	"fmt"
	"net/http"
	"strings"
)

// InferHttpStatusText takes the already configured HTTP Status Code and infers
// the HTTP status text from it using the net/http package. If the status code
// is not known to the package the HTTP status text will stay empty
func (e *WISdoMError) InferHttpStatusText() {
	e.HttpStatusText = http.StatusText(e.HttpStatusCode)
}

// WrapError takes a native golang error as parameter and wraps it into a
// WISdoMError. The WISdoMError instance will overwrite every field already
// present on the error and set the status code to 500 to indicate a internal
// error occurred. It optionally takes a service name as argument, if multiple
// names are supplied they are joined together using a dot (.)
func (e *WISdoMError) WrapError(err error, serviceName ...string) {
	// create the error code prefix
	errorCodePrefix := strings.Join(serviceName, ".")
	// now build full the full error code used for wrapping internal errors
	e.ErrorCode = fmt.Sprintf("%s.%s", errorCodePrefix, "INTNERNAL_ERROR")
	// set the title to a generic error title
	e.ErrorTitle = "Internal Error in Microservice"
	// set the error description to the external error
	e.ErrorDescription = err.Error()
	// set the http code to 500 and infer the text from this
	e.HttpStatusCode = http.StatusInternalServerError
	e.HttpStatusText = http.StatusText(e.HttpStatusCode)
}

// Send takes a response writer as agrument and sends the error to the request
// origin using JSON. This function is only callable, if no other response
// headers have been written, due to the fact that only one http status code
// may be sent in a HTTP response
func (e WISdoMError) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")
}
