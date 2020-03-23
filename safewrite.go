package safewrite

import (
	"io"
	"sync"
)

type SafeAppend struct {
	io.WriteCloser
}
type SafeWriter struct {
	mutex  *sync.Mutex
	writer io.WriteCloser
}

func NewWriter(w io.WriteCloser) *SafeWriter {
	return &SafeWriter{writer: w, mutex: &sync.Mutex{}}
}
func (w *SafeWriter) Write(b []byte) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.writer.Write(b)
}
func (w *SafeWriter) WriteString(b string) (int, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.writer.Write([]byte(b))
}
func (w *SafeWriter) Close() error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return w.writer.Close()
}
func NewAppender(w io.WriteCloser) SafeAppend {
	return SafeAppend{w}
}
func (w *SafeAppend) Append(b []byte) error {
	_, err := w.Write(append(b, '\n'))
	return err
}

func (w *SafeAppend) Close() error {
	return w.Close()
}
