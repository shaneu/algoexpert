# Validate Subsequence

## Problem

Does the target subsequence exist in our input array? 

## Solution

We must traverse both arrays because we need to determine if our subsequence is in the main array, and it is possible
that the last item in our subsequence array is the last item in our input array in which case we will have traversed
both arrays completely. We can use a for loop to iterate over the main input array and maintain a pointer into the
subsequence array to track our position in that array. At each iteration the for loop handles incrementing our position
in the main array, but we will only increment our position in the subsequence array if we have a match. We know that if
we have traversed the entire subsequence array we have found the sequence in the main array, so we can break out of the
loop early if that is the case.

## Time Complexity

O(n), linear time. The greater the length of the array we are trying to find the sequence in, the longer our algorithm
will run.

## Space Complexity

O(1), constant space. We are not creating any additional data structures therefore we are not using any appreciable
amount of auxiliary memory.

## Shane's takeaway

What I find most interesting about this solution is this notion of using the information we have to infer our answer. If
we've found that the target subsequence exists in our array it means that we have traversed every value in our
subsequence array and by that fact we are able to know the sequence exists. If we made it to the end of our input array,
and we still haven't finished traversing the subsequence array we know we haven't found the sequence. We don't need any
more complicated logic than that. There's a stripped down simplicity and elegance in that sort of solution that I find
very appealing.

