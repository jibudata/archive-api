// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: library_manager.proto

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

// Validate checks the field values on CreateLibraryManagerParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateLibraryManagerParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateLibraryManagerParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateLibraryManagerParamsMultiError, or nil if none found.
func (m *CreateLibraryManagerParams) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateLibraryManagerParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSpec()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateLibraryManagerParamsValidationError{
					field:  "Spec",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateLibraryManagerParamsValidationError{
					field:  "Spec",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateLibraryManagerParamsValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateLibraryManagerParamsMultiError(errors)
	}

	return nil
}

// CreateLibraryManagerParamsMultiError is an error wrapping multiple
// validation errors returned by CreateLibraryManagerParams.ValidateAll() if
// the designated constraints aren't met.
type CreateLibraryManagerParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateLibraryManagerParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateLibraryManagerParamsMultiError) AllErrors() []error { return m }

// CreateLibraryManagerParamsValidationError is the validation error returned
// by CreateLibraryManagerParams.Validate if the designated constraints aren't met.
type CreateLibraryManagerParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateLibraryManagerParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateLibraryManagerParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateLibraryManagerParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateLibraryManagerParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateLibraryManagerParamsValidationError) ErrorName() string {
	return "CreateLibraryManagerParamsValidationError"
}

// Error satisfies the builtin error interface
func (e CreateLibraryManagerParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateLibraryManagerParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateLibraryManagerParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateLibraryManagerParamsValidationError{}

// Validate checks the field values on UpdateLibraryManagerParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateLibraryManagerParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateLibraryManagerParams with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateLibraryManagerParamsMultiError, or nil if none found.
func (m *UpdateLibraryManagerParams) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateLibraryManagerParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResourceKey()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateLibraryManagerParamsValidationError{
					field:  "ResourceKey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateLibraryManagerParamsValidationError{
					field:  "ResourceKey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResourceKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateLibraryManagerParamsValidationError{
				field:  "ResourceKey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSpec()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateLibraryManagerParamsValidationError{
					field:  "Spec",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateLibraryManagerParamsValidationError{
					field:  "Spec",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateLibraryManagerParamsValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateLibraryManagerParamsMultiError(errors)
	}

	return nil
}

// UpdateLibraryManagerParamsMultiError is an error wrapping multiple
// validation errors returned by UpdateLibraryManagerParams.ValidateAll() if
// the designated constraints aren't met.
type UpdateLibraryManagerParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateLibraryManagerParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateLibraryManagerParamsMultiError) AllErrors() []error { return m }

