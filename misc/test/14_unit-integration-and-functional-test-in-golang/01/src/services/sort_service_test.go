package services

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	got := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	Sort(got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Arrays sould be equal. want: %v, got: %v", want, got)
	}
}
