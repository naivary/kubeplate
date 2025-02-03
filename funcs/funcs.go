package funcs

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"google.golang.org/protobuf/types/known/structpb"
)

func Get(data map[string]*structpb.Struct) func(filename, key string) (any, error) {
	return func(filename, key string) (any, error) {
		var scope any
		const delim = "."
		strct, isFileExisting := data[filename]
		if !isFileExisting {
			return nil, fmt.Errorf("variable file `%s` does not exist", filename)
		}
		accessPath := strings.Split(key, delim)
		if len(accessPath) == 0 {
			return nil, errors.New("empty key in `get` function")
		}
		m := strct.AsMap()
		scope, isKeyExisting := m[accessPath[0]]
		if !isKeyExisting {
			return nil, fmt.Errorf("key `%s` does not exist in map", accessPath[0])
		}
		for _, el := range accessPath[1:] {
			if i, err := strconv.Atoi(el); err == nil {
				// its an integer, array is getting accessed
				arr, isConverted := scope.([]any)
				if !isConverted {
					return nil, fmt.Errorf("`%s` is trying to access the element of an array (`%s`) but its not an array", strings.Join(accessPath, ","), el)
				}
				scope = arr[i]
				continue
			}

			if unicode.IsLetter(rune(el[0])) || unicode.IsDigit(rune(el[0])) {
				// its a map access
				obj, isConverted := scope.(map[string]any)
				if !isConverted {
					return nil, fmt.Errorf("`%s` is not an object", el)
				}
				scope = obj[el]
			}
		}
		return scope, nil
	}
}
