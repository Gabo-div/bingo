package converters

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func StringToTimestamp(date string) (pgtype.Timestamp, error) {
	t, err := time.Parse(time.RFC3339, date)

	if err != nil {
		return pgtype.Timestamp{}, err
	}

	return pgtype.Timestamp{
		Time:  t,
		Valid: true,
	}, nil
}

func Float64ToNumeric(float float64) (pgtype.Numeric, error) {
	var numeric pgtype.Numeric

	err := numeric.Scan(fmt.Sprintf("%f", float))

	if err != nil {
		return pgtype.Numeric{}, err
	}

	return numeric, nil
}
