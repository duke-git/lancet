# Tree
树是树节点的集合。 每个树节点都有一个值，一个指向左节点的左指针和一个指向右节点的右指针。

<div STYLE="page-break-after: always;"></div>

## 源码

- [https://github.com/duke-git/lancet/blob/main/datastructure/tree/bstree.go](https://github.com/duke-git/lancet/blob/main/datastructure/tree/bstree.go)


<div STYLE="page-break-after: always;"></div>

## 用法
```go
import (
    tree "github.com/duke-git/lancet/v2/datastructure/tree"
)
```

<div STYLE="page-break-after: always;"></div>

## 目录

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
- [Print](#BSTree_Print)



<div STYLE="page-break-after: always;"></div>

## 文档

## 1. BSTree
BSTree是一种二叉搜索树数据结构，其中每个节点有两个孩子，分别称为左孩子和右孩子。 在 BSTree 中：leftNode < rootNode < rightNode。 T类型应该实现lancetconstraints.Comparator。

### <span id="NewBSTree">NewBSTree</span>
<p>返回BSTree指针实例</p>

<b>函数签名:</b>

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
<b>示例:</b>

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
<p>将值插入二叉搜索树</p>

<b>函数签名:</b>

```go
func (t *BSTree[T]) Insert(data T)
```
<b>示例:</b>

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
<p>删除插入二叉搜索树中指定的值</p>

<b>函数签名:</b>

```go
func (t *BSTree[T]) Delete(data T)
```
<b>示例:</b>

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
<p>按前序遍历树节点</p>

<b>函数签名:</b>

```go
func (t *BSTree[T]) PreOrderTraverse() []T
```
<b>示例:</b>

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
<p>按中序遍历树节点</p>

<b>函数签名:</b>

```go
func (t *BSTree[T]) InOrderTraverse() []T
```
<b>示例:</b>

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
<p>按后序遍历树节点</p>

<b>函数签名:</b>

```go
func (t *BSTree[T]) PostOrderTraverse() []T
```
<b>示例:</b>

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
<p>按节点层次遍历树节点</p>

<b>函数签名:</b>

```go
func (t *BSTree[T]) LevelOrderTraverse() []T
```
<b>示例:</b>

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
<p>获取树的深度</p>

<b>函数签名:</b>

```go
func (t *BSTree[T]) Depth() int
```
<b>示例:</b>

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
}
```




### <span id="BSTree_HasSubTree">HasSubTree</span>
<p>判断给定树是否是子树</p>

<b>函数签名:</b>

```go
func (t *BSTree[T]) HasSubTree(subTree *BSTree[T]) bool
```
<b>示例:</b>

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
}
```




### <span id="BSTree_Print">Print</span>
<p>打印树结构</p>

<b>函数签名:</b>

```go
func (t *BSTree[T]) Print()
```
<b>示例:</b>

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

    fmt.Println(bstree.Print())
//        6
//       / \
//      /   \
//     /     \
//    /       \
//    5       7
//   /
//  /
//  2
//   \
//    4
}
```