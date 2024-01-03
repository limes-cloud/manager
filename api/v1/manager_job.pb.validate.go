// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: manager_job.proto

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

// Validate checks the field values on Job with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Job) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Job with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in JobMultiError, or nil if none found.
func (m *Job) ValidateAll() error {
	return m.validate(true)
}

func (m *Job) validate(all bool) error {
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
		return JobMultiError(errors)
	}

	return nil
}

// JobMultiError is an error wrapping multiple validation errors returned by
// Job.ValidateAll() if the designated constraints aren't met.
type JobMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JobMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JobMultiError) AllErrors() []error { return m }

// JobValidationError is the validation error returned by Job.Validate if the
// designated constraints aren't met.
type JobValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JobValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JobValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JobValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JobValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JobValidationError) ErrorName() string { return "JobValidationError" }

// Error satisfies the builtin error interface
func (e JobValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJob.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JobValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JobValidationError{}

// Validate checks the field values on GetJobRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetJobRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetJobRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetJobRequestMultiError, or
// nil if none found.
func (m *GetJobRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetJobRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetJobRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetJobRequestMultiError(errors)
	}

	return nil
}

// GetJobRequestMultiError is an error wrapping multiple validation errors
// returned by GetJobRequest.ValidateAll() if the designated constraints
// aren't met.
type GetJobRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetJobRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetJobRequestMultiError) AllErrors() []error { return m }

