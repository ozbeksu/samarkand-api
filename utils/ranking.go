package utils

import (
	"math"
	"time"
)

var epoch = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
var offset = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

func HotScore(upVotes, downVotes int, date time.Time) float64 {
	votes := float64(upVotes - downVotes)
	order := math.Log10(math.Max(math.Abs(votes), 1))
	sign := 0.0

	if votes > 0 {
		sign = 1
	} else if votes < 0 {
		sign = -1
	}

	seconds := epochSeconds(date) - epochSeconds(offset)
	return sign*order + seconds/45000
}

const confidence float64 = 1.96

func BestScore(upVotes, totalVotes int) float64 {
	if totalVotes == 0 {
		return 0
	}

	z := confidence
	phat := float64(upVotes) / float64(totalVotes)
	score := (phat + z*z/(2*float64(totalVotes)) - z*math.Sqrt((phat*(1-phat)+z*z/(4*float64(totalVotes)))/float64(totalVotes))) / (1 + z*z/float64(totalVotes))

	return score
}

func epochSeconds(date time.Time) float64 {
	return date.Sub(epoch).Seconds()
}
