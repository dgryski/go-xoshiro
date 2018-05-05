package xoshiro

import "testing"

func TestNext(t *testing.T) {

	// generated from reference C code
	reference := []uint64{
		11520,
		0,
		1509978240,
		1215971899390074240,
		1216172134540287360,
		607988272756665600,
		16172922978634559625,
		8476171486693032832,
		10595114339597558777,
		2904607092377533576,
	}

	s := State{1, 2, 3, 4}

	for i, want := range reference {
		got := s.Next()
		if got != want {
			t.Errorf("Next() mismatch at offset %v: got %v, want %v", i, got, want)
		}
	}
}

func TestJump(t *testing.T) {

	// generated from reference C code
	reference := []uint64{
		13534147089533256664,
		5049895679018676702,
		10357798701826169776,
		14797244810105073478,
		18137915528557693442,
		8136292926577591182,
		3941356121782675177,
		545106238668810688,
		11541040995294198212,
		4705590444263342743,
	}

	s := State{1, 2, 3, 4}

	for i, want := range reference {
		s.Jump()
		got := s.Next()
		if got != want {
			t.Errorf("Jump()/Next() mismatch at offset %v: got %v, want %v", i, got, want)
		}
	}
}
