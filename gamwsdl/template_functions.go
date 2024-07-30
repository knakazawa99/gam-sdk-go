package gamwsdl

import (
	"strings"
	"unicode"
)

func generatePackageName(serviceName string) string {
	return strings.ReplaceAll(strings.ToLower(serviceName), "service", "")
}

var xsdTypeMap = map[string]string{
	"string":       "string",
	"int":          "int",
	"double":       "float64",
	"long":         "int64",
	"boolean":      "bool",
	"base64Binary": "[]byte",
}

func getGoBuiltinType(elementType string) string {
	goType := xsdTypeMap[elementType]
	if goType == "" {
		return removeDotFromString(elementType)
	}
	return goType
}

func getOperationTypeFromOperationName(operationName string) string {
	if strings.HasPrefix(operationName, "perform") {
		return "Perform"
	}

	return "Default"
}

func getTypeFromElementType(elementType string) string {
	typeKind, typeStr := strings.Split(elementType, ":")[0], strings.Split(elementType, ":")[1]
	if typeKind == "xsd" || typeKind == "tns" {
		return typeStr
	} else {
		return ""
	}
}

func getTypeKind(typeName string) string {
	if !strings.Contains(typeName, ".") {
		return "Enum"
	}

	typeHint := strings.Split(typeName, ".")[1]
	if typeHint == "Reason" {
		return "Error"
	} else {
		return "Type"
	}
}

func getTypeName(typeName string) string {
	return strings.Split(typeName, ".")[0]
}

func isElementTypeError(elementType string) bool {
	typeNames := strings.Split(elementType, ".")
	if len(typeNames) < 2 {
		return false
	}

	if typeNames[1] == "Reason" {
		return true
	}

	return false
}

func makeComment(text string) string {
	lines := strings.Split(text, "\n")

	var output string
	if len(lines) == 1 && lines[0] == "" {
		return ""
	}

	// Helps to determine if there is an actual comment without screwing newlines
	// in real comments.
	hasComment := false

	for _, line := range lines {
		line = strings.TrimLeftFunc(line, unicode.IsSpace)
		if line != "" {
			hasComment = true
		}
		output += "\n// " + line
	}

	if hasComment {
		return output
	}
	return ""
}

func removeDotFromString(str string) string {
	return strings.ReplaceAll(str, ".", "")
}

// toUpperCamelFromCamelCase converts a string from camelCase to UpperCamelCase.
func toUpperCamelFromCamelCase(camelCase string) string {
	for i, v := range camelCase {
		return string(unicode.ToUpper(v)) + camelCase[i+1:]
	}
	return ""
}

func toCamelCaseFromUpperCamelCase(upperCamelCase string) string {
	if upperCamelCase == "" {
		return ""
	}
	return strings.ToLower(upperCamelCase[:1]) + upperCamelCase[1:]
}

// toUpperCamelFromUpperSnakeCase converts a string from UPPER_SNAKE_CASE to UpperCamelCase.
func toUpperCamelFromUpperSnakeCase(s string) string {
	words := strings.Split(s, "_")
	for i := range words {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}
	return strings.Join(words, "")
}
