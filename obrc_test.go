package obrc

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestComputeStats(t *testing.T) {
	cases := []struct {
		in   string
		want map[string]Stats
	}{
		{
			"a;1.1\nb;2.2\n",
			map[string]Stats{
				"a": {Min: 1.1, Max: 1.1, Sum: 1.1, Count: 1},
				"b": {Min: 2.2, Max: 2.2, Sum: 2.2, Count: 1},
			},
		},
		{
			"a;1.1\na;2.2\n",
			map[string]Stats{
				"a": {Min: 1.1, Max: 2.2, Sum: 3.3, Count: 2},
			},
		},
	}

	for _, tc := range cases {
		got, err := ComputeStats(strings.NewReader(tc.in))
		if err != nil {
			t.Errorf("%q: %v", tc.in, err)
		}
		if diff := cmp.Diff(tc.want, got, cmpopts.EquateApprox(0, 0.1)); diff != "" {
			t.Errorf("%q: mismatch (-want +got):\n%s", tc.in, diff)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	for _, n := range []int{0, 1, 1e3, 1e6} {
		var buf bytes.Buffer
		GenerateCSV(&buf, n)
		_, err := ComputeStats(strings.NewReader(buf.String()))
		if err != nil {
			t.Errorf("n=%d: %v", n, err)
		}
	}
}
