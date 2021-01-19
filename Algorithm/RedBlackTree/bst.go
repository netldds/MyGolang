package main

//二叉搜索树
type BST struct {
	Parent *BST
	Left   *BST
	Right  *BST
	Key    int
	Color  int //1=black 2=red
}

const (
	BLACK = 1
	RED   = 2
)

//最大
func TreeMaximum(T *BST) *BST {
	if T.Right == nil {
		return T
	}
	return TreeMaximum(T)
}

//最小
func TreeMinimum(T *BST) *BST {
	if T.Left == nil {
		return T
	}
	return TreeMinimum(T)
}

//替换u为v
func TreeTransPlant(T, u, v *BST) *BST {
	//u就是根
	if u.Parent == nil {
		T.Parent = v
	}
	if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
	return T
}

//插入
func RBInsert(T, z *BST) *BST {
	x := T
	var y *BST
	for x != nil {
		y = x
		if z.Key < y.Key {
			x = y.Left
		}
		if z.Key > y.Key {
			x = y.Right
		}
		if z.Key == y.Key {
			break
		}
	}
	z.Parent = y
	if y == nil {
		T = z
		return T
	}
	if z.Key < y.Key {
		y.Left = z
	}
	if z.Key > y.Key {
		y.Right = z
	}
	//不同
	z.Left = nil
	z.Right = nil
	z.Color = RED
	RBInsertFixup(T, z)
	return T
}
func RBInsertFixup(T, z *BST) {
}

//左旋转
func LeftRotation(T, x *BST) {
	y := x.Right
	x.Right = y.Left
	if y.Left != nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent != nil {
		if x.Left.Parent == x {
			x.Left.Parent = y
		} else {
			x.Right.Parent = y
		}
	}
	y.Left = x
	x.Parent = y

}

//右旋转
func RightRotation(T, y *BST) {
	x := y.Left
	y.Left = x.Right
	if x.Right != nil {
		x.Right.Parent = y
	}
	x.Parent = y.Parent
	if y.Parent != nil {
		if y.Parent.Left == y {
			y.Parent.Left = x
		} else {
			y.Parent.Right = x
		}
	}
	x.Right = y
	y.Parent = x
}

//搜索
func TreeSearch(T *BST, k int) *BST {
	if T == nil || T.Key == k {
		return T
	}
	if k > T.Key {
		return TreeSearch(T.Right, k)
	} else {
		return TreeSearch(T.Left, k)
	}
}

//删除
func TreeDelete(T, z *BST) *BST {
	return nil
}

//successor
func TreeSuccessor(T *BST) *BST {
	if T.Right != nil {
		return TreeMinimum(T.Right)
	}
	x := T
	y := T.Parent
	for y != nil && x == y.Right {
		x = y
		y = y.Parent
	}
	return y
}

//predecessor
func TreePredecessor(T *BST) *BST {
	if T.Left != nil {
		return TreeMaximum(T.Left)
	}
	x := T
	y := T.Parent
	for y != nil && x == y.Left {
		x = y
		y = y.Parent
	}
	return y

}
