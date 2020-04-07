package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func logPrefix() string {
	return time.Now().Format("2006-01-02 15:04:05.99999")
}

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf("%s %s", logPrefix(), string(bytes))
}

type webLogWriter struct {
	Logger  middleware.LoggerInterface
	NoColor bool
}

func (writer *webLogWriter) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &tinkLogEntry{
		webLogWriter: writer,
		request:      r,
		buf:          &bytes.Buffer{},
	}

	entry.buf.WriteString(logPrefix())

	reqID := middleware.GetReqID(r.Context())
	if reqID != "" {
		unsafeFprintf(entry.buf, " [%s] ", reqID)
	}

	entry.buf.WriteString(r.RemoteAddr)

	// "GET http://url:port/path PROTOCOL"
	unsafeFprintf(entry.buf, " \"")
	unsafeFprintf(entry.buf, "%s ", r.Method)

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	unsafeFprintf(entry.buf, "%s://%s%s %s\" ", scheme, r.Host, r.RequestURI, r.Proto)

	return entry
}

func unsafeFprintf(w io.Writer, format string, a ...interface{}) {
	_, _ = fmt.Fprintf(w, format, a...)
}

type tinkLogEntry struct {
	*webLogWriter
	request *http.Request
	buf     *bytes.Buffer
}

func (l *tinkLogEntry) Write(status, bytes int, elapsed time.Duration) {
	unsafeFprintf(l.buf, "- %03d ", status)

	unsafeFprintf(l.buf, "[%dB", bytes)
	unsafeFprintf(l.buf, "/%s]", elapsed)

	l.Logger.Print(l.buf.String())
}

func (l *tinkLogEntry) Panic(v interface{}, stack []byte) {
	panicEntry := l.NewLogEntry(l.request).(*tinkLogEntry)
	unsafeFprintf(panicEntry.buf, "panic: %+v", v)
	l.Logger.Print(panicEntry.buf.String())
	l.Logger.Print(string(stack))
}

func init() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{})
}
