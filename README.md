# Too Many Monkeys™ -- simulating the game from Gamewright

This is a simulation of the game [Too Many Monkeys](https://gamewright.com/product/Too-Many-Monkeys). This is a little different from other games I've simulated in that there are decisions, so alternate strategies are theoretically allowed. As a practical matter, there appears to be one correct strategy, which is what I've implemented here. So I've deftly ducked the need for ML once again!

One area where there is an opportunity for real-life players to beat the odds here is in ganging up -- Sending all the negative cards to a single player. This simulation gives each player the same heuristic (number of showing cards) to determine which player poses the greatest threat. It is that player who gets negative cards. In my house, I am apparently perceived is the greatest threat because I seem to get the lion's share of the negative cards. I am, in fact, a lion, so perhaps that makes sense.

There aren't any parameters to the code at this time, so you will need to modify `main` to change the number of players or iterations.

## Results (on my laptop):

```
100000 matches in 4627.063ms (0.046ms/match)
Player 0: 17670 wins (17.67%)
Player 1: 17753 wins (17.75%)
Player 2: 17346 wins (17.35%)
Player 3: 16908 wins (16.91%)
Player 4: 15624 wins (15.62%)
Player 5: 14699 wins (14.70%)
```
