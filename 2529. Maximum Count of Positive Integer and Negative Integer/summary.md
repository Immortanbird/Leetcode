The brute force solution has O(n) time complexity.

Use the non-descending condition. Thus, the array is divided by into three possible sectors: negative integers, zeroes, positive integers.

We may come to a solution that we can use the index of the first zero and the last zero.

Use binary search. Watch for the boundary conditions. When looking for the lower boundary, shrink the right index if the value is less than or equal to zero. When looking for the upper boundary, shrink the left index if the value is greater than or equal to zero.