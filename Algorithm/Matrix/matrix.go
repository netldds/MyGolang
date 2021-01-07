package matrix

type Matrix struct {
	row, column int
	m           [][]float32
}

type Error struct {
	content string
	err     error
}

//Multiply for matrix
func (e *Matrix) Multiply(dst *Matrix) (*Matrix, error) {
	//check value
	if e.column != dst.row {
		return &Matrix{}, Error{content: "cannot multiply due to incorrect matrix struct."}
	}
	finalMatrix := make([][]float32, e.row)
	for temp := 0; temp < e.row; temp++ {
		finalMatrix[temp] = make([]float32, dst.column)
	}
	for i := 0; i < e.row; i++ {
		for j := 0; j < dst.column; j++ {
			// finalMatrix[i][j] = subMultiply(*e, dst, i, j)
			var sum float32
			length := len(e.m[i])
			for x := 0; x < length; x++ {
				sum += e.m[i][x] * dst.m[x][j]
			}
			finalMatrix[i][j] = sum
		}
	}
	res, err := New(finalMatrix)
	return res, err
}
func subMultiply(src, dst Matrix, row, column int) float32 {
	var sum float32
	length := len(src.m[row])
	for i := 0; i < length; i++ {
		sum += src.m[row][i] * dst.m[i][column]
	}
	return sum
}
func (e Error) Error() string {
	return "struct is wrong"
}

//NewBlank return a fresh matrix
func NewBlank(row, column int) *Matrix {
	mt := make([][]float32, column)
	for i := 0; i < len(mt); i++ {
		mt[i] = make([]float32, row)
	}
	return &Matrix{row: row, column: column, m: mt}
}
func New(src [][]float32) (*Matrix, error) {
	ok := validate(src)
	var boo Error
	if ok {
		boo = Error{}
		return &Matrix{row: len(src), column: len(src[0]), m: src}, boo
	}
	boo = Error{content: "struct wrong"}
	return &Matrix{}, boo

}
func validate(src [][]float32) bool {
	for i := 0; i < len(src); i++ {
		length := len(src[i])
		if length == 0 {
			return false
		}
		for j := 0; j < len(src); j++ {
			if i == j {
				continue
			}
			if len(src[i]) != len(src[j]) {
				return false
			}
		}
	}
	return true
}
