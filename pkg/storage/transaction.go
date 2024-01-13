package storage

import (
	"database/sql"
	"sync/atomic"
)

var _isolationLevel atomic.Int32

func init() {
	_isolationLevel.Store(int32(sql.LevelDefault))
}

func SetIsolationLevel(isolationLevel sql.IsolationLevel) {
	_isolationLevel.Store(int32(isolationLevel))
}

func IsolationLevel() sql.IsolationLevel {
	return sql.IsolationLevel(_isolationLevel.Load())
}

func TxOptions() *sql.TxOptions {
	return &sql.TxOptions{
		Isolation: IsolationLevel(),
		ReadOnly:  false,
	}
}
