// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/post.proto

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

// Validate checks the field values on CreatePostRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreatePostRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePostRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePostRequestMultiError, or nil if none found.
func (m *CreatePostRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePostRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Content

	if len(errors) > 0 {
		return CreatePostRequestMultiError(errors)
	}

	return nil
}

// CreatePostRequestMultiError is an error wrapping multiple validation errors
// returned by CreatePostRequest.ValidateAll() if the designated constraints
// aren't met.
type CreatePostRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePostRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePostRequestMultiError) AllErrors() []error { return m }

// CreatePostRequestValidationError is the validation error returned by
// CreatePostRequest.Validate if the designated constraints aren't met.
type CreatePostRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePostRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePostRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePostRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePostRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePostRequestValidationError) ErrorName() string {
	return "CreatePostRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePostRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePostRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePostRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePostRequestValidationError{}

// Validate checks the field values on CreatePostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreatePostResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreatePostResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreatePostResponseMultiError, or nil if none found.
func (m *CreatePostResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreatePostResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Content

	if len(errors) > 0 {
		return CreatePostResponseMultiError(errors)
	}

	return nil
}

// CreatePostResponseMultiError is an error wrapping multiple validation errors
// returned by CreatePostResponse.ValidateAll() if the designated constraints
// aren't met.
type CreatePostResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreatePostResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreatePostResponseMultiError) AllErrors() []error { return m }

// CreatePostResponseValidationError is the validation error returned by
// CreatePostResponse.Validate if the designated constraints aren't met.
type CreatePostResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePostResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePostResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePostResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePostResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePostResponseValidationError) ErrorName() string {
	return "CreatePostResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePostResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePostResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePostResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePostResponseValidationError{}
