package services

import (
	"context"
	"gen_server/src/db"
	api "gen_server/src/generated"
)

type SessionService struct {
	queries *db.Queries
}

func NewSessionService(queries *db.Queries) *SessionService {
	return &SessionService{queries}
}

func (cc *SessionService) GetSessions(ctx context.Context, userId int64) (list *api.SessionsList, err error) {
	openedSessions, err := cc.queries.SelectOpenedSessions(ctx, userId)
	if err != nil {
		return nil, err
	}
	list = &api.SessionsList{}
	for _, item := range openedSessions {
		list.Sessions = append(list.Sessions, api.Session{
			SessionID: api.ID(item.ID),
			CameraID:  api.ID(item.CameraID),
			GymID:     api.ID(item.GymID),
		})
	}
	return list, err
}

func (cc *SessionService) StartSession(ctx context.Context, userId int64, sessionToStart *api.SessionToStart) (startedSession *api.StartedSession, err error) {
	sessionId, err := cc.queries.InsertSession(ctx, db.InsertSessionParams{
		UserID:   userId,
		GymID:    int64(sessionToStart.GymID),
		CameraID: int64(sessionToStart.CameraID),
	})
	if err != nil {
		return nil, err
	}
	return &api.StartedSession{SessionID: api.ID(sessionId)}, nil
}

func (cc *SessionService) FinishSession(ctx context.Context, userId int64, finishSessionParams *api.FinishSessionParams) (err error) {
	// TODO: Проверять что корректный пользователь
	return cc.queries.CloseSession(ctx, int64(finishSessionParams.SessionId))
}
