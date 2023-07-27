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
	// В этой бд все коллекции из data. btcAddress и btcTx - документы, остальные ребра. Также создан граф "graphNext", где вершины из btcTx, а ребра btcNext
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

func TestNmotifAddr(t *testing.T) {
	// arguments of function you want to test
	type args struct {
		addr1 string
		addr2 string
		n     int
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
		expected [][]string
	}{
		// TODO: Add test cases.
		{
			name: "test for addr1 btcAddress/1, addr2 btcAddress/2",
			args: args{
				addr1: "btcAddress/1",
				addr2: "btcAddress/2",
				n:     1,
			},
			wantErr: false,
			expected: [][]string{
				{"1", "btcTx/0b879ca09946fd4f30ea838e322727e0e7b48b828a8eddff1ea2f908af7c7a9f", "2"},
				{"1", "btcTx/FBxWSvAIjhi8lAUo8gFB4ld1DYH7qkLUNmZXiOgaB4uUFE3Nyn1BdQ1CLCw1Wcgw", "2"},
				{"1", "btcTx/7LDC3SyFlJLQp0Wbhc5sC4Sv9FzxtPd8iBaSohzM9RFpXRe8sDV5mKDoc5aGVNcz", "2"},
			},
		},
		{
			name: "test for addr Daslw12eascaCaawWAsadlasd (not exist)",
			args: args{
				addr1: "Daslw12eascaCaawWAsadlasd",
				addr2: "dITWeUoEbaxbmiVXpM1TbmFlmXJP2ZEe4QR7RqAL7M8BcMrWwiq2jkgsVwBCW5Ot",
				n:     2,
			},
			wantErr:  true,
			expected: nil,
		},
		{
			name: "test for addr btcAddress/21324 (not exist)",
			args: args{
				addr1: "btcAddress/1",
				addr2: "btcAddress/21324",
				n:     1,
			},
			wantErr:  true,
			expected: nil,
		},
		{
			name: "test for addr1 btcAddress/4, addr2 btcAddress/3 (0 path)",
			args: args{
				addr1: "btcAddress/4",
				addr2: "btcAddress/3",
				n:     2,
			},
			wantErr:  false,
			expected: nil,
		},
	}
	// iterate all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if path, err := NmotifAddr(ctx, db, tt.args.addr1, tt.args.addr2, tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("NmotifAddr error = %v, wantErr %v", err, tt.wantErr)
			} else {
				for i, arr := range path {
					for j, val := range arr {
						if val != tt.expected[i][j] {
							t.Errorf("NmotifAddr error = %v, wantErr %v, received = %s, expected = %s", err, tt.wantErr, path, tt.expected)
						}
					}
				}
			}
		})
	}
}
