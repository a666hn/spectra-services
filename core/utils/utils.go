package utils

import (
	// "crypto/sha256"
	// "database/sql"
	// "encoding/hex"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/jinzhu/copier"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"

	appConfig "github.com/skinnyguy/spectra-services/core/config"
	connection "github.com/skinnyguy/spectra-services/core/connection"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

const (
	// Default Date Format ...
	DateFomat = "2008-08-08"
	// Default time format ...
	TimeFormat = "23:59:59"
	// Default time format without second
	TimeFormatWithoutSecond = "23:59"
	// Core utils ...
	CoreUtils = "core.utils"
	// Postgresql ...
	Postgres = "postgresql"
	// Float type 32
	Float32 = 32
	// Float type 64
	Float64 = 64
)

type User struct {
	ID          string
	FamilyOrgID string
	Role        string
}

var (
	UserData = User{
		ID:          "id",
		FamilyOrgID: "familyOrgID",
	}
)

// GetDateFromTimestamp ...
func GetDateFromTimestamp(input time.Time) string {
	return input.Format(DateFomat)
}

// GetTimeFromTimestamp ...
func GetTimeFromTimestamp(input time.Time) string {
	return input.Format(TimeFormat)
}

// GetTimeFromTimestampWithoutSecond ...
func GetTimeFromTimestampWithoutSecond(input time.Time) string {
	return input.Format(TimeFormatWithoutSecond)
}

// StringToTime ...
func StringToTime(input string) (time.Time, error) {
	return time.Parse(time.RFC3339, input)
}

// StringToBoolean ...
func StringToBoolean(input string) (bool, error) {
	result, _ := strconv.ParseBool(input)
	return result, nil
}

// StringToInt ...
func StringToInt(input string) (int, error) {
	result, _ := strconv.Atoi(input)
	return result, nil
}

// StringToInt8 ...
func StringToInt8(input string) (int8, error) {
	result, _ := strconv.Atoi(input)
	return int8(result), nil
}

// StringToInt16 ...
func StringToInt16(input string) (int16, error) {
	result, _ := strconv.Atoi(input)
	return int16(result), nil
}

// StringToInt32 ...
func StringToInt32(input string) (int32, error) {
	result, _ := strconv.Atoi(input)
	return int32(result), nil
}

// StringToInt64 ...
func StringToInt64(input string) (int64, error) {
	result, _ := strconv.Atoi(input)
	return int64(result), nil
}

// StringToFloat32 ...
func StringToFloat32(input string) (float32, error) {
	result, _ := strconv.ParseFloat(input, Float32)
	return float32(result), nil
}

// StringToFloat64 ...
func StringToFloat64(input string) (float64, error) {
	result, _ := strconv.ParseFloat(input, Float64)
	return float64(result), nil
}

// CopyObject ...
func CopyObject(from interface{}, to interface{}) error {
	if from == nil {
		to = nil
		return nil
	}

	return copier.Copy(to, from)
}

// SendLogError ...
func SendLogError(id string, err error) error {
	log.Log(errors.BadRequest(id, err.Error()))
	return err
}

// SendLogInfo ...
func SendLogInfo(msg ...string) {
	log.Info(msg)
}

// SendLogDebug ...
func SendLogDebug(msg ...string) {
	log.Debug(msg)
}

var (
	conn    connection.Connection
	oneConn sync.Once
)

// GetPostgresHandler ...
func GetPostgresHandler() connection.Connection {
	oneConn.Do(func() {
		conn = connection.NewConnection(appConfig.GetSpectraConfig())
	})

	return conn
}

// GetDatasourceInfo ...
func GetDatasourceInfo() string {
	if typ := os.Getenv("REPO_TYPE"); typ != "" {
		return typ
	}

	return Postgres
}
