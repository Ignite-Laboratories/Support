package testing

import (
	"github.com/ignite-laboratories/support"
	"github.com/ignite-laboratories/support/operate"
	"testing"
)

func Test_Operate_OnEach(t *testing.T) {
	data := support.RandomBytes(32)
	startAvg := operate.GetAverage(data...)

	delta := byte(5)

	// Add to each

	data = operate.OnEach(data, operate.Add, delta)
	avg := operate.GetAverage(data...)

	result := avg - startAvg
	if result != delta { // Should be the delta value
		t.Errorf("Expected %d, got %d", delta, result)
	}

	// Subtract from each

	data = operate.OnEach(data, operate.Subtract, delta)
	avg = operate.GetAverage(data...)

	result = avg - startAvg
	if result != 0 { // Should be zero
		t.Errorf("Expected %d, got %d", delta, result)
	}
}

func Test_Operate_GetAverage(t *testing.T) {
	data := support.RandomBytes(32)
	average := 0

	for _, d := range data {
		average += int(d)
	}
	average /= 32
	toTest := operate.GetAverage(data...)

	if average != int(toTest) {
		t.Errorf("Expected %d, got %d", average, toTest)
	}
}
