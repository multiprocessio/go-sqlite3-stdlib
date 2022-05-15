package stdlib

import (
	"fmt"
	"testing"
)

func Test_date(t *testing.T) {
	tests := []struct {
		in  string
		fn  string
		out int
	}{
		{"2021-04-05", "date_year", 2021},
		{"May 6, 2020", "date_month", 5},
		{"May 6, 2020", "date_day", 6},
		{"May 6, 2020", "date_yearday", 127},
		{"May 6, 2020 4:50", "date_hour", 4},
		{"May 6, 2020 4:50", "date_minute", 50},
		{"May 6, 2020 4:50:20", "date_second", 20},
		{"May 6, 2020 4:50:20", "date_unix", 1588740620},
	}

	for _, test := range tests {
		assertQuery(t, fmt.Sprintf("SELECT %s('%s')", test.fn, test.in), fmt.Sprintf("%d", test.out))
		assertQuery(t, fmt.Sprintf("SELECT %s(' total garbage')", test.fn), "-1")
	}

	assertQuery(t, "SELECT date_rfc3339('May 6, 2020 4:50:20')", "2020-05-06T04:50:20Z")
	assertQuery(t, "SELECT date_rfc3339(' total garbage ')", "")
}
