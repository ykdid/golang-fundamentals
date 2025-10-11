package main

func main() {

	StructsDemo()
	MethodsDemo()
	EmbeddingDemo()
	
	quad := Quadrilateral{Width: 5, Height: 3}
	circ := Circle{Radius: 4}

	PrintShapeInfo(quad)
	PrintShapeInfo(circ)
}