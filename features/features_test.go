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

	db, err = client.Database(ctx, "third")
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
		{
			name: "TestBalanceAddr for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
		{
			name: "TestBalanceAddr for address (not exist)",
			args: args{
				btcAddress: "",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := BalanceAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
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
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := TotalGetAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("TotalGetAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
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
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := FirstTimeAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("FirstTimeAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestLastTimeAddr(t *testing.T) {
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
			expected: 1514946668,
		},
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := LastTimeAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("LastTimeAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountOutTx(t *testing.T) {
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
			expected: 5,
		},
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountOutTx(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountOutTx error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountInTx(t *testing.T) {
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
			expected: 2,
		},
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountInTx(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountInTx error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountInAddr(t *testing.T) {
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
			expected: 1,
		},
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountInAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountInAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountOutAddr(t *testing.T) {
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
			expected: 2,
		},
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountOutAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountOutAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountSharedAddr(t *testing.T) {
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
			expected: 1,
		},
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountSharedAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountSharedAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestTotalCountAddr(t *testing.T) {
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
			expected: 2,
		},
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := TotalCountAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("TotalCountAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountUniqueAddr(t *testing.T) {
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
				btcAddress: "btcAddress/2",
			},
			wantErr:  false,
			expected: 1,
		},
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountUniqueAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountUniqueAddr error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestAverageCountOutAddr(t *testing.T) {
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
		expected float64
	}{
		// TODO: Add test cases.
		{
			name: "test for address btcAddress/1",
			args: args{
				btcAddress: "btcAddress/1",
			},
			wantErr:  false,
			expected: 1,
		},
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := AverageCountOutAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("AverageCountOutAddr error = %v, wantErr %v, received = %f, expected = %f", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestAverageCountInAddr(t *testing.T) {
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
		expected float64
	}{
		// TODO: Add test cases.
		{
			name: "test for address btcAddress/2",
			args: args{
				btcAddress: "btcAddress/2",
			},
			wantErr:  false,
			expected: 8.0 / 3.0,
		},
		{
			name: "test for address btcAddress/123 (not exist)",
			args: args{
				btcAddress: "btcAddress/123",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := AverageCountInAddr(ctx, db, tt.args.btcAddress); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("AverageCountInAddr error = %v, wantErr %v, received = %f, expected = %f", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestBalanceClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "TestBalanceAddr for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 2512407,
		},
		{
			name: "TestBalanceAddr for walletId ysadwaasdasdwq (not exist)",
			args: args{
				walletId: "ysadwaasdasdwq",
			},
			wantErr:  true,
			expected: 0,
		},
		{
			name: "TestBalanceAddr for walletId (not exist)",
			args: args{
				walletId: "",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := BalanceClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("BalanceClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestTotalGetClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 4407621,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := TotalGetClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("TotalGetClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestFirstTimeClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 1514946660,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := FirstTimeClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("FirstTimeClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestLastTimeClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 1514996668,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := LastTimeClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("LastTimeClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountOutTxClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 4,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountOutTxClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountOutTxClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountInTxClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 2,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountInTxClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountInTxClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountInClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 4,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountInClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountInClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountOutClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 3,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountOutClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountOutClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountSharedClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 3,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountSharedClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountSharedClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestTotalCountClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 4,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := TotalCountClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("TotalCountClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestCountUniqueClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 1,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := CountUniqueClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("CountUniqueClust error = %v, wantErr %v, received = %d, expected = %d", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestAverageCountOutClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
		expected float64
	}{
		// TODO: Add test cases.
		{
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 1.0,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := AverageCountOutClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("AverageCountOutClust error = %v, wantErr %v, received = %f, expected = %f", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestAverageCountInClust(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId string
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
		expected float64
	}{
		// TODO: Add test cases.
		{
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			},
			wantErr:  false,
			expected: 3.0,
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId: "Daslw12eascaCaawWAsadlasd",
			},
			wantErr:  true,
			expected: 0,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := AverageCountInClust(ctx, db, tt.args.walletId); (err != nil) != tt.wantErr || val != tt.expected {
				t.Errorf("AverageCountInClust error = %v, wantErr %v, received = %f, expected = %f", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}

func TestNmotif(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		walletId1 string
		walletId2 string
		n         int
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
		expected [][]interface{}
	}{
		// TODO: Add test cases.
		{
			name: "test for walletId PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
			args: args{
				walletId2: "PlVuhq2enKKdVpdrXI3IC6SfaNlPwS5F6HlRErl8FPM1YiyYfszMPRrXC7KDpPeX",
				walletId1: "dITWeUoEbaxbmiVXpM1TbmFlmXJP2ZEe4QR7RqAL7M8BcMrWwiq2jkgsVwBCW5Ot",
				n:         1,
			},
			wantErr: false,
			expected: [][]interface{}{
				{"1", "btcTx/VjeSsLWUwp6kejVz3hWr2Kau8BnMYmBfgm25ceK3CtZcT2q63D4nUzJZ8h5KkUUZ", "3"},
				{"1", "btcTx/VjeSsLWUwp6kejVz3hWr2Kau8BnMYmBfgm25ceK3CtZcT2q63D4nUzJZ8h5KkUUZ", "4"},
				{"2", "btcTx/qtNMLgUxe7vitlc4rkzirW3PYKMNYM32Y6Kk4dQEvItEIKUX7P0wxmlTDtJ36vR8", "3"},
				{"2", "btcTx/MmfJzrWA38XWSqA0NfNgF9Do5yJSA7P0fqR30LFUPuReMDqxqu97vplOyad4Wg8Q", "4"},
			},
		},
		{
			name: "test for walletId Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				walletId1: "Daslw12eascaCaawWAsadlasd",
				walletId2: "dITWeUoEbaxbmiVXpM1TbmFlmXJP2ZEe4QR7RqAL7M8BcMrWwiq2jkgsVwBCW5Ot",
				n:         2,
			},
			wantErr:  true,
			expected: nil,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if val, err := Nmotif(ctx, db, tt.args.walletId1, tt.args.walletId2, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Nmotif error = %v, wantErr %v, received = %s, expected = %s", err, tt.wantErr, val, tt.expected)
			}
		})
	}
}
