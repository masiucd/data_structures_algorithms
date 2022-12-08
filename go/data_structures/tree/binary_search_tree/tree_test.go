package binary_search_tree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func makeBstTree(doRecursiveInsert bool) *Bst {
	if doRecursiveInsert {
		bst := NewBst()
		bst.InsertRec(10)
		bst.InsertRec(2)
		bst.InsertRec(1)
		bst.InsertRec(20)
		return bst
	} else {
		bst := NewBst()
		bst.Insert(10)
		bst.Insert(2)
		bst.Insert(1)
		bst.Insert(20)
		return bst
	}
}

func TestBst_Insert(t *testing.T) {
	bst := makeBstTree(false)

	want := 10
	got := bst.Root.Value
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	want = 2
	got = bst.Root.Left.Value
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	want = 1
	got = bst.Root.Left.Left.Value
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	want = 20
	got = bst.Root.Right.Value
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))
}

func TestBst_InsertRec(t *testing.T) {
	bst := makeBstTree(true)

	want := 10
	got := bst.Root.Value
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	want = 2
	got = bst.Root.Left.Value
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	want = 1
	got = bst.Root.Left.Left.Value
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	want = 20
	got = bst.Root.Right.Value
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))
}

func TestBst_Search(t *testing.T) {
	bst := makeBstTree(false)

	got := bst.Search(2).Value
	want := 2
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	got = bst.Search(20).Value
	want = 20
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	assert.Nil(t, bst.Search(200))
}

func TestBst_MaxValue(t *testing.T) {
	bst := makeBstTree(false)
	got := bst.MaxValue(bst.Root).Value
	want := 20
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	got = bst.MaxValue(bst.Root.Right).Value
	want = 20
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	got = bst.MaxValue(bst.Root.Left).Value
	want = 2
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))
}

func TestBst_MinValue(t *testing.T) {
	bst := makeBstTree(false)
	got := bst.MinValue(bst.Root).Value
	want := 1
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	got = bst.MinValue(bst.Root.Left).Value
	want = 1
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))

	got = bst.MinValue(bst.Root.Right).Value
	want = 20
	assert.Equal(t, got, want, fmt.Sprintf("Expected %d  got %d", want, got))
}

func TestBst_Delete(t *testing.T) {
	bst := makeBstTree(false)
	bst.Remove(bst.Root, 2)
	expected := 1
	assert.Equal(t, bst.Root.Left.Value, expected, fmt.Sprintf("Expected %d got %d", expected, bst.Root.Left.Value))
}

func TestBst_inorder(t *testing.T) {
	bst := makeBstTree(false)
	got := bst.Inorder()
	want := 4
	assert.Equal(t, len(got), want, fmt.Sprintf("Expected %d got %d", want, got))

	want = 1
	assert.Equal(t, got[0], want, fmt.Sprintf("Expected %d got %d", want, got))

	want = 2
	assert.Equal(t, got[1], want, fmt.Sprintf("Expected %d got %d", want, got))
}

func TestBst_preorder(t *testing.T) {
	bst := makeBstTree(false)
	got := bst.Preorder()

	want := 4
	assert.Equal(t, len(got), want, fmt.Sprintf("Expected %d got %d", want, got))

	want = 10
	assert.Equal(t, got[0], want, fmt.Sprintf("Expected %d got %d", want, got))

	want = 2
	assert.Equal(t, got[1], want, fmt.Sprintf("Expected %d got %d", want, got))
}

func TestBst_postorder(t *testing.T) {
	bst := makeBstTree(false)
	got := bst.Postorder()
	fmt.Println(got)

	want := 4
	assert.Equal(t, len(got), want, fmt.Sprintf("Expected %d got %d", want, got))

	want = 1
	assert.Equal(t, got[0], want, fmt.Sprintf("Expected %d got %d", want, got))

	want = 2
	assert.Equal(t, got[1], want, fmt.Sprintf("Expected %d got %d", want, got))
}

func TestBst_BreadthFirstSearch(t *testing.T) {
	bst := makeBstTree(false)
	got := bst.BreadthFirstSearch()
	want := 4
	assert.Equal(t, len(got), want, fmt.Sprintf("Expected %d got %d", want, got))
	expectedResult := []int{10, 2, 20, 1}
	for i := range got {
		assert.Equal(t, expectedResult[i], got[i], fmt.Sprintf("Expected %d got %d", expectedResult[i], got[i]))
	}
}
