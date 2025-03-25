package jsonUtils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNumberedJSON_GetString(t *testing.T) {
	tests := []struct {
		name    string
		nj      *numberedJSON
		want    string
		wantErr bool
	}{
		{
			name: "Basic test",
			nj: &numberedJSON{
				items:      &[]string{"Item 1", "Item 2"},
				key:        stringPtr("Key"),
				generalKey: stringPtr("General"),
			},
			want:    `{"General":{"Key1":"Item 1","Key2":"Item 2"}}`,
			wantErr: false,
		},
		{
			name: "Empty items",
			nj: &numberedJSON{
				items:      &[]string{},
				key:        stringPtr("Key"),
				generalKey: stringPtr("General"),
			},
			want:    `{"General":{}}`,
			wantErr: false,
		},
		{
			name: "Error case - nil Items",
			nj: &numberedJSON{
				items:      nil,
				key:        stringPtr("Key"),
				generalKey: stringPtr("General"),
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.nj.GetString()
			if (err != nil) != tt.wantErr {
				t.Errorf("numberedJSON.GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("numberedJSON.GetString() = %v, want %v", got, tt.want)
			} else {
				fmt.Printf("Successfully. numberedJSON.GetString() = %v\n", got)
			}
		})
	}
}

func TestNumberedJSON_GetJSONObject(t *testing.T) {
	tests := []struct {
		name string
		nj   *numberedJSON
		want *map[string]map[string]string
	}{
		{
			name: "Basic test",
			nj: &numberedJSON{
				items:      &[]string{"Item 1", "Item 2"},
				key:        stringPtr("Key"),
				generalKey: stringPtr("General"),
			},
			want: &map[string]map[string]string{
				"General": {
					"Key1": "Item 1",
					"Key2": "Item 2",
				},
			},
		},
		{
			name: "Empty items",
			nj: &numberedJSON{
				items:      &[]string{},
				key:        stringPtr("Key"),
				generalKey: stringPtr("General"),
			},
			want: &map[string]map[string]string{
				"General": {},
			},
		},
		{
			name: "Nil items",
			nj: &numberedJSON{
				items:      nil,
				key:        stringPtr("Key"),
				generalKey: stringPtr("General"),
			},
			want: &map[string]map[string]string{
				"General": {},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.nj.GetMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberedJSON.GetJSONObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNumberedJSON(t *testing.T) {
	items := []string{"Item 1", "Item 2"}
	key := "Key"
	generalKey := "General"

	got := CreateNumberedJSON(&items, key, generalKey)

	if !reflect.DeepEqual(*got.items, items) {
		t.Errorf("CreateNumberedJSON().Items = %v, want %v", *got.items, items)
	}
	if *got.key != key {
		t.Errorf("CreateNumberedJSON().Key = %v, want %v", *got.key, key)
	}
	if *got.generalKey != generalKey {
		t.Errorf("CreateNumberedJSON().GeneralKey = %v, want %v", *got.generalKey, generalKey)
	}
}

func stringPtr(s string) *string {
	return &s
}
