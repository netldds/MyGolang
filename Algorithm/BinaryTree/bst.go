package main

//二叉搜索树
type BST struct {
	Parent *BST
	Left   *BST
	Right  *BST
	Key    int
}

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
func TreeInsert(T, z *BST) *BST {
	x := T
	var y *BST
	for x != nil {
		y = x
		if z.Key == x.Key {
			break
		}
		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
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
	return T
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
	if z.Left == nil {
		TreeTransPlant(T, z, z.Right)
	} else if z.Right == nil {
		TreeTransPlant(T, z, z.Left)
	} else {
		y := TreeSuccessor(z)
		if y.Parent != z {
			TreeTransPlant(T, y, y.Right)
			y.Right = z.Right
			z.Right.Parent = y
		}
		TreeTransPlant(T, z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
	return T
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
