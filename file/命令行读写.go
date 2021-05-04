// os_args.go
package file

import (
    "fmt"
    "os"
    "strings"
)

func File4() {
    who := "Alice "
    if len(os.Args) > 1 {
        who += strings.Join(os.Args[1:], " ")
    }
    fmt.Println("Good Morning", who)
}