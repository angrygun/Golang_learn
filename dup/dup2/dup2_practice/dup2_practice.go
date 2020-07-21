package main

import (
    "bufio"
    "fmt"
    "os"
)

type LNFile struct {
    Count     int
    FileNames []string
}

func main() {
    counts := make(map[string]*LNFile)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {

                fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n.Count > 1 {
            fmt.Printf("%d %v\n%s\n", n.Count, n.FileNames, line)
        }
    }
}

func countLines(f *os.File, counts map[string]*LNFile) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        key := input.Text()
        _, ok := counts[key]
        if ok {
            counts[key].Count++
            counts[key].FileNames = append(counts[key].FileNames, f.Name())
        } else {
            counts[key] = new(LNFile)
            counts[key].Count = 1
            counts[key].FileNames = append(counts[key].FileNames, f.Name())
        }
    }
}