// UpdateLibraryManagerParamsValidationError is the validation error returned
// by UpdateLibraryManagerParams.Validate if the designated constraints aren't met.
type UpdateLibraryManagerParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateLibraryManagerParamsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateLibraryManagerParamsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateLibraryManagerParamsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateLibraryManagerParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateLibraryManagerParamsValidationError) ErrorName() string {
	return "UpdateLibraryManagerParamsValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateLibraryManagerParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateLibraryManagerParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateLibraryManagerParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateLibraryManagerParamsValidationError{}

// Validate checks the field values on LibraryManagerObject with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LibraryManagerObject) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LibraryManagerObject with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LibraryManagerObjectMultiError, or nil if none found.
func (m *LibraryManagerObject) ValidateAll() error {
	return m.validate(true)
}

func (m *LibraryManagerObject) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSpec()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LibraryManagerObjectValidationError{
					field:  "Spec",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LibraryManagerObjectValidationError{
					field:  "Spec",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSpec()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LibraryManagerObjectValidationError{
				field:  "Spec",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLibraryInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LibraryManagerObjectValidationError{
					field:  "LibraryInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LibraryManagerObjectValidationError{
					field:  "LibraryInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLibraryInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LibraryManagerObjectValidationError{
				field:  "LibraryInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSystem()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LibraryManagerObjectValidationError{
					field:  "System",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LibraryManagerObjectValidationError{
					field:  "System",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSystem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LibraryManagerObjectValidationError{
				field:  "System",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDataManagerInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LibraryManagerObjectValidationError{
					field:  "DataManagerInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LibraryManagerObjectValidationError{
					field:  "DataManagerInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDataManagerInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LibraryManagerObjectValidationError{
				field:  "DataManagerInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LibraryManagerObjectMultiError(errors)
	}

	return nil
}

// LibraryManagerObjectMultiError is an error wrapping multiple validation
// errors returned by LibraryManagerObject.ValidateAll() if the designated
// constraints aren't met.
type LibraryManagerObjectMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LibraryManagerObjectMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LibraryManagerObjectMultiError) AllErrors() []error { return m }

// LibraryManagerObjectValidationError is the validation error returned by
// LibraryManagerObject.Validate if the designated constraints aren't met.
type LibraryManagerObjectValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LibraryManagerObjectValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LibraryManagerObjectValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LibraryManagerObjectValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LibraryManagerObjectValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LibraryManagerObjectValidationError) ErrorName() string {
	return "LibraryManagerObjectValidationError"
}

// Error satisfies the builtin error interface
func (e LibraryManagerObjectValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLibraryManagerObject.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LibraryManagerObjectValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LibraryManagerObjectValidationError{}

// Validate checks the field values on LibraryManagersReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LibraryManagersReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LibraryManagersReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LibraryManagersReplyMultiError, or nil if none found.
func (m *LibraryManagersReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LibraryManagersReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LibraryManagersReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LibraryManagersReplyValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LibraryManagersReplyValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalSize

	if len(errors) > 0 {
		return LibraryManagersReplyMultiError(errors)
	}

	return nil
}

// LibraryManagersReplyMultiError is an error wrapping multiple validation
// errors returned by LibraryManagersReply.ValidateAll() if the designated
// constraints aren't met.
type LibraryManagersReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LibraryManagersReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LibraryManagersReplyMultiError) AllErrors() []error { return m }

// LibraryManagersReplyValidationError is the validation error returned by
// LibraryManagersReply.Validate if the designated constraints aren't met.
type LibraryManagersReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LibraryManagersReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LibraryManagersReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LibraryManagersReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LibraryManagersReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LibraryManagersReplyValidationError) ErrorName() string {
	return "LibraryManagersReplyValidationError"
}

// Error satisfies the builtin error interface
func (e LibraryManagersReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLibraryManagersReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LibraryManagersReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LibraryManagersReplyValidationError{}

// Validate checks the field values on TapeLibraryInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TapeLibraryInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TapeLibraryInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TapeLibraryInfoMultiError, or nil if none found.
func (m *TapeLibraryInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *TapeLibraryInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Vendor

	// no validation rules for SerialNumber

	// no validation rules for ProductId

	// no validation rules for LibraryName

	// no validation rules for FirmwareVersion

	// no validation rules for State

	if len(errors) > 0 {
		return TapeLibraryInfoMultiError(errors)
	}

	return nil
}

// TapeLibraryInfoMultiError is an error wrapping multiple validation errors
// returned by TapeLibraryInfo.ValidateAll() if the designated constraints
// aren't met.
type TapeLibraryInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TapeLibraryInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TapeLibraryInfoMultiError) AllErrors() []error { return m }

// TapeLibraryInfoValidationError is the validation error returned by
// TapeLibraryInfo.Validate if the designated constraints aren't met.
type TapeLibraryInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TapeLibraryInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TapeLibraryInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TapeLibraryInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TapeLibraryInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TapeLibraryInfoValidationError) ErrorName() string { return "TapeLibraryInfoValidationError" }

// Error satisfies the builtin error interface
func (e TapeLibraryInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTapeLibraryInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TapeLibraryInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TapeLibraryInfoValidationError{}

// Validate checks the field values on SystemInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SystemInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SystemInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SystemInfoMultiError, or
// nil if none found.
func (m *SystemInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *SystemInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Revision

	// no validation rules for CapacityBytes

	// no validation rules for TotalCatridges

	// no validation rules for LicensedCapacityBytes

	// no validation rules for NumberCopies

	// no validation rules for Worm

	// no validation rules for FaultyCatridges

	// no validation rules for CleaningCatridges

	// no validation rules for NumberFiles

	if len(errors) > 0 {
		return SystemInfoMultiError(errors)
	}

	return nil
}

// SystemInfoMultiError is an error wrapping multiple validation errors
// returned by SystemInfo.ValidateAll() if the designated constraints aren't met.
type SystemInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SystemInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SystemInfoMultiError) AllErrors() []error { return m }

// SystemInfoValidationError is the validation error returned by
// SystemInfo.Validate if the designated constraints aren't met.
type SystemInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SystemInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SystemInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SystemInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SystemInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SystemInfoValidationError) ErrorName() string { return "SystemInfoValidationError" }

// Error satisfies the builtin error interface
func (e SystemInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSystemInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SystemInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SystemInfoValidationError{}

// Validate checks the field values on LTFSDataManagerInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LTFSDataManagerInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LTFSDataManagerInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LTFSDataManagerInfoMultiError, or nil if none found.
func (m *LTFSDataManagerInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *LTFSDataManagerInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Revision

	// no validation rules for State

	if len(errors) > 0 {
		return LTFSDataManagerInfoMultiError(errors)
	}

	return nil
}

// LTFSDataManagerInfoMultiError is an error wrapping multiple validation
// errors returned by LTFSDataManagerInfo.ValidateAll() if the designated
// constraints aren't met.
type LTFSDataManagerInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LTFSDataManagerInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LTFSDataManagerInfoMultiError) AllErrors() []error { return m }

// LTFSDataManagerInfoValidationError is the validation error returned by
// LTFSDataManagerInfo.Validate if the designated constraints aren't met.
type LTFSDataManagerInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LTFSDataManagerInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LTFSDataManagerInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LTFSDataManagerInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LTFSDataManagerInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LTFSDataManagerInfoValidationError) ErrorName() string {
	return "LTFSDataManagerInfoValidationError"
}

