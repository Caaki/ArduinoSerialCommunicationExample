package handlers

import (
	"net/http"
)

// Response structure
type Response struct {
	Code     int         `json:"code"`
	Response interface{} `json:"data"`
	Errors   interface{} `json:"errors"`
}

// Chan ...
type Chan struct {
	Code int
	Data interface{}
}

// Create channel interface (can be used instead of ch <- Chan{...} )
func (Chan) Create(code int, data interface{}) Chan {
	return Chan{
		Code: code,
		Data: data,
	}
}

// CreateChanResponse - Create response based on channel data.
func CreateChanResponse(response Chan) Response {
	return CreateResponse(response.Code, response.Data)
}

// CreateResponse - Create response. This function will filter response by code, and create response based on
// code and data provided. Added new check to reduce unlimited interface creation when response interface "nil"
func CreateResponse(code int, data interface{}) Response {
	if data == nil {
		data = map[string]interface{}{}
	}

	switch code {
	case http.StatusOK:
		fallthrough
	case http.StatusCreated:
		return StatusOK(data)
	case http.StatusBadRequest:
		return StatusBadRequest(data)
	case http.StatusUnauthorized:
		return StatusUnauthorized(data)
	case http.StatusNotFound:
		return StatusNotFound(data)
	case http.StatusForbidden:
		return StatusForbidden(data)
	case http.StatusMethodNotAllowed:
		return StatusMethodNotAllowed(data)
	case http.StatusConflict:
		return StatusConflict(data)
	case http.StatusServiceUnavailable:
		return StatusServiceUnavailable(data)
	case http.StatusInternalServerError:
		return StatusInternalServerError(data)
	default:
		return StatusInternalServerError(map[string]interface{}{"0": "error_internal_server_error"})
	}
}

// StatusOK - Create OK response - 200
func StatusOK(data interface{}) Response {
	response := Response{
		Code:     http.StatusOK,
		Response: data,
	}

	return response
}

// StatusBadRequest - Create bad request response - 400
func StatusBadRequest(data interface{}) Response {
	response := Response{
		Code:   http.StatusBadRequest,
		Errors: data,
	}

	return response
}

// StatusUnauthorized - Create unauthorized response - 401
func StatusUnauthorized(data interface{}) Response {
	response := Response{
		Code:   http.StatusUnauthorized,
		Errors: data,
	}

	return response
}

// StatusForbidden - Create forbidden response - 403
func StatusForbidden(data interface{}) Response {
	response := Response{
		Code:   http.StatusForbidden,
		Errors: data,
	}

	return response
}

// StatusNotFound - Create method not found response - 404
func StatusNotFound(data interface{}) Response {
	response := Response{
		Code:   http.StatusNotFound,
		Errors: data,
	}

	return response
}

// StatusMethodNotAllowed - Create method not allowed response - 405
func StatusMethodNotAllowed(data interface{}) Response {
	response := Response{
		Code:   http.StatusMethodNotAllowed,
		Errors: data,
	}

	return response
}

// StatusConflict - Create conflict response - 409
func StatusConflict(data interface{}) Response {
	response := Response{
		Code:   http.StatusConflict,
		Errors: data,
	}

	return response
}

// StatusInternalServerError - Create internal server error response - 500
func StatusInternalServerError(data interface{}) Response {
	response := Response{
		Code:   http.StatusInternalServerError,
		Errors: data,
	}

	return response
}

// StatusServiceUnavailable - Create service unavailable response - 503
func StatusServiceUnavailable(data interface{}) Response {
	response := Response{
		Code:   http.StatusServiceUnavailable,
		Errors: data,
	}

	return response
}
