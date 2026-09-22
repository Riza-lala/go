//go:build !solution

package dedup

// Deduplicate удаляет повторы, сохраняя порядок первых появлений.
func Deduplicate(input []string) []string {
	if input == nil {
		return nil
	}
	seen := make(map[string]bool)
	result := []string{}
	for _, s := range input {
		if seen[s] {
			continue
		}
		seen[s] = true
		result = append(result, s)
	}

	return result
}
