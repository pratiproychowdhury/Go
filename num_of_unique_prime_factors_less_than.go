package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the primeCount function below.
 */
func primeCount(n int64) int32 {
    /*
     * Write your code here.
     */
    var zzz int64
    zzz = n
    _ = zzz
    result := int32(1)
    for i := int64(2); i <= zzz; i++ {
        if zzz > i {
            zzz = zzz / i
            result++
            fmt.Println(i)
        }
    }
    return result
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        n, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)

        result := primeCount(n)

        fmt.Fprintf(writer, "%d\n", result)
    }

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
