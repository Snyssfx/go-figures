# Go Figures!
Watch cool figures in your terminal and relax 💆‍💆‍♂️

[![asciicast](https://asciinema.org/a/6XgLYETWWBBWghK1rreE3zq2B.svg)](https://asciinema.org/a/6XgLYETWWBBWghK1rreE3zq2B)

### Running
Install docker and then type in the terminal:
`docker run --rm -ti snyssfx/go-figures:latest`

or `git clone` and `go run .`

### Control
"1"-"9" for changing figures

"+", "-" for changing length of history points

"<", ">" for changin coefficient in figures

### Supported figures
see [./figures.go](./figures.go) for full info
1) circle, coef = radius
2) symmetric cosinus, coef = period
3) sinus with axis, coef = period
4) epicycloid, coef = k (see https://en.wikipedia.org/wiki/Epicycloid)
5) hypocycloid, coef = k (see https://en.wikipedia.org/wiki/Hypocycloid)
