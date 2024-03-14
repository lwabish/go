package util

import "testing"

var (
	retryCounter = 0
)

func TestRetry(t *testing.T) {
	for retry := Start(DefaultRetryOptions); retry.Next(); {
		if !succeedAt(3) {
			t.Logf("failed for the %d time", retryCounter)
			continue
		} else {
			t.Logf("done")
			break
		}
	}
}

func succeedAt(c int) bool {
	retryCounter++
	if retryCounter == c {
		return true
	}
	return false
}
