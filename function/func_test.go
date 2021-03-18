package function

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

const (
	N int = 100
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestDivideAndRemainder(t *testing.T) {
	a, b := rand.Intn(N), rand.Intn(N)

	want1, want2 := a / b, a % b

	if got1, got2 := DivideAndRemainder(a, b); got1 != want1 || got2 != want2 {
		t.Errorf("DivideAndRemainder(%d, %d) = %d, %d; got = %d, %d\n", a, b, want1, want2, got1, got2)
	}
}

func TestSum(t *testing.T) {
	a, b := rand.Intn(N), rand.Intn(N)

	want := a + b
	if got := Sum(a, b); got != want {
		t.Errorf("Sum(%d, %d) = %d; got = %d\n", a, b, want, got)
	}
}

func TestSwap(t *testing.T) {
	a, b := rand.Intn(N), rand.Intn(N)

	want1, want2 := b, a
	if got1, got2 := Swap(a, b); got1 != want1 || got2 != want2 {
		t.Errorf("Swap(%d, %d) = %d, %d; got = %d, %d\n", a, b, want1, want2, got1, got2)
	}
}

func TestSqrt(t *testing.T) {
	a := rand.Float64() * float64(N)

	want := math.Sqrt(a)
	if got := Sqrt(a); (got - want) / want > 0.1 {
		t.Errorf("Sqrt(%g) = %g; got = %g\n", a, want, got)
	}
}



func TestVarArgs(t *testing.T) {
	VarArgs(1,"2", 3)
}
