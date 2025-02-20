package helper_test

import (
	"testing"

	helper "github.com/umarali-nagoor/go-unit-test-coverage-tool"
)

func TestDate(t *testing.T) {

	// Input Data
	tcs := []struct {
		team, entity int
		out          string
		err          error
	}{
		{1, 1, "Welcome to Workspace of Schematics", nil},
		{1, 2, "Welcome to Actions of Schematics", nil},
		{2, 6, "Welcome to events of Event Notification", nil},
		//{3, 5, "Welcome to config of AppConfig", nil},
	}

	for _, tc := range tcs {
		msg, err := helper.Org(tc.team, tc.entity)
		if err != tc.err {
			t.Errorf("expected error to be %v got %v", tc.err, err)
		}
		if msg != tc.out {
			t.Errorf("expected %s but got %s", tc.out, msg)
		}
	}
}
