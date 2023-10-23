package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (pl *ProcessingLog) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	pl.errLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (pl *ProcessingLog) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (pl *ProcessingLog) notFound(w http.ResponseWriter) {
	pl.clientError(w, http.StatusNotFound)
}
