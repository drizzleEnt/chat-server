// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user.proto

package user_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUserCreate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "UserCreate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "UserCreate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserCreate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "UserCreate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on UserCreate with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserCreate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserCreate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserCreateMultiError, or
// nil if none found.
func (m *UserCreate) ValidateAll() error {
	return m.validate(true)
}

func (m *UserCreate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUserUpdate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UserCreateValidationError{
					field:  "UserUpdate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UserCreateValidationError{
					field:  "UserUpdate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserUpdate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserCreateValidationError{
				field:  "UserUpdate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Password

	if len(errors) > 0 {
		return UserCreateMultiError(errors)
	}

	return nil
}

// UserCreateMultiError is an error wrapping multiple validation errors
// returned by UserCreate.ValidateAll() if the designated constraints aren't met.
type UserCreateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserCreateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserCreateMultiError) AllErrors() []error { return m }

// UserCreateValidationError is the validation error returned by
// UserCreate.Validate if the designated constraints aren't met.
type UserCreateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserCreateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserCreateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserCreateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserCreateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserCreateValidationError) ErrorName() string { return "UserCreateValidationError" }

// Error satisfies the builtin error interface
func (e UserCreateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserCreate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserCreateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserCreateValidationError{}

// Validate checks the field values on UserUpdate with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UserUpdate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserUpdate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UserUpdateMultiError, or
// nil if none found.
func (m *UserUpdate) ValidateAll() error {
	return m.validate(true)
}

func (m *UserUpdate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 50 {
		err := UserUpdateValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 50 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Email

	// no validation rules for Role

	if len(errors) > 0 {
		return UserUpdateMultiError(errors)
	}

	return nil
}

// UserUpdateMultiError is an error wrapping multiple validation errors
// returned by UserUpdate.ValidateAll() if the designated constraints aren't met.
type UserUpdateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserUpdateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserUpdateMultiError) AllErrors() []error { return m }

// UserUpdateValidationError is the validation error returned by
// UserUpdate.Validate if the designated constraints aren't met.
type UserUpdateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserUpdateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserUpdateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserUpdateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserUpdateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserUpdateValidationError) ErrorName() string { return "UserUpdateValidationError" }

// Error satisfies the builtin error interface
func (e UserUpdateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserUpdate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserUpdateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserUpdateValidationError{}

// Validate checks the field values on CreateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateRequestMultiError, or
// nil if none found.
func (m *CreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateRequestValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateRequestMultiError(errors)
	}

	return nil
}

// CreateRequestMultiError is an error wrapping multiple validation errors
// returned by CreateRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateRequestMultiError) AllErrors() []error { return m }

// CreateRequestValidationError is the validation error returned by
// CreateRequest.Validate if the designated constraints aren't met.
type CreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequestValidationError) ErrorName() string { return "CreateRequestValidationError" }

// Error satisfies the builtin error interface
func (e CreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequestValidationError{}

// Validate checks the field values on CreateResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CreateResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CreateResponseMultiError,
// or nil if none found.
func (m *CreateResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateResponseMultiError(errors)
	}

	return nil
}

// CreateResponseMultiError is an error wrapping multiple validation errors
// returned by CreateResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateResponseMultiError) AllErrors() []error { return m }

// CreateResponseValidationError is the validation error returned by
// CreateResponse.Validate if the designated constraints aren't met.
type CreateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResponseValidationError) ErrorName() string { return "CreateResponseValidationError" }

// Error satisfies the builtin error interface
func (e CreateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResponseValidationError{}

// Validate checks the field values on GetRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetRequestMultiError, or
// nil if none found.
func (m *GetRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetRequestMultiError(errors)
	}

	return nil
}

// GetRequestMultiError is an error wrapping multiple validation errors
// returned by GetRequest.ValidateAll() if the designated constraints aren't met.
type GetRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRequestMultiError) AllErrors() []error { return m }

// GetRequestValidationError is the validation error returned by
// GetRequest.Validate if the designated constraints aren't met.
type GetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRequestValidationError) ErrorName() string { return "GetRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRequestValidationError{}

// Validate checks the field values on GetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetResponseMultiError, or
// nil if none found.
func (m *GetResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetResponseMultiError(errors)
	}

	return nil
}

// GetResponseMultiError is an error wrapping multiple validation errors
// returned by GetResponse.ValidateAll() if the designated constraints aren't met.
type GetResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetResponseMultiError) AllErrors() []error { return m }

// GetResponseValidationError is the validation error returned by
// GetResponse.Validate if the designated constraints aren't met.
type GetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetResponseValidationError) ErrorName() string { return "GetResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetResponseValidationError{}

// Validate checks the field values on UpdateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateRequestMultiError, or
// nil if none found.
func (m *UpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateRequestValidationError{
					field:  "Info",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRequestValidationError{
				field:  "Info",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateRequestMultiError(errors)
	}

	return nil
}

// UpdateRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRequestMultiError) AllErrors() []error { return m }

// UpdateRequestValidationError is the validation error returned by
// UpdateRequest.Validate if the designated constraints aren't met.
type UpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRequestValidationError) ErrorName() string { return "UpdateRequestValidationError" }

// Error satisfies the builtin error interface
func (e UpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRequestValidationError{}

// Validate checks the field values on DeleteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DeleteRequestMultiError, or
// nil if none found.
func (m *DeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteRequestMultiError(errors)
	}

	return nil
}

// DeleteRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteRequestMultiError) AllErrors() []error { return m }

// DeleteRequestValidationError is the validation error returned by
// DeleteRequest.Validate if the designated constraints aren't met.
type DeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteRequestValidationError) ErrorName() string { return "DeleteRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteRequestValidationError{}
