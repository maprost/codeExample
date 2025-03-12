## Go Interview Questions by Nina Pakshina
Source: https://medium.com/@ninucium

Part1: https://medium.com/@ninucium/go-interview-questions-part-1-pointers-channels-and-range-67c61345cf3c

Part2: https://medium.com/@ninucium/go-interview-questions-part-2-slices-87f5289fb7eb

Part3: https://medium.com/@ninucium/go-interview-questions-part-3-size-of-slices-and-int-25080e51cf72

### Folders:
- `channel01`
- `channel02`
- `range01`
- `slice01`
- `slice02`
- `slice03`

### Summary to remember in order to crack these interview questions
Source: https://medium.com/@ninucium/go-interview-questions-part-2-slices-87f5289fb7eb
- A slice is a reference data type. Inside there is a pointer to the first element of the slice. This factor is what determines how certain operations, even when performed on copies of the slice, can affect the original slice.
- A slice has a length, which describes the number of elements currently stored in the slice, and a capacity, which indicates how many elements can be added to this memory area.
- If the inequality len(x) + 1 <= cap(x) is not met when adding a new element, the slice expands into a new area of memory, and capacity doubles (until it reaches the size of 1024, after which they increase by 25% with each expansion).
- When you pass a slice as an argument to a function as a copy (not via a pointer), you should remember that the slice contains a pointer to the first element, which allows for modifications to the original slice.
- The length and capacity values are passed by copy. If you pass a slice to a function and then the same slice is modified elsewhere in the code (e.g., by adding a new element), it will not affect the length and capacity of the copied slice within the function.

## Other pages
- https://www.interviewbit.com/golang-interview-questions/#is-it-possible-to-return-multiple-values-from-a-function