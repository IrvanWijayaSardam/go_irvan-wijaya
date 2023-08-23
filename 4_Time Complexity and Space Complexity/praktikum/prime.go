package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func power(x, y, p *big.Int) *big.Int {
	res := big.NewInt(1)
	zero := big.NewInt(0)

	x.Mod(x, p)

	for y.Cmp(zero) > 0 {
		if y.Bit(0) == 1 {
			res.Mul(res, x)
			res.Mod(res, p)
		}
		y.Rsh(y, 1)
		x.Mul(x, x)
		x.Mod(x, p)
	}

	return res
}

func isCompositeWitness(a, n *big.Int, t, u *big.Int) bool {
	x := power(a, u, n)
	one := big.NewInt(1)
	nMinusOne := new(big.Int).Sub(n, one)

	if x.Cmp(one) == 0 || x.Cmp(nMinusOne) == 0 {
		return false
	}

	for i := big.NewInt(0); i.Cmp(t) < 0; i.Add(i, one) {
		x.Mul(x, x)
		x.Mod(x, n)
		if x.Cmp(nMinusOne) == 0 {
			return false
		}
	}

	return true
}

func isPrimeMillerRabin(n *big.Int, k int) bool {
	if n.Cmp(big.NewInt(2)) < 0 {
		return false
	}

	if n.Cmp(big.NewInt(2)) == 0 || n.Cmp(big.NewInt(3)) == 0 {
		return true
	}

	if new(big.Int).Mod(n, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return false
	}

	one := big.NewInt(1)
	nMinusOne := new(big.Int).Sub(n, one)
	t := big.NewInt(0)
	u := new(big.Int).Set(nMinusOne)

	for u.Bit(0) == 0 {
		t.Add(t, one)
		u.Rsh(u, 1)
	}

	for i := 0; i < k; i++ {
		a := big.NewInt(rand.Int63n(n.Int64() - 2))
		a.Add(a, big.NewInt(2))
		if isCompositeWitness(a, n, t, u) {
			return false
		}
	}

	return true
}

func primeNumber(number int) bool {
	bigNumber := big.NewInt(int64(number))
	return isPrimeMillerRabin(bigNumber, 5) // You can adjust the number of iterations (5 in this case) for a trade-off between accuracy and speed.
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(primeNumber(1000000007)) // Output: true
	fmt.Println(primeNumber(13))         // Output: true
	fmt.Println(primeNumber(17))         // Output: true
	fmt.Println(primeNumber(20))         // Output: false
	fmt.Println(primeNumber(35))         // Output: false

}
