package main

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// MockRedis is a mock implementation of the Redis struct.
type MockRedis struct {
	mock.Mock
}

func (m *MockRedis) Connect() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRedis) Get(key string) ([]byte, time.Duration, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Get(1).(time.Duration), args.Error(2)
}

func (m *MockRedis) Set(key string, value interface{}, ttl time.Duration) error {
	args := m.Called(key, value, ttl)
	return args.Error(0)
}

func (m *MockRedis) Increment(key string) error {
	args := m.Called(key)
	return args.Error(0)
}

func (m *MockRedis) Close() error {
	args := m.Called()
	return args.Error(0)
}
