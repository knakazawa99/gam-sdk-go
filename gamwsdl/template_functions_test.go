package gamwsdl

import "testing"

func Test_generatePackageName(t *testing.T) {
	type args struct {
		serviceName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success/return lineitem",
			args: args{
				serviceName: "LineItemService",
			},
			want: "lineitem",
		},
		{
			name: "success/return company",
			args: args{
				serviceName: "CompanyService",
			},
			want: "company",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePackageName(tt.args.serviceName); got != tt.want {
				t.Errorf("generatePackageName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGoBuiltinType(t *testing.T) {
	type args struct {
		elementType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success/return string type",
			args: args{
				elementType: "string",
			},
			want: "string",
		},
		{
			name: "success/return long type",
			args: args{
				elementType: "long",
			},
			want: "int64",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGoBuiltinType(tt.args.elementType); got != tt.want {
				t.Errorf("getGoBuiltinType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTypeFromElementType(t *testing.T) {
	type args struct {
		elementType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success/return string type",
			args: args{
				elementType: "xsd:string",
			},
			want: "string",
		},
		{
			name: "success/return LineItem",
			args: args{
				elementType: "tns:LineItem",
			},
			want: "LineItem",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTypeFromElementType(tt.args.elementType); got != tt.want {
				t.Errorf("getTypeFromElementType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTypeKind(t *testing.T) {
	type args struct {
		typeName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success/return Enum",
			args: args{
				typeName: "LineItem",
			},
			want: "Enum",
		},
		{
			name: "success/return Reason",
			args: args{
				typeName: "NotNullError.Reason",
			},
			want: "Error",
		},
		{
			name: "success/return Type",
			args: args{
				typeName: "LineItem.Type",
			},
			want: "Type",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTypeKind(tt.args.typeName); got != tt.want {
				t.Errorf("getTypeKind() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTypeName(t *testing.T) {
	type args struct {
		typeName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success/return Company",
			args: args{
				typeName: ":Company.Type",
			},
			want: ":Company",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTypeName(tt.args.typeName); got != tt.want {
				t.Errorf("getTypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isElementTypeError(t *testing.T) {
	type args struct {
		elementType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success/return true",
			args: args{
				elementType: "NotNullError.Reason",
			},
			want: true,
		},
		{
			name: "success/return false",
			args: args{
				elementType: "LineItem.Type",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isElementTypeError(tt.args.elementType); got != tt.want {
				t.Errorf("isElementTypeError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeComment(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success/return comment",
			args: args{
				text: "This is a comment",
			},
			want: "\n// This is a comment",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeComment(tt.args.text); got != tt.want {
				t.Errorf("makeComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toUpperCamelFromCamelCase(t *testing.T) {
	type args struct {
		camelCase string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success/return UpperCamelCase",
			args: args{
				camelCase: "lineItem",
			},
			want: "LineItem",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toUpperCamelFromCamelCase(tt.args.camelCase); got != tt.want {
				t.Errorf("toUpperCamelFromCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toUpperCamelFromUpperSnakeCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success/return UpperCamelCase",
			args: args{
				s: "PENDING_GOOGLE_APPROVAL",
			},
			want: "PendingGoogleApproval",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toUpperCamelFromUpperSnakeCase(tt.args.s); got != tt.want {
				t.Errorf("toUpperCamelFromUpperSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
