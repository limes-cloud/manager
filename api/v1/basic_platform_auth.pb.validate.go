// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: basic_platform_auth.proto

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

// Validate checks the field values on LoginCaptchaReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginCaptchaReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginCaptchaReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginCaptchaReplyMultiError, or nil if none found.
func (m *LoginCaptchaReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginCaptchaReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	// no validation rules for Captcha

	// no validation rules for Expire

	if len(errors) > 0 {
		return LoginCaptchaReplyMultiError(errors)
	}

	return nil
}

// LoginCaptchaReplyMultiError is an error wrapping multiple validation errors
// returned by LoginCaptchaReply.ValidateAll() if the designated constraints
// aren't met.
type LoginCaptchaReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginCaptchaReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginCaptchaReplyMultiError) AllErrors() []error { return m }

// LoginCaptchaReplyValidationError is the validation error returned by
// LoginCaptchaReply.Validate if the designated constraints aren't met.
type LoginCaptchaReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginCaptchaReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginCaptchaReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginCaptchaReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginCaptchaReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginCaptchaReplyValidationError) ErrorName() string {
	return "LoginCaptchaReplyValidationError"
}

// Error satisfies the builtin error interface
func (e LoginCaptchaReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginCaptchaReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginCaptchaReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginCaptchaReplyValidationError{}

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUsername()) < 6 {
		err := LoginRequestValidationError{
			field:  "Username",
			reason: "value length must be at least 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 6 {
		err := LoginRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCaptchaId()) < 1 {
		err := LoginRequestValidationError{
			field:  "CaptchaId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCaptcha()) < 1 {
		err := LoginRequestValidationError{
			field:  "Captcha",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReplyMultiError, or
// nil if none found.
func (m *LoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return LoginReplyMultiError(errors)
	}

	return nil
}

// LoginReplyMultiError is an error wrapping multiple validation errors
// returned by LoginReply.ValidateAll() if the designated constraints aren't met.
type LoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReplyMultiError) AllErrors() []error { return m }

// LoginReplyValidationError is the validation error returned by
// LoginReply.Validate if the designated constraints aren't met.
type LoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReplyValidationError) ErrorName() string { return "LoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReplyValidationError{}

// Validate checks the field values on ParseTokenReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ParseTokenReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ParseTokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ParseTokenReplyMultiError, or nil if none found.
func (m *ParseTokenReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ParseTokenReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for RoleId

	if len(errors) > 0 {
		return ParseTokenReplyMultiError(errors)
	}

	return nil
}

// ParseTokenReplyMultiError is an error wrapping multiple validation errors
// returned by ParseTokenReply.ValidateAll() if the designated constraints
// aren't met.
type ParseTokenReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ParseTokenReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ParseTokenReplyMultiError) AllErrors() []error { return m }

// ParseTokenReplyValidationError is the validation error returned by
// ParseTokenReply.Validate if the designated constraints aren't met.
type ParseTokenReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParseTokenReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParseTokenReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParseTokenReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParseTokenReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParseTokenReplyValidationError) ErrorName() string { return "ParseTokenReplyValidationError" }

// Error satisfies the builtin error interface
func (e ParseTokenReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParseTokenReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParseTokenReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParseTokenReplyValidationError{}

// Validate checks the field values on RefreshTokenReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RefreshTokenReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RefreshTokenReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RefreshTokenReplyMultiError, or nil if none found.
func (m *RefreshTokenReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RefreshTokenReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return RefreshTokenReplyMultiError(errors)
	}

	return nil
}

// RefreshTokenReplyMultiError is an error wrapping multiple validation errors
// returned by RefreshTokenReply.ValidateAll() if the designated constraints
// aren't met.
type RefreshTokenReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefreshTokenReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefreshTokenReplyMultiError) AllErrors() []error { return m }

// RefreshTokenReplyValidationError is the validation error returned by
// RefreshTokenReply.Validate if the designated constraints aren't met.
type RefreshTokenReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefreshTokenReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefreshTokenReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefreshTokenReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefreshTokenReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefreshTokenReplyValidationError) ErrorName() string {
	return "RefreshTokenReplyValidationError"
}

// Error satisfies the builtin error interface
func (e RefreshTokenReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRefreshTokenReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefreshTokenReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefreshTokenReplyValidationError{}

// Validate checks the field values on AuthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthRequestMultiError, or
// nil if none found.
func (m *AuthRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Path

	// no validation rules for Method

	if len(errors) > 0 {
		return AuthRequestMultiError(errors)
	}

	return nil
}

// AuthRequestMultiError is an error wrapping multiple validation errors
// returned by AuthRequest.ValidateAll() if the designated constraints aren't met.
type AuthRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthRequestMultiError) AllErrors() []error { return m }

// AuthRequestValidationError is the validation error returned by
// AuthRequest.Validate if the designated constraints aren't met.
type AuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthRequestValidationError) ErrorName() string { return "AuthRequestValidationError" }

// Error satisfies the builtin error interface
func (e AuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthRequestValidationError{}
