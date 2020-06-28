package memory

import (
	"fmt"

	"github.com/jayjzheng/amazon/auth/internal/domain"
)

func (c *Client) CreateUser(u *domain.User) error {
	c.Lock()
	c.users[u.Login] = u
	c.Unlock()

	return nil
}

func (c *Client) FindUser(login string) (*domain.User, error) {
	c.RLock()
	u := c.users[login]
	c.RUnlock()

	return u, nil
}

func (c *Client) UpdateUser(u *domain.User) error {
	old, err := c.FindUser(u.Login)
	if err != nil {
		return fmt.Errorf("FindUser %s: %w", u.Login, err)
	}

	if old == nil {
		return domain.ErrUserNotFound
	}

	c.Lock()
	c.users[u.Login] = u
	c.Unlock()

	return nil
}
