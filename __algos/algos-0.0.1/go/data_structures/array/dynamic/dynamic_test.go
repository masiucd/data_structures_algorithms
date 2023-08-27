package dynamic

import (
	"testing"
)

func insertHelper() *Array {
	arr := newArray(3)
	arr.push(20)
	arr.push(10)
	arr.push(2)

	return arr
}

func TestPush(t *testing.T) {
	capacity := 3
	arr := newArray(capacity)

	arr.push(20)
	s := arr.size
	last := arr.data[s-1]
	want := 1

	if s != want {
		t.Errorf("Got = %v want = %v ", s, want)
	}

	want = 20
	if last != want {
		t.Errorf("Got = %v want = %v ", last, want)
	}

	arr.push(30)
	s = arr.size
	last = arr.data[s-1]
	want = 2

	if s != want {
		t.Errorf("Got = %v want = %v ", s, want)
	}

	want = 30
	if last != want {
		t.Errorf("Got = %v want = %v ", last, want)
	}

	arr.push(20)
	arr.push(30)
	arr.push(40)
	want = 40
	last = arr.data[arr.size-1]

	if last != want {
		t.Errorf("Got = %v want = %v ", last, want)
	}

	capacity = arr.capacity
	want = 6
	if capacity != want {
		t.Errorf("Got = %v want = %v ", last, want)

	}

}

func TestPop(t *testing.T) {
	arr := insertHelper()

	size := arr.size
	want := 3
	if size != want {
		t.Errorf("Got = %v want = %v ", size, want)
	}

	popped := arr.pop()
	size = arr.size
	want = 2
	if size != want {
		t.Errorf("Got = %v want = %v ", size, want)
	}

	want = 2
	if popped != want {
		t.Errorf("Got = %v want = %v ", popped, want)
	}
}

func TestResize(t *testing.T) {
	arr := insertHelper()

	got := arr.capacity
	want := 3
	if got != want {
		t.Errorf("Got = %v want = %v ", got, want)
	}

	arr.resize()

	got = arr.capacity
	want = 6
	if got != want {
		t.Errorf("Got = %v want = %v ", got, want)
	}

}

func TestGet(t *testing.T) {
	arr := insertHelper()

	got := arr.get(1)
	want := 10

	if got != want {
		t.Errorf("Got = %v want = %v ", got, want)
	}

	got = arr.get(100)
	want = -1

	if got != want {
		t.Errorf("Got = %v want = %v ", got, want)
	}
}
