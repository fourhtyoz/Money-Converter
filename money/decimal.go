package money

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	ErrInvalidDecimal = Error("unable to convert the decimal")
	ErrTooLarge = Error("quantity over 10ˆ12 is too large")
)

// Decimal is capable of storing a floating-point value.
type Decimal struct {
	subunits int64
	precision byte
}

func ParseDecimal(amount string) (Decimal, error) {
	intPart, fracPart, _ := strings.Cut(amount, ".")

	const maxDecimal = 12

	if len(intPart) > maxDecimal {
		return Decimal{}, ErrTooLarge
	}

	subunits, err := strconv.ParseInt(intPart+fracPart, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err)
	}

	precision := byte(len(fracPart))

	return Decimal{subunits: subunits, precision: precision}, nil
}

func (d *Decimal) simplify() {
	for d.subunits % 10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}
