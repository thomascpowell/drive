package tests

type MockRedis struct {
	SetFunc   func(key, value string) error
	GetFunc   func(key string) (string, error)
	SetexFunc func(key, value string, ttl int) error
	TTLFunc   func(key string) (string, error)
}

func (m *MockRedis) Set(key, value string) error {
	return m.SetFunc(key, value)
}

func (m *MockRedis) Get(key string) (string, error) {
	return m.GetFunc(key)
}

func (m *MockRedis) Setex(key, value string, ttl int) error {
	return m.SetexFunc(key, value, ttl)
}

func (m *MockRedis) TTL(key string) (string, error) {
	return m.TTLFunc(key)
}
