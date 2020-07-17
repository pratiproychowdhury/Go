/*
Luke is daydreaming in Math class. He has a sheet of graph paper with n rows and m columns, and he imagines that there is an army base 
in each cell for a total of n*m bases. He wants to drop supplies at strategic points on the sheet, marking each drop point with a red dot. 
If a base contains at least one package inside or on top of its border fence, 
then it's considered to be supplied
*/

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
 * Complete the gameWithCells function below.
 */
func gameWithCells(n int32, m int32) int32 {
    /*
     * Write your code here.
     */
    var drops int32

    if n%2 == 0 {
        if m%2 == 0 {
            drops = n*m/4
        } else {
            drops = (n/2)*(m/2) + n/2
        }
    } else {
        if m%2 == 0 {
            drops = (n/2)*(m/2) + m/2
        } else {
            drops = (n/2)*(m/2) + n/2 + m/2 + 1
        }
    }
    return drops
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nm := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nm[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(nm[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    result := gameWithCells(n, m)

    fmt.Fprintf(writer, "%d\n", result)

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
