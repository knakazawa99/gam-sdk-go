package xml

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDeepMarshal(t *testing.T) {

	type Address struct {
		Street string `xml:"street"`
		City   string `xml:"city"`
	}
	type Village struct {
		Name      string     `xml:"name"`
		Age       *int       `xml:"age"`
		Addresses *[]Address `xml:"addresses"`
	}

	age := 10
	type args struct {
		v         interface{}
		ignoreNil bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success/no ignore nil",
			args: args{
				v: Village{
					Name: "John Doe",
					Age:  nil,
					Addresses: &[]Address{
						{
							Street: "123 Main St",
							City:   "Springfield",
						},
						{
							Street: "124 Main St",
							City:   "Springfield",
						},
					},
				},
			},
			want:    "<name>John Doe</name><age>0</age><addresses><street>123 Main St</street><city>Springfield</city></addresses><addresses><street>124 Main St</street><city>Springfield</city></addresses>",
			wantErr: false,
		},
		{
			name: "success/ignore nil",
			args: args{
				v: Village{
					Name:      "John Doe",
					Age:       &age,
					Addresses: nil,
				},
				ignoreNil: true,
			},
			want:    "<name>John Doe</name><age>10</age>",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeepMarshal(tt.args.v, tt.args.ignoreNil)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeepMarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(string(got), tt.want); diff != "" {
				t.Errorf("DeepMarshal() diff = %v got = %v, want %v", diff, string(got), string(tt.want))
			}
		})
	}
}
