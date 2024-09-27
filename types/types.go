package types

type BinaryTree struct {
    Value int         `json:"value"`
    Left  *BinaryTree `json:"left"`
    Right *BinaryTree `json:"right"`
}

type Node struct {
    ID    string `json:"id"`
    Left  string `json:"left"`
    Right string `json:"right"`
    Value int    `json:"value"`
}

type TreeInput struct {
    Nodes []Node `json:"nodes"`
    Root  string `json:"root"`
}

type Input struct {
    Tree TreeInput `json:"tree"`
}

type Output struct {
    MaxPathSum int `json:"maxPathSum"`
}
