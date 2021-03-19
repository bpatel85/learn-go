package algo

import (
	"errors"
	"fmt"
	"strings"
)

/*
# Zookeper namespace API
#
# * Zookeeper uses a hierarchical namespace that resembles a standard file system.
#
# * In the namespace hierarchy each node can contain data as well as children.
#
# * A "node_path" is a sequence of path elements separated by a slash (/),
#   for example "/broker1/topics/sensorA"
#
# * Clients should be able to create nodes with some data, set the data values on
#   the nodes and retrieve the data.
#
#
# Please start by implementing the following simplified namespace API:
#
#     create(String node_path, String data)
#     set_data(String node_path, String data)
#     get_data(String node_path)
#	  delete(String node_path) -- deletes a node and all its children
#     exists(String node_path) -- tests if a node exists at a location
#     depth() -- returns the namespace depth, defined as the max node depth, where node_depth('/') = 0, node_depth('/broker1') = 1, etc.
#
#
# NOTE: Think about an implementation that is open for extension, based on the
#       kind of uses you think users might have for a hierarchical namespace.
*/

const pathSeperator = "/"

type ZKNode struct {
	key, value string
	children   map[string]*ZKNode
}

func (zNode ZKNode) Depth() int {
	if len(zNode.children) == 0 {
		return 0
	}

	maxDepth := 0
	for _, v := range zNode.children {
		if cDepth := v.Depth(); cDepth > maxDepth {
			maxDepth = cDepth
		}
	}

	return maxDepth + 1
}

type ZK struct {
	root *ZKNode
}

func NewZK() ZK {
	return ZK{root: &ZKNode{children: make(map[string]*ZKNode), key: pathSeperator}}
}

func (node ZK) findNode(path string, crateIfNotPresent bool) (*ZKNode, *ZKNode, error) {
	tokens := strings.Split(path, pathSeperator)
	currNode := node.root
	var parent *ZKNode

	for _, dirName := range tokens {
		if dirName == "" {
			continue
		}

		nextNode, present := currNode.children[dirName]
		if !present {
			if !crateIfNotPresent {
				return nil, nil, errors.New(fmt.Sprintf("%v is not present as path", fmt.Sprintf("%v%v", pathSeperator, dirName)))
			}

			nextNode = &ZKNode{children: make(map[string]*ZKNode), key: dirName}
			currNode.children[dirName] = nextNode
		}

		parent = currNode
		currNode = nextNode
	}

	return currNode, parent, nil
}

func (node ZK) Create(path, data string) error {
	zknode, _, err := node.findNode(path, true)
	if err != nil {
		fmt.Printf("this is not expected. Something went wrong")
		return err
	}

	zknode.value = data
	return nil
}

func (node ZK) SetData(path, data string) error {
	zknode, _, err := node.findNode(path, false)
	if err != nil {
		fmt.Printf("\nNode with path not found. Use create before setting data for path: %v.", path)
		return err
	}

	zknode.value = data
	return nil
}

func (node ZK) GetData(path string) (string, error) {
	zknode, _, err := node.findNode(path, false)
	if err != nil {
		fmt.Printf("\nNode with path not found. Path: %v", path)
		return "", err
	}

	return zknode.value, nil
}

func (node ZK) Exists(path string) bool {
	_, _, err := node.findNode(path, false)
	return err == nil
}

func (node ZK) Depth() int {
	return node.root.Depth()
}

func (node ZK) Delete(path string) bool {
	n, parent, err := node.findNode(path, false)
	if err != nil {
		return false
	}

	delete(parent.children, n.key)
	return true
}
