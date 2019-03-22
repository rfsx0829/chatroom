package mysql

func (c *Conn) AddEmail(uid int, email string) error {
	return c.updateEmail(uid, email)
}

func (c *Conn) updateEmail(uid int, email string) error {
	_, err := c.db.Exec(Update+"email=? where uid=?", email, uid)
	return err
}
