package catalog

import (
	"testing"
)

func TestValidateSlug(t *testing.T) {
	tests := []struct {
		name    string
		slug    string
		wantErr bool
	}{
		{"valid lowercase", "my-service", false},
		{"valid with numbers", "api-v2", false},
		{"valid single word", "database", false},
		{"valid with multiple hyphens", "my-super-cool-service-1", false},
		{"empty slug", "", true},
		{"uppercase letters", "My-Service", true},
		{"spaces", "my service", true},
		{"underscore", "my_service", true},
		{"leading hyphen", "-myservice", true},
		{"trailing hyphen", "myservice-", true},
		{"double hyphen", "my--service", true},
		{"special characters", "my-service!", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateSlug(tt.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateSlug() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
