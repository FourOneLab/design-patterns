package specification_refactoring

import "testing"

func TestText_ToNumber(t1 *testing.T) {
	type fields struct {
		content string
	}

	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{"number", fields{content: "123"}, 123, false},
		{"empty", fields{content: ""}, 0, true},
		{"one prefix empty", fields{content: " 123"}, 123, false},
		{"one suffix empty", fields{content: "123 "}, 123, false},
		{"one empty", fields{content: " 123 "}, 123, false},
		{"more empty", fields{content: "  123  "}, 123, false},
		{"with character", fields{content: "123a4"}, 0, true},
	}

	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Text{
				content: tt.fields.content,
			}
			got, err := t.ToNumber()
			if (err != nil) != tt.wantErr {
				t1.Errorf("ToNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("ToNumber() got = %v, want %v", got, tt.want)
			}
		})
	}
}
