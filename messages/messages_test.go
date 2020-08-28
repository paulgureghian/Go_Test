package messages

import "testing"

func TestGreet(t *testing.T) {

	got := Greet("Gopher")
	expected := "Hello, Gopher \n"

	if got != expected {

		t.Errorf("Did not get expected result. Expected %q, got: %q \n", expected, got)
	}

}

func TestGreetTableDriven(t *testing.T) {

	scenarios := []struct {
		input    string
		expected string
	}{
		{input: "Gopher", expected: "Hello, Gopher \n"},
		{input: "", expected: "Hello, \n"},
	}

	for _, s := range scenarios {

		got := Greet(s.input)

		if got != s.expected {
			t.Errorf("Did not get expected result for input '%v'. Expected %q, got %q",
				s.input, got, s.expected)
		}
	}
}

func TestDepart(t *testing.T) {

	got := depart("Gopher")
	expected := "Goodbye, Gopher \n"

	if got != expected {

		t.Errorf("Did not get expected result. Expected %q, got: %q \n", expected, got)
	}
}

func TestFailureTypes(t *testing.T) {

	t.Error("Error signals a failed test, but the program keeps running")
	t.Fatal("Fatal will fail the test and stop its execution")
	t.Error("This will never print since a failure came right before it")
}
