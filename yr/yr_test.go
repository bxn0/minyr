package yr_test

import (
	"testing"

	"github.com/bxn0/minyr/yr"
)



func TestCountLines(t *testing.T) {
	type test struct {
		input string
		want  int
	}

	tests := []test{
		{input: "kjevik-temp-celsius-20220318-20230318.csv", want: 16756},
	}

	for _, tc := range tests {
		got := yr.CountLines(tc.input)
		if got != tc.want {
			t.Errorf("%v: want %v, got %v,", tc.input, tc.want, got)
		}
	}
}



func TestConvertLines(t *testing.T) {

	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
			want: "The date might not be up to date"},
	}

	for _, tc := range tests {
		got := yr.ProcessLine(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestGetAverageTemperature(t *testing.T) {
	actualAvg, err := yr.GetAverageTemperature("kjevik-temp-celsius-20220318-20230318.csv", "celsius")
	if err != nil {
		t.Fatal(err)
	}

	expectedAvg := "8.56"
	if actualAvg != expectedAvg {
		t.Errorf("expected average temperature %v, but got %v", expectedAvg, actualAvg)
	}
}
