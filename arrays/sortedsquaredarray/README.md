# Sorted Squared Array

## Problem

Write a function that takes in a non-empty array of integers that are sorted in ascending order and returns a new array
of the same length with the squares of the original integers also sorted in ascending order.

## Solution

The brute force solution would be to simply square all the values in the input array, assign them to the output array
and sort the output array in place, giving us a nlog(n) time complexity.

The more optimal solution can be done in linear time, thanks to the fact that our input array is sorted. The absolute
values at either end of the array are compared and whichever is greater has its square inserted into the output array
from right to left.

## Time Complexity

O(n)

The time complexity is linear, we only iterate over the input array once.

## Space Complexity

O(n)

The space complexity is linear, we must create an output array equal to the length of the input array.

## Shane's takeaway

The key thing in this problem is to notice that the input array is sorted. That is often a clue that we can achieve a
linear time complexity in out implementation.
