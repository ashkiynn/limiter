package limiter

import (
	"fmt"
	"testing"
)

func TestBaseAddLimiter(t *testing.T) {
	type args struct {
		name        string
		limiterType int8
		duration    uint32
		count       uint32
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试1",
			args: args{
				name:        "固定窗口1",
				limiterType: Fixed,
				duration:    20,
				count:       100,
			},
		}, {
			name: "测试2",
			args: args{
				name:        "固定窗口2",
				limiterType: Fixed,
				duration:    20,
				count:       100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BaseAddLimiter(tt.args.name, tt.args.limiterType, tt.args.duration, tt.args.count); (err != nil) != tt.wantErr {
				t.Errorf("BaseAddLimiter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBaseAllow(t *testing.T) {
	type args struct {
		name string
		id   int32
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BaseAllow(tt.args.name, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseAllow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("BaseAllow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestF(t *testing.T) {
	name := "固定窗口"
	err := BaseAddLimiter("", Fixed, 86400, 100000)
	// err = BaseAddLimiter(name, Fixed, 10, 20)
	if err != nil {
		fmt.Printf("BaseAddLimiter() error = %v", err)
		return
	}
	fmt.Println(1111)

	testloop(name)
	fmt.Println("--------------")

	// time.Sleep(time.Second * 5)/
	// testloop(name)
	// fmt.Println("--------------")
	//
	// time.Sleep(time.Second * 2)
	// testloop(name)
	// fmt.Println("--------------")

}

func testloop(name string) {
	for i := 0; i < 10000000; i++ {
		// go func() {
		_, err := BaseAllow(name, 1)
		if err != nil {
			panic(err)
		}
		// fmt.Println(allow, err)
		// }()
	}
}

func TestFNum(t *testing.T) {
	name1 := "固定窗口1"
	err := BaseAddLimiter(name1, Fixed, 10, 20)
	if err != nil {
		fmt.Printf("BaseAddLimiter() error = %v", err)
		return
	}

	name2 := "固定窗口2"
	err = BaseAddLimiter(name2, Fixed, 10, 20)
	if err != nil {
		fmt.Printf("BaseAddLimiter() error = %v", err)
		return
	}

	fmt.Println(BaseAllowNum(name1, 1, 3))
	fmt.Println(BaseAllowNum(name1, 1, 3))
	fmt.Println(BaseAllowNum(name1, 1, 4))
	fmt.Println(BaseAllowNum(name1, 1, 3))
	fmt.Println(BaseAllowNum(name1, 1, 3))
	fmt.Println(BaseAllowNum(name1, 1, 3))
	fmt.Println(BaseAllowNum(name1, 1, 3))
	fmt.Println(BaseAllowNum(name1, 1, 3))
	fmt.Println(BaseAllowNum(name1, 1, 3))

	fmt.Println(BaseAllowNum(name2, 1, 3))
	fmt.Println(BaseAllowNum(name2, 1, 3))
	fmt.Println(BaseAllowNum(name2, 1, 4))
	fmt.Println(BaseAllowNum(name2, 1, 3))
	fmt.Println(BaseAllowNum(name2, 1, 3))
	fmt.Println(BaseAllowNum(name2, 1, 3))
	fmt.Println(BaseAllowNum(name2, 1, 3))
	fmt.Println(BaseAllowNum(name2, 1, 3))
	fmt.Println(BaseAllowNum(name2, 1, 3))
}
