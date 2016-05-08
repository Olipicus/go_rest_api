package rest

import "net/http"

var (
	resultSuccess        Result = Result{StatusCode: http.StatusOK, Description: "Success"}
	resultInsertComplete Result = Result{StatusCode: http.StatusCreated, Description: "Insert Success"}
	resultUpdateComplete Result = Result{StatusCode: http.StatusOK, Description: "Update Success"}
	resultDeleteComplete Result = Result{StatusCode: http.StatusOK, Description: "Delete Success"}
	resultDataNotFound   Result = Result{StatusCode: http.StatusNoContent, Description: "Data Not Found"}
)

//Result Struct for print json result
type Result struct {
	StatusCode  int    `json:"status_code"`
	Description string `json:"description"`
}

//ResultSuccess for print json result
type ResultSuccess struct {
	Result
	Data interface{} `json:"data"`
}

//Error error message
type Error struct {
	Code         int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

//ResultError for print json result
type ResultError struct {
	Result
	Error Error `json:"error"`
}
