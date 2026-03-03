package models

// TreeNode represents a node in the customer navigator tree.
// Used for both "by country" and "by sales rep" tree modes.
type TreeNode struct {
	ID       string      `json:"id"`
	Label    string      `json:"label"`
	Value    interface{} `json:"value"`
	Children []TreeNode  `json:"children,omitempty"`
}
