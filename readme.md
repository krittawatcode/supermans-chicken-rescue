# Superman's Chicken Rescue

This is a Go project that aims to solve an  problem: How can Superman protect as many chickens as possible from a rainstorm using a roof of limited length?



In this one-dimensional world, Superman has a roof of a certain length and he needs to use it to shield chickens from a heavy rainstorm. The positions of the chickens are represented as an array of integers, each integer representing a position on a one-dimensional axis.

The goal of this project is to determine the maximum number of chickens that Superman can protect given the length of the roof and the positions of the chickens.

## Problem Statement

- `n` represents the number of chickens
- `k` denotes the length of the roof Superman can carry
- The positions of the chickens on the one-dimensional axis are represented as an array of integers

## Input

The input consists of two integers n and k (1 <= n,k <= 1,000,000), where n represents the number of chickens
and k denotes the length of the roof Superman can carry. The next line contains n integers (1 <= x <=
1,000,000,000) representing the positions of the chickens on the one-dimensional axis.

## Output

The program outputs a single integer, denoting the maximum number of chickens Superman can protect with the given roof length.

Note:
- Superman can position the roof starting at any point on the axis.
- The roof can cover chickens whose positions are within k units from its starting point.
- It's not required to cover all chickens, but to maximize the number of chickens protected.
- It’s guaranteed that the given positions of the chickens will be sorted from lowest to highest.

### Example:
| Input    | Output   |
|----------|----------|
| 5 5<br>2 5 10 12 15  | 2 |
| 6 10<br>1 11 30 34 35 37  | 4 |


Explanation:
In the first example, superman can position the roof starting at position 2 (roof at 2 - 6), covering chickens at
positions 2 and 5. Thus, he can protect a maximum of 2 chickens.
In the second example, superman can position the roof starting at position 30 (roof at 30 - 39), covering
chickens at positions 30, 34, 35, and 37. Thus, he can protect a maximum of 4 chickens.

### The main files in this project are:

- `rescue.go`: Contains the main logic for solving the problem.
- `rescue_test.go`: Contains the tests for the `rescue.go` file.
- `main.go`: The entry point of the application.

The project uses Go modules for dependency management, as specified in the `go.mod` file.