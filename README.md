# go-binary-tree-server

A web server that accepts a Binary Tree and returns its max path sum.
A path is a collection of connected nodes in a tree, where no node is connected to more than
two other nodes; a path sum is the sum of the values of the nodes in a particular path.

Sample Input:
```
tree =
    1
    / \
   2   3
  / \ / \
 4  5 6  7
```

Sample Output:
18 // 5 + 2 + 1 + 3 + 7

- - -

Example Request Body for 'POST /max-path-sum'
```
    {
        "tree": {
            "nodes": [
                {"id": "1", "left": "2", "right": "3", "value": 1},
                {"id": "3", "left": "6", "right": "7", "value": 3},
                {"id": "7", "left": null, "right": null, "value": 7},
                {"id": "6", "left": null, "right": null, "value": 6},
                {"id": "2", "left": "4", "right": "5", "value": 2},
                {"id": "5", "left": null, "right": null, "value": 5},
                {"id": "4", "left": null, "right": null, "value": 4}
            ],
            "root": "1"
        }
    }
```

Response Format: 

```
{
    "maxPathSum": 18
}
```
