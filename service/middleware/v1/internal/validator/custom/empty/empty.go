package empty

func Validate(m map[string]string) bool {
	for k, v := range m {
		switch k {
		case "paginationPage", "paginationCount":
			continue
		case "level", "message", "resourceId",
			"timestampStart", "timestampEnd", "traceId",
			"spanId", "commit", "metadateParentResourceId":
			if v == "" {
				return true
			}
		}
	}
	return false
}
