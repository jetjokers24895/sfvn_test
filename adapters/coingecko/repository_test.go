package adapter

import (
	"testing"
	"time"

	"github.com/redis/rueidis"
	"gorm.io/gorm"
)

func Test_repository_IsLeakData(t *testing.T) {
	type fields struct {
		db        *gorm.DB
		redis     rueidis.Client
		domainApi string
		apiKey    string
	}
	type args struct {
		period       string
		startDate    int64
		startDateRes int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "Leak Data",
			args: args{
				period:       "30m",
				startDate:    time.Now().UnixMilli(),
				startDateRes: time.Now().UnixMilli() + 3600*1000,
			},
			fields: fields{
				db:        nil,
				redis:     nil,
				domainApi: "",
				apiKey:    "",
			},
			want: true,
		},

		{
			name: "Not Leak Data",
			args: args{
				period:       "30m",
				startDate:    time.Now().UnixMilli(),
				startDateRes: time.Now().UnixMilli() + 20*60*1000,
			},
			fields: fields{
				db:        nil,
				redis:     nil,
				domainApi: "",
				apiKey:    "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repository{
				db:        tt.fields.db,
				redis:     tt.fields.redis,
				domainApi: tt.fields.domainApi,
				apiKey:    tt.fields.apiKey,
			}
			if got := r.IsLeakData(tt.args.period, tt.args.startDate, tt.args.startDateRes); got != tt.want {
				t.Errorf("repository.IsLeakData() = %v, want %v", got, tt.want)
			}
		})
	}
}
