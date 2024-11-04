# Setter Guessing Game

This is a simple guessing game meant to display how strings are stored in memory as sequences of bytes.

The game is written in Go to take advantage of its `rune` type, which is a UTF-8 encoded Unicode code point. For more information, see my write-up [here]().

## How to Play

1. Clone the repository
2. Run the game with `go run .`
3. Enter characters to try to guess the secret letter. The game will tell you if you are too high or too low.

## Future Work

This game is a simple demonstration of how strings are stored in memory. I plan to expand on this concept by creating a more complex, PVP word guessing game using websockets.
