package platform

type room struct {
	messages []message
	online   map[int]struct{}
	token    string
}
