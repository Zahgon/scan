package scan

import (
	"errors"
	"io"
	"reflect"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	// ErrTooManyColumns indicates that a select query returned multiple columns and
	// attempted to bind to a slice of a primitive type. For example, trying to bind
	// `select col1, col2 from mutable` to []string
	ErrTooManyColumns = errors.New("too many columns returned for primitive slice")

	// ErrSliceForRow occurs when trying to use Row on a slice
	ErrSliceForRow = errors.New("cannot scan Row into slice")

	// AutoClose is true when scan should automatically close Scanner when the scan
	// is complete. If you set it to false, then you must defer rows.Close() manually
	AutoClose = true

	// OnAutoCloseError can be used to log errors which are returned from rows.Close()
	// By default this is a NOOP function
	OnAutoCloseError = func(error) {}

	// ScannerMapper transforms database field names into struct/map field names
	// E.g. you can set function for convert snake_case into CamelCase
	ScannerMapper = func(name string) string { return cases.Title(language.English).String(name) }
)

// Row scans a single row into a single variable. It requires that you use
// db.Query and not db.QueryRow, because QueryRow does not return column names.
// There is no performance impact in using one over the other. QueryRow only
// defers returning err until Scan is called, which is an unnecessary
// optimization for this library.
func Row(v interface{}, r RowsScanner) error { _ = "STUB: not implemented"; return nil }

// RowStrict scans a single row into a single variable. It is identical to
// Row, but it ignores fields that do not have a db tag
func RowStrict(v interface{}, r RowsScanner) error { _ = "STUB: not implemented"; return nil }

func row(v interface{}, r RowsScanner, strict bool) error { _ = "STUB: not implemented"; return nil }

// Rows scans sql rows into a slice (v)
func Rows(v interface{}, r RowsScanner) (outerr error) { _ = "STUB: not implemented"; return nil }

// RowsStrict scans sql rows into a slice (v) only using db tags
func RowsStrict(v interface{}, r RowsScanner) (outerr error) { _ = "STUB: not implemented"; return nil }

func rows(v interface{}, r RowsScanner, strict bool) (outerr error) {
	_ = "STUB: not implemented"
	return nil
}

// Initialization the tags from struct.
func initFieldTag(sliceItem reflect.Value, fieldTagMap *map[string]reflect.Value) {
	_ = "STUB: not implemented"
	return
}

// found an embedded struct

func structPointers(sliceItem reflect.Value, cols []string, strict bool) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

// have to add if we found a column because Scan() requires
// len(cols) arguments or it will error. This way we can scan to
// a useless pointer

func closeRows(c io.Closer) { _ = "STUB: not implemented"; return }
