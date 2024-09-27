package handlers

import (
    "net/http"
    . "golang-binary-tree/types"
	"encoding/json"
	helpers "golang-binary-tree/helpers"
)

func HandleMaxPathSum(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

    var input Input
    err := json.NewDecoder(r.Body).Decode(&input)
    if err != nil {
        http.Error(w, "Invalid JSON input", http.StatusBadRequest)
        return
    }

    nodes := make(map[string]Node)
    for _, node := range input.Tree.Nodes {
        nodes[node.ID] = node
    }

    root := helpers.BuildTree(nodes, input.Tree.Root)
    _, maxSum := helpers.MaxPathSum(root)

    output := Output{MaxPathSum: maxSum}
	
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(output)
}
