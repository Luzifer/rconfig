package rconfig

import (
	"fmt"
	"reflect"
	"strconv"
)

func parseIntForType(s string, base int, fieldType reflect.Kind) (i int64, err error) {
	switch fieldType {
	case reflect.Int:
		return strconv.ParseInt(s, base, strconv.IntSize) //nolint:wrapcheck

	case reflect.Int8:
		return strconv.ParseInt(s, base, 8) //nolint:wrapcheck

	case reflect.Int16:
		return strconv.ParseInt(s, base, 16) //nolint:wrapcheck

	case reflect.Int32:
		return strconv.ParseInt(s, base, 32) //nolint:wrapcheck

	case reflect.Int64:
		return strconv.ParseInt(s, base, 64) //nolint:wrapcheck

	default:
		return 0, fmt.Errorf("unsupported type: %v", fieldType)
	}
}

func parseUintForType(s string, base int, fieldType reflect.Kind) (uint64, error) {
	switch fieldType {
	case reflect.Uint:
		return strconv.ParseUint(s, base, strconv.IntSize) //nolint:wrapcheck

	case reflect.Uint8:
		return strconv.ParseUint(s, base, 8) //nolint:wrapcheck

	case reflect.Uint16:
		return strconv.ParseUint(s, base, 16) //nolint:wrapcheck

	case reflect.Uint32:
		return strconv.ParseUint(s, base, 32) //nolint:wrapcheck

	case reflect.Uint64:
		return strconv.ParseUint(s, base, 64) //nolint:wrapcheck

	default:
		return 0, fmt.Errorf("unsupported type: %v", fieldType)
	}
}
