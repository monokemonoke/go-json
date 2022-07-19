package myjson

import "testing"

func TestParserGetNum(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{
			input:  "42",
			expect: 42,
		},
		{
			input:  "31",
			expect: 31,
		},
	}

	for _, test := range tests {
		p := NewParser(test.input)
		actual := p.getNumber()
		if actual != test.expect {
			t.Errorf("Expect %d but got %d\n", test.expect, actual)
		}
	}
}
