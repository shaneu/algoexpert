# Tournament Winner

## Problem

Given a list of competitions and the results of each round, determine the overall winner.

## Solution

Iterate through each round of competition using a map to keep track of team scores. At each round compare the round
winner to the overall current winning team, and update the current leader if the round winner is now leading the
tournament.

## Time Complexity

O(n)

The time complexity is linear. The larger the input array n, the longer our algorithm will run in the worst case.

## Space Complexity

O(k) where k is the number of teams in the competition. In the worst case we will have one entry in the map per team.

## Shane's takeaway

The takeaway here is because the result is binary, only one team can win, we were able to leverage a "currentLeader"
variable to keep track of the team that was ahead, and update the value of that variable when a new team took the lead.
I'm wondering if a similar approach can be used in algorithms that aggregate several values to find one single value
that is the greatest.