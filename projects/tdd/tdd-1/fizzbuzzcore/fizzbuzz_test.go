package fizzbuzzcore

import "testing"

func TestFizzBuzz_When1_Returns1(t *testing.T)  {
	// Arrange
	input := 1
	// Act
	output := GetValue(input)
	// Assert
	if output != "1" {
		t.Error("input != output")
	}
}

func TestFizzBuzz_When2_Returns2(t *testing.T) {
	input := 2
	output := GetValue(input)
	if output != "2" {
		t.Error("output != 2")
	}
}

func TestFizzBuzz_When3_ReturnsFizz(t *testing.T) {
	input := 3
	output := GetValue(input)
	if output != "fizz" {
		t.Error("output != fizz")
	}
}

func TestFizzBuzz_When5_ReturnsBuzz(t *testing.T) {
	input := 5
	output := GetValue(input)
	if output != "buzz" {
		t.Error("output != buzz")
	}
}

func TestFizzBuzz_WhenDiv3_ReturnsFizz(t *testing.T) {
	got := GetValue(6)
	want := "fizz"
	if got != want {
		t.Errorf("got %s, wanted %s", got,want)
	}
}

// Writing table-driven tests in Go
type fizzbuzzTest struct {
	input int
	expected string
}

var fizzTests = []fizzbuzzTest{
	fizzbuzzTest{3,"fizz"},
	fizzbuzzTest{6,"fizz"},
	fizzbuzzTest{9,"fizz"},
	fizzbuzzTest{12,"fizz"},
	fizzbuzzTest{18,"fizz"},
}

func TestFizzBuzz_WhenDiv3_ReturnsFizz_TableDriven(t *testing.T) {
	for _, test := range fizzTests {
		if output := GetValue(test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

var buzzTests = []fizzbuzzTest{
	fizzbuzzTest{5,"buzz"},
	fizzbuzzTest{10,"buzz"},
}

func TestFizzBuzz_WhenDiv5_ReturnsBuzz(t *testing.T) {
	for _, test := range buzzTests {
		if output := GetValue(test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

var dafaultTests = []fizzbuzzTest{
	fizzbuzzTest{1,"1"},
	fizzbuzzTest{2,"2"},
	fizzbuzzTest{4,"4"},
	fizzbuzzTest{7,"7"},
	fizzbuzzTest{8,"8"},
	fizzbuzzTest{11,"11"},
	fizzbuzzTest{13,"13"},
	fizzbuzzTest{14,"14"},
	fizzbuzzTest{16,"16"},
	fizzbuzzTest{17,"17"},
	fizzbuzzTest{19,"19"},
}
func TestFizzBuzz_WhenDefaults_ReturnsInput(t *testing.T) {
	for _, test := range dafaultTests {
		if output := GetValue(test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

var fizzbuzzTests = []fizzbuzzTest{
	fizzbuzzTest{15,"fizzbuzz"},
}

func TestFizzBuzz_WhenDiv3AndDiv5_ReturnsFizzbuzz(t *testing.T) {
	for _, test := range fizzbuzzTests {
		if output := GetValue(test.input); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}