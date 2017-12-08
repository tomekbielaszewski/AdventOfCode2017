package Day_7

import (
	"container/list"
	"strings"
	"strconv"
)

type NodeToProcess struct {
	parent    string
	childName string
	childRow  []string
}

type Node struct {
	parent   *Node
	children *list.List

	name        string
	weight      int
	totalWeight int
}

func GetCorrectWeight(input *list.List) int {
	rows := prepareRows(input)
	rootName, _ := GetBottomNodeName(input)
	rootNode, _ := buildTree(rows, rootName)
	calculateWeights(rootNode)
	unbalancedNode, difference := findUnbalancedWeight(rootNode)
	return unbalancedNode.weight - difference
}

func findUnbalancedWeight(node *Node) (*Node, int) {
	return nil, 0
}

func isDifferent(node *Node) bool {
	siblings := node.parent.children
	isDifferent := true
	for e := siblings.Front(); e != nil; e = e.Next() {
		sibling := e.Value.(*Node)
		if node.name != sibling.name {
			isDifferent = isDifferent && (node.totalWeight-sibling.totalWeight == 0)
		}
	}
	return isDifferent
}

func areTotalWeightsTheSame(children *list.List) bool {
	if children.Len() < 2 {
		return true
	}

	areSame := true
	lastChecked := children.Front().Value.(*Node)

	for e := children.Front(); e != nil; e = e.Next() {
		this := e.Value.(*Node)
		areSame = areSame && this.totalWeight == lastChecked.totalWeight
		lastChecked = this
	}

	return areSame
}

func calculateWeights(root *Node) {
	recursiveCalcWeights(root)
}

func recursiveCalcWeights(node *Node) int {
	if node.children.Len() == 0 {
		return node.weight
	}

	sum := node.weight
	for childNode := node.children.Front(); childNode != nil; childNode = childNode.Next() {
		sum += recursiveCalcWeights(childNode.Value.(*Node))
	}

	node.totalWeight = sum
	return node.totalWeight
}

func prepareRows(input *list.List) map[string][]string {
	rows := make(map[string][]string)

	for e := input.Front(); e != nil; e = e.Next() {
		row := strings.Replace(e.Value.(string), ",", "", -1)
		row = strings.Replace(row, "(", "", 1)
		row = strings.Replace(row, ")", "", 1)

		rowElements := strings.Fields(row)
		rows[rowElements[0]] = rowElements
	}

	return rows
}

func buildTree(rows map[string][]string, rootName string) (*Node, map[string]*Node) {
	nodes := make(map[string]*Node)

	rootNode := createDetachedNode(rows[rootName])
	nodes[rootName] = rootNode

	childrenToProcess := list.New()
	childrenToProcess.PushBackList(getChildrenNodesToProcess(rootName, rows))

	for e := childrenToProcess.Front(); e != nil; e = e.Next() {
		child := e.Value.(*NodeToProcess)

		childNode := createDetachedNode(child.childRow)
		attachNode(childNode, nodes[child.parent])
		nodes[child.childName] = childNode

		if len(getChildrenFromRow(child.childName, rows)) > 0 {
			childrenToProcess.PushBackList(getChildrenNodesToProcess(child.childName, rows))
		}
	}

	return rootNode, nodes
}

func createDetachedNode(row []string) *Node {
	weight, _ := strconv.Atoi(row[1])
	return &Node{name: row[0], weight: weight, children: list.New()}
}

func attachNode(child *Node, parent *Node) {
	parent.children.PushFront(child)
	if child.parent != nil {
		panic("Node was already attached, something's fucky!")
	}
	child.parent = parent
}

func getChildrenNodesToProcess(parent string, rows map[string][]string) *list.List {
	row := rows[parent]
	childrenNames := row[3:]
	children := list.New()

	for _, childName := range childrenNames {
		children.PushFront(&NodeToProcess{parent: parent, childRow: rows[childName], childName: childName})
	}

	return children
}

func getChildrenFromRow(parent string, rows map[string][]string) []string {
	return rows[parent][2:]
}