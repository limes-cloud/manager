// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: manager_role_menu.proto

package v1

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

// Validate checks the field values on RoleMenu with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RoleMenu) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RoleMenu with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RoleMenuMultiError, or nil
// if none found.
func (m *RoleMenu) ValidateAll() error {
	return m.validate(true)
}

func (m *RoleMenu) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MenuId

	// no validation rules for RoleId

	if all {
		switch v := interface{}(m.GetMenu()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RoleMenuValidationError{
					field:  "Menu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RoleMenuValidationError{
					field:  "Menu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMenu()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RoleMenuValidationError{
				field:  "Menu",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RoleMenuMultiError(errors)
	}

	return nil
}

// RoleMenuMultiError is an error wrapping multiple validation errors returned
// by RoleMenu.ValidateAll() if the designated constraints aren't met.
type RoleMenuMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RoleMenuMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RoleMenuMultiError) AllErrors() []error { return m }

// RoleMenuValidationError is the validation error returned by
// RoleMenu.Validate if the designated constraints aren't met.
type RoleMenuValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoleMenuValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoleMenuValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoleMenuValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoleMenuValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoleMenuValidationError) ErrorName() string { return "RoleMenuValidationError" }

// Error satisfies the builtin error interface
func (e RoleMenuValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoleMenu.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoleMenuValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoleMenuValidationError{}

// Validate checks the field values on CurrentRoleMenuTreeReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CurrentRoleMenuTreeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CurrentRoleMenuTreeReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CurrentRoleMenuTreeReplyMultiError, or nil if none found.
func (m *CurrentRoleMenuTreeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CurrentRoleMenuTreeReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CurrentRoleMenuTreeReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CurrentRoleMenuTreeReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CurrentRoleMenuTreeReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CurrentRoleMenuTreeReplyMultiError(errors)
	}

	return nil
}

// CurrentRoleMenuTreeReplyMultiError is an error wrapping multiple validation
// errors returned by CurrentRoleMenuTreeReply.ValidateAll() if the designated
// constraints aren't met.
type CurrentRoleMenuTreeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CurrentRoleMenuTreeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CurrentRoleMenuTreeReplyMultiError) AllErrors() []error { return m }

// CurrentRoleMenuTreeReplyValidationError is the validation error returned by
// CurrentRoleMenuTreeReply.Validate if the designated constraints aren't met.
type CurrentRoleMenuTreeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrentRoleMenuTreeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrentRoleMenuTreeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrentRoleMenuTreeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrentRoleMenuTreeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrentRoleMenuTreeReplyValidationError) ErrorName() string {
	return "CurrentRoleMenuTreeReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CurrentRoleMenuTreeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrentRoleMenuTreeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrentRoleMenuTreeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrentRoleMenuTreeReplyValidationError{}

// Validate checks the field values on GetRoleMenuIdsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetRoleMenuIdsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRoleMenuIdsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetRoleMenuIdsRequestMultiError, or nil if none found.
func (m *GetRoleMenuIdsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRoleMenuIdsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetRoleId() <= 0 {
		err := GetRoleMenuIdsRequestValidationError{
			field:  "RoleId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetRoleMenuIdsRequestMultiError(errors)
	}

	return nil
}

// GetRoleMenuIdsRequestMultiError is an error wrapping multiple validation
// errors returned by GetRoleMenuIdsRequest.ValidateAll() if the designated
// constraints aren't met.
type GetRoleMenuIdsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRoleMenuIdsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRoleMenuIdsRequestMultiError) AllErrors() []error { return m }

// GetRoleMenuIdsRequestValidationError is the validation error returned by
// GetRoleMenuIdsRequest.Validate if the designated constraints aren't met.
type GetRoleMenuIdsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleMenuIdsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleMenuIdsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleMenuIdsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleMenuIdsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleMenuIdsRequestValidationError) ErrorName() string {
	return "GetRoleMenuIdsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetRoleMenuIdsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleMenuIdsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleMenuIdsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleMenuIdsRequestValidationError{}

// Validate checks the field values on GetRoleMenuIdsReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetRoleMenuIdsReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetRoleMenuIdsReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetRoleMenuIdsReplyMultiError, or nil if none found.
func (m *GetRoleMenuIdsReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetRoleMenuIdsReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetRoleMenuIdsReplyMultiError(errors)
	}

	return nil
}

// GetRoleMenuIdsReplyMultiError is an error wrapping multiple validation
// errors returned by GetRoleMenuIdsReply.ValidateAll() if the designated
// constraints aren't met.
type GetRoleMenuIdsReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetRoleMenuIdsReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetRoleMenuIdsReplyMultiError) AllErrors() []error { return m }

// GetRoleMenuIdsReplyValidationError is the validation error returned by
// GetRoleMenuIdsReply.Validate if the designated constraints aren't met.
type GetRoleMenuIdsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetRoleMenuIdsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetRoleMenuIdsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetRoleMenuIdsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetRoleMenuIdsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetRoleMenuIdsReplyValidationError) ErrorName() string {
	return "GetRoleMenuIdsReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetRoleMenuIdsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetRoleMenuIdsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetRoleMenuIdsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetRoleMenuIdsReplyValidationError{}

// Validate checks the field values on UpdateRoleMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateRoleMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateRoleMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateRoleMenuRequestMultiError, or nil if none found.
func (m *UpdateRoleMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateRoleMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetRoleId() <= 0 {
		err := UpdateRoleMenuRequestValidationError{
			field:  "RoleId",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetMenuIds()) < 1 {
		err := UpdateRoleMenuRequestValidationError{
			field:  "MenuIds",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateRoleMenuRequestMultiError(errors)
	}

	return nil
}

// UpdateRoleMenuRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateRoleMenuRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateRoleMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateRoleMenuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateRoleMenuRequestMultiError) AllErrors() []error { return m }

// UpdateRoleMenuRequestValidationError is the validation error returned by
// UpdateRoleMenuRequest.Validate if the designated constraints aren't met.
type UpdateRoleMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRoleMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRoleMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRoleMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRoleMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRoleMenuRequestValidationError) ErrorName() string {
	return "UpdateRoleMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRoleMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRoleMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRoleMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRoleMenuRequestValidationError{}
