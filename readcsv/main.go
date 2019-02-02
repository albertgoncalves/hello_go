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

func process(r *csv.Reader) []Game {
    cropHeader(r)
    var games []Game

    for {
        l, err := r.Read()
        if err == io.EOF {
            break
        }
        panicIf(err)

        games = append(games, Game{
            Date:      l[0],
            AwayTeam:  l[1],
            AwayGoals: readInt(l[2]),
            HomeTeam:  l[3],
            HomeGoals: readInt(l[4]),
        })
    }

    return games
}

func main() {
    csvFile, err := os.Open("playoffs_2018.csv")
    panicIf(err)

    games := process(csv.NewReader(bufio.NewReader(csvFile)))
    for _, g := range games {
        fmt.Println(g)
    }
}
