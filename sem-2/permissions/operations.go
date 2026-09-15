//go:build !solution

package permissions

// Grant добавляет права added к current.
func Grant(current, added Permission) Permission {
	_, _ = current, added

	return 0
}

// Revoke удаляет права removed из current.
func Revoke(current, removed Permission) Permission {
	_, _ = current, removed

	return 0
}

// Has проверяет наличие всех прав required в current.
func Has(current, required Permission) bool {
	_, _ = current, required

	return false
}
