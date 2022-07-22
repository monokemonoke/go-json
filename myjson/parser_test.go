package myjson

import "testing"

func TestParseSymbol(t *testing.T) {
	cases := []struct {
		input        string
		expectSymbol string
		expectErr    error
	}{
		{
			input:        "symbol",
			expectSymbol: "symbol",
			expectErr:    nil,
		},
		{
			input:        "symbol ",
			expectSymbol: "symbol",
			expectErr:    nil,
		},
		{
			input:        "   symbol",
			expectSymbol: "symbol",
			expectErr:    nil,
		},
		{
			input:        "symbol123",
			expectSymbol: "symbol123",
			expectErr:    nil,
		},
		{
			input:        "あかさたな",
			expectSymbol: "あかさたな",
			expectErr:    nil,
		},
	}

	for _, tc := range cases {
		p := NewParser(tc.input)
		actualSymbol, actualErr := p.parseSymbol()

		if actualErr != tc.expectErr {
			t.Errorf("want error `%s` but got `%s`\n", tc.expectErr, actualErr)
		}

		if actualSymbol != tc.expectSymbol {
			t.Errorf("want symbol `%s` but got `%s`\n", tc.expectSymbol, actualSymbol)
		}
	}
}
