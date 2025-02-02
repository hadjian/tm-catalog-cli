package utils

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

// ReadRequiredFile reads the file. Returns expanded absolute representation of the filename and file contents.
// Removes Byte-Order-Mark from the content
func ReadRequiredFile(name string) (string, []byte, error) {
	var log = slog.Default()

	abs, err := filepath.Abs(name)
	if err != nil {
		log.Error("error expanding file name", "filename", name, "error", err)
		return "", nil, err
	}
	log.Debug("reading file", "filename", abs)

	stat, err := os.Stat(abs)
	if err != nil {
		log.Error("error reading file", "filename", abs, "error", err)
		return "", nil, err
	}
	if stat.IsDir() {
		err = errors.New("not a file")
		return "", nil, err
	}
	raw, err := os.ReadFile(abs)
	if err != nil {
		log.Error("error reading file", "filename", abs, "error", err)
		return "", nil, err
	}
	raw = removeBOM(raw)
	log.Debug(fmt.Sprintf("read %d bytes from %s beginning with %s", len(raw), abs, string(raw[:100])))
	return abs, raw, nil
}

func removeBOM(bytes []byte) []byte {
	if len(bytes) > 2 && bytes[0] == 0xef && bytes[1] == 0xbb && bytes[2] == 0xbf {
		bytes = bytes[3:]
	}
	return bytes
}

// ExpandHome expands ~ in path with user's home directory, but only if path begins with ~ or /~
// Otherwise, returns path unchanged
func ExpandHome(path string) (string, error) {
	if !strings.HasPrefix(path, "~") && !strings.HasPrefix(path, "/~") {
		return path, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		slog.Default().Error("cannot expand user home directory", "error", err)
		return "", fmt.Errorf("cannot expand user home directory: %w", err)
	}
	_, rest, found := strings.Cut(path, "~")
	if !found {
		panic(errors.New("should have checked for ~ before"))
	}
	return filepath.Join(home, rest), nil
}

func ToTrimmedLower(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

func NormalizeLineEndings(bytes []byte) []byte {
	res := make([]byte, 0, len(bytes))
	var prevB byte
	for _, b := range bytes {
		switch b {
		case '\n':
			if prevB != '\r' {
				res = append(res, '\n')
			}
		case '\r':
			res = append(res, '\n')
		default:
			res = append(res, b)
		}
		prevB = b
	}
	return res
}

func JsGetBool(js map[string]any, key string) *bool {
	if v, ok := js[key]; ok {
		if b, ok := v.(bool); ok {
			return &b
		}
	}
	return nil
}
func JsGetString(js map[string]any, key string) *string {
	if v, ok := js[key]; ok {
		if s, ok := v.(string); ok {
			return &s
		}
	}
	return nil
}
func JsGetMap(js map[string]any, key string) map[string]any {
	if v, ok := js[key]; ok {
		if m, ok := v.(map[string]any); ok {
			return m
		}
	}
	return nil
}

// ConvertToNativeLineEndings converts all instances of '\n' to native line endings for the platform.
// Assumes that line endings are normalized, i.e. there are no '\r' or "\r\n" line endings in the data
// See NormalizeLineEndings
func ConvertToNativeLineEndings(b []byte) []byte {
	return convertToNativeLineEndings(b)
}

// AtomicWriteFile writes data to the named file quasi-atomically, creating it if necessary.
// On unix-like systems, the function uses github.com/google/renameio.
// On Windows, it has a simpler implementation using os.Rename(), which is believed to be atomic on NTFS,
// but there is no hard guarantee from Microsoft on that.
func AtomicWriteFile(name string, data []byte, perm os.FileMode) error {
	return atomicWriteFile(name, data, perm)
}

func ParseAsList(list, separator string, trim bool) []string {
	ret := make([]string, 0)

	for _, entry := range strings.Split(list, separator) {
		if trim {
			entry = strings.TrimSpace(entry)
		}
		if entry != "" {
			ret = append(ret, entry)
		}
	}
	return ret
}
