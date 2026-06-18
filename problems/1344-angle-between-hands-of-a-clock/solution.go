package anglebetweenhandsofclock

import "math"

func angleClock(hour int, minutes int) float64 {
	minuteHandPosition := float64(minutes) / 5
	hourHandPosition := float64(hour%12) + float64(minutes)/60

	positionDifference := math.Abs(minuteHandPosition - hourHandPosition)
	angle := positionDifference * 30

	return math.Min(angle, 360-angle)
}
