package memory

import (
	"sync"

	"github.com/jayjzheng/amazon/auth/internal/domain"
)

type Client struct {
	users map[string]*domain.User

	sync.RWMutex
}

func NewClient() *Client {
	return &Client{
		users: make(map[string]*domain.User),
	}
}
