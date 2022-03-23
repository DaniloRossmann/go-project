package main

import (
	"fmt"

	operation "github.com/DaniloRossmann/go-project/pkg/operations"
)

func main() {
	valorDiff := operation.Diff(6, 2)
	valorSum := operation.Sum(15, 10)
	vari := operation.Variable
	fmt.Printf("%v %v \n %v", valorDiff, valorSum, vari)

}
