# Two Number Sum

## Problem

Is there a pair of numbers in the input array that sum up to the target sum?

## Solution

One possible solution could be to implement with two nested for loops and compare each number to every other number, but
this results in a time complexity of O(n^2).

We can improve this by iterating over the input array only once and using a hash map to store the values we have already
traversed.


## Time Complexity

O(n)

The time complexity is linear. The larger the input array n, the longer our algorithm will run in the worst case.

## Space Complexity


O(n)

The space complexity is linear. The larger the input array n, the more memory our algorithm will use storing previously
traversed values in the hashmap.

## Shane's takeaway

The thing I find most interesting about this algorithm is the solution can be found by iterating over the array just a
single time. At each iteration we need to know what the value is that we are trying to find, and we can determine that
by subtracting the current value from our target value, but we can also know what values we have seen by storing already
traversed numbers in the map. There's a clever elegance to this type of solution that I find very satisfying.