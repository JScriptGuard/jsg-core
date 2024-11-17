//go:build darwin
// +build darwin

package cenvxcore

import (
	"path/filepath"
)

var (
	// Gibt das Standard Config Verzeichniss an
	//LINUX_HOST_CONFIG_DIR_PATH  types.HOST_CONFIG_PATH = types.HOST_CONFIG_PATH(filepath.Join("/", "etc", "CustodiaJS"))
	HOST_CONFIG_DIR_PATH HOST_CONFIG_PATH = HOST_CONFIG_PATH(filepath.Join("/", "Library", "Application Support", "CustodiaJS"))

	// Log Pfade
	//LINUX_DEFAULT_LOGGING_DIR_PATH  types.LOG_DIR = types.LOG_DIR(filepath.Join("var", "log", "CustodiaJS"))
	DEFAULT_LOGGING_DIR_PATH LOG_DIR = LOG_DIR(filepath.Join("/", "Library", "Logs", "CustodiaJS"))

	// Gibt die Sockets für den Hypervisor an, wird verwendet damit der Hypervisor mit dem Host Kommunizieren kann
	//LINUX_CNH_SOCKET_PATH  types.CHN_CORE_SOCKET_PATH = types.CHN_CORE_SOCKET_PATH("")
	DARWIN_CNH_SOCKET_PATH CHN_CORE_SOCKET_PATH = CHN_CORE_SOCKET_PATH("")

	// Legt die Dateipfade für z.b Unix Sockets fest
	NONE_ROOT_UNIX_SOCKET SOCKET_PATH = SOCKET_PATH(filepath.Join("/", "tmp", "cusjs_none_root_sock"))
	ROOT_UNIX_SOCKET      SOCKET_PATH = SOCKET_PATH(filepath.Join("/", "tmp", "cusjs_root_sock"))
)
