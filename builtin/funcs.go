package builtin

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"google.golang.org/protobuf/types/known/structpb"
)

func Get(data map[string]*structpb.Struct) func(filename, key string) (any, error) {
	// name.old
	// name: {old: 3}

	// name.0, name.1
	// name: [first_element, second_element]

	// array = []any
	// map = map[string]any
	return func(filename, key string) (any, error) {
		var value any
		const delim = "."
		str, ok := data[filename]
		if !ok {
			return nil, fmt.Errorf("variable file `%s` does not exist", filename)
		}
		accessPath := strings.Split(key, delim)
		if len(accessPath) == 0 {
			return nil, errors.New("empty key in `get` function")
		}
		m := str.AsMap()
		value = m[accessPath[0]]
		for _, el := range accessPath[1:] {
			if i, err := strconv.Atoi(el); err == nil {
				// its an integer, array is getting accessed

				arr, ok := value.([]any)
				if !ok {
					return nil, fmt.Errorf("`%s` is trying to access the element of an array (`%s`) but its not an array", strings.Join(accessPath, ","), el)
				}
				value = arr[i]
			}

			if isLetter := unicode.IsLetter(rune(el[0])); isLetter {
				// its a map access
				obj, ok := value.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("`%s` is not an object", el)
				}
				value = obj[el]
			}
		}
		return value, nil
	}
}
