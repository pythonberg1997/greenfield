package errors

import (
	"cosmossdk.io/errors"
)

const RootCodespace = "greenfiled"

var (
	ErrInvalidBucketName     = errors.Register(RootCodespace, 1000, "Invalid bucket name")
	ErrInvalidObjectName     = errors.Register(RootCodespace, 1001, "Invalid object name")
	ErrInvalidGroupName      = errors.Register(RootCodespace, 1002, "Invalid group name")
	ErrInvalidChecksum       = errors.Register(RootCodespace, 1003, "Invalid checksum")
	ErrInvalidContentType    = errors.Register(RootCodespace, 1004, "Invalid content type")
	ErrInvalidSPSignature    = errors.Register(RootCodespace, 1005, "Invalid sp signature")
	ErrInvalidSPAddress      = errors.Register(RootCodespace, 1006, "Invalid sp address")
	ErrInvalidPrincipal      = errors.Register(RootCodespace, 1007, "Invalid principal")
	ErrInvalidGRN            = errors.Register(RootCodespace, 1008, "Not a standard greenfield resource name format")
	ErrInvalidParameter      = errors.Register(RootCodespace, 1009, "Invalid parameter")
	ErrInvalidVisibilityType = errors.Register(RootCodespace, 1010, "Invalid public type")

	ErrGRNTypeMismatch = errors.Register(RootCodespace, 2000, "Greenfield resource type mismatch")
)
