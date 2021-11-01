package sourceutils

import (
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

func TestGetAllPages(t *testing.T) {
	type args struct {
		getOnePage GetOnePage
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success without pages",
			args: args{
				getOnePage: func(page int) ([]string, error) { return nil, nil },
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "success with pages",
			args: args{
				getOnePage: func() GetOnePage {
					pages := [][]string{
						{"one", "two"},
						{"three", "four"},
						{"five", "six"},
						{"seven", "eight"},
						{"nine", "ten"},
					}
					return func(page int) ([]string, error) {
						pageIndex := page - 1
						if pageIndex > len(pages)-1 {
							return nil, nil
						}
						return pages[pageIndex], nil
					}
				}(),
			},
			want: []string{
				"one", "two",
				"three", "four",
				"five", "six",
				"seven", "eight",
				"nine", "ten",
			},
			wantErr: assert.NoError,
		},
		{
			name: "error",
			args: args{
				getOnePage: func() GetOnePage {
					pages := [][]string{
						{"one", "two"},
						{"three", "four"},
						{"five", "six"},
						{"seven", "eight"},
						{"nine", "ten"},
					}
					return func(page int) ([]string, error) {
						pageIndex := page - 1
						if pageIndex > len(pages)/2 {
							return nil, iotest.ErrTimeout
						}
						return pages[pageIndex], nil
					}
				}(),
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllPages(tt.args.getOnePage)

			assert.Equal(t, tt.want, got)
			tt.wantErr(t, err)
		})
	}
}
