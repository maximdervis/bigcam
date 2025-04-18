// Code generated by ogen, DO NOT EDIT.

package client

// Ref: #/components/schemas/ActionParams
type ActionParams struct {
	Velocity ActionParamsVelocity `json:"velocity"`
	Deadline string               `json:"deadline"`
}

// GetVelocity returns the value of Velocity.
func (s *ActionParams) GetVelocity() ActionParamsVelocity {
	return s.Velocity
}

// GetDeadline returns the value of Deadline.
func (s *ActionParams) GetDeadline() string {
	return s.Deadline
}

// SetVelocity sets the value of Velocity.
func (s *ActionParams) SetVelocity(val ActionParamsVelocity) {
	s.Velocity = val
}

// SetDeadline sets the value of Deadline.
func (s *ActionParams) SetDeadline(val string) {
	s.Deadline = val
}

type ActionParamsVelocity struct {
	Pan  float32 `json:"pan"`
	Tilt float32 `json:"tilt"`
	Zoom float32 `json:"zoom"`
}

// GetPan returns the value of Pan.
func (s *ActionParamsVelocity) GetPan() float32 {
	return s.Pan
}

// GetTilt returns the value of Tilt.
func (s *ActionParamsVelocity) GetTilt() float32 {
	return s.Tilt
}

// GetZoom returns the value of Zoom.
func (s *ActionParamsVelocity) GetZoom() float32 {
	return s.Zoom
}

// SetPan sets the value of Pan.
func (s *ActionParamsVelocity) SetPan(val float32) {
	s.Pan = val
}

// SetTilt sets the value of Tilt.
func (s *ActionParamsVelocity) SetTilt(val float32) {
	s.Tilt = val
}

// SetZoom sets the value of Zoom.
func (s *ActionParamsVelocity) SetZoom(val float32) {
	s.Zoom = val
}

// Local server cameras info.
// Ref: #/components/schemas/Camera
type Camera struct {
	CameraID ID `json:"camera_id"`
	// Camera area description.
	Description string `json:"description"`
}

// GetCameraID returns the value of CameraID.
func (s *Camera) GetCameraID() ID {
	return s.CameraID
}

// GetDescription returns the value of Description.
func (s *Camera) GetDescription() string {
	return s.Description
}

// SetCameraID sets the value of CameraID.
func (s *Camera) SetCameraID(val ID) {
	s.CameraID = val
}

// SetDescription sets the value of Description.
func (s *Camera) SetDescription(val string) {
	s.Description = val
}

// Local server cameras list.
// Ref: #/components/schemas/CamerasList
type CamerasList struct {
	Cameras []Camera `json:"cameras"`
}

// GetCameras returns the value of Cameras.
func (s *CamerasList) GetCameras() []Camera {
	return s.Cameras
}

// SetCameras sets the value of Cameras.
func (s *CamerasList) SetCameras(val []Camera) {
	s.Cameras = val
}

type ID int64

type Ok struct{}
