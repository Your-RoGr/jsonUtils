package jsonUtils

import (
	"reflect"
	"testing"
)

func TestFindAndParseJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *[]map[string]interface{}
		wantErr bool
	}{
		{
			name:  "Single valid JSON",
			input: "This is some text {\"key\": \"value\"} and more text",
			want: &[]map[string]interface{}{
				{"key": "value"},
			},
			wantErr: false,
		},
		{
			name:  "Multiple valid JSONs",
			input: "First JSON: {\"id\": 1} Second JSON: {\"name\": \"John\"}",
			want: &[]map[string]interface{}{
				{"id": float64(1)},
				{"name": "John"},
			},
			wantErr: false,
		},
		{
			name:    "No JSON in text",
			input:   "This text contains no JSON objects",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Invalid JSON",
			input:   "This is an invalid JSON: {\"key\": \"value\",}",
			want:    nil,
			wantErr: true,
		},
		{
			name:  "Nested JSON",
			input: "Nested JSON: {\"outer\": {\"inner\": \"value\"}}",
			want: &[]map[string]interface{}{
				{"outer": map[string]interface{}{"inner": "value"}},
			},
			wantErr: false,
		},
		{
			name:  "JSON with array",
			input: "JSON with array: {\"items\": [1, 2, 3]}",
			want: &[]map[string]interface{}{
				{"items": []interface{}{float64(1), float64(2), float64(3)}},
			},
			wantErr: false,
		},
		{
			name:  "Valid and invalid JSON",
			input: "Valid: {\"valid\": true} Invalid: {\"invalid\": false,}",
			want: &[]map[string]interface{}{
				{"valid": true},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindAndParseJSON(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAndParseJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAndParseJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
