If sum of gas is equal or greater than sum of cost, then there is a solution. Otherwise, no solution.

The solution is guranteed to be unique, so if there is a solution, we can certainly go from the start point to the end of the array (because it says clockwise, if we cant even go to the end of the array, how can we go clockwise).

So, just find the start point that gets us to the end and calculate the sums meanwhile.