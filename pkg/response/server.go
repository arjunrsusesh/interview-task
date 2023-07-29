package response

import (
	"encoding/json"
	"log"
	"net/http"
	"task/pkg/e"
)

type Response struct {
	Status string          `json:"status"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *ErrorRes       `json:"error,omitempty"`
}

type ErrorRes struct {
	ErrorCode int      `json:"error_code"`
	Messaage  string   `json:"message"`
	Details   []string `json:"details"`
}

const (
	StatusOK    = "ok"
	StatusNotOK = "nok"
)

func Success(w http.ResponseWriter, status int, result interface{}) {
	var r *Response
	if result != nil {
		res, err := json.Marshal(result)
		if err != nil {
			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
			return
		}
		r = &Response{
			Status: StatusOK,
			Result: res,
		}
	}

	j, err := json.Marshal(r)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(j)
	if err != nil {
		log.Println("failed to write response")
	}
}

func Fail(w http.ResponseWriter, status, errCode int, details ...string) {
	var r *Response

	msg := e.ErrorMap[errCode].Error()

	r = &Response{
		Status: StatusNotOK,
		Error: &ErrorRes{
			ErrorCode: errCode,
			Messaage:  msg,
			Details:   details,
		},
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(j)
	if err != nil {
		log.Println("failed to write response")
	}
}
