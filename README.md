# Data structures and algorithms

## About

This repository contains implementations of data structures and algorithms using `(GO)[https://golang.org/]`.

Data structures and algorithms are the building blocks of computer programming. They are the tools that help us to solve complex problems in an efficient way. Understanding data structures and algorithms is essential for every programmer. It is a common practice to ask data structure and algorithm problems in interviews. So, it is important to have a good understanding of data structures and algorithms.

## Table of contents

- [Data structures and algorithms](#data-structures-and-algorithms)
  - [About](#about)
  - [Table of contents](#table-of-contents)
  - [Data structures](#data-structures)
    - [Array](#array)
    - [Linked List](#linked-list)
    - [Stack](#stack)
    - [Queue](#queue)
    - [Tree](#tree)
    - [Graph](#graph)
    - [Hash Table](#hash-table)
  - [Algorithms](#algorithms)
    - [Searching](#searching)
      - [Linear Search](#linear-search)
      - [Binary Search](#binary-search)
    - [Sorting](#sorting)
      - [Bubble Sort](#bubble-sort)
      - [Selection Sort](#selection-sort)
      - [Insertion Sort](#insertion-sort)
      - [Merge Sort](#merge-sort)
      - [Quick Sort](#quick-sort)
  - [Problem Solving](#problem-solving)

## Data structures

Data structures are a way of organizing and storing data so that they can be accessed and worked with efficiently. They define the relationship between the data, and the operations that can be performed on the data. There are many different types of data structures, generally built upon simpler primitive data types:

- **Primitive** data types - includes integer, character, boolean, floating point and pointer
- **Non-primitive** data types - includes array, list, file, pointer, stack, queue, tree, graph, etc.

Data structures provide a means to manage large amounts of data efficiently for uses such as large databases and internet indexing services. Usually, efficient data structures are key to designing efficient algorithms. Some formal design methods and programming languages emphasize data structures, rather than algorithms, as the key organizing factor in software design. Storing and retrieving can be carried out on data stored in both main memory and in secondary memory.

### Array

An array is a collection of items stored at contiguous memory locations. The idea is to store multiple items of the same type together. This makes it easier to calculate the position of each element by simply adding an offset to a base value, i.e., the memory location of the first element of the array (generally denoted by the name of the array).

For simplicity, we can think of an array a fleet of stairs where on each step is placed a value (let’s say one of your friends). Here, you can identify the location of any of your friends by simply knowing the count of the step they are on.

![Array](https://www.geeksforgeeks.org/wp-content/uploads/array-2.png)

### Linked List

A linked list is a linear data structure, in which the elements are not stored at contiguous memory locations. The elements in a linked list are linked using pointers as shown in the below image:

![Linked List](https://www.geeksforgeeks.org/wp-content/uploads/gq/2013/03/Linkedlist.png)

In simple words, a linked list consists of nodes where each node contains a data field and a reference(link) to the next node in the list.

### Stack

A stack is an Abstract Data Type (ADT), commonly used in most programming languages. It is named stack as it behaves like a real-world stack, for example – a deck of cards or a pile of plates, etc.

A real-world stack allows operations at one end only. For example, we can place or remove a card or plate from the top of the stack only. Likewise, Stack ADT allows all data operations at one end only. At any given time, we can only access the top element of a stack.

This feature makes it LIFO data structure. LIFO stands for Last-in-first-out. Here, the element which is placed (inserted or added) last, is accessed first. In stack terminology, insertion operation is called PUSH operation and removal operation is called POP operation.

![Stack](https://www.geeksforgeeks.org/wp-content/uploads/gq/2013/03/stack.png)

### Queue

Like Stack, Queue is a linear structure which follows a particular order in which the operations are performed. The order is First In First Out (FIFO). A good example of queue is any queue of consumers for a resource where the consumer that came first is served first.

![Queue](https://www.geeksforgeeks.org/wp-content/uploads/gq/2013/03/queue.png)

### Tree

A tree is a hierarchical data structure defined as a collection of nodes. Nodes represent the structure of the tree and are often arranged in levels. The first node of the tree is called the root. Each node can have a number of children. Nodes with no children are called leaves.

![Tree](https://www.geeksforgeeks.org/wp-content/uploads/gq/2013/03/BasicTree.png)

### Graph

A graph is a data structure that consists of the following two components:

1. A finite set of vertices also called as nodes.
2. A finite set of ordered pair of the form (u, v) called as edge. The pair is ordered because (u, v) is not the same as (v, u) in case of a directed graph(di-graph). The pair of the form (u, v) indicates that there is an edge from vertex u to vertex v. The edges may contain weight/value/cost.

![Graph](https://www.geeksforgeeks.org/wp-content/uploads/undirectedgraph.png)

### Hash Table

Hashing is an important Data Structure which is designed to use a special function called the Hash function which is used to map a given value with a particular key for faster access of elements. The efficiency of mapping depends of the efficiency of the hash function used.

![Hash Table](https://www.geeksforgeeks.org/wp-content/uploads/HashingDataStructure-min-1.png)

## Algorithms

An algorithm is a set of instructions designed to perform a specific task. This can be a simple process, such as multiplying two numbers. Search engines use proprietary algorithms to display the most relevant results from their search index for specific queries.

### Searching

Searching is the process of finding some particular element in the list. If the element is found, then the search is successful, otherwise unsuccessful. The list may be ordered or unordered.

#### Linear Search

Linear search is a very simple search algorithm. In this type of search, a sequential search is made over all items one by one. Every item is checked and if a match is found then that particular item is returned, otherwise the search continues till the end of the data collection.

![Linear Search](https://www.geeksforgeeks.org/wp-content/uploads/Linear-Search.png)

#### Binary Search

Binary search is a fast search algorithm with run-time complexity of Ο(log n). This search algorithm works on the principle of divide and conquer. For this algorithm to work properly, the data collection should be in the sorted form.

![Binary Search](https://www.geeksforgeeks.org/wp-content/uploads/Binary-Search.png)

### Sorting

Sorting is the process of arranging the elements in a list in increasing or decreasing order of some property. Sorting algorithms are used to optimize the performance and resources usage in computer science.

#### Bubble Sort

Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in wrong order.

![Bubble Sort](https://www.geeksforgeeks.org/wp-content/uploads/bubble-sort-300x300.png)

#### Selection Sort

The selection sort algorithm sorts an array by repeatedly finding the minimum element (considering ascending order) from unsorted part and putting it at the beginning.

![Selection Sort](https://www.geeksforgeeks.org/wp-content/uploads/Selection-sort-flowchart.jpg)

#### Insertion Sort

Insertion sort is a simple sorting algorithm that works the way we sort playing cards in our hands.

![Insertion Sort](https://www.geeksforgeeks.org/wp-content/uploads/insertionsort.png)

#### Merge Sort

Like QuickSort, Merge Sort is a Divide and Conquer algorithm. It divides input array in two halves, calls itself for the two halves and then merges the two sorted halves.

![Merge Sort](https://www.geeksforgeeks.org/wp-content/uploads/Merge-Sort-Tutorial.png)

#### Quick Sort

Like Merge Sort, QuickSort is a Divide and Conquer algorithm. It picks an element as pivot and partitions the given array around the picked pivot.

![Quick Sort](https://www.geeksforgeeks.org/wp-content/uploads/gq/2014/01/QuickSort2.png)

## Problem Solving

Problem solving consists of using generic or ad hoc methods in an orderly manner to find solutions to problems. Some of the problem-solving techniques developed and used in philosophy, artificial intelligence, computer science, engineering, mathematics, or medicine are related to mental problem-solving techniques studied in psychology.

This problems is a combination of using data structures and algorithms to solve a problem.
Common problems for example:

- Reverse string.
- Find the first non-repeated character in a string.
- Find the first repeated character in a string.
- Find the first repeated word in a string.
- Find the first non-repeated word in a string.
- Two sum problem.
- Find the missing number in a given integer array of 1 to 100.
- And many more...
