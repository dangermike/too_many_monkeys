package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/urfave/cli"
)

var verbose = false

func playerPower(player *Player) int {
	p := len(player.cards)
	for _, card := range player.cards {
		if card.IsHidden() {
			p--
		}
	}
	return p
}

func main() {
	app := cli.NewApp()

	app.Action = appMain

	app.Flags = append(app.Flags,
		cli.UintFlag{
			Name:  "p, players",
			Usage: "Number of players in the tournament",
			Value: 6,
		},
		cli.UintFlag{
			Name:  "g, games",
			Usage: "Number of games to play",
			Value: 100000,
		},
		cli.BoolFlag{
			Name:        "v, verbose",
			Usage:       "Verbose logging",
			Destination: &verbose,
		},
	)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func appMain(c *cli.Context) error {
	iterations := c.Int("games")
	workers := runtime.NumCPU()
	games := make([]int, iterations)
	turns := make([]int, iterations)
	wins := make([]uint, c.Int("players"))
	var winLock sync.Mutex
	var wg sync.WaitGroup
	startTime := time.Now()
	for wIx := 0; wIx < workers; wIx++ {
		wg.Add(1)
		go func(wIx int) {
			defer wg.Done()
			var winner int
			localWins := make([]uint, len(wins))
			for i := wIx; i < iterations; i += workers {
				winner, games[i], turns[i] = match(len(wins))
				localWins[winner]++
			}
			winLock.Lock()
			for i := 0; i < len(localWins); i++ {
				wins[i] += localWins[i]
			}
			winLock.Unlock()
		}(wIx)
	}
	wg.Wait()
	durationMs := time.Since(startTime).Seconds() * 1000
	fmt.Printf("%d matches in %0.3fms (%0.3fms/match)\n", iterations, durationMs, durationMs/float64(iterations))
	for ix, winCnt := range wins {
		fmt.Printf("Player %d: %d wins (%0.2f%%)\n", ix, winCnt, float64(100*winCnt)/float64(iterations))
	}
	return nil
}

func match(numPlayers int) (winner int, games int, turns int) {
	players := make([]*Player, numPlayers)
	for i := 0; i < len(players); i++ {
		players[i] = &Player{id: i, cards: make([]Card, 6)}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	i := r.Intn(numPlayers)
	for len(players[i].cards) > 0 {
		games++
		game := NewGame(players, r)
		for i = r.Intn(numPlayers); !players[i].IsDone(); i = (i + 1) % numPlayers {
			turns++
			players[i].Turn(game)
		}
		printf("Player %d has won the round!\n", i)
		players[i].cards = players[i].cards[1:]
	}
	// printf("Player %d has won the match!!!\n", i)
	return i, games, turns
}

func print(a ...interface{}) (n int, err error) {
	if !verbose {
		return 0, nil
	}
	return fmt.Print(a...)
}

func println(a ...interface{}) (n int, err error) {
	if !verbose {
		return 0, nil
	}
	return fmt.Println(a...)
}

func printf(format string, a ...interface{}) (n int, err error) {
	if !verbose {
		return 0, nil
	}
	return fmt.Printf(format, a...)
}
