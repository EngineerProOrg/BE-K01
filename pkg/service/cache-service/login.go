package cache_service

import (
	"context"
	"strconv"
)

func (s *cacheService) Login(userName string) (string, error) {
	ctx := context.Background()
	sessionId, err := s.redis.Do(ctx, "incr", "sessionId").Int64()
	if err != nil {
		return "", err
	}

	ctx = context.Background()
	err = s.redis.HMSet(ctx, "session", sessionId, userName).Err()
	if err != nil {
		return "", err
	}

	sessionIdStr := strconv.Itoa(int(sessionId))

	return sessionIdStr, nil
}
