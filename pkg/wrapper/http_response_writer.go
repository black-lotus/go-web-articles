package wrapper

import "net/http"

//WrapResponseWriter structure
type WrapResponseWriter struct {
	StatusCode int
	http.ResponseWriter
}

//NewWrapResponseWriter cunstructor
func NewWrapResponseWriter(res http.ResponseWriter) *WrapResponseWriter {
	// Default the status code to 200
	return &WrapResponseWriter{200, res}
}

//Status Give a way to get the status
func (w WrapResponseWriter) Status() int {
	return w.StatusCode
}

//Header Satisfy the http.ResponseWriter interface get header
func (w WrapResponseWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

//WriteSatisfy the http.ResponseWriter interface write response
func (w WrapResponseWriter) Write(data []byte) (int, error) {
	return w.ResponseWriter.Write(data)
}

//WriteHeader the http.ResponseWriter interface write header
func (w WrapResponseWriter) WriteHeader(statusCode int) {
	// Store the status code
	w.StatusCode = statusCode

	// Write the status code onward.
	w.ResponseWriter.WriteHeader(statusCode)
}
