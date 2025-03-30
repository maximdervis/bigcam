package services

import (
	"context"
	"database/sql"
	"gen_server/src/db"
	api "gen_server/src/generated"

	"github.com/google/uuid"
)

type GymService struct {
	queries *db.Queries
}

func NewGymService(queries *db.Queries) *GymService {
	return &GymService{queries}
}

func (cc *GymService) GetGymInfo(ctx context.Context, params *api.GetGymByIdParams) (info *api.GymInfo, err error) {
	gym, err := cc.queries.SelectGymInfo(ctx, int64(params.GymId))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	info = &api.GymInfo{
		Name: api.Name(gym),
	}
	return info, nil
}

func (cc *GymService) CreateGym(ctx context.Context, info *api.GymInfo) (authInfo *api.GymAuthInfo, err error) {
	uuid := uuid.New()
	authKey := uuid.String()
	insert_params := db.InsertGymParams{
		Name:    string(info.Name),
		AuthKey: authKey,
	}
	err = cc.queries.InsertGym(ctx, insert_params)
	if err != nil {
		return nil, err
	}

	authInfo = &api.GymAuthInfo{
		AuthKey: api.AuthKey(authKey),
	}
	return authInfo, nil
}