// GetJobRequestValidationError is the validation error returned by
// GetJobRequest.Validate if the designated constraints aren't met.
type GetJobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetJobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetJobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetJobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetJobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetJobRequestValidationError) ErrorName() string { return "GetJobRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetJobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetJobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetJobRequestValidationError{}

// Validate checks the field values on GetJobReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetJobReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetJobReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetJobReplyMultiError, or
// nil if none found.
func (m *GetJobReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetJobReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetJob()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetJobReplyValidationError{
					field:  "Job",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetJobReplyValidationError{
					field:  "Job",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetJob()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetJobReplyValidationError{
				field:  "Job",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetJobReplyMultiError(errors)
	}

	return nil
}

// GetJobReplyMultiError is an error wrapping multiple validation errors
// returned by GetJobReply.ValidateAll() if the designated constraints aren't met.
type GetJobReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetJobReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetJobReplyMultiError) AllErrors() []error { return m }

// GetJobReplyValidationError is the validation error returned by
// GetJobReply.Validate if the designated constraints aren't met.
type GetJobReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetJobReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetJobReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetJobReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetJobReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetJobReplyValidationError) ErrorName() string { return "GetJobReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetJobReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetJobReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetJobReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetJobReplyValidationError{}

// Validate checks the field values on GetUserJobRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUserJobRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserJobRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserJobRequestMultiError, or nil if none found.
func (m *GetUserJobRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserJobRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := GetUserJobRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetUserJobRequestMultiError(errors)
	}

	return nil
}

// GetUserJobRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserJobRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserJobRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserJobRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserJobRequestMultiError) AllErrors() []error { return m }

// GetUserJobRequestValidationError is the validation error returned by
// GetUserJobRequest.Validate if the designated constraints aren't met.
type GetUserJobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserJobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserJobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserJobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserJobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserJobRequestValidationError) ErrorName() string {
	return "GetUserJobRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserJobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserJobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserJobRequestValidationError{}

// Validate checks the field values on GetUserJobReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUserJobReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserJobReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserJobReplyMultiError, or nil if none found.
func (m *GetUserJobReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserJobReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetJobs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetUserJobReplyValidationError{
						field:  fmt.Sprintf("Jobs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetUserJobReplyValidationError{
						field:  fmt.Sprintf("Jobs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetUserJobReplyValidationError{
					field:  fmt.Sprintf("Jobs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetUserJobReplyMultiError(errors)
	}

	return nil
}

// GetUserJobReplyMultiError is an error wrapping multiple validation errors
// returned by GetUserJobReply.ValidateAll() if the designated constraints
// aren't met.
type GetUserJobReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserJobReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserJobReplyMultiError) AllErrors() []error { return m }

// GetUserJobReplyValidationError is the validation error returned by
// GetUserJobReply.Validate if the designated constraints aren't met.
type GetUserJobReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserJobReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserJobReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserJobReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserJobReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserJobReplyValidationError) ErrorName() string { return "GetUserJobReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUserJobReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserJobReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserJobReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserJobReplyValidationError{}

// Validate checks the field values on PageJobRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageJobRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageJobRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageJobRequestMultiError,
// or nil if none found.
func (m *PageJobRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageJobRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPage() <= 0 {
		err := PageJobRequestValidationError{
			field:  "Page",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPageSize(); val <= 0 || val > 50 {
		err := PageJobRequestValidationError{
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

	if len(errors) > 0 {
		return PageJobRequestMultiError(errors)
	}

	return nil
}

// PageJobRequestMultiError is an error wrapping multiple validation errors
// returned by PageJobRequest.ValidateAll() if the designated constraints
// aren't met.
type PageJobRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageJobRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageJobRequestMultiError) AllErrors() []error { return m }

// PageJobRequestValidationError is the validation error returned by
// PageJobRequest.Validate if the designated constraints aren't met.
type PageJobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageJobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageJobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageJobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageJobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageJobRequestValidationError) ErrorName() string { return "PageJobRequestValidationError" }

// Error satisfies the builtin error interface
func (e PageJobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageJobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageJobRequestValidationError{}

// Validate checks the field values on PageJobReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageJobReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageJobReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageJobReplyMultiError, or
// nil if none found.
func (m *PageJobReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PageJobReply) validate(all bool) error {
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
					errors = append(errors, PageJobReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PageJobReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PageJobReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PageJobReplyMultiError(errors)
	}

	return nil
}

// PageJobReplyMultiError is an error wrapping multiple validation errors
// returned by PageJobReply.ValidateAll() if the designated constraints aren't met.
type PageJobReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageJobReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageJobReplyMultiError) AllErrors() []error { return m }

// PageJobReplyValidationError is the validation error returned by
// PageJobReply.Validate if the designated constraints aren't met.
type PageJobReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageJobReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageJobReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageJobReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageJobReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageJobReplyValidationError) ErrorName() string { return "PageJobReplyValidationError" }

// Error satisfies the builtin error interface
func (e PageJobReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageJobReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageJobReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageJobReplyValidationError{}

// Validate checks the field values on AddJobRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddJobRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddJobRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddJobRequestMultiError, or
// nil if none found.
func (m *AddJobRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddJobRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKeyword()) < 1 {
		err := AddJobRequestValidationError{
			field:  "Keyword",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := AddJobRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := AddJobRequestValidationError{
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
		return AddJobRequestMultiError(errors)
	}

	return nil
}

// AddJobRequestMultiError is an error wrapping multiple validation errors
// returned by AddJobRequest.ValidateAll() if the designated constraints
// aren't met.
type AddJobRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddJobRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddJobRequestMultiError) AllErrors() []error { return m }

// AddJobRequestValidationError is the validation error returned by
// AddJobRequest.Validate if the designated constraints aren't met.
type AddJobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddJobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddJobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddJobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddJobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddJobRequestValidationError) ErrorName() string { return "AddJobRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddJobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddJobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddJobRequestValidationError{}

// Validate checks the field values on UpdateJobRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateJobRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateJobRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateJobRequestMultiError, or nil if none found.
func (m *UpdateJobRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateJobRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateJobRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := UpdateJobRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 1 {
		err := UpdateJobRequestValidationError{
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
			err := UpdateJobRequestValidationError{
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
		return UpdateJobRequestMultiError(errors)
	}

	return nil
}

// UpdateJobRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateJobRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateJobRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateJobRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateJobRequestMultiError) AllErrors() []error { return m }

// UpdateJobRequestValidationError is the validation error returned by
// UpdateJobRequest.Validate if the designated constraints aren't met.
type UpdateJobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateJobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateJobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateJobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateJobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateJobRequestValidationError) ErrorName() string { return "UpdateJobRequestValidationError" }

// Error satisfies the builtin error interface
func (e UpdateJobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateJobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateJobRequestValidationError{}

// Validate checks the field values on DeleteJobRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteJobRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteJobRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteJobRequestMultiError, or nil if none found.
func (m *DeleteJobRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteJobRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteJobRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteJobRequestMultiError(errors)
	}

	return nil
}

// DeleteJobRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteJobRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteJobRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteJobRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteJobRequestMultiError) AllErrors() []error { return m }

// DeleteJobRequestValidationError is the validation error returned by
// DeleteJobRequest.Validate if the designated constraints aren't met.
type DeleteJobRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteJobRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteJobRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteJobRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteJobRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteJobRequestValidationError) ErrorName() string { return "DeleteJobRequestValidationError" }

// Error satisfies the builtin error interface
func (e DeleteJobRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteJobRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteJobRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteJobRequestValidationError{}