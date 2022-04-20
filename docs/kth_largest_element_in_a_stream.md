# 703. Kth Largest Element in a Stream

## Approach: Heap

This problem is a perfect one to be solved with a heap, also known as a priority queue. If you don't know what a heap is, then this article should help get you up to speed. In this article, we'll discuss the basics of what a heap does and how to use one. If you're interested in the theory behind the implementation of a heap, [here](https://en.wikipedia.org/wiki/Heap_\(data_structure\)\#Implementation) is a resource for that.

In short, a heap is a data structure that is capable of giving you the smallest (or largest) element (by some criteria) in constant time, while also being able to add elements and remove the smallest (or largest) element in only logarithmic time, `O(log(n))`. Imagine if you wanted to replicate this functionality naively with an array. To make sure we can find the smallest element in constant time, let's just keep our array sorted, so that the last element is always the largest (or smallest, depending on if we're sorting in ascending or descending order). Removing the largest/smallest element will take `O(1)` time as we are popping from the end of the array. However, to add a new element, we first need to find where the element should be inserted and then insert it by shifting the array, which requires `O(n)` time. Now, there are potential improvements to this approach, like using a deque for removals and insertions and binary searching to find insertion points, but the point is that a heap makes it so we don't need to worry about any of that.

In summary, a heap:
* Stores elements, and can find the smallest (min heap) or largest (max heap) element stored in `O(1)`.
* Can add elements and remove the smallest (min heap) or largest (max heap) element in `O(log(n))`.
* Can perform insertions and removals while always maintaining the first property.

The capability to remove and insert elements in `O(log(n))` time makes heaps extremely useful. For example, many problems that can be naively solved in `O(n^2)` time, can be solved in `O(n⋅log(n))` time by using a heap. To put this in perspective, for an input size of `n=10^5` elements, `n⋅log(n)` is over 6000 times smaller than `n^2`.

So, now that we know what a heap does, how does it help solve this problem? Let's say we have some stream of numbers, `nums = [6, 2, 3, 1, 7]`, and `k = 3`. Because the input is small, we can clearly see the kth smallest element is 3. Although, earlier we said that a heap can only find an element in `O(1)` time if it's a minimum or maximum (depending on choice of implementation). Well, a heap is also capable of removing the smallest element quickly, so what if we just keep removing the smallest element from `nums` until `nums.length == k`? In this case, we would have `nums = [3, 6, 7]`, and a heap can now give us our answer in `O(1)` time.

That's the key to solving this problem - use a min heap (*min* means that the heap will remove/find the smallest element, a *max* heap is the same thing, but for the largest element) and keep the heap at size `k`. That way, the smallest element in the heap (the one we can access in `O(1)`) will always be the kth largest element. This way, when adding a number to the heap with `add()`, we can do it very quickly in `log(n)` time. If our heap exceeds size `k`, then we can also remove it very quickly. In the end, the smallest element in the heap will be the answer.

## Algorithm

1. In the constructor, create a min heap using the elements from `nums`. Then, pop from the heap until `heap.length == k`.

2. For every call to `add()`:
* First, push `val` into heap.
* Next, check if `heap.length > k`. If so, pop from the heap.
* Finally, return the smallest value from the heap, which we can get in `O(1)` time.

## Implementation

Most languages have built-in implementations for heaps.

## Complexity Analysis

Given `n` as the length of `nums` and `m` as the number of calls to `add()`:
* Time complexity: `O(n⋅log(n) + m⋅log(k))`. The time complexity is split into two parts. First, the constructor needs to turn `nums` into a heap of size `k`. This requires `O(n)`. Then, we need to remove from the heap until there are only `k` elements in it, which means removing `n - k` elements. Since `k` can be, say 1, in terms of big O this is `n` operations, with each operation costing `log(n)`. Therefore, the constructor costs `O(n + n⋅log(n)) = O(n⋅log(n)`. Next, every call to `add()` involves adding an element to heap and potentially removing an element from heap. Since our heap is of size `k`, every call to `add()` at worst costs `O(2∗log(k)) = O(log(k))`. That means `m` calls to `add()` costs `O(m⋅log(k))`.
* Space complexity: `O(n)`. The only extra space we use is the heap. During `add()` calls we limit the size of the heap to `k`. In the constructor we start by converting `nums` into a heap, which means the heap will initially be of size `n`.
