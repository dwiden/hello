package main

import ("fmt"
        "github.com/dwiden/stringstuff"
)

// this process is weird, but thats okay
func main() {
    fmt.Printf(stringstuff.Reverse(stringstuff.Reverse("hello, world\n")))
    fmt.Printf(stringstuff.Reverse("\nfladnaG-\t\n.su ot nevig si taht emit eht htiw od ot tahw si ediced ot evah ew llA"))
}
