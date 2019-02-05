package worker

import (
	"errors"
	"testing"
)

func TestResultString(t *testing.T) {
	url := "http://ya.ru"
	result := NewResult("http://ya.ru")

	expected := url + " d41d8cd98f00b204e9800998ecf8427e"
	actual := result.String()
	if expected != actual {
		t.Fatalf("Expected string representation %s, got %s", expected, actual)
	}
}

func TestResultStringWithError(t *testing.T) {
	url := "http://ya.ru"
	result := NewResult("http://ya.ru")
	result.Err = errors.New("something went wrong")

	expected := url + " Error while fetching url: " + result.Err.Error()
	actual := result.String()
	if expected != actual {
		t.Fatalf("Expected string representation %s, got %s", expected, actual)
	}
}
