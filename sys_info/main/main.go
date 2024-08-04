package main

import (
	"fmt"
	"runtime"
)

func main() {
    var mem runtime.MemStats
    runtime.ReadMemStats(&mem)

    fmt.Printf("Number of CPUs : %d \n", runtime.NumCPU())
    fmt.Printf("Total Allocated memory : %d \n", mem.TotalAlloc)
    fmt.Printf("Number of memory allocations : %d \n", mem.Mallocs)
}

