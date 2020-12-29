package function

import "testing"

func TestTypo(t *testing.T) {
	res := Typo(4, 3)
	if res != 5 {
		t.Error("ERROR")
	} else {
		t.Log("SUCCESS")
	}
}

func BenchmarkDuration(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Duration()
	}
}
