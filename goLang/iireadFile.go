
package main
import (
    "bufio"
    "fmt"
  //  "io"
//    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
//You’ll often want more control over how and what parts of a file are read. For these tasks, start by Opening a file to obtain an os.File value.

    f, err := os.Open("dat")
    check(err)
//Read some bytes from the beginning of the file. Allow up to 5 to be read but also note how many actually were read.

    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))
//The bufio package implements a buffered reader that may be useful both for its efficiency with many small reads and because of the additional reading methods it provides.

    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))
//Close the file when you’re done (usually this would be scheduled immediately after Opening with defer).

    f.Close()
}
