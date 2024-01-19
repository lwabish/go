package structs

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestIsValidCPUCores(t *testing.T) {
	casesRaw := [][]interface{}{
		{1, 2, 3},
		{"1-2", "3-4"},
		{"1-2", 3},
		{},
		{"1.2"},
		{"1-2", 3.4},
	}
	cases := make([]string, 0)
	for _, c := range casesRaw {
		j, err := json.Marshal(c)
		if err != nil {
			t.Fatalf("can't marshal to json: %v", err)
		}
		cases = append(cases, string(j))
	}

	wanted := []interface {
	}{
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3},
		nil,
		nil,
		nil,
	}
	for i, s := range cases {
		var c []interface{}
		t.Logf("case %d (str): %s", i, s)
		if err := json.Unmarshal([]byte(s), &c); err != nil {
			t.Fatalf("failed to unmarshal json %v", err)
		}
		got, err := IsValidCPUCores(c)
		if err != nil && wanted[i] == nil {
			t.Logf("case %d: passed", i)
			continue
		}
		if !reflect.DeepEqual(got, wanted[i]) {
			t.Fatalf("wanted %v, got %v", wanted[i], got)
		}
		t.Logf("case %d: passed", i)
	}
}
