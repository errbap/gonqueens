package gonqueens

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/errbap/gonqueens/nqueen"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	size, _ := strconv.Atoi(os.Args[1])

	q := nqueen.Make(size)
	res := q.Solve()

	fmt.Println(res)
	fmt.Println(res.Heuristic())
}