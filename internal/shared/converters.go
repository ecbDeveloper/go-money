package shared

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertNumericToFloat(n pgtype.Numeric) (float64, error) {
	f, err := n.Float64Value()
	if err != nil || !f.Valid {
		return 0, err
	}

	return f.Float64, nil
}

func ConvertFloatToNumeric(f float64) (pgtype.Numeric, error) {
	fStr := fmt.Sprintf("%f", f)

	var n pgtype.Numeric
	err := n.Scan(fStr)
	if err != nil {
		return pgtype.Numeric{}, err
	}

	return n, nil
}
