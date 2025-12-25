package sublist

// Relation is the comparison between lists
type Relation string

// Possible relations
const (
	// 两个列表相等
	RelationEqual Relation = "equal"
	// 列表 A 包含 B
	RelationSublist Relation = "sublist"
	// 列表 B 包含 A
	RelationSuperlist Relation = "superlist"
	// 两个列表不相等
	RelationUnequal Relation = "unequal"
)
