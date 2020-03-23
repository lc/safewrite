package safewrite

import (
	"os"
	"strconv"
	"testing"
)

func TestSafeWriter(t *testing.T) {
	fp, err := os.Create("hello.txt")
	if err != nil {
		t.Error(err)
	}
	w := NewWriter(fp)
	ap := NewAppender(w)

	count := 100
	done := make(chan bool, count)

	for i := 0; i < count; i++ {
		f := strconv.Itoa(i)
		go func(i string) {
			ap.Append([]byte(i))
			done <- true
		}(f)
	}
	for i := 0; i < count; i++ {
		<-done
	}

}