// Error satisfies the builtin error interface
func (e LTFSDataManagerInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLTFSDataManagerInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LTFSDataManagerInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LTFSDataManagerInfoValidationError{}

// Validate checks the field values on TapeLibrarySetting with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TapeLibrarySetting) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TapeLibrarySetting with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TapeLibrarySettingMultiError, or nil if none found.
func (m *TapeLibrarySetting) ValidateAll() error {
	return m.validate(true)
}

func (m *TapeLibrarySetting) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NumberCopies

	// no validation rules for Enable_WORM

	if len(errors) > 0 {
		return TapeLibrarySettingMultiError(errors)
	}

	return nil
}

// TapeLibrarySettingMultiError is an error wrapping multiple validation errors
// returned by TapeLibrarySetting.ValidateAll() if the designated constraints
// aren't met.
type TapeLibrarySettingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TapeLibrarySettingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TapeLibrarySettingMultiError) AllErrors() []error { return m }

// TapeLibrarySettingValidationError is the validation error returned by
// TapeLibrarySetting.Validate if the designated constraints aren't met.
type TapeLibrarySettingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TapeLibrarySettingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TapeLibrarySettingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TapeLibrarySettingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TapeLibrarySettingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TapeLibrarySettingValidationError) ErrorName() string {
	return "TapeLibrarySettingValidationError"
}

// Error satisfies the builtin error interface
func (e TapeLibrarySettingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTapeLibrarySetting.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TapeLibrarySettingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TapeLibrarySettingValidationError{}

// Validate checks the field values on HostInfo with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HostInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HostInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in HostInfoMultiError, or nil
// if none found.
func (m *HostInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *HostInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IpAddress

	// no validation rules for HostName

	if len(errors) > 0 {
		return HostInfoMultiError(errors)
	}

	return nil
}

// HostInfoMultiError is an error wrapping multiple validation errors returned
// by HostInfo.ValidateAll() if the designated constraints aren't met.
type HostInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HostInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HostInfoMultiError) AllErrors() []error { return m }

// HostInfoValidationError is the validation error returned by
// HostInfo.Validate if the designated constraints aren't met.
type HostInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HostInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HostInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HostInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HostInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HostInfoValidationError) ErrorName() string { return "HostInfoValidationError" }

// Error satisfies the builtin error interface
func (e HostInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHostInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HostInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HostInfoValidationError{}

// Validate checks the field values on LibraryManagerSpec with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LibraryManagerSpec) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LibraryManagerSpec with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LibraryManagerSpecMultiError, or nil if none found.
func (m *LibraryManagerSpec) ValidateAll() error {
	return m.validate(true)
}

func (m *LibraryManagerSpec) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for LibraryType

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetNodeInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LibraryManagerSpecValidationError{
					field:  "NodeInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LibraryManagerSpecValidationError{
					field:  "NodeInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNodeInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LibraryManagerSpecValidationError{
				field:  "NodeInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSettings()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LibraryManagerSpecValidationError{
					field:  "Settings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LibraryManagerSpecValidationError{
					field:  "Settings",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSettings()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LibraryManagerSpecValidationError{
				field:  "Settings",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LibraryManagerSpecMultiError(errors)
	}

	return nil
}

// LibraryManagerSpecMultiError is an error wrapping multiple validation errors
// returned by LibraryManagerSpec.ValidateAll() if the designated constraints
// aren't met.
type LibraryManagerSpecMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LibraryManagerSpecMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LibraryManagerSpecMultiError) AllErrors() []error { return m }

// LibraryManagerSpecValidationError is the validation error returned by
// LibraryManagerSpec.Validate if the designated constraints aren't met.
type LibraryManagerSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LibraryManagerSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LibraryManagerSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LibraryManagerSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LibraryManagerSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LibraryManagerSpecValidationError) ErrorName() string {
	return "LibraryManagerSpecValidationError"
}

// Error satisfies the builtin error interface
func (e LibraryManagerSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLibraryManagerSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LibraryManagerSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LibraryManagerSpecValidationError{}