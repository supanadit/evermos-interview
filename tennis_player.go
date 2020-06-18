package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/** This is the first test, and the challenge description based on PDF look like this

Rahman is a tennis player.
Rahman wants ball container to be full of balls before he’s going out to play.
Rahman will load ball randomly into the ball container.
If:
a) Rahman has many ball containers (indefinite).
b) Rahman can put one ball into one container at the time.
c) Only one container will be mark as verified, and Rahman will stop put/remove ball until it verified.
d) Verified mark only when the ball container is full.
e) You must create functional test, either the ball container is full or not.

How you can handle this in on programming?
Can you write the API based on those problem?

*/

// I create struct for player
type tennisPlayer struct {
	// CustomerName of player
	Name string `json:"name"`
	// Based on point (a) Rahman has many ball containers, which is rahman has many ball containers meaning the player can have more ball container
	BallContainer []ballContainer `json:"ball_container"`
}

func (player *tennisPlayer) addEmptyBallContainer() {
	player.BallContainer = append(player.BallContainer, ballContainer{})
}

// This function will tells point (b) Rahman can put one ball into one container at the time. which is randomly correct ?
func (player *tennisPlayer) putBallIntoContainer() {
	// So because the point (b) says rahman can put one ball into one container at the time, meaning i don't need write parameter for total ball because just one ball and randomly
	successPutBallIntoContainer := false
	for i, x := range player.BallContainer {
		if !isFull(x) {
			x.putBall()
			successPutBallIntoContainer = true
			player.BallContainer[i] = x
			break
		}
	}
	if !successPutBallIntoContainer {
		fmt.Printf("Sorry %s you can't put any ball, because no empty ball container is available \n", player.Name)
	}
}

// This function tells Rahman wants ball container to be full of balls before he’s going out to play.
func (player *tennisPlayer) canPlay() bool {
	return isAllBallContainerFull(player.BallContainer)
}

func (player *tennisPlayer) verifiedBallContainer() {
	for _, x := range player.BallContainer {
		if !isVerified(x) && isFull(x) {
			x.verified()
		}
	}
}

type ballContainer struct {
	Verified  bool `json:"is_verified"`
	TotalBall int  `json:"total_ball"`
}

func (b *ballContainer) putBall() {
	b.TotalBall += 1
}

func (b *ballContainer) verified() {
	b.Verified = true
}

func main() {
	// First question is "How you can handle this in on programming"

	// I create a player based on struct of tennisPlayer, so it clear tell rahman is a "Tennis Player"
	rahman := tennisPlayer{
		Name: "Rahman",
	}

	fmt.Printf("CustomerName of Tennis Player : %s \n", rahman.Name)

	// Rahman wants ball container to be full of balls before he’s going out to play, so it tells rahman should fill the ball container
	// Maximum total ball in a Ball Container based on real life is 3
	rahman.BallContainer = []ballContainer{
		{
			TotalBall: 3,
			Verified:  true,
		},
	}

	fmt.Printf("Can %s play tennis now ? %t \n", rahman.Name, rahman.canPlay())

	/* Rahman will load ball randomly into the ball container.
	If:
	a) Rahman has many ball containers (indefinite).
	b) Rahman can put one ball into one container at the time.
	c) Only one container will be mark as verified, and Rahman will stop put/remove ball until it verified.
	d) Verified mark only when the ball container is full.
	e) You must create functional test, either the ball container is full or not.
	*/

	// Meaning i should create a function that can answer point a till point e

	// Point a, rahman has many ball containers, meaning i must add more ball container
	rahman.addEmptyBallContainer()
	rahman.addEmptyBallContainer()
	rahman.addEmptyBallContainer()

	fmt.Printf("Total Ball Container own by %s is %d \n", rahman.Name, len(rahman.BallContainer))

	// Point b,c and d
	for !isAllBallContainerFull(rahman.BallContainer) {
		// Add ball into empty ball container
		rahman.putBallIntoContainer()
		rahman.verifiedBallContainer()
	}

	// Second question is "Can you write the API based on those problem?"
	r := gin.Default()
	r.GET("/player/information", func(c *gin.Context) {
		c.JSON(200, rahman)
	})
	r.GET("/player/information/can/play", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": rahman.canPlay(),
		})
	})
	r.POST("/player/container/add", func(c *gin.Context) {
		rahman.addEmptyBallContainer()
		c.JSON(200, gin.H{
			"status": "Success",
		})
	})
	r.POST("/player/container/ball/add", func(c *gin.Context) {
		rahman.putBallIntoContainer()
		c.JSON(200, gin.H{
			"status": "Success",
		})
	})
	r.PUT("/container/mark/verified", func(c *gin.Context) {
		rahman.verifiedBallContainer()
		c.JSON(200, gin.H{
			"status": "Success",
		})
	})
	_ = r.Run()
}

// This function will tells is ball container is full or not
func isFull(ballContainer ballContainer) bool {
	// Simple is better you know!
	return ballContainer.TotalBall >= 3
}

// Point (e) This function will tells if the ball is verified or not
func isVerified(ballContainer ballContainer) bool {
	return ballContainer.Verified
}

func isAllBallContainerFull(container []ballContainer) bool {
	isAllFull := true
	for _, x := range container {
		if !isFull(x) {
			isAllFull = false
			break
		}
	}
	return isAllFull
}
