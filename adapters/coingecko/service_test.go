package adapter

import (
	"reflect"
	"testing"
	"time"

	DTOS "sfvn_test/adapters/coingecko/dtos"
	_mock "sfvn_test/adapters/coingecko/mocks"
)

func Test_coingecko_GetHOLC(t *testing.T) {
	type fields struct {
		repo Repository
	}
	type args struct {
		symbol    string
		period    string
		startDate int64
		endDate   int64
	}

	mockRepo := _mock.NewRepository(t)
	startDate := time.Now().UnixMilli() - 3600*24*30*1000
	endDate := time.Now().UnixMilli()
	mockRepo.EXPECT().GetHOLCData("bitcoin", "4h", startDate, endDate).Return(nil, nil)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*DTOS.DTOGetHOLCResponse
		wantErr bool
	}{
		{
			name: "simple case",
			fields: fields{
				repo: mockRepo,
			},
			args: args{
				symbol:    "bitcoin",
				period:    "4h",
				startDate: startDate,
				endDate:   endDate,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &coingecko{
				repo: tt.fields.repo,
			}
			got, err := c.GetHOLC(tt.args.symbol, tt.args.period, tt.args.startDate, tt.args.endDate)
			if (err != nil) != tt.wantErr {
				t.Errorf("coingecko.GetHOLC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("coingecko.GetHOLC() = %v, want %v", got, tt.want)
			}
		})
	}
}
