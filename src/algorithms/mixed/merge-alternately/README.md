# Merge alternately

Given two strings `word1` and `word2`, merge the strings by adding letters in alternating order, starting with `word1`. If a string is longer than the other, append the additional letters onto the end of the merged string.

## Example 1:

```
Input: word1 = "abc", word2 = "pqr"
Output: "apbqcr"
Explanation: The merged string will be merged as so:
word1:  a   b   c
word2:    p   q   r
merged: a p b q c r
```

## Example 2:

```
Input: word1 = "ab", word2 = "pqrs"
Output: "apbqrs"
Explanation: The merged string will be merged as so:
word1:  a   b
word2:    p   q   r   s
merged: a p b q r s
```

## Example 3:

```
Input: word1 = "abcd", word2 = "pq"
Output: "apbqcd"
Explanation: The merged string will be merged as so:
word1:  a   b   c   d
word2:    p   q
merged: a p b q c d
```
