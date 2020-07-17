package main
import ( 
    "fmt"
    "math"
)

func main() {
 
    var a, b int
    var h float64

    fmt.Scan(&b)
    fmt.Scan(&a)

    if a%b == 0 {
        h = float64((a/b) * 2)
    } else {
        h = math.Ceil((float64(a)/float64(b))*2)
    }

    fmt.Println(int64(h))
}
