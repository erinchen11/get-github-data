package getgithubdata_test

import (
	"testing"

	gt "github.com/get-github-data"
)

func TestGetAverageStar(t *testing.T) {

	// // input test data

	// var tests = []struct{
	// 	name string
	// 	want float64
	// }{
	// 	// true data
	// 	{"tj",489.8},
	// 	{"erinchen", 0},
	// }

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := GetAverageStar(tt.name)
	// 		if err != nil {
	// 			t.Errorf("Should be nil error, but got %v ", err)
	// 		}

	// 		if got != tt.want {
	// 			t.Errorf("Should be %f, got %f", tt.want, got)
	// 		}
	// 	})
	// }

	// use Mock to test API
	// 模擬資料
	var tests = []struct {
		name string
		want float64
	}{}

	mock := new(gt.Mock)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gt.GetAverageStar(mock, tt.name)
			if err != nil {
				t.Errorf("Should be nil error, but got %v ", err)
			}

			if got != tt.want {
				t.Errorf("Should be %f, got %f", tt.want, got)
			}
		})
	}

}
