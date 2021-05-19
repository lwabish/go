package tree

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeByLevel 由特定形式的字符串生成一颗二叉树
// 空格作为分隔符，n代表跳过该位置的子树，对于不存在的子树，不需要继续写其子树
func NewTree(valueString string) (root *TreeNode) {
	values := strings.Split(valueString, " ")
	rootValue, _ := strconv.Atoi(values[0])

	// 待装配的node队列
	q := []*TreeNode{{Val: rootValue}}
	root = q[0]
	// node队列的指针，值指针
	qIndex, vIndex := 0, 1

	for vIndex < len(values) {
		currentNode := q[qIndex]
		qIndex += 1

		currentValue := values[vIndex]
		vIndex += 1
		// 尝试装配左子树
		if currentValue != "n" {
			v, _ := strconv.Atoi(currentValue)
			currentNode.Left = &TreeNode{Val: v}
			// 把新的左子树加入队列
			q = append(q, currentNode.Left)
		}

		if vIndex >= len(values) {
			break
		}

		currentValue = values[vIndex]
		vIndex += 1
		if currentValue != "n" {
			v, _ := strconv.Atoi(currentValue)
			currentNode.Right = &TreeNode{Val: v}
			q = append(q, currentNode.Right)
		}
	}
	return
}

// Draw 画出二叉树的图形结构
func (t *TreeNode) Draw() {
	t.draw("", true)
}

func (t *TreeNode) draw(prefix string, isLeft bool) {
	if t == nil {
		return
	}

	var rightPrefix string
	if t.Right != nil {
		if isLeft {
			rightPrefix = prefix + "│   "
		} else {
			rightPrefix = prefix + "    "
		}
		t.Right.draw(rightPrefix, false)
	}

	var thisPrefix string
	if isLeft {
		thisPrefix = prefix + "└── "
	} else {
		thisPrefix = prefix + "┌── "
	}
	fmt.Println(thisPrefix + strconv.Itoa(t.Val))

	var leftPrefix string
	if t.Left != nil {
		if isLeft {
			leftPrefix = prefix + "    "
		} else {
			leftPrefix = prefix + "│   "
		}
		t.Left.draw(leftPrefix, true)
	}
}

// LevelOrder 层次遍历
func (t *TreeNode) LevelOrder() []int {
	return []int{}
}
