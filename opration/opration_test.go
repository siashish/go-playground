package opration

import (
	"fmt"
	"log"
	"os"
	"testing"
	"unicode/utf8"
)

type oprationTest struct {
	arg1, arg2, want int
}

var multplyTests = []oprationTest{
	oprationTest{2, 3, 6},
	oprationTest{3, 4, 12},
	oprationTest{5, 4, 20},
	oprationTest{7, 3, 21},
}

var addTests = []oprationTest{
	oprationTest{2, 3, 5},
	oprationTest{3, 4, 7},
	oprationTest{5, 4, 9},
	oprationTest{7, 3, 10},
}

func TestMultply(t *testing.T) {
	// got := Multply(10, 20)

	// want := 200
	// if want != got {
	// 	t.Errorf("expected %d and got %d", want, got)
	// }
	for _, test := range multplyTests {
		if output := Multply(test.arg1, test.arg2); output != test.want {
			t.Errorf("want %d; but got %d", test.want, output)
		}
	}
}

func TestAdd(t *testing.T) {
	// got := Add(10, 20)
	// want := 30
	// if want != got {
	// 	t.Errorf("expected %d and got %d", want, got)
	// }
	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.want {
			t.Errorf("expected %d; but got %d", test.want, output)
		}
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(20, 10)
	want := 10
	if want != got {
		t.Errorf("expected %d and got %d", want, got)
	}
}

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello"
	if got != want {
		t.Errorf("expected %s and got %s", want, got)
	}
}

func BenchmarkMultply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multply(100, 999)
	}
}

func ExampleMultply() {
	fmt.Println(Multply(10, 20))
	// Output: 200
}

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

func TestMain(m *testing.M) {
	log.Println("DO STUFF BEFORE TEST")
	osexit := m.Run()
	log.Println("DO STUFF AFTER TEST")
	os.Exit(osexit)
}
