package channel

import (
	"database/sql"

	"github.com/faiz-gh/enshitradar-api/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetChannels() ([]types.Channel, error) {
	rows, err := s.db.Query("SELECT * FROM channels")
	if err != nil {
		return nil, err
	}

	channels := make([]types.Channel, 0)
	for rows.Next() {
		p, err := scanRowsIntoChannel(rows)
		if err != nil {
			return nil, err
		}

		channels = append(channels, *p)
	}

	return channels, nil
}

func (s *Store) GetChannelByID(channelID string) (*types.Channel, error) {
	rows, err := s.db.Query("SELECT * FROM channels WHERE channel_id = $1", channelID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	channel, err := scanRowsIntoChannel(rows)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func (s *Store) GetChannelByName(name string) (*types.Channel, error) {
	rows, err := s.db.Query("SELECT * FROM channels WHERE name = $1", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	channel, err := scanRowsIntoChannel(rows)
	if err != nil {
		return nil, err
	}

	return channel, nil
}

func (s *Store) AddChannel(channel types.AddChannelPayload) (*types.Channel, error) {
	var newChannel types.Channel
	err := s.db.QueryRow(
		"INSERT INTO channels (name, channel_id, description, level) VALUES ($1, $2, $3, $4) RETURNING id, name, description, level, date_added",
		channel.Name, channel.ChannelID, channel.Description, channel.Level,
	).Scan(&newChannel.ID, &newChannel.Name, &newChannel.ChannelID, &newChannel.Description, &newChannel.Level, &newChannel.DateAdded)
	if err != nil {
		return nil, err
	}
	return &newChannel, nil
}

func scanRowsIntoChannel(rows *sql.Rows) (*types.Channel, error) {
	channel := new(types.Channel)

	err := rows.Scan(
		&channel.ID,
		&channel.Name,
		&channel.ChannelID,
		&channel.Description,
		&channel.Level,
		&channel.DateAdded,
	)
	if err != nil {
		return nil, err
	}

	return channel, nil
}
