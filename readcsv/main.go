package main

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "strconv"
)

type Game struct {
    Date      string
    AwayTeam  string
    AwayGoals int
    HomeTeam  string
    HomeGoals int
}

func panicIf(err error) {
    if err != nil {
        panic(err)
    }
}

func readInt(s string) int {
    x, err := strconv.Atoi(s)
    panicIf(err)
    return x
}

func typeOf(v interface{}) string {
    return fmt.Sprintf("%T", v)
}

func readCsv(fn string) *os.File {
    csvFile, err := os.Open(fn)
    panicIf(err)

    return csvFile
}

func cropHeader(r *csv.Reader) {
    _, err := r.Read()
    panicIf(err)
}

func parseRow(row []string) Game {
    return Game{
        Date:      row[0],
        AwayTeam:  row[1],
        AwayGoals: readInt(row[2]),
        HomeTeam:  row[3],
        HomeGoals: readInt(row[4]),
    }
}

func process(r *csv.Reader) []Game {
    cropHeader(r)

    var games []Game
    for {
        row, err := r.Read()
        if err != nil {
            if err == io.EOF {
                break
            } else {
                panic(err)
            }
        }

        games = append(games, parseRow(row))
    }

    return games
}

func extractGoals(games []Game) ([]int, []int) {
    n := len(games)
    awayGoals := make([]int, n)
    homeGoals := make([]int, n)

    for i, g := range games {
        awayGoals[i] = g.AwayGoals
        homeGoals[i] = g.HomeGoals
    }

    return awayGoals, homeGoals
}

func printAll(xs ...interface{}) {
    for _, x := range xs {
        fmt.Println(x)
    }
}

func main() {
    csvFile := readCsv("playoffs_2018.csv")
    defer csvFile.Close()

    printAll(extractGoals(process(csv.NewReader(bufio.NewReader(csvFile)))))
}
