package pipe_filter

// Request is the input of ther filter
type Request interface{}

// Request is ther output of ther filter
type Response interface{}

// Filter inter is the definittion of the data processing componets
// Pipe-Fiter structure
type Filter interface {
	Process(data Request) (Response, error)
}
