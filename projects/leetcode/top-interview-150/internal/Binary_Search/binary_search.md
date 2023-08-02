# Binary Search

> Is there anybody who gets confuse while setting condition for binary search like whethere it should be (l<r) or (l<=r) or like whether it should be l = mid-1 or l = mid or r = mid+1 or r = mid.
This stuff confuses me a lot.
Someone please share your thoughts to clear.
Thanks in advance.


bparanj
Jul 04, 2023
Yes, setting the right conditions in a binary search can be tricky and it's one of the most common sources of confusion when implementing it. Here's a simple guideline to follow:

l < r vs l <= r: The condition l < r is generally used when you're doing a binary search where the search space reduces on each step, meaning that at some point l and r will be the same. When they are the same, you know you've found the element or it does not exist. l <= r is used when you're doing a binary search where you need to consider the middle element as well.

l = mid + 1 vs l = mid and r = mid - 1 vs r = mid: These decisions are often problem-specific and depend on whether you are looking for the first occurrence, last occurrence, or any occurrence. Usually, when the middle element is not the target, you move one step to the left or right to skip the current middle element.

Remember, the best way to get comfortable with these decisions is to practice problems involving binary search and try to understand the reasoning behind the conditions.

In this problem, we use l < r because we are reducing the search space on each step, and l = mid + 1 or r = mid because we have already checked the mid element, and we are sure that it's not the answer so we need to move one step further.

Visualizing these steps can be very helpful. Let's consider an array: [4, 5, 6, 7, 0, 1, 2].

Step 1. l < r vs l <= r:

A binary search algorithm repeatedly divides the search space (the portion of the array we are interested in) in half. l and r represent the bounds of this search space. Initially, l is at the start of the array, and r is at the end.

When we say l < r, it means the search continues as long as the left boundary is strictly less than the right boundary, indicating that there is more than one element left in the search space.

Here is a quick visualization:

Array:  4  5  6  7  0  1  2
Index:  0  1  2  3  4  5  6
         ^           ^
         |           |
         l           r

After a few iterations:

Array:  4  5  6  7  0  1  2
Index:  0  1  2  3  4  5  6
                  ^  ^
                  |  |
                  l  r
When l and r point to the same element, we stop, because we've narrowed down the search space to a single element.

Step 2. l = mid + 1 vs l = mid and r = mid - 1 vs r = mid:

Let's say we're at a point where l is at index 3 (value 7), r is at index 6 (value 2), and mid is at index 4 (value 0). Here we find that nums[mid] < nums[r]. So, the minimum must be in the left part of the array. We then set r to mid.

Why not mid - 1? Because the mid element is a potential answer (it could be the minimum), so we include it in our next search space.

Array:  4  5  6  7  0  1  2
Index:  0  1  2  3  4  5  6
                  ^  ^  ^
                  |  |  |
                  l mid r
                  
Next iteration:

Array:  4  5  6  7  0  1  2
Index:  0  1  2  3  4  5  6
                  ^  ^
                  |  |
                  l/r(mid)    
We can see here that including mid in our search space allows us to find the minimum element in the array.


```
First, let us understand the concept of Binary Search (BS). BS is a method of optimizing a search by reducing the search space by half in every iteration. This approach can be applied whenever there is a condition based on which the search space can be reduced by half. One such condition is when the array is sorted. However, it is not necessary for the array to be sorted always. If we can find any other condition to divide the search space in two, we can still apply BS. Lets use this information to determine when to use L<=R vs L<R.

I typically use L<=R when I know precisely which value I am searching for. In this scenario, if the value is not present in the array and the loop breaks when R<L, then R becomes the floor of the value and L becomes the ceiling of the value we are searching for. On the other hand, I use L<R when I do not know when the loop will break - because I do not know the value I am searching for in the array. In this case, I only know which side of the search space to look into in the next iteration. A practical example of this scenario is searching for a pivot in a mountain array. We do not know what the peak or pivot is, but we do know on which side the peak will be. Therefore, when L==R, the loop breaks, and we get our solution.

When we use L<=R ---->> L or R is always updated as L = mid +1 or R = mid-1 because we have already included nums[mid] in our search space by doing L<=R.

```
