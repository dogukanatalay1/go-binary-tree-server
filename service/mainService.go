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

/* This function returns sum of max path of binary tree */
func MaxPathSum(node *BinaryTree) (int, int) {
    if node == nil {
        return 0, math.MinInt32
    }

    leftPathSum, leftMaxSum := MaxPathSum(node.Left)
    rightPathSum, rightMaxSum := MaxPathSum(node.Right)

    maxChildPathSum := math.Max(float64(leftPathSum), float64(rightPathSum))
    maxPathSumAsRoot := math.Max(float64(node.Value), float64(node.Value) + maxChildPathSum)
    maxSumWithSplit := math.Max(maxPathSumAsRoot, float64(node.Value) + float64(leftPathSum) + float64(rightPathSum))
    maxOverallSum := math.Max(maxSumWithSplit, math.Max(float64(leftMaxSum), float64(rightMaxSum)))

    return int(maxPathSumAsRoot), int(maxOverallSum)
}
