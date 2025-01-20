package utils

func MergeKeys(maps ...map[string]float64) map[string]struct{} {
	merged := make(map[string]struct{})
	for _, m := range maps {
		for key := range m {
			merged[key] = struct{}{}
		}
	}
	return merged
}
