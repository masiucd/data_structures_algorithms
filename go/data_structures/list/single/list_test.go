package single

import "testing"

func TestPushFront(t *testing.T) {
	var list Listable[int] = NewList[int]()

	got := list.Len()
	expected := 0

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}

	list.PushFront(10)
	list.PushFront(20)
	list.PushFront(1)

	got = list.Len()
	expected = 3

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}

	got = list.GetAt(0).KeyValue()
	expected = 1

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)

	}
}

func TestPushBack(t *testing.T) {
	var list Listable[int] = NewList[int]()

	got := list.Len()
	expected := 0

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}

	list.PushBack(10)
	list.PushBack(20)
	list.PushBack(1)

	got = list.Len()
	expected = 3

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}

	got = list.GetAt(0).KeyValue()
	expected = 10

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}
}
