package timerange

import (
	"errors"
	"fmt"
	"time"

	"github.com/hardeepnarang10/query/service/common/timestring"
)

func Validate(m map[string]string) error {
	tsstr, tsok := m["timestampStart"]
	testr, teok := m["timestampEnd"]

	if tsok != teok {
		return errors.New("one of timestampStart or timestampEnd is missing. Both must either be present or absent.")
	}

	if tsok && teok {
		ts, err := time.Parse(timestring.Layout, tsstr)
		if err != nil {
			return fmt.Errorf("unable to parse %q with layout %q: %w", "timestampStart", timestring.Layout, err)
		}
		te, err := time.Parse(timestring.Layout, testr)
		if err != nil {
			return fmt.Errorf("unable to parse %q with layout %q: %w", "timestampEnd", timestring.Layout, err)
		}

		if ts.Before(MinTimestampValue) || ts.After(MaxTimestampValue) {
			return fmt.Errorf("timestampStart must be after %q and before %q", MinTimestampValue.Format(timestring.Layout), MaxTimestampValue.Format(timestring.Layout))
		}
		if te.Before(MinTimestampValue) || te.After(MaxTimestampValue) {
			return fmt.Errorf("timestampEnd must be after %q and before %q", MinTimestampValue.Format(timestring.Layout), MaxTimestampValue.Format(timestring.Layout))
		}
	}
	return nil
}
