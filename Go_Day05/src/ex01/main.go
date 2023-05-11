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

func unrollGarland(root *Node) [][]bool {
	if root == nil {
		return [][]bool{}
	}

	var result [][]bool
	queue1 := []*Node{root}
	queue2 := []*Node{}
	level := 1

	for len(queue1) > 0 {
		var levelNodes []bool

		for len(queue1) > 0 {
			node := queue1[0]
			queue1 = queue1[1:]

			levelNodes = append(levelNodes, node.Value)

			if node.Left != nil {
				queue2 = append(queue2, node.Left)
			}
			if node.Right != nil {
				queue2 = append(queue2, node.Right)
			}
		}

		if level%2 == 0 {
			reverse(levelNodes)
		}

		result = append(result, levelNodes)
		queue1, queue2 = queue2, queue1
		queue2 = []*Node{}
		level++
	}

	return result
}

func reverse(arr []bool) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
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
	qwe := unrollGarland(root)
	fmt.Println(qwe)
}
