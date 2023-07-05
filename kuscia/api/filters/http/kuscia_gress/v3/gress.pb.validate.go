// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: kuscia/api/filters/http/kuscia_gress/v3/gress.proto

package v3

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

// Validate checks the field values on Gress with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Gress) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Gress with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in GressMultiError, or nil if none found.
func (m *Gress) ValidateAll() error {
	return m.validate(true)
}

func (m *Gress) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Instance

	// no validation rules for SelfNamespace

	for idx, item := range m.GetRewriteHostConfig() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GressValidationError{
						field:  fmt.Sprintf("RewriteHostConfig[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GressValidationError{
						field:  fmt.Sprintf("RewriteHostConfig[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GressValidationError{
					field:  fmt.Sprintf("RewriteHostConfig[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for AddOriginSource

	// no validation rules for MaxLoggingBodySizePerReqeuest

	if len(errors) > 0 {
		return GressMultiError(errors)
	}

	return nil
}

// GressMultiError is an error wrapping multiple validation errors returned by
// Gress.ValidateAll() if the designated constraints aren't met.
type GressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GressMultiError) AllErrors() []error { return m }

// GressValidationError is the validation error returned by Gress.Validate if
// the designated constraints aren't met.
type GressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GressValidationError) ErrorName() string { return "GressValidationError" }

// Error satisfies the builtin error interface
func (e GressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GressValidationError{}

// Validate checks the field values on Gress_RewriteHostByHeader with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Gress_RewriteHostByHeader) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Gress_RewriteHostByHeader with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Gress_RewriteHostByHeaderMultiError, or nil if none found.
func (m *Gress_RewriteHostByHeader) ValidateAll() error {
	return m.validate(true)
}

func (m *Gress_RewriteHostByHeader) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RewritePolicy

	// no validation rules for Header

	for idx, item := range m.GetPathMatchers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, Gress_RewriteHostByHeaderValidationError{
						field:  fmt.Sprintf("PathMatchers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, Gress_RewriteHostByHeaderValidationError{
						field:  fmt.Sprintf("PathMatchers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Gress_RewriteHostByHeaderValidationError{
					field:  fmt.Sprintf("PathMatchers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for SpecifiedHost

	if len(errors) > 0 {
		return Gress_RewriteHostByHeaderMultiError(errors)
	}

	return nil
}

// Gress_RewriteHostByHeaderMultiError is an error wrapping multiple validation
// errors returned by Gress_RewriteHostByHeader.ValidateAll() if the
// designated constraints aren't met.
type Gress_RewriteHostByHeaderMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Gress_RewriteHostByHeaderMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Gress_RewriteHostByHeaderMultiError) AllErrors() []error { return m }

// Gress_RewriteHostByHeaderValidationError is the validation error returned by
// Gress_RewriteHostByHeader.Validate if the designated constraints aren't met.
type Gress_RewriteHostByHeaderValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Gress_RewriteHostByHeaderValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Gress_RewriteHostByHeaderValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Gress_RewriteHostByHeaderValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Gress_RewriteHostByHeaderValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Gress_RewriteHostByHeaderValidationError) ErrorName() string {
	return "Gress_RewriteHostByHeaderValidationError"
}

// Error satisfies the builtin error interface
func (e Gress_RewriteHostByHeaderValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGress_RewriteHostByHeader.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Gress_RewriteHostByHeaderValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Gress_RewriteHostByHeaderValidationError{}
