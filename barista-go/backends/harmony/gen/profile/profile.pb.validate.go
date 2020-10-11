// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: profile/v1/profile.proto

package profilev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _profile_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	return nil
}

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserName

	// no validation rules for UserAvatar

	// no validation rules for UserStatus

	return nil
}

// GetUserResponseValidationError is the validation error returned by
// GetUserResponse.Validate if the designated constraints aren't met.
type GetUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserResponseValidationError) ErrorName() string { return "GetUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserResponseValidationError{}

// Validate checks the field values on GetUserMetadataRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetUserMetadataRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AppId

	return nil
}

// GetUserMetadataRequestValidationError is the validation error returned by
// GetUserMetadataRequest.Validate if the designated constraints aren't met.
type GetUserMetadataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserMetadataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserMetadataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserMetadataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserMetadataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserMetadataRequestValidationError) ErrorName() string {
	return "GetUserMetadataRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserMetadataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserMetadataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserMetadataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserMetadataRequestValidationError{}

// Validate checks the field values on GetUserMetadataResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetUserMetadataResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Metadata

	return nil
}

// GetUserMetadataResponseValidationError is the validation error returned by
// GetUserMetadataResponse.Validate if the designated constraints aren't met.
type GetUserMetadataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserMetadataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserMetadataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserMetadataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserMetadataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserMetadataResponseValidationError) ErrorName() string {
	return "GetUserMetadataResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserMetadataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserMetadataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserMetadataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserMetadataResponseValidationError{}

// Validate checks the field values on UsernameUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UsernameUpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserName

	return nil
}

// UsernameUpdateRequestValidationError is the validation error returned by
// UsernameUpdateRequest.Validate if the designated constraints aren't met.
type UsernameUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UsernameUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UsernameUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UsernameUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UsernameUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UsernameUpdateRequestValidationError) ErrorName() string {
	return "UsernameUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UsernameUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUsernameUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UsernameUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UsernameUpdateRequestValidationError{}

// Validate checks the field values on StatusUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StatusUpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for NewStatus

	return nil
}

// StatusUpdateRequestValidationError is the validation error returned by
// StatusUpdateRequest.Validate if the designated constraints aren't met.
type StatusUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatusUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusUpdateRequestValidationError) ErrorName() string {
	return "StatusUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StatusUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatusUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatusUpdateRequestValidationError{}