# Max Product

### About

Given a list of integers, find the starting index of the subsequence of length `n` that has the largest product. The list will have at least `n` integers.

### Level 1

In Level 1 n is always 2. So for example, given the list [2, 5, 8, 0, 9, 5, 3, 7], the largest product of 2 consecutive numbers is 9 \* 5 = 45. In this case, we should return 4 as this is the starting index (the index of the 9 in the list).

### Level 2

In Level 2 you get the list of integers and n as inputs. For example, given the list [3, 5, 8, 1, 6, 4, 6, 8, 3, 8, 1] and n=3, the largest product of 3 consecutive numbers is 4 _ 6 _ 8 = 192. In this case, we should return 5 as this is the starting index (the index of the 4 in the list). However, index 7 is also a valid solution as 8 _ 3 _ 8 also gives the product 192.
