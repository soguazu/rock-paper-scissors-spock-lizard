# Boilerplate

Company service manages everything that has to do with a company.

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