package main

import "testing"

func TestHello(t *testing.T) {
	result := hello("山澤さん")
	expext := "Hello, 山澤さん"
	if result != expext {
		t.Error("\n実際： ", result, "\n理想： ", expext)
	}

	t.Log("TestHello終了")
}
