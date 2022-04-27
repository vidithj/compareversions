package api

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCompareVersions(t *testing.T) {

	tests := []struct {
		name            string
		requestversion1 string
		requestversion2 string
		response        int
	}{
		{
			name:            "valid input,v1 is greater than v2",
			requestversion1: "0.2.1.1.1",
			requestversion2: "0.1.2",
			response:        1,
		},
		{
			name:            "valid input,v2 is greater than v1",
			requestversion1: "0.2.1.1.1",
			requestversion2: "1.1.2",
			response:        -1,
		},
		{
			name:            "valid input,v2 is same as v1",
			requestversion1: "0.1.1.1.1",
			requestversion2: "0.1.1.1.1",
			response:        0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := CompareVersions(test.requestversion1, test.requestversion2)
			if diff := cmp.Diff(response, test.response); diff != "" {
				t.Errorf("%s failed - differs: (-want +got)\n %s", test.name, diff)
			}
		})
	}

}

func TestMatchSize(t *testing.T) {

	tests := []struct {
		name             string
		version1         []string
		version2         []string
		responseversion1 []string
		responseversion2 []string
	}{
		{
			name:             "valid input,length v1 is greater than v2",
			version1:         []string{"0", "2", "1", "1"},
			version2:         []string{"0", "1", "1"},
			responseversion1: []string{"0", "2", "1", "1"},
			responseversion2: []string{"0", "1", "1", "0"},
		},
		{
			name:             "valid input,length of v2 is greater than v1",
			version1:         []string{"0", "2", "1"},
			version2:         []string{"1", "1", "1", "2"},
			responseversion1: []string{"0", "2", "1", "0"},
			responseversion2: []string{"1", "1", "1", "2"},
		},
		{
			name:             "valid input,length of v2 is same as v1",
			version1:         []string{"0", "2", "1", "1"},
			version2:         []string{"1", "1", "1", "2"},
			responseversion1: []string{"0", "2", "1", "1"},
			responseversion2: []string{"1", "1", "1", "2"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			responsev1, responsev2 := matchSize(test.version1, test.version2)
			if diff := cmp.Diff(responsev1, test.responseversion1); diff != "" {
				if diff := cmp.Diff(responsev2, test.responseversion2); diff != "" {
					t.Errorf("%s failed - differs: (-want +got)\n %s", test.name, diff)
				}
			}
		})
	}

}
