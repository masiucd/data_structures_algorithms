# Single Linked List

In computer science, a linked list is a linear collection of data elements, in which linear order is not given by their physical placement in memory. Instead, each element points to the next. It is a data structure consisting of a collection of nodes which together represent a sequence. In its most basic form, each node contains: data, and a reference (in other words, a link) to the next node in the sequence. This structure allows for efficient insertion or removal of elements from any position in the sequence during iteration. More complex variants add additional links, allowing more efficient insertion or removal of nodes at arbitrary positions. A drawback of linked lists is that access time is linear (and difficult to pipeline). Faster access, such as random access, is not feasible. Arrays have better cache locality compared to linked lists.

## Methods

- `Append(value int)` - Add a node to the end of the list
- `Prepend(value int)` - Add a node to the beginning of the list
- `InsertAt(value, index int)` - Add a node at a given index
- `Search(value int) bool` - Search for a node
- `Get(index int) *SingleNode` - Get a node at a given index
- `Delete(value int)` - Delete a node
- `Traverse()` - Traverse the list
- `Reverse()` - Reverse the list
- `Print()` - Print the list
