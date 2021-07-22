package main
import (
"fmt"
"os"
"cap2/cmd"
)
func main() {
wd, err := os.Getwd()
if err != nil {
fmt.Fprintf(os.Stderr, "unable to get working directory: %s",
err.Error())
}
res, err := cmd.ExecuteLs(wd)
if err != nil {
fmt.Fprintf(os.Stderr, "unable list files in %s, error: %s",
wd, err.Error())
}
fmt.Printf("%s\n", res)
}
