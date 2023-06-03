package cache_service

import "context"

const PING_SCOREBOARD = "ping-scoreboard"
const PING_TOTAL = "ping-total"

func PingSessionKey(sessionId string) string {
	return "ping-" + sessionId
}

func (s *cacheService) Top10Ping() ([]string, error) {
	ctx := context.Background()
	r, err := s.redis.ZRevRange(ctx, PING_SCOREBOARD, 0, 9).Result()
	if err != nil {
		return []string{}, err
	}
	return r, nil
}

func (s *cacheService) Count() (int64, error) {
	ctx := context.Background()
	r, err := s.redis.PFCount(ctx, PING_TOTAL).Result()
	if err != nil {
		return 0, err
	}
	return r, nil
}

func (s *cacheService) CountBySessionId(sessionId string) (int64, error) {
	ctx := context.Background()
	r, err := s.redis.Get(ctx, PingSessionKey(sessionId)).Int64()
	if err != nil {
		return 0, err
	}
	return r, nil
}
