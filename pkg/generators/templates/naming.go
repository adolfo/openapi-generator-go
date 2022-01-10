package templates

import (
	"fmt"
	"strings"
)

// TypeDisplayName returns the value with special characters replaced with
// there semantic meaning and then converted to pascal case. For example,
//     *string
// becomes
//     NullableString
func TypeDisplayName(value string) string {
	out := value

	if strings.HasPrefix(value, "optional.Value") {
		var typ string
		fmt.Sscanf(value, "optional.Value[%s]", &typ)
		out = fmt.Sprintf("optional_%s", typ)
	}

	if strings.HasPrefix(value, "*") {
		out = strings.Replace(out, "*", "nullable_", 1)
	}

	if strings.Contains(value, "[]") {
		out = strings.ReplaceAll(out, "[]", "")

		sliceCount := strings.Count(value, "[]")
		for sliceCount > 0 {
			out = fmt.Sprintf("%s_slice", out)
			sliceCount--
		}
	}

	return ToPascalCase(out)
}
