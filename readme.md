# Rock, Paper, Scissor, Spock, Lizard

Sam invented this game (with Karen Bryla) because it seems like when you know someone well enough, 
75-80% of any Rock-Paper-Scissors games you play with that person end up in a tie. Well, here is a slight 
variation that reduces that probability. (Note that for those of you who like to swing your fist back and 
forth and say, "Rock, Paper, Scissors, GO!", might want to continue to do that, replacing "Rock" with "One," 
"Paper" with "Two," and "Scissors" with "Three.") This version is also nice because it satisfies the Law of
Fives.

![](./img.png)

### Links
- [How to play the game](https://www.youtube.com/watch?v=zjoVuV8EeOU)
- [Big Bang Theory](https://www.youtube.com/watch?v=x5Q6-wMx-K8&t=49s)


## NOTE 
A user id is generated a returned to the client, please pass that id in the header as Playid, 
it is used to keep track of the player

## Built with:
- [Golang](https://go.dev/dl/)
- [Docker](https://www.docker.com/products/docker-desktop/)

## Installation

Use the package manager [go modules](https://go.dev/blog/using-go-modules) to install all dependencies.

```bash
git clone https://github.com/soguazu/rock-paper-scissors-spock-lizard.git
```

```bash
cd rock-paper-scissors-spock-lizard
```

```bash
go mod download
```

```bash
touch .env
```
Copy the value inside .env.sample into the .env and fill the values for the necessary config


## Usage with Makefile

```bash
# To run build swagger docs and run the server
make swagger
```

### Visit [swagger docs](http://localhost:8085/swagger/index.html)

## Usage with docker compose
### Note make sure you have docker running
```bash
# To run build swagger docs and run the server
make start
```
### Visit [swagger docs](http://localhost:8085/swagger/index.html)

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)