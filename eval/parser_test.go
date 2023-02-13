package eval

import (
	"container/list"
	"fmt"
	"testing"
)

func TestInfixToPostFix(t *testing.T) {
	lst1 := list.New()
	lst1.PushBack("1")
	lst1.PushBack("2")
	lst1.PushBack("+")

	lst2 := list.New()
	lst2.PushBack("1")
	lst2.PushBack("2")
	lst2.PushBack("+")
	lst2.PushBack("3")
	lst2.PushBack("*")

	lst3 := list.New()
	lst3.PushBack("1")
	lst3.PushBack("2")
	lst3.PushBack("3")
	lst3.PushBack("*")
	lst3.PushBack("+")

	testCases := []struct {
		input    string
		expected *list.List
		err      error
	}{
		{"1 + 2", lst1, nil},
		{"( 1 + 2 ) * 3", lst2, nil},
		{"1 + 2 * 3", lst3, nil},
		{"", nil, fmt.Errorf("error empty input")},
		{"1 + 2 +10", nil, fmt.Errorf("error invalid input, numbers can be between 0-9")},
	}

	for _, tc := range testCases {
		actual, err := InfixToPostFix(tc.input)
		if tc.err != nil {
			if err == nil {
				t.Errorf("Expected an error, but got nil")
			} else if err.Error() != tc.err.Error() {
				t.Errorf("Expected error: %s, but got %s", tc.err, err)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
			continue
		}

		if actual.Len() != tc.expected.Len() {
			t.Errorf("Expected length %d, but got %d", tc.expected.Len(), actual.Len())
			continue
		}

		for e1, e2 := actual.Front(), tc.expected.Front(); e1 != nil; e1, e2 = e1.Next(), e2.Next() {
			if e1.Value.(string) != e2.Value.(string) {
				t.Errorf("Expected value %s, but got %s", e2.Value.(string), e1.Value.(string))
				break
			}
		}
	}
}
