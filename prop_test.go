package gopter

import (
	"sync/atomic"
	"testing"
)

func TestPropUndecided(t *testing.T) {
	var called int64
	prop := Prop(func(genParams *GenParameters) *PropResult {
		atomic.AddInt64(&called, 1)

		return &PropResult{
			Status: PropUndecided,
		}
	})

	parameters := DefaultTestParameters()
	result := prop.Check(parameters)

	if result.Status != TestExhausted || result.Succeeded != 0 {
		t.Errorf("Invalid result: %#v", result)
	}
	if called != int64(parameters.MinSuccessfulTests)+1 {
		t.Errorf("Invalid number of calls: %d", called)
	}
}

func TestPropPassed(t *testing.T) {
	var called int64
	prop := Prop(func(genParams *GenParameters) *PropResult {
		atomic.AddInt64(&called, 1)

		return &PropResult{
			Status: PropTrue,
		}
	})

	parameters := DefaultTestParameters()
	result := prop.Check(parameters)

	if result.Status != TestPassed || result.Succeeded != parameters.MinSuccessfulTests {
		t.Errorf("Invalid result: %#v", result)
	}
	if called != int64(parameters.MinSuccessfulTests) {
		t.Errorf("Invalid number of calls: %d", called)
	}
}

func TestPropProof(t *testing.T) {
	var called int64
	prop := Prop(func(genParams *GenParameters) *PropResult {
		atomic.AddInt64(&called, 1)

		return &PropResult{
			Status: PropProof,
		}
	})

	parameters := DefaultTestParameters()
	result := prop.Check(parameters)

	if result.Status != TestProved || result.Succeeded != 1 {
		t.Errorf("Invalid result: %#v", result)
	}
	if called != 1 {
		t.Errorf("Invalid number of calls: %d", called)
	}
}

func TestPropFalse(t *testing.T) {
	var called int64
	prop := Prop(func(genParams *GenParameters) *PropResult {
		atomic.AddInt64(&called, 1)

		return &PropResult{
			Status: PropFalse,
		}
	})

	parameters := DefaultTestParameters()
	result := prop.Check(parameters)

	if result.Status != TestFailed || result.Succeeded != 0 {
		t.Errorf("Invalid result: %#v", result)
	}
	if called != 1 {
		t.Errorf("Invalid number of calls: %d", called)
	}
}

func TestPropError(t *testing.T) {
	var called int64
	prop := Prop(func(genParams *GenParameters) *PropResult {
		atomic.AddInt64(&called, 1)

		return &PropResult{
			Status: PropError,
		}
	})

	parameters := DefaultTestParameters()
	result := prop.Check(parameters)

	if result.Status != TestError || result.Succeeded != 0 {
		t.Errorf("Invalid result: %#v", result)
	}
	if called != 1 {
		t.Errorf("Invalid number of calls: %d", called)
	}
}

func TestPropPassedMulti(t *testing.T) {
	var called int64
	prop := Prop(func(genParams *GenParameters) *PropResult {
		atomic.AddInt64(&called, 1)

		return &PropResult{
			Status: PropTrue,
		}
	})

	parameters := DefaultTestParameters()
	parameters.Workers = 10
	result := prop.Check(parameters)

	if result.Status != TestPassed || result.Succeeded != parameters.MinSuccessfulTests {
		t.Errorf("Invalid result: %#v", result)
	}
	if called != int64(parameters.MinSuccessfulTests) {
		t.Errorf("Invalid number of calls: %d", called)
	}
}
