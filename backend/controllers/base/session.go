package base

func (c *Controller) SetUserSession(userID int, userName, nickName string) error {
	if err := c.SetSession("user_id", userID); err != nil {
		return err
	}

	if err := c.SetSession("user_name", userName); err != nil {
		return err
	}

	if err := c.SetSession("nick_name", nickName); err != nil {
		return err
	}

	return nil
}

func (c *Controller) GetMyUserID() int {
	userID, _ := c.GetSession("user_id").(int)
	return userID
}

func (c *Controller) GetMyUserName() string {
	userName, _ := c.GetSession("user_name").(string)
	return userName
}

func (c *Controller) GetMyNickName() string {
	nickName, _ := c.GetSession("nick_name").(string)
	return nickName
}

func (c *Controller) IsLogined() bool {
	return c.GetMyUserName() == ""
}
