// Code generated by ogen, DO NOT EDIT.

package client

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// GetCameras implements getCameras operation.
	//
	// GET /cameras
	GetCameras(ctx context.Context) (*CamerasList, error)
	// StartCameraAction implements startCameraAction operation.
	//
	// POST /cameras/{cameraId}/ptz
	StartCameraAction(ctx context.Context, req *ActionParams, params StartCameraActionParams) error
	// StopCameraAction implements stopCameraAction operation.
	//
	// DELETE /cameras/{cameraId}/ptz
	StopCameraAction(ctx context.Context, params StopCameraActionParams) error
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
