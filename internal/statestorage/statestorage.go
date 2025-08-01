package statestorage

import (
	"context"
	"encoding/json"
	"fmt"
	"telecalendar/internal/config"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/redis/go-redis/v9"
)

type UserState struct {
	State    string                 `json:"state"`
	TempData map[string]interface{} `json:"temp_data"`
}

type StateStorage struct {
	client *redis.Client
	ctx    context.Context
}

func Init(cfg *config.Config) *StateStorage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       0,
	})

	return &StateStorage{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (s *StateStorage) GetState(userID int64) (*UserState, error) {
	key := fmt.Sprintf("user:%d:state", userID)
	data, err := s.client.Get(s.ctx, key).Result()
	if err == redis.Nil {
		s.SetState(userID, initState)
		return deepcopy.Copy(initState).(*UserState), nil
	} else if err != nil {
		return nil, err
	}

	var state UserState
	err = json.Unmarshal([]byte(data), &state)
	if err != nil {
		return nil, err
	}

	return &state, nil
}

func (s *StateStorage) SetState(userID int64, state *UserState) error {
	key := fmt.Sprintf("user:%d:state", userID)
	data, err := json.Marshal(state)
	if err != nil {
		return err
	}

	return s.client.Set(s.ctx, key, data, 24*time.Hour).Err()
}

func (s *StateStorage) ClearState(userID int64) error {
	key := fmt.Sprintf("user:%d:state", userID)
	return s.client.Del(s.ctx, key).Err()
}
