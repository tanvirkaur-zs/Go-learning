package employeesql

import (
	"errors"
	"reflect"
	"testing"
)

func TestPost(t *testing.T) {
	successEmployee := employee{1, "tanvir", "sde", 1, 30000}
	testcase := []struct {
		desc   string
		input  employee
		output employee
		err    error
	}{
		{"Valid", successEmployee, successEmployee, nil},
		{"Id can't be zero", employee{0, "tanvir", "sde", 1, 30000}, employee{}, errors.New("id should be 1 and above")},
		{"Id can't be -ve", employee{-1, "tanvir", "sde", 1, 30000}, employee{}, errors.New("id should be 1 and above")},
		{"Name can't be empty", employee{1, "", "sde", 1, 30000}, employee{}, errors.New("name can't be empty")},
	}

	for i, tc := range testcase {
		ans, err := insert(tc.input)
		if ans != tc.output {
			t.Errorf("failure %d", i)
		}
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("failure %d", i)
		}

	}
}

func TestGet(t *testing.T) {
	testcase := []struct {
		desc   string
		input  int
		output employee
		expErr error
	}{
		{"Valid", 1, employee{1, "tanvir", "sde", 1, 30000}, nil},
		{"Id can't be zero", 0, employee{}, errors.New("id should be 1 and above")},
		{"Id can't be -ve", -1, employee{}, errors.New("id should be 1 and above")},
	}

	for i, tc := range testcase {
		ans, err := fetch(tc.input)
		if ans != tc.output {
			t.Errorf("failure %d, DESC %v ,expected output %v, actual output %v", i, tc.desc, tc.output, ans)
		}

		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("failure %d", i)
		}
	}
}
