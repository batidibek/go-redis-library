package cache

import "time"

type (
	// Cache is the top-level cache interface
	Cache interface {

		// Contains check if a cached key exists
		Contains(key string) (int64, error)

		// Delete remove the cached key
		Delete(key string) error

		// Get retrieve the cached key value
		Get(key string) (string, error)

		// HGet returns the value associated with field in the hash stored at key
		HGet(key string, field string) (string, error)

		// HGetAll retrieve multiple cached keys value
		HGetAll(key string) (map[string]string, error)

		// Set Save cache a value by key
		Set(key string, value interface{}, lifeTime time.Duration) error

		// HSet sets field in the hash stored at key to value
		HSet(key string, values ...interface{}) error
	}
)
