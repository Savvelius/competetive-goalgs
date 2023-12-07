# What strategies do i know?

1. Dynamic programming
2. Divide and Conquer
3. Greedy - identify a function that should be minimized with each iteration

* hashmap - results for some values repeat (cashing)
* Prefix sums - find sum of subarray. Summing can be generalized to any reversible operation (xor, ...)
* Prefix polynomial hash sums - contains subarray?
* Two pointer - some contiguous parts of array are continuously re-evaluated
* prefix function/analogues (dynamic programming subcase)
* heap - need dynamic access to min/max elements
* Manachen algorithm - any palindrome problems
* Sliding window - given iterable find "window size", or given iterable and "window size" find something
* Binary search - find min/max x such that some f(x) is true 

# Binary Search

## Have

f(x) on x in [l, r] where
f(x) is monotonic

## Want

x: f(x) == target

## Cases

1. Range contains target:
    * lsearch and rsearch will return it
2. Range contains many targets:
    * lsearch : first such
    * rsearch : last such
3. Range doesn't contain target and target is:
    * Less than min element:
        * lsearch : first index,
        * rsearch : first index - 1
    * More than max element:
        * lsearch : last index + 1
        * rsearch : last index
    * In between:
        * lsearch : index of element greater
        * rsearch : index of element less

# Greedy Algorithm Strategy:
0. formalize the problem
1. edge cases
2. get a solution implying edge cases
3. Improvise???

Exchange arguments - 
every solution can be improved by modifying it to look more like the greedy solution

# The Dynamic Programming Paradigm
1. Identify a relatively small collection of subproblems.
2. Show how to quickly and correctly solve “larger” sub-
problems given the solutions to “smaller” ones.
3. Show how to quickly and correctly infer the final solu-
tion from the solutions to all of the subproblems

