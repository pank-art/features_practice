package features

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"log"
	"testing"
)

var (
	ctx context.Context
	db  driver.Database
)

func init() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		log.Fatal(err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "artem"),
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx = context.Background()

	db, err = client.Database(ctx, "second")
	if err != nil {
		log.Fatal(err)
	}
}

func TestBalanceAddr(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		btcAddress string
	}
	// test cases
	tests := []struct {
		// name of test case
		name string
		// arguments for function
		args args
		// if function shoud return an error
		wantErr bool
		//
		expected int64
	}{
		// TODO: Add test cases.
		{
			name: "TestBalanceAddr for address btcAddress/1",
			args: args{
				btcAddress: "btcAddress/1",
			},
			wantErr:  false,
			expected: 4494681,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := BalanceAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr && val == tt.expected {
				t.Errorf("BalanceAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestTotalGetAddr(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		btcAddress string
	}
	// test cases
	tests := []struct {
		// name of test case
		name string
		// arguments for function
		args args
		// if function shoud return an error
		wantErr bool
		//
		expected int64
	}{
		// TODO: Add test cases.
		{
			name: "test for address btcAddress/1",
			args: args{
				btcAddress: "btcAddress/1",
			},
			wantErr:  false,
			expected: 5189895,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := TotalGetAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr && val == tt.expected {
				t.Errorf("BalanceAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestFirstTimeAddr(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		btcAddress string
	}
	// test cases
	tests := []struct {
		// name of test case
		name string
		// arguments for function
		args args
		// if function shoud return an error
		wantErr bool
		//
		expected int64
	}{
		// TODO: Add test cases.
		{
			name: "test for address btcAddress/1",
			args: args{
				btcAddress: "btcAddress/1",
			},
			wantErr:  false,
			expected: 1514046608,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := FirstTimeAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr && val == tt.expected {
				t.Errorf("BalanceAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}
