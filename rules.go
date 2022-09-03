package lsys

func Algae(b byte) []byte {
	switch b {
	case 'A':
		return []byte("AB")
	case 'B':
		return []byte{'A'}
	}
	panic(b)
}

func BinaryTree(b byte) []byte {
	switch b {
	case '0':
		return []byte("1[0]0")
	case '1':
		return []byte("11")
	case '[', ']':
		return []byte{b}
	}
	panic(b)
}

func Cantor(b byte) []byte {
	switch b {
	case 'A':
		return []byte("ABA")
	case 'B':
		return []byte("BBB")
	}
	panic(b)
}

func Koch(b byte) []byte {
	switch b {
	case 'F':
		return []byte("F+F-F-F+F")
	case '+', '-':
		return []byte{b}
	}
	panic(b)
}

func Sierpinski(b byte) []byte {
	switch b {
	case 'F':
		return []byte("F-G+F+G-F")
	case 'G':
		return []byte("GG")
	case '+', '-':
		return []byte{b}
	}
	panic(b)
}
