package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"telecalendar/internal/config"
	"telecalendar/internal/database/models"
	"time"

	"github.com/mohae/deepcopy"
	"github.com/redis/go-redis/v9"
)

type UserState struct {
	State    StateType     `json:"state"`
	Event    *models.Event `json:"event,omitempty"`
	Calendar string        `json:"calendar_name,omitempty"`
}

type StateStorage struct {
	client *redis.Client
	ctx    context.Context
}

func (s *StateStorage) GetState(userID int64) (UserState, error) {
	key := fmt.Sprintf("user:%d:state", userID)
	data, err := s.client.Get(s.ctx, key).Result()
	if err == redis.Nil {
		return s.initState(userID)
	} else if err != nil {
		return UserState{}, err
	}

	var state UserState
	err = json.Unmarshal([]byte(data), &state)
	if err != nil {
		return UserState{}, err
	}

	return state, nil
}

func (s *StateStorage) initState(userId int64) (UserState, error) {
	return deepcopy.Copy(InitState).(UserState), s.SetState(userId, InitState)
}

func (s *StateStorage) SetState(userID int64, state UserState) error {
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
