package services

import (
	"context"
	"gen_server/src/db"
	api "gen_server/src/generated"
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

type LocalGymService struct {
	queries *db.Queries
	redis   *redis.Client
}

func NewLocalGymService(queries *db.Queries, redis *redis.Client) *LocalGymService {
	return &LocalGymService{queries, redis}
}

func (cc *LocalGymService) GymAssign(ctx context.Context, params *api.GymAuthInfo, ipAddress *string) (err error) {
	gymId, err := cc.queries.SelectGymIdByAuthKey(ctx, string(params.AuthKey))
	if err != nil {
		return err
	}

	cc.redis.Set(strconv.FormatInt(gymId, 10), *ipAddress, 0)
	log.Printf("Set client_ip %s for gym_id %d", *ipAddress, gymId)
	return nil
}
