///Serves as the struct for a single ball within the clock.
///Ball will keep track of its own original position.

package ball

type Ball struct {
	//Starting position for the ball, within the Ball Holder.
	Id uint8
}

func New(id uint8) Ball {
    return Ball{id}
}
