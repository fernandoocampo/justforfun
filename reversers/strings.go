package reversers

import (
	"context"
	"strings"
)

func ReverseString(ctx context.Context, input string) string {
	runeStream := reverseString(ctx, input)
	var builder strings.Builder
	for v := range runeStream {
		builder.WriteByte(v)
	}
	return builder.String()
}

func reverseString(ctx context.Context, input string) <-chan byte {
	stream := make(chan byte)
	go func() {
		defer close(stream)
		for i := len(input) - 1; i >= 0; i-- {
			select {
			case <-ctx.Done():
				return
			default:
				stream <- input[i]
			}
		}
	}()
	return stream
}

func ReverseNormalString(input string) string {
	var builder strings.Builder
	for i := len(input) - 1; i >= 0; i-- {
		builder.WriteByte(input[i])
	}
	return builder.String()
}
