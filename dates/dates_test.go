package dates_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/bir/iken/dates"
)

func nowFunc(now time.Time) func() time.Time {
	return func() time.Time {
		return now
	}
}

func TestToTime(t *testing.T) {

	la, _ := time.LoadLocation("America/Los_Angeles")

	tests := []struct {
		name     string
		now      time.Time
		duration string
		location *time.Location
		want     time.Time
		wantErr  bool
	}{
		{
			name:     "72h",
			now:      time.Date(2020, 1, 1, 11, 11, 11, 0, time.UTC),
			duration: "72h",
			location: nil,
			want:     time.Date(2020, 1, 4, 11, 11, 11, 0, time.UTC),
			wantErr:  false,
		}, {
			name:     "EOD In LA ",
			now:      time.Date(2020, 1, 1, 2, 11, 11, 0, time.UTC),
			duration: "EOD",
			location: la,
			want:     time.Date(2019, 12, 31, 23, 59, 0, 0, la),
			wantErr:  false,
		}, {
			name:     "EOD",
			now:      time.Date(2020, 1, 1, 11, 11, 11, 0, time.UTC),
			duration: "EOD",
			location: nil,
			want:     time.Date(2020, 1, 1, 23, 59, 0, 0, time.UTC),
			wantErr:  false,
		}, {
			name:     "EOY",
			now:      time.Date(2020, 1, 1, 11, 11, 11, 0, time.UTC),
			duration: "EOY",
			location: nil,
			want:     time.Date(2020, 12, 31, 23, 59, 0, 0, time.UTC),
			wantErr:  false,
		}, {
			name:     "EOD+24h",
			now:      time.Date(2020, 1, 1, 11, 11, 11, 0, time.UTC),
			duration: "EOD+24h",
			location: nil,
			want:     time.Date(2020, 1, 2, 23, 59, 0, 0, time.UTC),
			wantErr:  false,
		}, {
			name:     "Error",
			now:      time.Date(2020, 1, 1, 11, 11, 11, 0, time.UTC),
			duration: "EOD+InvalidDuration",
			location: nil,
			want:     time.Time{},
			wantErr:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dates.NowFunc = nowFunc(test.now)

			got, err := dates.ToTime(test.duration, test.location)
			if (err != nil) != test.wantErr {
				t.Errorf("ToTime() error = %v, wantErr %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("ToTime() got = %v, want %v", got, test.want)
			}
		})
	}
}
