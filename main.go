package main
				//woring example don't change
import (
    "fmt"
    "os"
    "io"
	"github.com/JJT100/greetings/hello" // works  github.com/JJT100/greetings -> has folder hello/greeting.go
//		"github.com/JJT100/greetings/world"
//    	"github.com/amitsaha/using-go-modules/greetings/hello"
//    	"github.com/amitsaha/using-go-modules/greetings/world"
)

func displayGreetings(w io.Writer) {
    fmt.Fprintln(w, greetings.Greet())
//    fmt.Fprintln(w, world.Greet())
}

func main() {   
    displayGreetings(os.Stdout)
}
