// Package routes is the route & midlleware service package
// @author Valentino <daud.darianus@kudo.co.id>
package routes

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/urfave/negroni"

	"github.com/nkristianto/tcp_server/viewmodels"
)

// RequestLogger is the structure
// passed to the template.
type RequestLogger struct {
	StartTime string
	Status    int
	Duration  time.Duration
	Hostname  string
	Method    string
	Path      string
	Request   *http.Request
}

// LoggerDefaultFormat is the format
// logged used by the default Logger instance.
var LoggerDefaultFormat = "{{.StartTime}} | {{.Status}} | \t {{.Duration}} | {{.Hostname}} | {{.Method}} {{.Path}} \n"

// LoggerDefaultDateFormat is the
// format used for date by the
// default Logger instance.
var LoggerDefaultDateFormat = time.RFC3339

// ALogger interface
type ALogger interface {
	Println(v ...interface{})
	Printf(format string, v ...interface{})
}

// LoggerMiddleware is a middleware handler that logs the request as it goes in and the response as it goes out.
type LoggerMiddleware struct {
	// ALogger implements just enough log.Logger interface to be compatible with other implementations
	ALogger
	dateFormat string
	template   *template.Template
}

// NewLoggerMiddleware returns a new Logger instance
func NewLoggerMiddleware() *LoggerMiddleware {
	logger := &LoggerMiddleware{ALogger: log.New(os.Stdout, "[biller - narindo] ", 0), dateFormat: LoggerDefaultDateFormat}
	logger.SetFormat(LoggerDefaultFormat)
	logger.SetDateFormat("2006-01-02 15:04:05")
	return logger
}

// SetFormat output log
func (l *LoggerMiddleware) SetFormat(format string) {
	l.template = template.Must(template.New("kudo_parser").Parse(format))
}

// SetDateFormat log time
func (l *LoggerMiddleware) SetDateFormat(format string) {
	l.dateFormat = format
}

// ServeHTTP the http serve
func (l *LoggerMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	defer func() {
		r := recover()
		if r != nil {
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusInternalServerError)

			rsp := viewmodels.ServiceResponse{
				Code:    2000,
				Message: "internal server error",
			}

			json.NewEncoder(rw).Encode(&rsp)

			return
		}
	}()

	start := time.Now()

	next(rw, r)

	res := rw.(negroni.ResponseWriter)
	log := RequestLogger{
		StartTime: start.Format(l.dateFormat),
		Status:    res.Status(),
		Duration:  time.Since(start),
		Hostname:  r.Host,
		Method:    r.Method,
		Path:      r.URL.Path,
		Request:   r,
	}

	buff := &bytes.Buffer{}
	l.template.Execute(buff, log)
	l.Printf(buff.String())
}
