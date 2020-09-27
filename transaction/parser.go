package transaction

import "fmt"

func ParseAmount(value float64, ot OperationType) (float64, error) {
	switch ot {
	case Payment:
		return value, nil
	case ForCash, Installment, Withdrawn:
		return value * -1, nil
	default:
		return 0.0, fmt.Errorf("invalid operation: %s", ot)
	}
}
