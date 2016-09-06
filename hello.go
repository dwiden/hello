package main

import ("fmt"
        "github.com/dwiden/stringstuff"
        "math/rand"
)

func swap(first, second int) (int, int) {
    return second, first
}

// better than using named return variables
// playing around with functions
func add(x int, y int) int {
    fmt.Printf("Performing %d + %d = %d\n", x, y, x + y)	
    return x + y
}

// this process is weird, but thats okay
func main() {
    var x, y int
    x = 3
    y = 5
    add(x, y)

    // apparently this is not legit
    // x, y := swap(x, y)
    // add(x, y)

    fmt.Printf(stringutil.Reverse(stringutil.Reverse("hello, world\n")))
    fmt.Printf(stringutil.Reverse("\nfladnaG-\t\n.su ot nevig si taht emit eht htiw od ot tahw si ediced ot evah ew llA"))
    fmt.Printf("Everyone likes the number: %d\n", rand.Intn(10))
}
