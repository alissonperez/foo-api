package teardown

import (
	"github.com/alissonperez/api-foo/contrib/testutils"
	"testing"
)

func TestCancelMustCallExpectedFunctions(t *testing.T) {
	td := newTearDown()

	callsFirst := 0
	callsSecond := 0

	firstTestFunc := func() {
		callsFirst++
	}
	secondTestFunc := func() {
		callsSecond++
	}

	td.Register(&firstTestFunc)
	td.Register(&secondTestFunc)

	td.Cancel()

	testutils.Equals(t, 1, callsFirst)
	testutils.Equals(t, 1, callsSecond)
}
