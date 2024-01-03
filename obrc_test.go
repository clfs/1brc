package obrc

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestTakeRecordings(t *testing.T) {
	cases := []struct {
		in   string
		want map[string]Recording
	}{
		{
			"a;1.1\nb;2.2\n",
			map[string]Recording{
				"a": {min: 1.1, max: 1.1, sum: 1.1, count: 1},
				"b": {min: 2.2, max: 2.2, sum: 2.2, count: 1},
			},
		},
		{
			"a;1.1\na;2.2\n",
			map[string]Recording{
				"a": {min: 1.1, max: 2.2, sum: 3.3, count: 2},
			},
		},
	}

	for _, tc := range cases {
		got, err := TakeRecordings(NewReader(strings.NewReader(tc.in)))
		if err != nil {
			t.Errorf("%q: %v", tc.in, err)
		}
		if diff := cmp.Diff(tc.want, got, cmpopts.EquateApprox(0, 0.1), cmp.AllowUnexported(Recording{})); diff != "" {
			t.Errorf("%q: mismatch (-want +got):\n%s", tc.in, diff)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	for _, n := range []int{0, 1, 1e3, 1e6} {
		t.Run(fmt.Sprintf("%d", n), func(t *testing.T) {
			var buf bytes.Buffer
			if err := Generate(NewWriter(&buf), n); err != nil {
				t.Fatalf("Generate: %v", err)
			}
			if _, err := TakeRecordings(NewReader(&buf)); err != nil {
				t.Errorf("TakeRecordings: %v", err)
			}
		})
	}
}
