package dog

import (
"fmt"
"testing"
)

func TestYears(t *testing.T) {
	s:= 7
	ans:= 49
	s = Years(s)
	if ans != s{
		t.Error("got", s, "expected",ans)
	}
}

func TestYearsTwo(t *testing.T) {
	s:= 7
	ans:= 49
	s = YearsTwo(s)
	if ans != s{
		t.Error("got", s, "expected",ans)
	}
}

func ExampleYears() {
	s := 7
	xs := Years(s)
	fmt.Println(xs)
	// Output:
	// 49
}

func ExampleYearsTwo() {
	s := 7
	xs := YearsTwo(s)
	fmt.Println(xs)
	// Output:
	// 49
}

func BenchmarkYears(b *testing.B) {
	xs := 7
	b.ResetTimer()
	for i:= 0; i<b.N; i++ {
		Years(xs)
	}
}


func BenchmarkYearsTwo(b *testing.B) {
	xs := 7
	b.ResetTimer()
	for i:= 0; i<b.N; i++ {
		YearsTwo(xs)
	}
}
