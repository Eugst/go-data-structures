package main

import "fmt"

type Knight struct {
	x int
	y int
}

type Board struct {
	xMin     int
	yMin     int
	xMax     int
	yMax     int
	path     []int
	path_len int
}

type Distanation struct {
	x int
	y int
}

func (k *Knight) moveTopLeft(b *Board) {
	if k.x-1 < b.xMin || k.y+2 > b.yMax {
		fmt.Println("Out of bounds")
		return
	}
	k.x -= 1
	k.y += 2
}

func (k *Knight) moveTopRight(b *Board) {
	if k.x+1 > b.xMax || k.y+2 > b.yMax {
		fmt.Println("Out of bounds")
		return
	}
	k.x += 1
	k.y += 2
}

func (k *Knight) moveBottomLeft(b *Board) {
	if k.x-1 < b.xMin || k.y-2 < b.yMin {
		fmt.Println("Out of bounds")
		return
	}
	k.x -= 1
	k.y -= 2
}

func (k *Knight) moveBottomRight(b *Board) {
	if k.x+1 > b.xMax || k.y-2 < b.yMin {
		fmt.Println("Out of bounds")
		return
	}
	k.x += 1
	k.y -= 2
}

func (k *Knight) moveLeftTop(b *Board) {
	if k.x-2 < b.xMin || k.y+1 > b.yMax {
		fmt.Println("Out of bounds")
		return
	}
	k.x -= 2
	k.y += 1
}

func (k *Knight) moveLeftBottom(b *Board) {
	if k.x-2 < b.xMin || k.y-1 < b.yMin {
		fmt.Println("Out of bounds")
		return
	}
	k.x -= 2
	k.y -= 1
}

func (k *Knight) moveRightTop(b *Board) {
	if k.x+2 > b.xMax || k.y+1 > b.yMax {
		fmt.Println("Out of bounds")
		return
	}
	k.x += 2
	k.y += 1
}

func (k *Knight) moveRightBottom(b *Board) {
	if k.x+1 > b.xMax || k.y-2 < b.yMin {
		fmt.Println("Out of bounds")
		return
	}
	k.x += 2
	k.y -= 1
}

func (b *Board) logMove(x, y int) {
	b.path[hashMove(x, y)] = 1
	b.path_len += 1
}

func hashMove(x, y int) int {
	return (x+y)*(x+y+1)/2 + y
}

func (k *Knight) chooseMove(d *Distanation, b *Board) {
	if d.x > k.x {
		if d.y > k.y {
			if (d.x - k.x) > (d.y - k.y) {
				k.moveRightTop(b)
				fmt.Println("Right Top")
			} else {
				k.moveTopRight(b)
				fmt.Println("Top Right")
			}
		} else {
			if (d.x - k.x) > (k.y - d.y) {
				k.moveRightBottom(b)
				fmt.Println("Right Bottom")
			} else {
				k.moveBottomRight(b)
				fmt.Println("Bottom Right")
			}
		}
	} else {
		if d.y > k.y {
			if (k.x - d.x) > (d.y - k.y) {
				k.moveLeftTop(b)
				fmt.Println("Left Top")
			} else {
				k.moveTopLeft(b)
				fmt.Println("Top Left")
			}
		} else {
			if (k.x - d.x) > (k.y - d.y) {
				k.moveLeftBottom(b)
				fmt.Println("Left Bottom")
			} else {
				k.moveBottomLeft(b)
				fmt.Println("Bottom Left")
			}
		}
	}
}

func (k *Knight) moveTo(d *Distanation, b *Board) int {
	if d.x == k.x && d.y == k.y {
		return 0
	}
	k.chooseMove(d, b)
	if b.path[hashMove(k.x, k.y)] == 1 {
		fmt.Println("Already visited")
		return 1
	}

	b.logMove(k.x, k.y)
	return k.moveTo(d, b)
}

func main() {
	var k Knight
	var d Distanation
	var b Board
	k.x = 1
	k.y = 1
	d.x = 7
	d.y = 7
	b.xMin = 1
	b.yMin = 1
	b.xMax = 8
	b.yMax = 8
	b.path = make([]int, b.xMax*b.xMax+b.yMax*b.yMax+b.xMax*b.yMax)
	b.path_len = 0
	if k.moveTo(&d, &b) == 0 {
		fmt.Println("")
		fmt.Print("Result: Reached in ")
		fmt.Println(b.path_len)
	} else {
		fmt.Println("")
		fmt.Println("Result: Not reached")
	}
}
