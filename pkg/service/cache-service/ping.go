package cache_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func (s *cacheService) Ping(sessionId string) error {

	ctx := context.Background()
	userName, err := s.redis.HGet(ctx, "session", sessionId).Result()
	if err != nil {
		return err
	}

	ctx = context.Background()
	key := fmt.Sprintf("last-ping-%s", sessionId)
	now := time.Now()

	selfBlockTime, err := s.redis.Get(ctx, key).Int64()
	if err == nil && selfBlockTime > now.Unix() {
		return fmt.Errorf("not allowed to ping twice in 60s")
	} else if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	ctx = context.Background()
	now = time.Now()
	globalBlockDuration := time.Second * 5
	globalBlockTime := now.Add(globalBlockDuration)
	globalOk, err := s.redis.SetNX(ctx, "global-block", time.Duration(globalBlockTime.Unix()+1), globalBlockDuration).Result()
	if err != nil {
		return err
	}
	if !globalOk {
		currentBlockTime, err := s.redis.Get(ctx, "global-block").Int64()
		if err != nil {
			return err
		}
		if currentBlockTime < now.Unix() {
			err = s.redis.Set(ctx, "global-block", globalBlockTime, globalBlockDuration).Err()
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unable to set lock")
		}
	}

	selfBlockDuration := time.Second * 60
	newSelfBlockTime := time.Now().Add(selfBlockDuration).Unix()
	err = s.redis.Set(ctx, key, newSelfBlockTime, selfBlockDuration).Err()
	if err != nil {
		return err
	}

	ctx = context.Background()
	sessionPingCount, err := s.redis.Incr(ctx, PingSessionKey(sessionId)).Result()
	if err != nil {
		return fmt.Errorf("unable to increment ping count")
	}

	ctx = context.Background()
	err = s.redis.ZAdd(ctx, PING_SCOREBOARD, &redis.Z{
		Score:  float64(sessionPingCount),
		Member: userName,
	}).Err()
	if err != nil {
		return fmt.Errorf("unable to update scoreboard")
	}

	ctx = context.Background()
	err = s.redis.PFAdd(ctx, PING_TOTAL, userName).Err()
	if err != nil {
		return fmt.Errorf("unable to update total")
	}

	return nil
}
