package Day_7

import (
	"container/list"
	"strings"
	"strconv"
	"github.com/tomekbielaszewski/AdventOfCode2017/day8"
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
	outOfBalanceNode := findOutOfBalanceNode(rootNode)
	difference := calculateTheDifference(outOfBalanceNode)
	return outOfBalanceNode.weight + difference
}

func calculateTheDifference(outOfbalanceNode *Node) int {
	outOfBalanceNodeParent := outOfbalanceNode.parent
	childrenSum := outOfBalanceNodeParent.totalWeight - outOfBalanceNodeParent.weight
	balancedChildrenSum := childrenSum - outOfbalanceNode.totalWeight
	balancedChildWeight := balancedChildrenSum / (outOfBalanceNodeParent.children.Len() - 1)
	difference := balancedChildWeight - outOfbalanceNode.totalWeight
	return difference
}

func findOutOfBalanceNode(node *Node) *Node {
	different := getDifferent(node.children)
	if different != nil {
		lastUnbalanced := findOutOfBalanceNode(different)
		return lastUnbalanced
	}
	return node
}

func getDifferent(children *list.List) *Node {
	if children == nil || areTheSame(children) {
		return nil
	}
	type CountedNode struct {
		node  *Node
		count int
	}
	weights := make(map[int]*CountedNode)

	for e := children.Front(); e != nil; e = e.Next() {
		child := e.Value.(*Node)
		weight := child.totalWeight
		if weights[weight] == nil {
			weights[weight] = &CountedNode{node: child, count: 1}
		} else {
			weights[weight].count++
		}
	}

	min := Day_8.MaxInt
	var differentNode *Node

	for _, v := range weights {
		if v.count < min {
			differentNode = v.node
			min = v.count
		}
	}

	return differentNode
}

func areTheSame(children *list.List) bool {
	same := true
	first := children.Front().Value.(*Node)
	for e := children.Front(); e != nil; e = e.Next() {
		same = same && first.totalWeight == e.Value.(*Node).totalWeight
	}
	return same
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
