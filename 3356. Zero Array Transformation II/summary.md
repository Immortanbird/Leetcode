Using brute force will result in TLE.

The key is to find out a method to check if the queries applied are sufficient in linear time complexity.

Since the problem is about applying queries to a specific slice, we should think of prefix sum. However, the tougher question is about how to calculate prefix sum in linear time complexity. The idea is to use a difference array to record each query applied. The query value is added at the start index and minused at after the ending index. We can think it as adding the influence at the beginning and remove it at the end. The idea is quite brilliant.

Then we can iterate through the given nums and check if the applied queries can reduce the current num to zero. If not, we apply more queries.
