package window

import "testing"

func TestContains(t *testing.T) {
	for _, tt := range []struct {
		name               string
		first, last, value int
		want               bool
	}{
		{name: "lower bound", first: -3, last: 3, value: -3, want: true},
		{name: "interior", first: -3, last: 3, value: 0, want: true},
		{name: "below", first: -3, last: 3, value: -4},
		{name: "above", first: -3, last: 3, value: 4},
		{name: "reversed bounds", first: 3, last: -3, value: 0},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.first, tt.last, tt.value); got != tt.want {
				t.Errorf("Contains(%d, %d, %d) = %t, want %t", tt.first, tt.last, tt.value, got, tt.want)
			}
		})
	}
}

func TestContainsClosedInterval(t *testing.T) {
	if !Contains(2, 5, 5) {
		t.Error("closed interval must contain upper endpoint: Contains(2, 5, 5) = false, want true")
	}

	const maxInt = int(^uint(0) >> 1)
	const minInt = -maxInt - 1
	for _, tt := range []struct {
		first, last, value int
		want               bool
	}{
		{first: 5, last: 5, value: 5, want: true},
		{first: 5, last: 5, value: 4},
		{first: 5, last: 5, value: 6},
		{first: minInt, last: maxInt, value: minInt, want: true},
		{first: minInt, last: maxInt, value: maxInt, want: true},
		{first: maxInt, last: maxInt, value: maxInt, want: true},
		{first: maxInt, last: minInt, value: 0},
	} {
		if got := Contains(tt.first, tt.last, tt.value); got != tt.want {
			t.Errorf("Contains(%d, %d, %d) = %t, want %t", tt.first, tt.last, tt.value, got, tt.want)
		}
	}
}
