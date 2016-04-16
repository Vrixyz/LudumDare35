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
const Wall byte = 'W'
const Trap byte = 'T'
const Empty byte = ' '

type Maze struct {
	Blocks []byte
	w int
	h int
}

var maze Maze

func GetI(x int, y int) byte {
	return maze.Blocks[x + y * maze.w]
}
func GetF(x float64, y float64) byte {
	return maze.Blocks[	int(round(x, 0.5, 10) + 
						round(y, 0.5, 10) *	float64(maze.w))]
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
}