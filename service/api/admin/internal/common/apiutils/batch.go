package apiutils

func ExtractFields[T any](list []T, extractor func(T) string) []string {
	seen := make(map[string]struct{})
	result := make([]string, 0, len(list))

	for _, item := range list {
		value := extractor(item)
		if value == "" {
			continue
		}

		if _, ok := seen[value]; ok {
			continue
		}

		seen[value] = struct{}{}
		result = append(result, value)
	}

	return result
}

func BatchQuery[T any, R any](list []T, extractor func(T) string, query func([]string) (R, error)) (R, error) {
	return query(ExtractFields(list, extractor))
}
