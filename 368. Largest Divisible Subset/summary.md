To simplify the condition that "answer[i] % answer[j] == 0, or answer[j] % answer[i] == 0", it is intuitive to think of sorting.

Then, we can use dynamic programming. We simply consider the longest length of the subsequence ending with the given index.