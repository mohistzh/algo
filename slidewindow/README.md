# Sliding Window

In many problems dealing with arrays or linkedlists, we are asked to find or calculate something among all the contiguous subarrays of a given size. For example, take a look at the following problem:

> Given an array, find the average of all contiguous subarrays of size 'K' in it.

Let's understand this problem with a real input:

```
Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
```

Here, we are asked to find the average of all contiguous subarrays of size '5' in the given array. Let's solve this:

1. For the first 5 numbers (subarray from index 0-4), the average is: (1 + 3 + 2 + 6 - 1) / 5 => 2.2
2. The average of next 5 numbers (subarray from index 1-5) is: (3 + 2 + 6 - 1 + 4) / 5 => 2.8
3. For the next 5 numbers (subarray from index 2-6), the average is: (2 + 6 − 1 + 4 + 1) / 5 => 2.4

......


Here is the final output containing the average of all contiguous subarrays of size 5:

```
Output: [2.2, 2.8, 2.4, 3.6, 2.8]
``` 

A brute-force algorithm will be to calculate the sum of every 5-elements contiguous subarray of the given array and divide the sum by '5' to find the average.

But the time complexity is too high, since for every element of the input array, we are calculating the sum of its next 'K' elements, the time complexity of the naive solution will be O(N * K) where 'N' is the number of elements in the input array.

The inefficiency is that for any two consecutive subarrays of size '5', the overlapping part (which will contain four elements) will be evaluate twice.

The efficient way to solve this problem would be **to visualize each contiguous subarray as a sliding window** of '5' elements. This means that when we move on to the next subarray, we will slide the window by one element. So, to reuse the *sum* from the previous subarray, we will subtract the element going out of the window and add the element now being included in the sliding window. This will save us from going through the whole subarray to find the *sum* and, as a result, the algorithm complexity will reduce to O(N).