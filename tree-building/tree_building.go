package tree

import (
	"fmt"
	"sort"
)

// Record entry
type Record struct {
	ID, Parent int
}

// Node (tree structure)
type Node struct {
	ID       int
	Children []*Node
}

// Build ([]Records) *Node; builds the tree
func Build(records []Record) (*Node, error) {
	if len(records) == 0 { // return for empty
		return nil, nil
	}
	sort.SliceStable(records, func(i, j int) bool { // sort records on ID value ascending keep duplicates
		return records[i].ID < records[j].ID
	})
	if records[0].ID != 0 || records[0].ID == 0 && records[0].Parent != 0 { // check for valid root
		return nil, fmt.Errorf("no valid root node")
	}
	// at least root present
	tree := make(map[int]*Node)

	// deal with rest of Records
	for i, rec := range records {
		if rec.ID != i || (rec.ID != 0 && rec.ID == rec.Parent) || rec.Parent > rec.ID {
			return nil, fmt.Errorf("invalid record or sequence detected")
		}
		// add valid record to tree
		tree[rec.ID] = &Node{ID: rec.ID}
		if rec.ID != 0 { // don't add root in its children ...cause ID == Parent == 0
			tree[rec.Parent].Children = append(tree[rec.Parent].Children, tree[rec.ID])
		}
	}
	return tree[0], nil
}
