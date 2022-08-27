package trickies

import "strings"

func GetStringFromChannel(stream chan int) string {
	var builder strings.Builder

	for done := false; !done; {
		select {
		default: // default is always the last resort
			builder.WriteString("1")
			done = true
		case <-stream:
			// despite of the fact the channel is buffered
			// the sender is not the one to get stuck here
			// so the sender has priority.
			builder.WriteString("2")
			stream = nil
		case stream <- 1: // the sender has priority, this will be the first.
			builder.WriteString("3")
		}
	}

	return builder.String()
}
