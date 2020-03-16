package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry 复合类型入口
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (cp CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range cp {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (cp CompositeEntry) String() string {
	strs := make([]string, len(cp))
	for i, entry := range cp {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
