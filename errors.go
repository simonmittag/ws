package ws

// RejectOption represents an option used to control the way connection is
// rejected.
type RejectOption func(*RejectConnectionErrorType)

// RejectionReason returns an option that makes connection to be rejected with
// given reason.
func RejectionReason(reason string) RejectOption {
	return func(err *RejectConnectionErrorType) {
		err.reason = reason
	}
}

// RejectionStatus returns an option that makes connection to be rejected with
// given HTTP status code.
func RejectionStatus(code int) RejectOption {
	return func(err *RejectConnectionErrorType) {
		err.code = code
	}
}

// RejectionHeader returns an option that makes connection to be rejected with
// given HTTP headers.
func RejectionHeader(h HandshakeHeader) RejectOption {
	return func(err *RejectConnectionErrorType) {
		err.header = h
	}
}

// RejectConnectionError constructs an error that could be used to control the way
// handshake is rejected by Upgrader.
func RejectConnectionError(options ...RejectOption) error {
	err := new(RejectConnectionErrorType)
	for _, opt := range options {
		opt(err)
	}
	return err
}

// RejectConnectionErrorType represents a rejection of upgrade error.
//
// It can be returned by Upgrader's On* hooks to control the way WebSocket
// handshake is rejected.
type RejectConnectionErrorType struct {
	reason string
	code   int
	header HandshakeHeader
}

// Error implements error interface.
func (r *RejectConnectionErrorType) Error() string {
	return r.reason
}
