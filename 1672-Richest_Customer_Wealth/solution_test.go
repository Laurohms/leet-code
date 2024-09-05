package richestcustomerwealth_test

import (
	"testing"

	richestcustomerwealth "github.com/Laurohms/leet-code/1672-Richest_Customer_Wealth"
)

func TestRichestCustomerWealth(t *testing.T) {
	cases := []struct {
		desc string
		got  [][]int
		want int
	}{
		{"Test Case 1", [][]int{{1, 2, 3}, {3, 2, 1}}, 6},
		{"Test Case 2", [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17},
		{"Test Case 3", [][]int{{1, 5}, {7, 3}, {3, 5}}, 10},
	}

	for _, tt := range cases {
		t.Run(tt.desc, func(t *testing.T) {
			got := richestcustomerwealth.MaximumWealth(tt.got)

			if got != tt.want {
				t.Errorf("RichestCustomerWealth(%v) = %v; want %v", tt.got, got, tt.want)
			}
		})
	}

}
