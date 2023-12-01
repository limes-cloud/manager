// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: manager_menu.proto

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

// Validate checks the field values on Menu with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Menu) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Menu with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MenuMultiError, or nil if none found.
func (m *Menu) ValidateAll() error {
	return m.validate(true)
}

func (m *Menu) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for ParentId

	// no validation rules for Title

	// no validation rules for Type

	// no validation rules for Icon

	// no validation rules for Path

	// no validation rules for Keyword

	// no validation rules for Permission

	// no validation rules for Component

	// no validation rules for Api

	// no validation rules for Method

	// no validation rules for App

	for idx, item := range m.GetChildren() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, MenuValidationError{
						field:  fmt.Sprintf("Children[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MenuValidationError{
					field:  fmt.Sprintf("Children[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if m.Redirect != nil {
		// no validation rules for Redirect
	}

	if m.Weight != nil {
		// no validation rules for Weight
	}

	if m.IsHidden != nil {
		// no validation rules for IsHidden
	}

	if m.IsCache != nil {
		// no validation rules for IsCache
	}

	if m.IsHome != nil {
		// no validation rules for IsHome
	}

	if m.IsAffix != nil {
		// no validation rules for IsAffix
	}

	if len(errors) > 0 {
		return MenuMultiError(errors)
	}

	return nil
}

// MenuMultiError is an error wrapping multiple validation errors returned by
// Menu.ValidateAll() if the designated constraints aren't met.
type MenuMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MenuMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MenuMultiError) AllErrors() []error { return m }

// MenuValidationError is the validation error returned by Menu.Validate if the
// designated constraints aren't met.
type MenuValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MenuValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MenuValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MenuValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MenuValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MenuValidationError) ErrorName() string { return "MenuValidationError" }

// Error satisfies the builtin error interface
func (e MenuValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMenu.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MenuValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MenuValidationError{}

// Validate checks the field values on GetMenuTreeReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetMenuTreeReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetMenuTreeReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetMenuTreeReplyMultiError, or nil if none found.
func (m *GetMenuTreeReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetMenuTreeReply) validate(all bool) error {
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
					errors = append(errors, GetMenuTreeReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetMenuTreeReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetMenuTreeReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetMenuTreeReplyMultiError(errors)
	}

	return nil
}

// GetMenuTreeReplyMultiError is an error wrapping multiple validation errors
// returned by GetMenuTreeReply.ValidateAll() if the designated constraints
// aren't met.
type GetMenuTreeReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetMenuTreeReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetMenuTreeReplyMultiError) AllErrors() []error { return m }

