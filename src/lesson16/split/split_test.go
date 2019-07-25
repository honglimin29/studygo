package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("期望:%v,实际:%v", want, got)
	}
}

func TestMoreSplit(t *testing.T) {
	got := Split("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("期望:%v,实际:%v", want, got)
	}
}

func TestSplit2(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]test{
		"test1": test{input: "a:b:c:d", sep: ":", want: []string{"a", "b", "c", "d"}},
		"test2": test{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"test3": test{input: "沙河又沙河", sep: "沙", want: []string{"河又", "河"}},
	}

	for _, test := range tests {
		ret := Split(test.input, test.sep)
		if !reflect.DeepEqual(ret, test.want) {
			t.Errorf("期望:%#v,实际:%#v", test.want, ret)
		}
	}
}
