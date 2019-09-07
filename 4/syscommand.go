package main

import "os/exec"
import "log"
import "fmt"

func main() {
    cmd := exec.Command("cmd", "/C", "dir")
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
    fmt.Printf("combined out:\n%s\n", string(out))
}