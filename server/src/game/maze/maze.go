package maze

import (
    "fmt"
    "io/ioutil"
    "os"
	"math"
)
 
func round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

const Key byte = 'K'
const Wall byte = '1'
const Trap byte = 'T'
const Empty byte = '0'

type Maze struct {
	blocks []byte
	w int
	h int
}

var maze Maze

func GetHeight() int {
	return maze.h
}
func GetWidth() int {
	return maze.w - 1 // '\n' doesn't count
}

func IsWalkable(c byte) bool {
	fmt.Printf("walkable: %b ; '0' == %b ;  '1' == %b\n", c, '0', '1')
	return (c == Empty || c == Trap)
}

func GetI(x int, y int) byte {
	return maze.blocks[x + y * (maze.w)]
}
func GetF(x float64, y float64) byte {
	return maze.blocks[	int(round(x, 0.5, 0) + 
						round(y, 0.5, 0) *	float64(maze.w))]
}

func GetWalkable(currentX float64, currentY float64, x float64, y float64) (float64, float64) {
	// FIXME: this is the ugliest move function I've ever done
	tryX := round(x, 0.5, 0)
	tryY := round(y, 0.5, 0)
	if (IsWalkable(maze.blocks[	int(tryX + 
						tryY * float64(maze.w))])) {
		return x, y
	} else {
		if (IsWalkable(maze.blocks[	int(currentX + 
						tryY *	float64(maze.w))])) {
			return currentX, y
		} else if (IsWalkable(maze.blocks[	int(x + 
						currentY *	float64(maze.w))])) {
			return x, currentY
		} else {
			return currentX, currentY
		}
	}
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func Parse(file string) {
	dat, err := ioutil.ReadFile(file)
    check(err)
    fmt.Print(string(dat))
	
	f, err := os.Open(file)
    check(err)
	maxLength := 1024
	b1 := make([]byte, maxLength)
    n1, err := f.Read(b1)
    check(err)
	if (n1 > maxLength) {
		panic("file too long")
	}
    fmt.Printf("%d bytes: %s\n", n1, string(b1))
	for i := 0; i < n1; i++ {
		fmt.Printf("bytes[%i]: %c\n", i, b1[i])
		
		if (b1[i] == '\n') {
			maze.w = i;
			break;
		}
	}
	maze.h = n1 / maze.w
	maze.w = maze.w + 1
	maze.blocks = b1
	fmt.Printf("height: %d, width: %d\n", maze.h, maze.w - 1)
}