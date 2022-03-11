package win_registry

import "syscall"

const (
	REG_OPTION_NON_VOLATILE = 0

	REG_CREATED_NEW_KEY     = 1
	REG_OPENED_EXISTING_KEY = 2

	ERROR_NO_MORE_ITEMS syscall.Errno = 259
)
