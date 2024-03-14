package util

import "testing"

var (
	retryCounter = 0
)

func TestRetry(t *testing.T) {
	for retry := Start(DefaultRetryOptions); retry.Next(); {
		if !succeedAt3rdRetry() {
			t.Logf("failed for the %d time", retryCounter)
			continue
		} else {
			t.Logf("done")
			break
		}
	}
}

func succeedAt3rdRetry() bool {
	if retryCounter == 3 {
		return true
	}
	retryCounter++
	return false
}
