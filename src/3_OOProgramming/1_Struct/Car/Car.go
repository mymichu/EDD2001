package Car

import "errors"

type position struct {
	x uint16
	y uint16
}

type Car struct {
	position    position
	positionLog []position
	color       string
}

func (r *Car) SetPosition(x uint16, y uint16) {
	newPosition := position{x: x, y: y}
	r.position = newPosition
	r.positionLog = append(r.positionLog, newPosition)
}

func (r *Car) SetColor(color string) {
	r.color = color
}

func (r Car) GetColor() string {
	return r.color
}

func (r Car) GetPositionLogEntry(lastLog int) (uint16, uint16, error) {
	var err error = nil
	position := position{x: 0, y: 0}
	if lastLog < 0 {
		position, err = r.readLogEntry(lastLog)
	} else {
		err = errors.New("lastlog has to be negative")
	}
	return position.x, position.y, err
}

func (r Car) readLogEntry(lastLog int) (position, error) {
	position := position{x: 0, y: 0}
	var err error = nil
	lastLog = lastLog * -1
	if lastLog < len(r.positionLog) {
		position = r.positionLog[lastLog]
	} else {
		err = errors.New("Log is not as long as requested")
	}
	return position, err
}
