package main
import (
  "fmt"
  "math"
)
func maxPoints(deck []int, k int) int{
    left := 0;
    right := len(deck) - k
    var total, best int
    total = 0
    for i := right; i < len(deck); i++ {
        total += deck[i]
    }
    best = total
    for i := 0; i < k; i++ {
        total += deck[left] - deck[right]
        best = int(math.Max(float64(best), float64(total)))
        left++
        right++
    }
    return best
}

func main() {
    deck := []int{5,3,4,4,2,3,2,6,3}
    k := 4
    fmt.Println(maxPoints(deck, k))
}