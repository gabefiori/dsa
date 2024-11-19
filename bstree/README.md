# Binary Search Tree (BST)

## Time Complexity

### 1. Search
- **Best Case**: *O(1)* - The element is found at the root.
- **Average Case**: *O(log n)* - The tree is balanced.
- **Worst Case**: *O(n)* - The tree is skewed (e.g., all elements are inserted in sorted order).

### 2. Insertion
- **Best Case**: *O(1)* - The new element is inserted as the root.
- **Average Case**: *O(log n)* - The tree is balanced.
- **Worst Case**: *O(n)* - The tree is skewed.

### 3. Deletion
- **Best Case**: *O(1)* - The element to be deleted is the root.
- **Average Case**: *O(log n)* - The tree is balanced.
- **Worst Case**: *O(n)* - The tree is skewed.

### 4. Traversal
- **In-order Traversal**: *O(n)* - Visits each node once.
- **Pre-order Traversal**: *O(n)* - Visits each node once.
- **Post-order Traversal**: *O(n)* - Visits each node once.
