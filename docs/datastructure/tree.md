# Tree
Tree is a collection of tree nodes. Each tree node has a value, a left pointer point to left node and a right pointer point to right node.

<div STYLE="page-break-after: always;"></div>

## Source

- [https://github.com/duke-git/lancet/blob/main/datastructure/tree/bstree.go](https://github.com/duke-git/lancet/blob/main/datastructure/tree/bstree.go)


<div STYLE="page-break-after: always;"></div>

## Usage
```go
import (
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)
```

<div STYLE="page-break-after: always;"></div>

## Index

### 1. BSTree

- [NewBSTree](#NewBSTree)
- [Insert](#BSTree_Insert)
- [Delete](#BSTree_Delete)
- [PreOrderTraverse](#BSTree_PreOrderTraverse)
- [InOrderTraverse](#BSTree_InOrderTraverse)
- [PostOrderTraverse](#BSTree_PostOrderTraverse)
  
- [LevelOrderTraverse](#BSTree_LevelOrderTraverse)
- [Depth](#BSTree_Depth)
- [HasSubTree](#BSTree_HasSubTree)



<div STYLE="page-break-after: always;"></div>

## Documentation

## 1. BSTree
BSTree is a binary search tree data structure in which each node has at two children, which are referred to as the left child and the right child. In BSTree: leftNode < rootNode < rightNode. Type T should implements Compare function in lancetconstraints.Comparator interface.

### <span id="NewBSTree">NewBSTree</span>
<p>Make a BSTree pointer instance</p>

<b>Signature:</b>

```go
func NewBSTree[T any](rootData T, comparator lancetconstraints.Comparator) *BSTree[T]

type BSTree[T any] struct {
	root       *datastructure.TreeNode[T]
	comparator lancetconstraints.Comparator
}

type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    bstree := tree.NewBSTree(6, &intComparator{})
    fmt.Println(bstree) //
}
```




### <span id="BSTree_Insert">Insert</span>
<p>Insert value into binary search tree</p>

<b>Signature:</b>

```go
func (t *BSTree[T]) Insert(data T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    bstree := tree.NewBSTree(6, &intComparator{})
    bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

    fmt.Println(bstree.PreOrderTraverse()) //6, 5, 2, 4, 7
}
```




### <span id="BSTree_Delete">Delete</span>
<p>Delete value of binary search tree</p>

<b>Signature:</b>

```go
func (t *BSTree[T]) Delete(data T)
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    bstree := tree.NewBSTree(6, &intComparator{})
    bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

    bstree.Delete(4)

    fmt.Println(bstree.PreOrderTraverse()) //2, 5, 6, 7
}
```




### <span id="BSTree_PreOrderTraverse">PreOrderTraverse</span>
<p>Traverse tree nodes in pre order</p>

<b>Signature:</b>

```go
func (t *BSTree[T]) PreOrderTraverse() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    bstree := tree.NewBSTree(6, &intComparator{})
    bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

    fmt.Println(bstree.PreOrderTraverse()) //6, 5, 2, 4, 7
}
```




### <span id="BSTree_InOrderTraverse">InOrderTraverse</span>
<p>Traverse tree nodes in middle order</p>

<b>Signature:</b>

```go
func (t *BSTree[T]) InOrderTraverse() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    bstree := tree.NewBSTree(6, &intComparator{})
    bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

    fmt.Println(bstree.InOrderTraverse()) //2, 4, 5, 6, 7
}
```




### <span id="BSTree_PostOrderTraverse">PostOrderTraverse</span>
<p>Traverse tree nodes in post order</p>

<b>Signature:</b>

```go
func (t *BSTree[T]) PostOrderTraverse() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    bstree := tree.NewBSTree(6, &intComparator{})
    bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

    fmt.Println(bstree.PostOrderTraverse()) //5, 2, 4, 7, 6
}
```




### <span id="BSTree_LevelOrderTraverse">LevelOrderTraverse</span>
<p>Traverse tree nodes in node level order</p>

<b>Signature:</b>

```go
func (t *BSTree[T]) LevelOrderTraverse() []T
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    bstree := tree.NewBSTree(6, &intComparator{})
    bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

    fmt.Println(bstree.LevelOrderTraverse()) //6, 5, 7, 2, 4
}
```




### <span id="BSTree_Depth">Depth</span>
<p>Get the depth of a binary saerch tree</p>

<b>Signature:</b>

```go
func (t *BSTree[T]) Depth() int
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    bstree := tree.NewBSTree(6, &intComparator{})
    bstree.Insert(7)
	bstree.Insert(5)
	bstree.Insert(2)
	bstree.Insert(4)

    fmt.Println(bstree.Depth()) //4
```




### <span id="BSTree_HasSubTree">HasSubTree</span>
<p>Check if the given tree is sub tree of origin tree or not</p>

<b>Signature:</b>

```go
func (t *BSTree[T]) HasSubTree(subTree *BSTree[T]) bool
```
<b>Example:</b>

```go
package main

import (
    "fmt"
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)

type intComparator struct{}

func (c *intComparator) Compare(v1, v2 any) int {
	val1, _ := v1.(int)
	val2, _ := v2.(int)

	if val1 < val2 {
		return -1
	} else if val1 > val2 {
		return 1
	}
	return 0
}

func main() {
    superTree := tree.NewBSTree(8, &intComparator{})
	superTree.Insert(4)
	superTree.Insert(5)
	superTree.Insert(6)
	superTree.Insert(9)
	superTree.Insert(4)

    subTree := tree.NewBSTree(5, &intComparator{})
	subTree.Insert(4)
	subTree.Insert(6)

    fmt.Println(superTree.HasSubTree(subTree)) //true
    fmt.Println(subTree.HasSubTree(superTree)) //false
```