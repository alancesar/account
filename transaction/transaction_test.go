package transaction

import "testing"

func TestParseAmount(t *testing.T) {
	operations := map[OperationType]float64{
		ForCash:     -100.0,
		Installment: -100.0,
		Withdrawn:   -100.0,
		Payment:     100.0,
	}

	for k, v := range operations {
		if a, _ := ParseAmount(100.0, k); a != v {
			t.Errorf("unexpected value at ParseAmount with %s", k)
		}
	}
}
