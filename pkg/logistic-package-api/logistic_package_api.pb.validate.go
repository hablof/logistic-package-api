// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: hablof/logistic_package_api/v1/logistic_package_api.proto

package logistic_package_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on MaybeTimestamp with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MaybeTimestamp) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MaybeTimestampValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MaybeTimestampValidationError is the validation error returned by
// MaybeTimestamp.Validate if the designated constraints aren't met.
type MaybeTimestampValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MaybeTimestampValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MaybeTimestampValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MaybeTimestampValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MaybeTimestampValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MaybeTimestampValidationError) ErrorName() string { return "MaybeTimestampValidationError" }

// Error satisfies the builtin error interface
func (e MaybeTimestampValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMaybeTimestamp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MaybeTimestampValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MaybeTimestampValidationError{}

// Validate checks the field values on Package with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Package) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ID

	// no validation rules for Title

	// no validation rules for Material

	// no validation rules for MaximumVolume

	// no validation rules for Reusable

	if v, ok := interface{}(m.GetCreated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PackageValidationError{
				field:  "Created",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdated()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PackageValidationError{
				field:  "Updated",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PackageValidationError is the validation error returned by Package.Validate
// if the designated constraints aren't met.
type PackageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PackageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PackageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PackageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PackageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PackageValidationError) ErrorName() string { return "PackageValidationError" }

// Error satisfies the builtin error interface
func (e PackageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPackage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PackageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PackageValidationError{}

// Validate checks the field values on MaybeBool with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *MaybeBool) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Reusable

	return nil
}

// MaybeBoolValidationError is the validation error returned by
// MaybeBool.Validate if the designated constraints aren't met.
type MaybeBoolValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MaybeBoolValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MaybeBoolValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MaybeBoolValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MaybeBoolValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MaybeBoolValidationError) ErrorName() string { return "MaybeBoolValidationError" }

// Error satisfies the builtin error interface
func (e MaybeBoolValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMaybeBool.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MaybeBoolValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MaybeBoolValidationError{}

// Validate checks the field values on CreatePackageV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreatePackageV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetTitle()); l < 1 || l > 32 {
		return CreatePackageV1RequestValidationError{
			field:  "Title",
			reason: "value length must be between 1 and 32 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetMaterial()); l < 1 || l > 32 {
		return CreatePackageV1RequestValidationError{
			field:  "Material",
			reason: "value length must be between 1 and 32 runes, inclusive",
		}
	}

	if m.GetMaximumVolume() <= 0 {
		return CreatePackageV1RequestValidationError{
			field:  "MaximumVolume",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Reusable

	return nil
}

// CreatePackageV1RequestValidationError is the validation error returned by
// CreatePackageV1Request.Validate if the designated constraints aren't met.
type CreatePackageV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePackageV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePackageV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePackageV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePackageV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePackageV1RequestValidationError) ErrorName() string {
	return "CreatePackageV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePackageV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePackageV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePackageV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePackageV1RequestValidationError{}

// Validate checks the field values on CreatePackageV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreatePackageV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ID

	return nil
}

// CreatePackageV1ResponseValidationError is the validation error returned by
// CreatePackageV1Response.Validate if the designated constraints aren't met.
type CreatePackageV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePackageV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePackageV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePackageV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePackageV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePackageV1ResponseValidationError) ErrorName() string {
	return "CreatePackageV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePackageV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePackageV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePackageV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePackageV1ResponseValidationError{}

// Validate checks the field values on DescribePackageV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribePackageV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPackageID() <= 0 {
		return DescribePackageV1RequestValidationError{
			field:  "PackageID",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribePackageV1RequestValidationError is the validation error returned by
// DescribePackageV1Request.Validate if the designated constraints aren't met.
type DescribePackageV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribePackageV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribePackageV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribePackageV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribePackageV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribePackageV1RequestValidationError) ErrorName() string {
	return "DescribePackageV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribePackageV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribePackageV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribePackageV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribePackageV1RequestValidationError{}

// Validate checks the field values on DescribePackageV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribePackageV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribePackageV1ResponseValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribePackageV1ResponseValidationError is the validation error returned by
// DescribePackageV1Response.Validate if the designated constraints aren't met.
type DescribePackageV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribePackageV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribePackageV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribePackageV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribePackageV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribePackageV1ResponseValidationError) ErrorName() string {
	return "DescribePackageV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribePackageV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribePackageV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribePackageV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribePackageV1ResponseValidationError{}

// Validate checks the field values on ListPackagesV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListPackagesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOffset() < 0 {
		return ListPackagesV1RequestValidationError{
			field:  "Offset",
			reason: "value must be greater than or equal to 0",
		}
	}

	return nil
}

// ListPackagesV1RequestValidationError is the validation error returned by
// ListPackagesV1Request.Validate if the designated constraints aren't met.
type ListPackagesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPackagesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPackagesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPackagesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPackagesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPackagesV1RequestValidationError) ErrorName() string {
	return "ListPackagesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListPackagesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPackagesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPackagesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPackagesV1RequestValidationError{}

// Validate checks the field values on ListPackagesV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListPackagesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListPackagesV1ResponseValidationError is the validation error returned by
// ListPackagesV1Response.Validate if the designated constraints aren't met.
type ListPackagesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListPackagesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListPackagesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListPackagesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListPackagesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListPackagesV1ResponseValidationError) ErrorName() string {
	return "ListPackagesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListPackagesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListPackagesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListPackagesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListPackagesV1ResponseValidationError{}

// Validate checks the field values on RemovePackageV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemovePackageV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPackageID() <= 0 {
		return RemovePackageV1RequestValidationError{
			field:  "PackageID",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemovePackageV1RequestValidationError is the validation error returned by
// RemovePackageV1Request.Validate if the designated constraints aren't met.
type RemovePackageV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemovePackageV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemovePackageV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemovePackageV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemovePackageV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemovePackageV1RequestValidationError) ErrorName() string {
	return "RemovePackageV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemovePackageV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemovePackageV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemovePackageV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemovePackageV1RequestValidationError{}

// Validate checks the field values on RemovePackageV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemovePackageV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Suc

	return nil
}

// RemovePackageV1ResponseValidationError is the validation error returned by
// RemovePackageV1Response.Validate if the designated constraints aren't met.
type RemovePackageV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemovePackageV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemovePackageV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemovePackageV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemovePackageV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemovePackageV1ResponseValidationError) ErrorName() string {
	return "RemovePackageV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemovePackageV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemovePackageV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemovePackageV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemovePackageV1ResponseValidationError{}

// Validate checks the field values on UpdatePackageV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdatePackageV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetPackageID() <= 0 {
		return UpdatePackageV1RequestValidationError{
			field:  "PackageID",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetTitle()) > 32 {
		return UpdatePackageV1RequestValidationError{
			field:  "Title",
			reason: "value length must be at most 32 runes",
		}
	}

	if utf8.RuneCountInString(m.GetMaterial()) > 32 {
		return UpdatePackageV1RequestValidationError{
			field:  "Material",
			reason: "value length must be at most 32 runes",
		}
	}

	if m.GetMaximumVolume() < 0 {
		return UpdatePackageV1RequestValidationError{
			field:  "MaximumVolume",
			reason: "value must be greater than or equal to 0",
		}
	}

	if v, ok := interface{}(m.GetReusable()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePackageV1RequestValidationError{
				field:  "Reusable",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdatePackageV1RequestValidationError is the validation error returned by
// UpdatePackageV1Request.Validate if the designated constraints aren't met.
type UpdatePackageV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePackageV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePackageV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePackageV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePackageV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePackageV1RequestValidationError) ErrorName() string {
	return "UpdatePackageV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePackageV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePackageV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePackageV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePackageV1RequestValidationError{}

// Validate checks the field values on UpdatePackageV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdatePackageV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Suc

	return nil
}

// UpdatePackageV1ResponseValidationError is the validation error returned by
// UpdatePackageV1Response.Validate if the designated constraints aren't met.
type UpdatePackageV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePackageV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePackageV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePackageV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePackageV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePackageV1ResponseValidationError) ErrorName() string {
	return "UpdatePackageV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePackageV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePackageV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePackageV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePackageV1ResponseValidationError{}
