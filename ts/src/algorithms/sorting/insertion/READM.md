# Insertion sort

## Introduction

Insertion sort is a sorting algorithm that builds a sorted array (or list) one element at a time by comparing each new element to the already-sorted elements and inserting the new element into the correct location.

## Mental model

Imagine you have a deck of cards and you want to sort them. You start with an empty left hand and the cards face down on the table. You then remove one card at a time from the table and insert it into the correct position in the left hand. To find the correct position for a card, you compare it with each of the cards already in the hand, from right to left. At all times, the cards held in the left hand are sorted, and these cards were originally the top cards of the pile on the table.

## Algorithm

1. Iterate from `arr[1]` to `arr[n]` over the array.
2. Compare the current element (`key`) to its predecessor.
3. If the key element is smaller than its predecessor, compare it to the elements before. Move the greater elements one position up to make space for the swapped element.

## Complexity

- **Time complexity**: `O(n^2)` - since in worst case we have to do `n` comparisons for `n` elements.
- **Space complexity**: `O(1)` - since we only use constant extra space.

## References

- [Wikipedia](https://en.wikipedia.org/wiki/Insertion_sort)
