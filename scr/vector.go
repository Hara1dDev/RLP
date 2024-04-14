package rlp

type Vec2 struct {
	X float32
	Y float32
}

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func ProcessVector(vector interface{}) {
	switch vector.(type) {
	case Vec2:
		print(vector, "2")
	case Vec3:
		print(vector, "3")
	default:
		print("No vector recieved")
	}
}
