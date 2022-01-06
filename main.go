/*package main

import (
    "io"
    "os"
)

func main() {
    // open input file
    fi, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    // close fi on exit and check for its returned error
    defer func() {
        if err := fi.Close(); err != nil {
            panic(err)
        }
    }()

    // open output file
    fo, err := os.Create("output.txt")
    if err != nil {
        panic(err)
    }
    // close fo on exit and check for its returned error
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    // make a buffer to keep chunks that are read
    buf := make([]byte, 1024)
    for {
        // read a chunk
        n, err := fi.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 {
            break
        }

        // write a chunk
        if _, err := fo.Write(buf[:n]); err != nil {
            panic(err)
        }
    }
}*/
package main

import (
    "io/ioutil"
    "fmt"
)

func main() {
    // read the whole file at once
    b, err := ioutil.ReadFile("foo.txt")
    if err != nil {
        panic(err)
    }
	
	/*bufferLength := cap(b)
	fmt.Println("ok: %d",bufferLength.len())*/
	
    // write the whole body at once
    err = ioutil.WriteFile("output.txt", b, 0644)
    if err != nil {
        panic(err)
    }
}
