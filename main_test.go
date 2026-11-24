package main

import (
	"reflect"
	"testing"
)

func TestMorrisInorder(t *testing.T) {
	root := tn(4, tn(2, tn(1, nil, nil), tn(3, nil, nil)), tn(6, tn(5, nil, nil), tn(7, nil, nil)))
	want := []int{1, 2, 3, 4, 5, 6, 7}
	if got := MorrisInorder(root); !reflect.DeepEqual(got, want) {
		t.Errorf("MorrisInorder = %v, want %v", got, want)
	}
	// Tree must be fully restored: all threads removed.
	if root.Left.Right.Right != nil || root.Left.Left.Right != nil {
		t.Error("temporary threads were not removed")
	}

	if got := MorrisInorder(nil); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("nil tree: %v", got)
	}
	chain := tn(3, tn(2, tn(1, nil, nil), nil), nil) // left-leaning
	if got := MorrisInorder(chain); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("left chain: %v", got)
	}
}
