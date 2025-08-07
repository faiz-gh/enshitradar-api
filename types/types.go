package types

import "time"

type UserStore interface {
	AddUser() (*User, error)
	GetUserByID(id string) (*User, error)
}

type ChannelStore interface {
	GetChannels() ([]Channel, error)
	AddChannel(channel AddChannelPayload) (*Channel, error)
	GetChannelByID(channelID string) (*Channel, error)
	GetChannelByName(name string) (*Channel, error)
}

type VoteStore interface {
	AddVote(vote Vote) error
	GetVotesByChannelID(channelID string) ([]Vote, error)
}

type User struct {
	ID string `json:"id"`
}

type Channel struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ChannelID   string    `json:"channel_id"`
	Level       string    `json:"level"`
	Description string    `json:"description"`
	DateAdded   time.Time `json:"date_added"`
}

type Vote struct {
	ID              string    `json:"id"`
	ChannelID       string    `json:"channel_id"`
	UserID          string    `json:"user_id"`
	Level           string    `json:"level"`
	PublicIPAddress string    `json:"public_ip_address"`
	DateAdded       time.Time `json:"date_added"`
}

type AddChannelPayload struct {
	Name        string `json:"name"`
	ChannelID   string `json:"channel_id"`
	Description string `json:"description"`
	Level       string `json:"level"`
}
