// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: manager_dict.proto

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

// Validate checks the field values on Dict with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Dict) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Dict with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DictMultiError, or nil if none found.
func (m *Dict) ValidateAll() error {
	return m.validate(true)
}

func (m *Dict) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Keyword

	// no validation rules for Name

	// no validation rules for Description

	if m.Weight != nil {
		// no validation rules for Weight
	}

	if m.CreatedAt != nil {
		// no validation rules for CreatedAt
	}

	if m.UpdatedAt != nil {
		// no validation rules for UpdatedAt
	}

	if len(errors) > 0 {
		return DictMultiError(errors)
	}

	return nil
}

// DictMultiError is an error wrapping multiple validation errors returned by
// Dict.ValidateAll() if the designated constraints aren't met.
type DictMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DictMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DictMultiError) AllErrors() []error { return m }

// DictValidationError is the validation error returned by Dict.Validate if the
// designated constraints aren't met.
type DictValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DictValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DictValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DictValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DictValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DictValidationError) ErrorName() string { return "DictValidationError" }

// Error satisfies the builtin error interface
func (e DictValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDict.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DictValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DictValidationError{}

// Validate checks the field values on PageDictRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PageDictRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageDictRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageDictRequestMultiError, or nil if none found.
func (m *PageDictRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageDictRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PageDictRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val > 50 {
		err := PageDictRequestValidationError{
			field:  "PageSize",
			reason: "value must be inside range (0, 50]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Keyword != nil {
		// no validation rules for Keyword
	}

	if m.Name != nil {
		// no validation rules for Name
	}

	if len(errors) > 0 {
		return PageDictRequestMultiError(errors)
	}

	return nil
}

// PageDictRequestMultiError is an error wrapping multiple validation errors
// returned by PageDictRequest.ValidateAll() if the designated constraints
// aren't met.
type PageDictRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageDictRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageDictRequestMultiError) AllErrors() []error { return m }

// PageDictRequestValidationError is the validation error returned by
// PageDictRequest.Validate if the designated constraints aren't met.
type PageDictRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageDictRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageDictRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageDictRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageDictRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageDictRequestValidationError) ErrorName() string { return "PageDictRequestValidationError" }

// Error satisfies the builtin error interface
func (e PageDictRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageDictRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageDictRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageDictRequestValidationError{}

// Validate checks the field values on PageDictReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageDictReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageDictReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageDictReplyMultiError, or
// nil if none found.
func (m *PageDictReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageDictReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PageDictReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PageDictReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PageDictReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PageDictReplyMultiError(errors)
	}

	return nil
}

// PageDictReplyMultiError is an error wrapping multiple validation errors
// returned by PageDictReply.ValidateAll() if the designated constraints
// aren't met.
type PageDictReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageDictReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageDictReplyMultiError) AllErrors() []error { return m }

// PageDictReplyValidationError is the validation error returned by
// PageDictReply.Validate if the designated constraints aren't met.
type PageDictReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageDictReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageDictReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageDictReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageDictReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageDictReplyValidationError) ErrorName() string { return "PageDictReplyValidationError" }

// Error satisfies the builtin error interface
func (e PageDictReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageDictReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageDictReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageDictReplyValidationError{}

// Validate checks the field values on AddDictRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddDictRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddDictRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddDictRequestMultiError,
// or nil if none found.
func (m *AddDictRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddDictRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKeyword()) < 1 {
		err := AddDictRequestValidationError{
			field:  "Keyword",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := AddDictRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := AddDictRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Weight != nil {
		// no validation rules for Weight
	}

	if len(errors) > 0 {
		return AddDictRequestMultiError(errors)
	}

	return nil
}

// AddDictRequestMultiError is an error wrapping multiple validation errors
// returned by AddDictRequest.ValidateAll() if the designated constraints
// aren't met.
type AddDictRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddDictRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddDictRequestMultiError) AllErrors() []error { return m }

// AddDictRequestValidationError is the validation error returned by
// AddDictRequest.Validate if the designated constraints aren't met.
type AddDictRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddDictRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddDictRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddDictRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddDictRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddDictRequestValidationError) ErrorName() string { return "AddDictRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddDictRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddDictRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddDictRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddDictRequestValidationError{}

// Validate checks the field values on UpdateDictRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateDictRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateDictRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateDictRequestMultiError, or nil if none found.
func (m *UpdateDictRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateDictRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateDictRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := UpdateDictRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := UpdateDictRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Weight != nil {

		if m.GetWeight() <= 0 {
			err := UpdateDictRequestValidationError{
				field:  "Weight",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return UpdateDictRequestMultiError(errors)
	}

	return nil
}

// UpdateDictRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateDictRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateDictRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateDictRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateDictRequestMultiError) AllErrors() []error { return m }

// UpdateDictRequestValidationError is the validation error returned by
// UpdateDictRequest.Validate if the designated constraints aren't met.
type UpdateDictRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDictRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDictRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDictRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDictRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDictRequestValidationError) ErrorName() string {
	return "UpdateDictRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDictRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDictRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDictRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDictRequestValidationError{}

// Validate checks the field values on DeleteDictRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteDictRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteDictRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteDictRequestMultiError, or nil if none found.
func (m *DeleteDictRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteDictRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteDictRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteDictRequestMultiError(errors)
	}

	return nil
}

// DeleteDictRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteDictRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteDictRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteDictRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteDictRequestMultiError) AllErrors() []error { return m }

// DeleteDictRequestValidationError is the validation error returned by
// DeleteDictRequest.Validate if the designated constraints aren't met.
type DeleteDictRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteDictRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteDictRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteDictRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteDictRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteDictRequestValidationError) ErrorName() string {
	return "DeleteDictRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteDictRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteDictRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteDictRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteDictRequestValidationError{}
