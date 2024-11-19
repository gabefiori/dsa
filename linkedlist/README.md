# Linked List

## Time Complexity

### 1. Search
- **Singly Linked List**: 
  - **Best Case**: *O(1)* - The element is found at the head.
  - **Average Case**: *O(n)* - The element is found somewhere in the list.
  - **Worst Case**: *O(n)* - The element is not found.
  
- **Doubly Linked List**: 
  - **Best Case**: *O(1)* - The element is found at the head.
  - **Average Case**: *O(n)* - The element is found somewhere in the list.
  - **Worst Case**: *O(n)* - The element is not found.

### 2. Insertion
- **Singly Linked List**: 
  - **Prepend**: *O(1)* - Insert at the head.
  - **Append**: *O(n)* - Insert at the tail (requires traversal).
  - **Insert At Position**: *O(n)* - Requires traversal to the position.
  
- **Doubly Linked List**: 
  - **Prepend**: *O(1)* - Insert at the head.
  - **Append**: *O(1)* - Insert at the tail (direct access to the tail).
  - **Insert At Position**: *O(n)* - Requires traversal to the position.

### 3. Deletion
- **Singly Linked List**: 
  - **Delete First**: *O(1)* - Remove the head.
  - **Delete Last**: *O(n)* - Requires traversal to find the second-to-last node.
  - **Delete At Position**: *O(n)* - Requires traversal to the position.
  
- **Doubly Linked List**: 
  - **Delete First**: *O(1)* - Remove the head.
  - **Delete Last**: *O(1)* - Remove the tail (direct access to the tail).
  - **Delete At Position**: *O(n)* - Requires traversal to the position.
