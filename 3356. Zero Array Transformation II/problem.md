# 3356. Zero Array Transformation II

You are given an integer array nums of length n and a 2D array queries where $queries[i] = [l_i, r_i, val_i]$.

Each queries[i] represents the following action on nums:

- Decrement the value at each index in the range $[l_i, r_i]$ in nums by at most vali.
- The amount by which each value is decremented can be chosen independently for each index.

A Zero Array is an array with all its elements equal to 0.

Return the minimum possible non-negative value of k, such that after processing the first k queries in sequence, nums becomes a Zero Array. If no such k exists, return -1.

**Constraints:**

- $1 \le nums.length \le 10^5$
- $0 \le nums[i] \le 5 * 10^5$
- $1 \le queries.length \le 10^5$
- $queries[i].length == 3$
- $0 \le l_i \le r_i < nums.length$
- $1 \le val_i \le 5$