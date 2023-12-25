package storage_logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/lowl11/flex"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

func (logger Logger) write(env, service, level, message string) error {
	if logger.connection == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	id := uuid.New()
	statement, err := logger.connection.Preparex(`
INSERT INTO logs (id, environment, service, level, message)
VALUES ($1, $2, $3, $4, $5)
`)
	if err != nil {
		return err
	}

	_, err = statement.ExecContext(ctx, id, env, service, level, message)
	if err != nil {
		return err
	}

	return nil
}

func build(args ...any) string {
	if len(args) == 0 {
		return ""
	}

	stringArgs := strings.Builder{}
	for _, arg := range args {
		stringArgs.WriteString(toString(arg, true))
		stringArgs.WriteString(" ")
	}
	return stringArgs.String()[:stringArgs.Len()-1]
}

func toString(anyValue any, memory bool) string {
	if anyValue == nil {
		return ""
	}

	// try cast to error
	if _, ok := anyValue.(error); ok {
		return anyValue.(error).Error()
	}

	// try cast to bytes
	if bytesBuffer, ok := anyValue.([]byte); ok {
		return bytesToString(bytesBuffer)
	}

	// try cast uuid
	if flex.Type(reflect.TypeOf(anyValue)).IsUUID() {
		uuidValue, ok := anyValue.(uuid.UUID)
		if ok {
			return uuidValue.String()
		}

		uuidPtrValue, ok := anyValue.(*uuid.UUID)
		if ok {
			return uuidPtrValue.String()
		}
	}

	value := reflect.ValueOf(anyValue)

	switch value.Kind() {
	case reflect.String:
		return anyValue.(string)
	case reflect.Bool:
		return strconv.FormatBool(anyValue.(bool))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32:
		return fmt.Sprintf("%f", value.Float())
	case reflect.Float64:
		return fmt.Sprintf("%g", value.Float())
	case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
		valueInBytes, err := json.Marshal(anyValue)
		if err != nil {
			return ""
		}
		return string(valueInBytes)
	case reflect.Ptr:
		if memory {
			return fmt.Sprintf("%v", value)
		}

		return toString(value.Elem().Interface(), true)
	default:
		return fmt.Sprintf("%v", value)
	}
}

func bytesToString(buffer []byte) string {
	return *(*string)(unsafe.Pointer(&buffer))
}
