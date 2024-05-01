package go_hello_tool

import "testing"

func TestHello(t *testing.T) {
	var hello string = "Hello, World!";
	var notHello string = "Not Hello, World!";

	result := Hello();

	if result != hello {
		t.Errorf("Hello() = %s; want %s", result, hello);
	}
	if result == notHello {
		t.Errorf("Hello() = %s; want not %s", result, notHello);
	}
}