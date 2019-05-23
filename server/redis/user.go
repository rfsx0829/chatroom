package redis

// Get user
func (c Client) Get(name string) (map[string]string, error) {
	return c.cli.HGetAll(name).Result()
}