// GetMenuTreeReplyValidationError is the validation error returned by
// GetMenuTreeReply.Validate if the designated constraints aren't met.
type GetMenuTreeReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetMenuTreeReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetMenuTreeReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetMenuTreeReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetMenuTreeReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetMenuTreeReplyValidationError) ErrorName() string { return "GetMenuTreeReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetMenuTreeReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetMenuTreeReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetMenuTreeReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetMenuTreeReplyValidationError{}

// Validate checks the field values on AddMenuRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddMenuRequestMultiError,
// or nil if none found.
func (m *AddMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ParentId

	if utf8.RuneCountInString(m.GetApp()) < 1 {
		err := AddMenuRequestValidationError{
			field:  "App",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := AddMenuRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _AddMenuRequest_Type_InLookup[m.GetType()]; !ok {
		err := AddMenuRequestValidationError{
			field:  "Type",
			reason: "value must be in list [R A M G BA]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Icon != nil {
		// no validation rules for Icon
	}

	if m.Path != nil {
		// no validation rules for Path
	}

	if m.Keyword != nil {
		// no validation rules for Keyword
	}

	if m.Permission != nil {
		// no validation rules for Permission
	}

	if m.Component != nil {
		// no validation rules for Component
	}

	if m.Api != nil {
		// no validation rules for Api
	}

	if m.Method != nil {
		// no validation rules for Method
	}

	if m.Redirect != nil {
		// no validation rules for Redirect
	}

	if m.Weight != nil {
		// no validation rules for Weight
	}

	if m.IsHidden != nil {
		// no validation rules for IsHidden
	}

	if m.IsCache != nil {
		// no validation rules for IsCache
	}

	if m.IsHome != nil {
		// no validation rules for IsHome
	}

	if m.IsAffix != nil {
		// no validation rules for IsAffix
	}

	if len(errors) > 0 {
		return AddMenuRequestMultiError(errors)
	}

	return nil
}

// AddMenuRequestMultiError is an error wrapping multiple validation errors
// returned by AddMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type AddMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddMenuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddMenuRequestMultiError) AllErrors() []error { return m }

// AddMenuRequestValidationError is the validation error returned by
// AddMenuRequest.Validate if the designated constraints aren't met.
type AddMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddMenuRequestValidationError) ErrorName() string { return "AddMenuRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddMenuRequestValidationError{}

var _AddMenuRequest_Type_InLookup = map[string]struct{}{
	"R":  {},
	"A":  {},
	"M":  {},
	"G":  {},
	"BA": {},
}

// Validate checks the field values on UpdateMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateMenuRequestMultiError, or nil if none found.
func (m *UpdateMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateMenuRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for ParentId

	if utf8.RuneCountInString(m.GetTitle()) < 1 {
		err := UpdateMenuRequestValidationError{
			field:  "Title",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _UpdateMenuRequest_Type_InLookup[m.GetType()]; !ok {
		err := UpdateMenuRequestValidationError{
			field:  "Type",
			reason: "value must be in list [R A M G BA]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Icon != nil {
		// no validation rules for Icon
	}

	if m.Path != nil {
		// no validation rules for Path
	}

	if m.Keyword != nil {
		// no validation rules for Keyword
	}

	if m.Permission != nil {
		// no validation rules for Permission
	}

	if m.Component != nil {
		// no validation rules for Component
	}

	if m.Api != nil {
		// no validation rules for Api
	}

	if m.Method != nil {
		// no validation rules for Method
	}

	if m.Redirect != nil {
		// no validation rules for Redirect
	}

	if m.Weight != nil {
		// no validation rules for Weight
	}

	if m.IsHidden != nil {
		// no validation rules for IsHidden
	}

	if m.IsCache != nil {
		// no validation rules for IsCache
	}

	if m.IsHome != nil {
		// no validation rules for IsHome
	}

	if m.IsAffix != nil {
		// no validation rules for IsAffix
	}

	if len(errors) > 0 {
		return UpdateMenuRequestMultiError(errors)
	}

	return nil
}

// UpdateMenuRequestMultiError is an error wrapping multiple validation errors
// returned by UpdateMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type UpdateMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateMenuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateMenuRequestMultiError) AllErrors() []error { return m }

// UpdateMenuRequestValidationError is the validation error returned by
// UpdateMenuRequest.Validate if the designated constraints aren't met.
type UpdateMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateMenuRequestValidationError) ErrorName() string {
	return "UpdateMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateMenuRequestValidationError{}

var _UpdateMenuRequest_Type_InLookup = map[string]struct{}{
	"R":  {},
	"A":  {},
	"M":  {},
	"G":  {},
	"BA": {},
}

// Validate checks the field values on DeleteMenuRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteMenuRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteMenuRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteMenuRequestMultiError, or nil if none found.
func (m *DeleteMenuRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteMenuRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := DeleteMenuRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteMenuRequestMultiError(errors)
	}

	return nil
}

// DeleteMenuRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteMenuRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteMenuRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteMenuRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteMenuRequestMultiError) AllErrors() []error { return m }

// DeleteMenuRequestValidationError is the validation error returned by
// DeleteMenuRequest.Validate if the designated constraints aren't met.
type DeleteMenuRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteMenuRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteMenuRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteMenuRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteMenuRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteMenuRequestValidationError) ErrorName() string {
	return "DeleteMenuRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteMenuRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteMenuRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteMenuRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteMenuRequestValidationError{}