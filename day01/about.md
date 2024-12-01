# AoC 2024 Day: 01

## Problem:
- Got 2 unsorted lists, A and B
- Have to find 'total_distance' between lists, where 'total_distance' is f(n) = sum(|A[i] - B[i]| for i in range(0, n))

## Steps:
1. Load input data, lists A and B
2. Sort both lists
3. Compare n-items - increment `total_distance` by |A[n] - B[n]|
4. Return `total_distance`

## What I've learned

## Final notes