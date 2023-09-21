package calculator_test

import (
	"log"
	"os"
	"testing"

	"github.com/cecardev/tdd-in-go/chapter02/calculator"
)

func TestAdd(t *testing.T) {
	defer func() {
		log.Printf("Defered tiering down.")
	}()

	//Arrange
	e := calculator.Engine{}
	x, y := 2.5, 3.0
	want := 5.5

	actAssert := func(x, y, want float64) {
		//Act
		result := e.Add(x, y)
		//Assert
		if result != want {
			t.Errorf("Add(%.2f, %.2f) = %.2f; want 5.5", x, y, result)
		}

	}

	t.Run("Add two positive numbers", func(t *testing.T) {
		actAssert(x, y, want)
	})

	t.Run("Add two negative numbers", func(t *testing.T) {
		actAssert(-x, -y, -want)
	})

}

func TestMain(m *testing.M) {
	// setup statements
	setup()
	// run the tests
	e := m.Run()
	// cleanup statements
	teardown()
	// report the exit code
	os.Exit(e)
}
func setup() {
	log.Println("Setting up.")
}
func teardown() {
	log.Println("Tearing down.")
}

func init() {
	log.Println("Init setup.")
}

func BenchmarkAdd(b *testing.B) {
	e := calculator.Engine{}
	// run the Add function b.N times
	for i := 0; i < b.N; i++ {
		e.Add(2, 3)
	}
}
