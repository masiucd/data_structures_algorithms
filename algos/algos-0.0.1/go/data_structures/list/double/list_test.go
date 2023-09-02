package double

import "testing"

func TestPushFrontDoubleList(t *testing.T) {
	var list = NewList[int]()

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

func TestRemoveStartDoubleList(t *testing.T) {
	var list = NewList[int]()

	got := list.Len()
	expected := 0

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}

	list.PushFront(10)
	list.PushFront(20)
	list.PushFront(1)

	list.RemoveStart()

	got = list.Len()
	expected = 2

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}

	got = list.GetAt(0).KeyValue()
	expected = 20

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)

	}
}

func TestLen(t *testing.T) {
	var list = NewList[int]()

	got := list.Len()
	expected := 0

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}

	list.PushFront(10)
	list.PushFront(20)
	list.PushFront(30)
	got = list.Len()
	expected = 3

	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)

	}
}

func TestPushBack(t *testing.T) {
	var list = NewList[int]()
	list.PushBack(100)
	list.PushBack(200)
	list.PushBack(300)
	got := list.tail.key
	expected := 300
	if got != expected {
		t.Errorf("Expected %d got %d", expected, got)
	}
}
