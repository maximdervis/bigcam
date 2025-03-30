package services

import (
	"context"
	"database/sql"
	"errors"
	"gen_server/src/db"
	api "gen_server/src/generated"
	"gen_server/src/utils"
	"strconv"
	"time"
)

type UserService struct {
	db      *sql.DB
	queries *db.Queries
}

func NewUserService(db *sql.DB, queries *db.Queries) *UserService {
	return &UserService{db, queries}
}

func (cc *UserService) GetUserInfo(ctx context.Context, userId int64) (info *api.UserInfo, err error) {
	userInfo, err := cc.queries.SelectUserInfo(ctx, userId)
	if err != nil {
		return nil, err
	}
	info = &api.UserInfo{
		Name:  api.Name(userInfo.Name),
		Email: api.Email(userInfo.Email),
	}
	if userInfo.AvatarID.Valid {
		info.AvatarID = api.OptAvatarId{Value: api.AvatarId(userInfo.AvatarID.String), Set: true}
	}
	if userInfo.Dob.Valid {
		info.Dob = api.OptDob{Value: api.Dob(userInfo.Dob.Time), Set: true}
	}
	return info, err
}

func (cc *UserService) UpdateUserInfo(ctx context.Context, userId int64, info *api.UserToUpdate) (err error) {
	tx, err := cc.db.Begin()
	txQueries := cc.queries.WithTx(tx)
	if err != nil {
		return err
	}
	if info.Email.IsSet() {
		err = txQueries.UpdateUserEmail(ctx, db.UpdateUserEmailParams{
			ID:    userId,
			Email: string(info.Email.Value),
		})
	}

	if info.Name.IsSet() {
		err = txQueries.UpdateUserName(ctx, db.UpdateUserNameParams{
			ID:   userId,
			Name: string(info.Name.Value),
		})
	}

	if info.Dob.IsSet() {
		err = txQueries.UpdateUserDob(ctx, db.UpdateUserDobParams{
			ID: userId,
			Dob: sql.NullTime{
				Time:  time.Time(info.Dob.Value),
				Valid: true,
			},
		})
	}

	if info.AvatarID.IsSet() {
		err = txQueries.UpdateUserAvatarId(ctx, db.UpdateUserAvatarIdParams{
			ID: userId,
			AvatarID: sql.NullString{
				String: string(info.AvatarID.Value),
				Valid:  true,
			},
		})
	}
	return err
}

func (cc *UserService) RegisterNewUser(ctx context.Context, info *api.SignUpInfo) (err error) {
	userAlreadyExists, err := cc.queries.ContainsUserWithEmail(ctx, string(info.Email))
	if err != nil {
		return err
	}

	if userAlreadyExists {
		return errors.New("user with specified email already exists")
	}

	user := db.InsertUserInfoParams{
		Name:     string(info.Name),
		Email:    string(info.Email),
		Password: string(info.Password),
	}
	user.Password, err = utils.GenerateHashPassword(user.Password)
	if err != nil {
		return err
	}

	err = cc.queries.InsertUserInfo(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (cc *UserService) LogInUser(ctx context.Context, info *api.SignInInfo) (tokens *api.AuthTokens, err error) {
	userInfo, err := cc.queries.SelectUserInfoByEmail(ctx, string(info.Email))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	passwordsMatched := utils.CompareHashPassword(string(info.Password), userInfo.Password)
	if !passwordsMatched {
		return nil, errors.New("incorrect password")
	}

	authData, err := utils.GetSignedTokens(strconv.FormatInt(userInfo.ID, 10))
	if err != nil {
		return nil, err
	}

	tokens = &api.AuthTokens{
		AccessToken:  authData.AccessKey,
		RefreshToken: authData.RefreshKey,
	}
	return tokens, nil
}

func (cc *UserService) RefreshAuthTokens(ctx context.Context, currentTokens *api.AuthTokens) (newTokens *api.AuthTokens, err error) {
	claims, err := utils.ParseRefreshToken(string(currentTokens.RefreshToken))
	if err != nil {
		return nil, err
	}
	accessToken, err := utils.GetAccessSignedToken(claims.Subject)
	if err != nil {
		return nil, err
	}
	newTokens = &api.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: currentTokens.RefreshToken,
	}
	return newTokens, nil
}
