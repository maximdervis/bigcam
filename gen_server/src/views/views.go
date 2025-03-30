package views

import (
	"context"
	"errors"
	"os"

	"gen_server/src/generated"
	"gen_server/src/services"
)

type Handlers struct {
	userService     *services.UserService
	gymService      *services.GymService
	localGymService *services.LocalGymService
	sessionService  *services.SessionService
	cameraService   *services.CameraService
}

func NewHandlers(
	userService *services.UserService,
	gymService *services.GymService,
	localGymService *services.LocalGymService,
	sessionService *services.SessionService,
	cameraService *services.CameraService,
) *Handlers {
	return &Handlers{userService, gymService, localGymService, sessionService, cameraService}
}

func (cc *Handlers) GetApiDocs(ctx context.Context) (r api.GetApiDocsOK, err error) {
	file, err := os.OpenFile("api.html", os.O_RDONLY, 0)
	if err != nil {
		return r, err
	}
	r.Data = file
	return r, nil
}

func (cc *Handlers) CreateGym(ctx context.Context, req *api.GymInfo) (*api.GymAuthInfo, error) {
	return cc.gymService.CreateGym(ctx, req)
}

func (cc *Handlers) FinishSession(ctx context.Context, params api.FinishSessionParams) (*api.Ok, error) {
	userId, err := cc.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	return &api.Ok{}, cc.sessionService.FinishSession(ctx, userId, &params)
}

func (cc *Handlers) UpdateUser(ctx context.Context, req *api.UserToUpdate) (*api.Ok, error) {
	userId, err := cc.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	return &api.Ok{}, cc.userService.UpdateUserInfo(ctx, userId, req)
}

func (cc *Handlers) GetUser(ctx context.Context) (*api.UserInfo, error) {
	userId, err := cc.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	return cc.userService.GetUserInfo(ctx, userId)
}

func (cc *Handlers) GetGymById(ctx context.Context, params api.GetGymByIdParams) (*api.GymInfo, error) {
	return cc.gymService.GetGymInfo(ctx, &params)
}

func (cc *Handlers) ListCameras(ctx context.Context, params api.ListCamerasParams) (*api.CameraInfos, error) {
	return cc.cameraService.GetCameras(ctx, &params)
}

func (cc *Handlers) ListSessions(ctx context.Context) (*api.SessionsList, error) {
	userId, err := cc.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	return cc.sessionService.GetSessions(ctx, userId)
}

func (cc *Handlers) LocalGymAssign(ctx context.Context, req *api.GymAuthInfo) (*api.Ok, error) {
	ipAddr, err := cc.getIncomingIdAddress(ctx)
	if err != nil {
		return nil, err
	}
	return &api.Ok{}, cc.localGymService.GymAssign(ctx, req, &ipAddr)
}

func (cc *Handlers) RefreshAuthTokens(ctx context.Context, req *api.AuthTokens) (*api.AuthTokens, error) {
	return cc.userService.RefreshAuthTokens(ctx, req)
}

func (cc *Handlers) SignIn(ctx context.Context, req *api.SignInInfo) (*api.AuthTokens, error) {
	return cc.userService.LogInUser(ctx, req)
}

func (cc *Handlers) SignUp(ctx context.Context, req *api.SignUpInfo) (*api.Ok, error) {
	return &api.Ok{}, cc.userService.RegisterNewUser(ctx, req)
}

func (cc *Handlers) StartCameraAction(ctx context.Context, req *api.CameraAction, params api.StartCameraActionParams) (*api.Ok, error) {
	return &api.Ok{}, cc.cameraService.StartCameraAction(ctx, &params, req)
}

func (cc *Handlers) StopCameraAction(ctx context.Context, params api.StopCameraActionParams) (*api.Ok, error) {
	return &api.Ok{}, cc.cameraService.StopCameraAction(ctx, &params)
}

func (cc *Handlers) StartSession(ctx context.Context, req *api.SessionToStart) (*api.StartedSession, error) {
	userId, err := cc.getUserId(ctx)
	if err != nil {
		return nil, err
	}
	return cc.sessionService.StartSession(ctx, userId, req)
}

func (cc *Handlers) getUserId(ctx context.Context) (int64, error) {
	userId := ctx.Value("userId")
	if userId == nil {
		return 0, errors.New("failed to obtain userId from context")
	}
	return userId.(int64), nil
}

func (cc *Handlers) getIncomingIdAddress(ctx context.Context) (string, error) {
	ipAddr := ctx.Value("ipAddr")
	if ipAddr == nil {
		return "", errors.New("failed to obtain id address from context")
	}
	return ipAddr.(string), nil
}

func (cc *Handlers) NewError(ctx context.Context, err error) (r *api.ErrorStatusCode) {
	r = new(api.ErrorStatusCode)
	return &api.ErrorStatusCode{
		StatusCode: 500,
		Response: api.Error{
			Code:    "INTERNAL_ERROR",
			Message: err.Error(),
		},
	}
}
