package main

import (
	"fmt"
)

/**
[매일매일 알고리즘 트레이닝]

lv.3

<그림판 색 채우기>

당신은 그림판의 '색 채우기' 기능을 구현하려한다.

이미지 크기는 제한이 없다. (처리속도 < 3s)

입력 설명
가로 세로
색을 채우기 시작할 점 과 색
이미지의 색상 데이터


입력 예시
10 10
5 5 3
0000000000
0000001000
0000110100
0011000010
0100000010
0100000010
0100000100
0010001000
0001011000
0000100000


출력 예시
0000000000
0000001000
0000113100
0011333310
0133333310
0133333310
0133333100
0013331000
0001331000
0000100000
*/

type Point struct {
	x int
	y int
}

func (p *Point) CheckBoundary() bool {
	if p.x < 0 {
		return false
	}
	if p.y < 0 {
		return false
	}
	if p.x >= width {
		return false
	}

	if p.y >= height {
		return false
	}

	return true
}

func (p *Point) CheckZeroCell() bool {
	if p.CheckBoundary() {
		if nodes[p.x][p.y] == 0 {
			return true
		}
	}
	return false
}
func bfs(x int, y int) {

	queue := make(chan *Point, 10*10)
	queue <- &Point{
		x: x,
		y: y,
	}
	for node := range queue {
		if nodes[node.x][node.y] == 0 {
			nodes[node.x][node.y] = 3
			fmt.Println(nodes)
		}

		up := &Point{
			x: node.x,
			y: node.y + 1,
		}

		down := &Point{
			x: node.x,
			y: node.y - 1,
		}

		left := &Point{
			x: node.x - 1,
			y: node.y,
		}

		right := &Point{
			x: node.x + 1,
			y: node.y,
		}

		if up.CheckZeroCell() {
			queue <- up
		}
		if down.CheckZeroCell() {
			queue <- down
		}
		if left.CheckZeroCell() {
			queue <- left
		}
		if right.CheckZeroCell() {
			queue <- right
		}

		if len(queue) == 0 {
			break
		}
	}
}

var width = 10
var height = 10
var nodes = [][]byte{
	[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	[]byte{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
	[]byte{0, 0, 0, 0, 1, 1, 0, 1, 0, 0},
	[]byte{0, 0, 1, 1, 0, 0, 0, 0, 1, 0},
	[]byte{0, 1, 0, 0, 0, 0, 0, 0, 1, 0},
	[]byte{0, 1, 0, 0, 0, 0, 0, 0, 1, 0},
	[]byte{0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
	[]byte{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
	[]byte{0, 0, 0, 1, 0, 1, 1, 0, 0, 0},
	[]byte{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
}

func main() {

	bfs(5, 5)
}
