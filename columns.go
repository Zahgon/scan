package scan

import (
	"errors"
	"reflect"
	"sync"
)

const dbTag = "db"

var (
	// ErrNotAPointer is returned when a non-pointer is received
	// when a pointer is expected.
	ErrNotAPointer = errors.New("not a pointer")

	// ErrNotAStructPointer is returned when a non-struct pointer
	// is received but a struct pointer was expected
	ErrNotAStructPointer = errors.New("not a struct pointer")

	// ErrNotASlicePointer is returned when receiving an argument
	// that is expected to be a slice pointer, but it is not
	ErrNotASlicePointer = errors.New("not a slice pointer")

	// ErrStructFieldMissing is returned when trying to scan a value
	// to a column which does not match a struct. This means that
	// the struct does not have a field that matches the column
	// specified.
	ErrStructFieldMissing = errors.New("struct field missing")

	// ColumnsMapper transforms struct/map field names
	// into the database column names.
	// E.g. you can set function for convert CamelCase into snake_case
	ColumnsMapper = func(name string) string { return name }
)

var columnsCache cache = &sync.Map{}

type cacheKey struct {
	Type   reflect.Type
	Strict bool
}

// Columns scans a struct and returns a list of strings
// that represent the assumed column names based on the
// db struct tag, or the field name. Any field or struct
// tag that matches a string within the excluded list
// will be excluded from the result.
func Columns(v interface{}, excluded ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ColumnsStrict is identical to Columns, but it only
// searches struct tags and excludes fields not tagged
// with the db struct tag.
func ColumnsStrict(v interface{}, excluded ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func columns(v interface{}, strict bool, excluded ...string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func columnNames(model reflect.Value, strict bool, excluded ...string) []string {
	_ = "STUB: not implemented"
	return nil
}

// there's no tag name and we're in strict mode so move on

func isExcluded(name string, excluded ...string) bool { _ = "STUB: not implemented"; return false }

func reflectValue(v interface{}) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func supportedColumnType(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

func isValidSqlValue(v reflect.Value) bool {
	_ = "STUB: not implemented"
	// This method covers two cases in which we know the Value can be converted to sql:
	//  1. It returns true for sql.driver's type check for types like time.Time
	//  2. It implements the driver.Valuer interface allowing conversion directly
	//     into sql statements
	return false
}
