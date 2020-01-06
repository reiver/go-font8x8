package font8x8_test

import (
	"github.com/reiver/go-font8x8"

	"fmt"
)

func ExampleType_Format_u0021() {

	fmt.Printf("%z\n", font8x8.U0021)

	// Output:
	//
	// ░░░██░░░
	// ░░████░░
	// ░░████░░
	// ░░░██░░░
	// ░░░██░░░
	// ░░░░░░░░
	// ░░░██░░░
	// ░░░░░░░░
}

func ExampleType_Format_u0022() {

	fmt.Printf("%z\n", font8x8.U0022)

	// Output:
	//
	// ░██░██░░
	// ░██░██░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
	// ░░░░░░░░
}

func ExampleType_Format_u0023() {

	fmt.Printf("%z\n", font8x8.U0023)

	// Output:
	//
	// ░██░██░░
	// ░██░██░░
	// ███████░
	// ░██░██░░
	// ███████░
	// ░██░██░░
	// ░██░██░░
	// ░░░░░░░░
}

func ExampleType_Format_u0024() {

	fmt.Printf("%z\n", font8x8.U0024)

	// Output:
	//
	// ░░██░░░░
	// ░█████░░
	// ██░░░░░░
	// ░████░░░
	// ░░░░██░░
	// █████░░░
	// ░░██░░░░
	// ░░░░░░░░
}
