package functions

import (
	"Appointy-Instagram/data"
	"fmt"
	"regexp"
)

func ValidatePost(post *data.InPost) error {
	if len(post.UserId) == 0 {
		return fmt.Errorf("userId is required")
	}
	if len(post.Caption) == 0 {
		return fmt.Errorf("caption is required")
	}
	if len(post.ImgUrl) == 0 {
		return fmt.Errorf("imgUrl is required")
	}
	return nil
}

func ValidateUser(user *data.InUser) error {
	if len(user.Id) == 0 {
		return fmt.Errorf("user id is required")
	}
	if len(user.Name) == 0 {
		return fmt.Errorf("user name is required")
	}
	if len(user.Email) == 0 {
		return fmt.Errorf("user email is required")
	} else {
		emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !emailRegex.MatchString(user.Email) {
			return fmt.Errorf("email is not valid")
		}
	}
	if len(user.Password) == 0 {
		return fmt.Errorf("password is required")
	}
	return nil
}
