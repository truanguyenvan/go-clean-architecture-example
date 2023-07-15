package conditions

import (
	"go-clean-architecture-example/pkg/db/mongo/utils"
	"testing"
)

func TestExactMatch(t *testing.T) {
	condition := Pipe(
		DateLessThanOrEqualTo(Condition{
			Key:   "endDate",
			Value: "2006-01-02T15:04:05.000Z",
		}),
		DateGreaterThanOrEqualTo(Condition{
			Key:   "startDate",
			Value: "2006-01-02T15:04:05.000Z",
		}),
	)
	utils.PrintJson(condition)
}
