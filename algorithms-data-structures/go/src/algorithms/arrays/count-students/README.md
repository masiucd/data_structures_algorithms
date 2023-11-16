# Number of Students Unable to Eat Lunch

The school cafeteria offers circular and square sandwiches at lunch break, referred to by numbers `0` and `1` respectively. All students stand in a queue. Each student either prefers square or circular sandwiches.

The number of sandwiches in the cafeteria is equal to the number of students. The sandwiches are placed in a **stack**. At each step:

- If the student at the front of the queue **prefers** the sandwich on the top of the stack, they will **take it** and leave the queue.

- Otherwise, they will **leave it** and go to the queue's end.

This continues until none of the queue students want to take the top sandwich and are thus unable to eat.

You are given two integer arrays `students` and `sandwiches` where `sandwiches[i]` is the type of the `i​​​​​​th` sandwich in the stack (`i = 0` is the top of the stack) and `students[j]` is the preference of the `j​​​​​​th` student in the initial queue (`j = 0` is the front of the queue). Return _the **number of students** that are **unable to eat**._



### Approach


1. Create an array of integers named counters with a length of 2. This array will keep track of the number of students who are unable to eat.
2. We receive two input arrays: students[] (containing 0s and 1s) and sandwiches[] (also containing 0s and 1s).
3. In the initial loop, tally the occurrences of 0s and 1s in the students[] array. Update the counters[] array accordingly. If the value is 0, increment counters[0]; if it's 1, increment counters[1].
4. In the second loop, traverse through the values in the sandwiches[] array. Decrease the corresponding counters[] value. If attempting to decrement a value that is already zero, break the loop since there are no more students desiring that specific sandwich. If there's no other student interested in that sandwich, the remaining students won't be able to eat.
5. Finally, return the sum of the remaining students from the counters[] array.