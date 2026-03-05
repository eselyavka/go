package p100

import (
	"github.com/eseliavka/go/util"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution100(t *testing.T) {
	assert := assert.New(t)

	root1 := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root1.Left = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root1.Right = &util.TreeNode{Val: 3, Left: nil, Right: nil}

	root2 := util.TreeNode{Val: 1, Left: nil, Right: nil}
	root2.Left = &util.TreeNode{Val: 2, Left: nil, Right: nil}
	root2.Right = &util.TreeNode{Val: 3, Left: nil, Right: nil}

	actual := isSameTree(&root1, &root2)

	assert.True(actual, "Solution100")

}
