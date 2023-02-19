# B+ tree

The B+ tree contains a list of keys and pointers to the next-level nodes in trees. During a search, recursion is used to search for an element by looking for the the adjacent node keys. B+ trees are used to store data in filesystems. B+ trees require fewer I/O operations to search for a node in the tree. Fan-out is defined as the number of nodes pointing to the child nodes of a node in a B+ tree. B+ trees were first described in a technical paper by Rudolf Bayer and Edward M. McCreight.

The block-oriented storage context in B+ trees helps with the storage and efficient retrieval of data. The space efficiency of a B+ tree can be enhanced by using compression techniques. B+ trees belong to a family of multiway search trees. For a b-order B+ tree, space usage is of the order O(n). Inserting, finding, and removing operations are of the order O(logbn).

# B-tree

The B-tree is a search tree with non-leaf nodes that only have keys, and the data is in the leaves. B-trees are used to reduce the number of disk accesses. The B-tree is a self-adjusting data structure that keeps data sorted. B-trees store keys in a sorted order for easy traversal. They can handle multiple insertions and deletions.

Knuth initially came up with the concept of this data structure. B-trees consist of nodes that have at most n children. Every non-leaf node in the tree has at least n/2 child nodes. Rudolf Bayer and Edward M. McCreight were the first to implement this data structure in their work. B-trees are used in HFS and Reiser4 filesystems to allow for quick access to any block in a file. On average, space usage is in the order of O(n). Insert, search, and delete operations are in the order of O(log n).

# T-tree

The T-tree is a balanced data structure that has both the index and actual data in memory. They are used in in-memory databases. T refers to the shape of the node. Each node consists of pointers to the parent node and the left and right child nodes. Each node in the tree node will have an ordered array of data pointers and extra control data. T-trees have similar performance benefits to in-memory tree structures. A T-tree is implemented on top of a self-balancing binary search tree. This data structure is good for ordered scanning of data. It supports various degrees of isolation.
