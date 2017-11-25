package coordinates

import (
	"fmt"
)

type coordinates struct {
	X, Y int
}

func (c coordinates) String() string {
	return fmt.Sprintf("x=%d, y=%d", c.X, c.Y)
}
