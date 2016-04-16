package maze

const codeKey char = 'K'
const codeWall char = 'W'
const codeTrap char = 'T'
const codeEmpty char = ' '

type Maze struct {
	Blocks []char
}