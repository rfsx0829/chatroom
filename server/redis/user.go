package redis

// Get user
func (c Client) Get(name string) (map[string]string, error) {
	return c.cli.HGetAll("user_" + name).Result()
}

// Set user field
func (c Client) Set(name, field string, value interface{}) (bool, error) {
	return c.cli.HSet("user_"+name, field, value).Result()
}
