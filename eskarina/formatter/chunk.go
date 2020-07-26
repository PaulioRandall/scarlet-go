package formatter

type chunk struct {
	head *line
	next *chunk
}
