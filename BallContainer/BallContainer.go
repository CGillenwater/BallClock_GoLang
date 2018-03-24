///Parts that would keep track of all things to deal with the collection of balls.
///Queue, location, etc.

package ballContainers

//Importing a ring container, from the GoLang library, in order to handle the circular
//List that will contain the collection of balls. 
import(
	"container/ring"
	"github.com/CGillenwater/BallClock_GoLang/Ball"
)

type BallContainer struct {
	capacity uint8
	numBalls uint8
} 

func NewBallContainer(capacity uint8, numBalls uint8) BallContainer {
	return BallContainer{capacity, numBalls}
}

func (bc BallContainer) isFull() bool {
	return bc.capacity == bc.numBalls
}

//Creating a ring pointer for our Ring container
type Queue struct {
	BallContainer
	ring *ring.Ring
}

//Filling the BallContainer
func NewQueue(capacity uint8) Queue {
	bc := NewBallContainer(capacity, capacity)
	r := ring.New(int(capacity))
	for i := uint8(0); i < capacity; i++ {
		r.Value = ball.New(i)
		r = r.Next()
	}

	return Queue{bc, r}
}

//Select the starting point for the Queue
func (queueP *Queue) Pop() ball.Ball {
	queueP.numBalls--
	ball := queueP.ring.Value.(ball.Ball)
	queueP.ring = queueP.ring.Next()
	return ball
}

//Check to see if the Container has made a full rotation. 
//(All balls are in their original position)
func (queueP *Queue) IsStartingPosition() bool {
	if !queueP.isFull() {
		return false;
	}

	tmp := queueP.ring
	for i := uint8(0); i < queueP.capacity; i++ {
		ball := queueP.ring.Value.(ball.Ball)
		queueP.ring = queueP.ring.Next()
		if ball.Id != i {
			queueP.ring = tmp
			return false
		}
	}

	return true
}

// //Test Queue
// // -1 = empty
// func (queueP *Queue) GetTestRepr() []int {
// 	repr := make([]int, queueP.capacity)
// 	for i := uint8(0); i < queueP.capacity; i++ {
// 		ball := queueP.ring.Value.(ball.Ball)
// 		queueP.ring = queueP.ring.Next()
// 		if i >= queueP.numBalls {
// 			repr[i] = -1 //Empty
// 		} else {
// 			repr[i] = int(ball.Id)
// 		}
// 	}
// 	return repr
// }

//Placing an Array of "ball" at the end of the queue
func (queueP *Queue) Push(balls []ball.Ball) {
	tmp := queueP.ring
	queueP.ring = queueP.ring.Move(int(queueP.numBalls))
	for i:= range balls {
		queueP.numBalls++
		queueP.ring.Value = balls[i]
		queueP.ring = queueP.ring.Next()
	}
	queueP.ring = tmp
}


//Creation of the Time Rails themselves
//This is where the balls are stored, and can be dropped to other rails
//In order to tell the time.
type Rail struct {
	BallContainer
	Balls []ball.Ball
}

//Instantiating a rail
func NewRail(capacity uint8) Rail {
	bc := NewBallContainer(capacity, 0)
	balls := make([]ball.Ball, capacity)
	return Rail{bc, balls}
}

//Emptying the rail
func (railP *Rail) spill() []ball.Ball {
	ballOverflow := make([]ball.Ball, railP.capacity)
	for i:= range railP.Balls {
		ballOverflow[railP.capacity-1-uint8(i)] = railP.Balls[i]
	}

	return ballOverflow
}

//Adding a single ball to the rail
//And all subsequent events that occur
func (railP *Rail) Push(ballInstance ball.Ball) []ball.Ball {
	if railP.isFull() {
		//Reset the rail state, and spill to the next rail
		railP.numBalls = 0
		return railP.spill()
	}

	railP.Balls[railP.numBalls] = ballInstance
	railP.numBalls++
	return []ball.Ball{}
}

// //Test for rail
// // -1 = empty
// func (railP *Rail) GetTestRepr() []int {
// 	repr := make([]int, railP.capacity)
// 	for i := uint8(0); i < railP.capacity; i++ {
// 		if i >= railP.numBalls {
// 			repr[i] = -1 //Empty
// 		} else {
// 			repr[i] = int(railP.Balls[i].Id)
// 		}
// 	}
// 	return repr
// }