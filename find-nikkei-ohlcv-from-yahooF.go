package nikke22go

import (
	"fmt"
	"strconv"
	"time"

	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
)

func GetOHLCvFromYahoo(code string) ([]finance.ChartBar, []finance.ChartMeta, error) {
	t := time.Now()
	ago5 := t.Add(-5 * 365 * 24 * time.Hour)
	p := &chart.Params{
		Symbol:   conversion(code),
		Start:    &datetime.Datetime{Month: int(ago5.Month()), Day: ago5.Day(), Year: ago5.Year()},
		End:      &datetime.Datetime{Month: int(t.Month()), Day: t.Day(), Year: t.Year()},
		Interval: datetime.OneDay,
	}
	iter := chart.Get(p)
	if iter == nil {
		return nil, nil, fmt.Errorf("does not gets data")
	}

	// Iterate over results. Will exit upon any error.
	var (
		results = make([]finance.ChartBar, 0)
		metas   = make([]finance.ChartMeta, 0)
	)
	for iter.Next() {
		b := iter.Bar()
		results = append(results, *b)

		// Meta-data for the iterator - (*finance.ChartMeta).
		metas = append(metas, iter.Meta())
	}

	// Catch an error, if there was one.
	if iter.Err() != nil {
		return nil, nil, fmt.Errorf("iterater error")
	}

	return results, metas, nil
}

func conversion(code string) string {
	_, err := strconv.Atoi(code)
	if err != nil {
		return code
	}

	return fmt.Sprintf("%s.T", code)

}
