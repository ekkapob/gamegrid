# Game Grid

Board Game with random game pieces in which a player can move left, right, up and down inside to collect game pieces' scores.

![Image of GameGrid](https://raw.githubusercontent.com/ekkapob/gamegrid/master/screen.gif)

### Game Rules
  - Grid size is defined in configuration file - config/config.json
  - Game pieces are randomly initialized into the board using game types set in the configuration
  - Player is randomly positioned in the grid and represented with *
  - The player can move (left, right, up, down) inside the board, restart the game, and exit
  - Total score represents accumulated game pieces' scores the player obtains when traveling inside the board
  - A game piece at the player's position will be removed from the board after the player moved to next cell

### Game Usage
  - Move command - [a]: left, [d]: right, [w]: up, [s]: down
  - Restart [r]
  - Exit [x]
 
### Technical Aspect
  - The application is written in GO Programming Language (compiled language), [GoLang]
  - Configuration 'config/config.json' is required
[GoLang]:https://golang.org/

### Game Configuration
**config/config.json**
```json
{
  "gridRowColumn": "8x8",
  "maxRow": 20,
  "maxColumn": 20,
  "allowEmptyGame": true,
  "games": [
    {
      "name": "A",
      "score": "8"
    },
    {
      "name": "B",
      "score": "y-pos"
    },
    {
      "name": "C",
      "score": "x-pos"
    }
  ]
}
```
 - *gridRowColumn* - Grid size row x column, default 8 x 8
 - *maxRow* - Maximum grid row
 - *maxColumn* - Maximum grid column
 - *allowEmptyGame* - Allow game board to have empty game pieces
 - *games* - List of game pieces which will be randomly selected to draw on the board
  - *name* - Game piece's name
  - *score* - Interger for fixed score, "y-pos" or "x-pos" to use game piece's row or column as the score

### How to Run the Game
As developed with GO, the application needs to be compiled before running.
##### 1. OS X 64-bit (Yosemite) Executable
  - https://github.com/ekkapob/gamegrid/raw/master/runnable/gamegrid-osx-64bit.zip
```
# gamegrid-osx-64bit directory
$ ./gamegrid
```

For other operating systems, source code compilation is needed. 

##### 2. Compile from source
1. Install GO from [Golang]
2. Create folder structure and Set up environment
```
$ mkdir $HOME/go
$ cd $HOME/go
$ mkdir bin pkg src
$ export GOPATH=$HOME/go
$ export GOBIN=$GOPATH/bin
$ export PATH=$PATH:$GOBIN
# Clone git repository
$ go get github.com/ekkapob/gamegrid
$ cd src/github.com/ekkapob/gamegrid
# Compile the application. The executable is installed at GOBIN path.
$ go install
$ gamegrid
```
