// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: manager_dictionary.proto

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

// Validate checks the field values on Dictionary with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Dictionary) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Dictionary with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DictionaryMultiError, or
// nil if none found.
func (m *Dictionary) ValidateAll() error {
	return m.validate(true)
}

func (m *Dictionary) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Keyword

	// no validation rules for Name

	// no validation rules for Description

	if m.CreatedAt != nil {
		// no validation rules for CreatedAt
	}

	if m.UpdatedAt != nil {
		// no validation rules for UpdatedAt
	}

	if len(errors) > 0 {
		return DictionaryMultiError(errors)
	}

	return nil
}

// DictionaryMultiError is an error wrapping multiple validation errors
// returned by Dictionary.ValidateAll() if the designated constraints aren't met.
type DictionaryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DictionaryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DictionaryMultiError) AllErrors() []error { return m }

// DictionaryValidationError is the validation error returned by
// Dictionary.Validate if the designated constraints aren't met.
type DictionaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DictionaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DictionaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DictionaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DictionaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DictionaryValidationError) ErrorName() string { return "DictionaryValidationError" }

// Error satisfies the builtin error interface
func (e DictionaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDictionary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DictionaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DictionaryValidationError{}

// Validate checks the field values on PageDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PageDictionaryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageDictionaryRequestMultiError, or nil if none found.
func (m *PageDictionaryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageDictionaryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PageDictionaryRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val > 50 {
		err := PageDictionaryRequestValidationError{
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
		return PageDictionaryRequestMultiError(errors)
	}

	return nil
}

// PageDictionaryRequestMultiError is an error wrapping multiple validation
// errors returned by PageDictionaryRequest.ValidateAll() if the designated
// constraints aren't met.
type PageDictionaryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageDictionaryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageDictionaryRequestMultiError) AllErrors() []error { return m }

// PageDictionaryRequestValidationError is the validation error returned by
// PageDictionaryRequest.Validate if the designated constraints aren't met.
type PageDictionaryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageDictionaryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageDictionaryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageDictionaryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageDictionaryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageDictionaryRequestValidationError) ErrorName() string {
	return "PageDictionaryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PageDictionaryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageDictionaryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageDictionaryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageDictionaryRequestValidationError{}

// Validate checks the field values on PageDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PageDictionaryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PageDictionaryReplyMultiError, or nil if none found.
func (m *PageDictionaryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageDictionaryReply) validate(all bool) error {
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
					errors = append(errors, PageDictionaryReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PageDictionaryReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PageDictionaryReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PageDictionaryReplyMultiError(errors)
	}

	return nil
}

// PageDictionaryReplyMultiError is an error wrapping multiple validation
// errors returned by PageDictionaryReply.ValidateAll() if the designated
// constraints aren't met.
type PageDictionaryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageDictionaryReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageDictionaryReplyMultiError) AllErrors() []error { return m }

// PageDictionaryReplyValidationError is the validation error returned by
// PageDictionaryReply.Validate if the designated constraints aren't met.
type PageDictionaryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageDictionaryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageDictionaryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageDictionaryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageDictionaryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageDictionaryReplyValidationError) ErrorName() string {
	return "PageDictionaryReplyValidationError"
}

// Error satisfies the builtin error interface
func (e PageDictionaryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageDictionaryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageDictionaryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageDictionaryReplyValidationError{}

// Validate checks the field values on AddDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddDictionaryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddDictionaryRequestMultiError, or nil if none found.
func (m *AddDictionaryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddDictionaryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKeyword()) < 1 {
		err := AddDictionaryRequestValidationError{
			field:  "Keyword",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := AddDictionaryRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := AddDictionaryRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddDictionaryRequestMultiError(errors)
	}

	return nil
}

// AddDictionaryRequestMultiError is an error wrapping multiple validation
// errors returned by AddDictionaryRequest.ValidateAll() if the designated
// constraints aren't met.
type AddDictionaryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddDictionaryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddDictionaryRequestMultiError) AllErrors() []error { return m }

// AddDictionaryRequestValidationError is the validation error returned by
// AddDictionaryRequest.Validate if the designated constraints aren't met.
type AddDictionaryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddDictionaryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddDictionaryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddDictionaryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddDictionaryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddDictionaryRequestValidationError) ErrorName() string {
	return "AddDictionaryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddDictionaryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddDictionaryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddDictionaryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddDictionaryRequestValidationError{}

// Validate checks the field values on AddDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddDictionaryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddDictionaryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddDictionaryReplyMultiError, or nil if none found.
func (m *AddDictionaryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AddDictionaryReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return AddDictionaryReplyMultiError(errors)
	}

	return nil
}

// AddDictionaryReplyMultiError is an error wrapping multiple validation errors
// returned by AddDictionaryReply.ValidateAll() if the designated constraints
// aren't met.
type AddDictionaryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddDictionaryReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddDictionaryReplyMultiError) AllErrors() []error { return m }

// AddDictionaryReplyValidationError is the validation error returned by
// AddDictionaryReply.Validate if the designated constraints aren't met.
type AddDictionaryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddDictionaryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddDictionaryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddDictionaryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddDictionaryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddDictionaryReplyValidationError) ErrorName() string {
	return "AddDictionaryReplyValidationError"
}

// Error satisfies the builtin error interface
func (e AddDictionaryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddDictionaryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddDictionaryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddDictionaryReplyValidationError{}

// Validate checks the field values on UpdateDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateDictionaryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateDictionaryRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateDictionaryRequestMultiError, or nil if none found.
func (m *UpdateDictionaryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateDictionaryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateDictionaryRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetKeyword()) < 1 {
		err := UpdateDictionaryRequestValidationError{
			field:  "Keyword",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := UpdateDictionaryRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := UpdateDictionaryRequestValidationError{
			field:  "Description",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateDictionaryRequestMultiError(errors)
	}

	return nil
}

// UpdateDictionaryRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateDictionaryRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateDictionaryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateDictionaryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateDictionaryRequestMultiError) AllErrors() []error { return m }

// UpdateDictionaryRequestValidationError is the validation error returned by
// UpdateDictionaryRequest.Validate if the designated constraints aren't met.
type UpdateDictionaryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDictionaryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDictionaryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDictionaryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDictionaryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDictionaryRequestValidationError) ErrorName() string {
	return "UpdateDictionaryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDictionaryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDictionaryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDictionaryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDictionaryRequestValidationError{}

// Validate checks the field values on DeleteDictionaryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteDictionaryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteDictionaryRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteDictionaryRequestMultiError, or nil if none found.
func (m *DeleteDictionaryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteDictionaryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteDictionaryRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteDictionaryRequestMultiError(errors)
	}

	return nil
}

// DeleteDictionaryRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteDictionaryRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteDictionaryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteDictionaryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteDictionaryRequestMultiError) AllErrors() []error { return m }

// DeleteDictionaryRequestValidationError is the validation error returned by
// DeleteDictionaryRequest.Validate if the designated constraints aren't met.
type DeleteDictionaryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteDictionaryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteDictionaryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteDictionaryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteDictionaryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteDictionaryRequestValidationError) ErrorName() string {
	return "DeleteDictionaryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteDictionaryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteDictionaryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteDictionaryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteDictionaryRequestValidationError{}