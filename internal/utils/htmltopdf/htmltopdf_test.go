package htmltopdf

import (
	"os"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	type args struct {
		url     string
		sel     string
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		// .vben-sign-off-content-done
		{
			name: "test1",
			args: args{
				url:     "http://localhost:5173/sign-off/overtime?uuid=5ce4c012-7fd0-47e0-bf39-4a105053c5d9",
				sel:     ".vben-overtime-sign-off-content-done",
				timeout: 1 * time.Minute,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Create(tt.args.url, tt.args.sel, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				err = os.WriteFile("example.pdf", got, 0644)
				if err != nil {
					t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}
