package types

import "fmt"

type Float float64

func (n Float) MarshalJSON() ([]byte, error) {
	if float64(n) == float64(int64(n)) {
		return []byte(fmt.Sprintf("%.1f", n)), nil
	} else {
		return []byte(fmt.Sprint(n)), nil
	}
}

