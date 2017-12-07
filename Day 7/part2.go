package Day_7

import (
	"container/list"
	"strings"
	"strconv"
)

func GetCorrectWeight(input *list.List) (int, error) {
	rows := make(map[string][]string)

	for e := input.Front(); e != nil; e = e.Next() {
		row := strings.Replace(e.Value.(string), ",", "", -1)
		row = strings.Replace(row, "(", "", 1)
		row = strings.Replace(row, ")", "", 1)

		rowElements := strings.Fields(row)
		rows[rowElements[0]] = rowElements
	}

	rootName, _ := GetBottomNodeName(input)
	/*rootNode, nodes := */ buildTree(rows, rootName)

	return 0, nil
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

type Node struct {
	parent   *Node
	children *list.List

	name   string
	weight int
}

type NodeToProcess struct {
	parent    string
	childName string
	childRow  []string
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
