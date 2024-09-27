package helpers

import (
    . "golang-binary-tree/types"
    "math"
)

func BuildTree(nodes map[string]Node, nodeID string) *BinaryTree {
    if nodeID == "" {
        return nil
    }
    
    node := nodes[nodeID]
    return &BinaryTree{
        Value: node.Value,
        Left:  BuildTree(nodes, node.Left),
        Right: BuildTree(nodes, node.Right),
    }
}

func MaxPathSum(node *BinaryTree) (int, int) {
    if node == nil {
        return 0, math.MinInt32
    }

    leftPathSum, leftMaxSum := MaxPathSum(node.Left)
    rightPathSum, rightMaxSum := MaxPathSum(node.Right)

    maxChildPathSum := max(leftPathSum, rightPathSum)
    maxPathSumAsRoot := max(node.Value, node.Value+maxChildPathSum)
    maxSumWithSplit := max(maxPathSumAsRoot, node.Value+leftPathSum+rightPathSum)
    maxOverallSum := max(maxSumWithSplit, max(leftMaxSum, rightMaxSum))

    return maxPathSumAsRoot, maxOverallSum
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
