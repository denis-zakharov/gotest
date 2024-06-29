package main

import "testing"

func TestFooer(t *testing.T) {
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("result was incorrect, got %s, want: %s", result, "Foo")
	}
}

func TestFooerTableDriven(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		{"9 should be Foo", 9, "Foo"},
		{"3 should be Foo", 3, "Foo"},
		{"1 is not Foo", 1, "1"},
		{"0 should be Foo", 0, "Foo"},
	}

	t.Cleanup(func() {
		t.Log("Table Driven Cleanup")
	})

	for _, tt := range tests {
		// define a subtest with a name
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {
				t.Logf("subtest \"%s\" cleanup\n", tt.name)
			})
			ans := Fooer(tt.input)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestFooerParallel(t *testing.T) {
	t.Run("Test 3 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(3)
		if result != "Foo" {
			t.Errorf("result was incorrect, got: %s, want: %s", result, "Foo")
		}
	})
	t.Run("Test 7 in Parallel", func(t *testing.T) {
		t.Parallel()
		result := Fooer(7)
		if result != "7" {
			t.Errorf("result was incorrect, got: %s, want: %s", result, "7")
		}
	})
}

func TestFooerSkipped(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}
	result := Fooer(3)
	if result != "Foo" {
		t.Errorf("result was incorrect, got: %s, want: %s", result, "Foo")
	}
}
