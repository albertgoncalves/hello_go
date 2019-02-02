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

func main() {
    csvFile, err := os.Open("playoffs_2018.csv")
    panicIf(err)
    defer csvFile.Close()

    games := process(csv.NewReader(bufio.NewReader(csvFile)))
    for _, g := range games {
        fmt.Println(g)
    }
}
