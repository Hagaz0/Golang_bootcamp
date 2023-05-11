package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Value bool
	Left  *Node
	Right *Node
}

func countTrueNodes(root *Node) int {
	if root == nil {
		return 0
	}
	count := 0
	if root.Value {
		count = 1
	}
	count += countTrueNodes(root.Left) + countTrueNodes(root.Right)
	return count
}

func printTree(root *Node, level int) {
	if root == nil {
		return
	}

	printTree(root.Right, level+1)

	for i := 0; i < level; i++ {
		fmt.Print("    ")
	}

	if root.Value {
		fmt.Println("1")
	} else {
		fmt.Println("0")
	}

	printTree(root.Left, level+1)
}

func buildTreeFromInput() *Node {
	var root *Node
	queue := []*Node{}

	fmt.Print("Введите корневое значение: ")
	var str string
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	str = sc.Text()

	rootValue, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Println("Введено не bool значение")
		os.Exit(1)
	}

	root = &Node{Value: rootValue}

	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left == nil {
			fmt.Printf("Введите правое дочернее значение (true/false или -): ")

			var str string
			sc := bufio.NewScanner(os.Stdin)
			sc.Scan()
			str = sc.Text()

			if str == "" {
				break
			} else if str != "-" {
				leftValue, err := strconv.ParseBool(str)
				if err != nil {
					fmt.Println("Введено не bool значение")
					os.Exit(1)
				} else {
					node.Left = &Node{Value: leftValue}
					queue = append(queue, node.Left)
				}
			}
		}

		if node.Right == nil {
			fmt.Printf("Введите правое дочернее значение (true/false или -): ")

			var str string
			sc := bufio.NewScanner(os.Stdin)
			sc.Scan()
			str = sc.Text()

			if str == "" {
				break
			} else if str != "-" {

				rightValue, err := strconv.ParseBool(str)
				if err != nil {
					fmt.Println("Введено не bool значение")
					os.Exit(1)
				} else {
					node.Right = &Node{Value: rightValue}
					queue = append(queue, node.Right)
				}
			}
		}
	}
	return root
}

func main() {
	root := buildTreeFromInput()
	printTree(root, 0)
	fmt.Println(countTrueNodes(root.Left) == countTrueNodes(root.Right))
}
