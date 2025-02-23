package misc

import (
	"log"
	"math/bits"
	"math/rand"
	"os"
	"strings"
)

const (
	RABBITMQ_CON_STRING = "amqp://guest:guest@rabbitmq:5672/"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func BodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func SeverityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s = os.Args[1]
	}
	return s
}

func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandInt(65, 90))
	}
	return string(bytes)
}

func RandInt(min int, max int) int {
	// src := rand.NewSource(time.Now().UnixNano())
	// rng := rand.New(src)
	return min + rand.Intn(max-min)
}

func FibRecursive(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return FibRecursive(n-1) + FibRecursive(n-2)
	}
}

// Memoization map to store computed Fibonacci numbers
var memo = make(map[int]int)

func FibMemoization(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	if val, found := memo[n]; found {
		return val
	}
	memo[n] = FibMemoization(n-1) + FibMemoization(n-2)
	return memo[n]
}

// using optimized tabulation
func Fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	n_2, n_1 := 0, 1
	var curr int

	for i := 2; i <= n; i++ {
		curr = n_1 + n_2
		n_2, n_1 = n_1, curr
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// OVerflow checks:
// If performance is a priority, // use bits.Add()/bits.Mul().
// If you need signed integers, use manual checks.
// If you need arbitrarily large numbers, use math/big.

func SafeAdd(a, b int) (int, bool) {
	sum, carry := bits.Add(uint(a), uint(b), 0)
	if carry != 0 {
		return 0, false
	}
	return int(sum), true
}

func SafeMul(a, b int) (int, bool) {
	product, overflow := bits.Mul(uint(a), uint(b))
	if overflow != 0 {
		return 0, false
	}
	return int(product), true
}

func SafeAdd2(a, b int) (int, bool) {
	if (b > 0 && a > (1<<31-1)-b) || (b < 0 && a < (-1<<31)-b) {
		return 0, false // Overflow detected
	}
	return a + b, true
}

func SafeMul2(a, b int) (int, bool) {
	if a > 0 && b > 0 && a > (1<<31-1)/b {
		return 0, false // Overflow detected
	}
	return a * b, true
}
