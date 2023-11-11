# Sliding window

## What is it?

A sliding window is a sub-list that runs over an underlying collection. The window is usually defined by two indices, `i` and `j`, where `i` is the start of the window and `j` is the end of the window. In other words, the window is the sub-list `arr[i:j]`. The window "slides" over the underlying collection, and the two indices change as the window moves.

## Why use it?

Sliding windows are used to solve problems where you need to find the longest/shortest substring, subarray, or a desired value that fulfills certain conditions. The window is used to keep track of the current substring, subarray, or value, and the two indices are used to keep track of the start and end of the window. The window is moved as the algorithm iterates through the collection, and the two indices are updated accordingly.
