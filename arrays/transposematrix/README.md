# Transpose Matrix

## Problem

Transpose the given matrix

## Solution

Traverse the matrix by column instead of by row, and collect the values into what will be a new row in the output
matrix.

## Time Complexity

O(w * h) where w is the width of the matrix and h is the height of the matrix

## Space Complexity

O(w * h)

## Shane's takeaway

The key was to iterate by column, then row, so starting at column zero we wanted to collect the values from row 0 at
column 0, row 1 at column 0, etc. Using out nested loops, the outer loop becomes the dimension we want to slowly iterate
over, so here columns, and the inner loop is the dimension we want to iterate quickly over, here rows.