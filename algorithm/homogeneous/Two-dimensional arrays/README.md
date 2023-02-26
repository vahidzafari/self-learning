# Two-dimensional arrays

A two-dimensional array, is a list of single-dimensional arrays. Every element in a two-dimensional array arr, is identified as arr[i][j], where arr is the name of the array and i and j represent rows and columns, and their values ranging from 0 to m and 0 to n, respectively. Traversing a two-dimensional array is of O(m*n) complexity.

The following code shows how to initialize an array:

```go
var arr = [4][5] int{
    {4,5,7,8,9},
    {1,2,4,5,6},
    {9,10,11,12,14},
    {3,5,6,8,9}
}
```

An element in a two-dimensional array is accessed with a row index and column index. In the following example, the array's value in row 2 and column 3 is retrieved as an integer value:

```go
var value int = arr[2][3]
```

Arrays can store a sequential collection of data elements of the same type. Homogeneous data structure arrays consist of contiguous memory address locations.

Two-dimensional matrices are modeled as two-dimensional arrays. A scalar is an element of a field that defines a vector space. A matrix can be multiplied by a scalar. You can divide a matrix by any non-zero real number.

The order of a matrix is the number of rows, m, by the number of columns, n. A matrix with rows m and columns n is referred to as an m x n matrix. There are multiple types of matrices, such as a row matrix, column matrix, triangular matrix, null matrix, and zero matrix;
