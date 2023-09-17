# Non Constructible Change

## Problem

Given a list of coins, find the smallest ammount of change that cannot be made

## Solution

Sort the list of coins, then iterate through the coins. At each iteration check if the coin is greater than the current
total of change plus one. We know that if the current coin is not greater than the total change pluse one we can make
change for every value equal to an less than the current total.


## Time Complexity

O(nlog(n)) The complete analysis would be nlog(n) + n, but because n is less than nlog(n) we can drop it.

## Space Complexity

O(1) constant space, we are sorting the array in place, thus not creating any additional data structures.

## Shane's takeaway

This was a very challenging one. The key was to notice that if the current coin was not greater than the total change
plus one, we could make change for every value <= the current change. That feels like something that a math person would
have known or grasped right away. For me that was not obvious. My hope is that by solving enough of these I'll see some
kind of repeating pattern that will allow me to intuit the solution more easily.