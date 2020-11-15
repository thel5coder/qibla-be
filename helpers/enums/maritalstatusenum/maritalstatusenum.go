package maritalstatusenum

var (
	maritalStatusEnums = []map[string]interface{}{
		{
			"key":  "single",
			"text": "Belum Menikah",
		},
		{
			"key":  "married",
			"text": "Menikah",
		},
		{
			"key":  "divorced",
			"text": "Cerai Hidup",
		},
		{
			"key":  "death-divorced",
			"text": "Cerai Mati",
		},
	}
)

// GetKey ...
func GetKey(index int) string {
	value := maritalStatusEnums[index]

	return value["key"].(string)
}

// GetText ...
func GetText(index int) string {
	value := maritalStatusEnums[index]

	return value["text"].(string)
}

// GetEnums ...
func GetEnums() []map[string]interface{} {
	return maritalStatusEnums
}

// GetEnum ...
func GetEnum(index int) map[string]interface{} {
	return maritalStatusEnums[index]
}
