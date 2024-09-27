package types

type Node struct {
    ID    string  `json:"id"`
    Left  *string `json:"left"`
    Right *string `json:"right"`
    Value int     `json:"value"`
}

type Tree struct {
    Nodes []Node `json:"nodes"`
    Root  string `json:"root"`
}
