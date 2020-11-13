# SwapiGo - Star Wars in Go


I wanted to learn Go. I also wanted to play with Tailwind CSS, so I built a website that fetches data from the Star Wars API and displays an HTML Page with Tailwind CSS classes. You can view pages of people, films, vehicles and starships. You can select a unit and view greater details and all associations of that unit. Initial pages are fetched using Go's WaitGroup and page associations are fetched concurrently using channels. I thought it was cool to build an entire website without a single line of JavaScript, haha. I like tailwind CSS and I really like Go.

## Installation

### `install go ` (I am running 1.14.6)

### `go run swapi.go` (from root directory)

### `visit localhost:8080`

![People List](https://user-images.githubusercontent.com/30199861/99013172-288c0680-250d-11eb-99dc-fc2d0b5ad547.png)
![Vehicles List](https://user-images.githubusercontent.com/30199861/99013183-2f1a7e00-250d-11eb-8ab8-982a7f6d2a4c.png)
![Film Detail Page](https://user-images.githubusercontent.com/30199861/99013184-30e44180-250d-11eb-9b4c-ad4547865b50.png)
