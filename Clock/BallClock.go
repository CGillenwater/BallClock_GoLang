package clock

import (
	"math"
	"github.com/CGillenwater/BallClock_GoLang/Ball"
	"github.com/CGillenwater/BallClock_GoLang/BallContainer"
)

//Max amount of balls per Rail
const HOUR_RAIL_CAP = 11
const FIVE_MIN_RAIL_CAP = 11
const ONE_MIN_RAIL_CAP = 4

var ballQueue ballContainers.Queue
var hourRail ballContainers.Rail
var fiveMinRail ballContainers.Rail
var oneMinRail ballContainers.Rail
var numClockRefreshes uint64

// //Function to determine if a rail is empty
// //Needs tested, Can call below with the following line: 
// //checkForBallOverflow(oneMinRail, clockBall);
// func checkForBallOverflow(specifiedRail ballContainer.Rail, currentBall ball.Ball) {
// 	var ballOverflow []ball.Ball
// 	ballOverflow = specifiedRail.Push(currentBall)
// 	if len(ballOverflow) == 0 {
// 		return
// 	}
// }


//Initialization for the new rails, in order to keep track of time.
func init() {
	oneMinRail = ballContainers.NewRail(ONE_MIN_RAIL_CAP)
	fiveMinRail = ballContainers.NewRail(FIVE_MIN_RAIL_CAP)
	hourRail = ballContainers.NewRail(HOUR_RAIL_CAP)
}

//Add a ball, to represent passing of time.
func updateClockState(clockBall ball.Ball) {
	var ballOverflow []ball.Ball

	//Determines if the One Minute Rail is full.
	//Then pushes the ballOverflow back to the queue.
	ballOverflow = oneMinRail.Push(clockBall)
	if len(ballOverflow) == 0 {
		return
	}
	
	ballQueue.Push(ballOverflow)

	//Determines if the Five Minute Rail is full.
	//Then pushes the ballOverflow back to the queue.
	ballOverflow = fiveMinRail.Push(clockBall)
	if len(ballOverflow) == 0 {
		return
	}

	ballQueue.Push(ballOverflow)

	//Determines if the Hour Rail is full.
	//Then pushes the ballOverflow back to the queue.
	ballOverflow = hourRail.Push(clockBall)
	if len(ballOverflow) == 0 {
		return
	}

	ballQueue.Push(append(ballOverflow, clockBall))
	// fmt.Print("Hour: ", ballOverflow)
	// fmt.Println()

}

//Determine a single cycle.
func calcCycle(queueCapacity uint8) {
	for {
		ball := ballQueue.Pop()
		updateClockState(ball)
		if ballQueue.IsFull() {
			numClockRefreshes++
			if ballQueue.IsStartingPosition() {
				break
			}
		}
	}
}

//Determines the number of days (total complete cycles / 2)
//The clock runs.
func CalcNumDaysInCycle(queueCapacity uint8) uint64 {
	numClockRefreshes = 0
	ballQueue = ballContainers.NewQueue(queueCapacity)
	calcCycle(queueCapacity)
	return uint64(math.Ceil(float64(numClockRefreshes) / 2.0))
}