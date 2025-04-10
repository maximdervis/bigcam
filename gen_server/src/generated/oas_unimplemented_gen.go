// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateGym implements createGym operation.
//
// POST /api/gym
func (UnimplementedHandler) CreateGym(ctx context.Context, req *GymInfo) (r *GymAuthInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// FinishSession implements finishSession operation.
//
// DELETE /api/session/{sessionId}
func (UnimplementedHandler) FinishSession(ctx context.Context, params FinishSessionParams) (r *Ok, _ error) {
	return r, ht.ErrNotImplemented
}

// GetApiDocs implements getApiDocs operation.
//
// Get api documentation.
//
// GET /api
func (UnimplementedHandler) GetApiDocs(ctx context.Context) (r GetApiDocsOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGymById implements getGymById operation.
//
// GET /api/gym/{gymId}
func (UnimplementedHandler) GetGymById(ctx context.Context, params GetGymByIdParams) (r *GymInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUser implements getUser operation.
//
// GET /api/user
func (UnimplementedHandler) GetUser(ctx context.Context) (r *UserInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// ListCameras implements listCameras operation.
//
// GET /api/gym/camera/{gymId}
func (UnimplementedHandler) ListCameras(ctx context.Context, params ListCamerasParams) (r *CameraInfos, _ error) {
	return r, ht.ErrNotImplemented
}

// ListSessions implements listSessions operation.
//
// GET /api/session
func (UnimplementedHandler) ListSessions(ctx context.Context) (r *SessionsList, _ error) {
	return r, ht.ErrNotImplemented
}

// LocalGymAssign implements localGymAssign operation.
//
// POST /api/local/gym/assign
func (UnimplementedHandler) LocalGymAssign(ctx context.Context, req *GymAuthInfo) (r *Ok, _ error) {
	return r, ht.ErrNotImplemented
}

// RefreshAuthTokens implements refreshAuthTokens operation.
//
// POST /api/auth/refresh
func (UnimplementedHandler) RefreshAuthTokens(ctx context.Context, req *AuthTokens) (r *AuthTokens, _ error) {
	return r, ht.ErrNotImplemented
}

// SignIn implements signIn operation.
//
// Sign in using email and password.
//
// POST /api/auth/sign-in
func (UnimplementedHandler) SignIn(ctx context.Context, req *SignInInfo) (r *AuthTokens, _ error) {
	return r, ht.ErrNotImplemented
}

// SignUp implements signUp operation.
//
// POST /api/auth/sign-up
func (UnimplementedHandler) SignUp(ctx context.Context, req *SignUpInfo) (r *Ok, _ error) {
	return r, ht.ErrNotImplemented
}

// StartCameraAction implements startCameraAction operation.
//
// POST /api/gym/camera/ptz/{gymId}/{cameraId}
func (UnimplementedHandler) StartCameraAction(ctx context.Context, req *CameraAction, params StartCameraActionParams) (r *Ok, _ error) {
	return r, ht.ErrNotImplemented
}

// StartSession implements startSession operation.
//
// POST /api/session
func (UnimplementedHandler) StartSession(ctx context.Context, req *SessionToStart) (r *StartedSession, _ error) {
	return r, ht.ErrNotImplemented
}

// StopCameraAction implements stopCameraAction operation.
//
// DELETE /api/gym/camera/ptz/{gymId}/{cameraId}
func (UnimplementedHandler) StopCameraAction(ctx context.Context, params StopCameraActionParams) (r *Ok, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateUser implements updateUser operation.
//
// PUT /api/user
func (UnimplementedHandler) UpdateUser(ctx context.Context, req *UserToUpdate) (r *Ok, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
