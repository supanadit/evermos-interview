# Evermos Interview Test

This is my interview test at Evermos which used Go programming language.

## Requirements
- Go 1.14+

## How to set up
- Clone repository with `git clone github.com/supanadit/evermos-interview'
- Install the required dependency with `go get`

### Quick start for "Tennis Player"
- Simply run `go run tennis_player.go`
- Available API
    - `GET` `/player/information` for getting information about rahman
    - `GET` `/player/information/can/play` for getting information wether rahman can play or not
    - `POST` `/player/container/add` for adding new empty ball container for rahman
    - `POST` `/player/container/ball/add` for adding one ball to random empty container for rahman
    - `PUT` `/container/mark/verified` for mark verify to any fully ball container
    
### Quick start for "Kirana Store"
- Simply run `go run kirana_store.go`
- Available API
    - `GET` `/product` for getting list available product
    - `POST` `/product` for create a new product including stock quantity
    - `PUT` `/product` edit existing product and update stock quantity
    - `DELETE` `/product` for delete existing product
    - `POST` `/order` for place new order
    - `GET` `/order` for getting list order
- Format POST, EDIT, DELETE must be JSON and provided at comment `kirana_store.go`
    
### Quick start for "Joni Home's Key"
- Simply run `go run find_key.go`

## Contribute
Please don't this is my own interview test at Evermos