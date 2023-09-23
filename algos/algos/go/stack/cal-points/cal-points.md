# LeetCode Problem 682: Baseball Game

You are now a point recorder for a baseball game. Given a list of strings, each string can belong to one of the following four types:

1. Integer (represents the score for the current round): Directly indicates the number of points earned in this round.
2. "+" (represents the sum of the last two valid round's points): The points for this round are the sum of the last two valid rounds.
3. "D" (represents the double of the last valid round's points): The points for this round are twice the points of the last valid round.
4. "C" (an operation, not a round's score): Indicates that the last valid round's points should be removed as they were invalid.

Each round's operation is permanent and can affect the rounds before and after it. Your task is to calculate and return the sum of the points earned in all the rounds.

**Example 1:**
Input: `["5","2","C","D","+"]`
Output: 30
Explanation:

- Round 1: Earned 5 points. The sum is: 5.
- Round 2: Earned 2 points. The sum is: 7.
- Operation 1: Round 2's data was invalid, so the sum is: 5.
- Round 3: Earned 10 points (round 2's data removed). The sum is: 15.
- Round 4: Earned 5 + 10 = 15 points. The sum is: 30.

**Example 2:**
Input: `["5","-2","4","C","D","9","+","+"]`
Output: 27
Explanation:

- Round 1: Earned 5 points. The sum is: 5.
- Round 2: Earned -2 points. The sum is: 3.
- Round 3: Earned 4 points. The sum is: 7.
- Operation 1: Round 3's data is invalid, so the sum is: 3.
- Round 4: Earned -4 points (round 3's data removed). The sum is: -1.
- Round 5: Earned 9 points. The sum is: 8.
- Round 6: Earned -4 + 9 = 5 points. The sum is: 13.
- Round 7: Earned 9 + 5 = 14 points. The sum is: 27.

**Note:**

- The size of the input list will be between 1 and 1000.
- Every integer represented in the list will be between -30000 and 30000.
