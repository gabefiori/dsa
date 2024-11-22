# Heap

## Time Complexity

### 1. Push
- **Best Case**: *O(1)* - The new element is added to the heap without needing to restructure.
- **Average Case**: *O(log n)* - The new element is added to the heap and may require restructuring.
- **Worst Case**: *O(log n)* - The new element is added to the heap and requires restructuring all the way to the root.

### 2. Pop
- **Best Case**: *O(log n)* - The root element is removed, and the heap is restructured efficiently.
- **Average Case**: *O(log n)* - The root element is removed, and the heap is restructured.
- **Worst Case**: *O(log n)* - The root element is removed, and the heap requires restructuring all the way to the bottom.

### 3. Peek
- **Best Case**: *O(1)* - The root element is accessed directly.
- **Average Case**: *O(1)* - The root element is accessed directly.
- **Worst Case**: *O(1)* - The root element is accessed directly.
