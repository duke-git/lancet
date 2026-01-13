package formatter

import (
	"encoding/json"
	"testing"
)

func TestParseCNAddress(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		withUser bool
		want     *AddressInfo
	}{
		{
			name:     "完整地址信息",
			input:    "张三 13800138000 北京市朝阳区建国路1号",
			withUser: true,
			want: &AddressInfo{
				Name:     "张三",
				Mobile:   "13800138000",
				Province: "北京",
				City:     "北京市",
				Region:   "朝阳区",
				Street:   "建国路1号",
			},
		},
		{
			name:     "带身份证和邮编",
			input:    "李四 18612345678 110101199001011234 100000 上海市浦东新区世纪大道100号",
			withUser: true,
			want: &AddressInfo{
				Name:     "李四",
				Mobile:   "18612345678",
				IDN:      "110101199001011234",
				Postcode: "100000",
			},
		},
		{
			name:     "仅地址不含用户信息",
			input:    "北京市海淀区中关村大街1号",
			withUser: false,
			want: &AddressInfo{
				Province: "北京",
				City:     "北京市",
				Region:   "海淀区",
				Street:   "中关村大街1号",
			},
		},
		{
			name:     "带收货关键词",
			input:    "收货人：王五 电话：13900139000 收货地址：天津市河西区友谊路20号",
			withUser: true,
			want: &AddressInfo{
				Name:     "王五",
				Mobile:   "13900139000",
				Province: "天津",
				City:     "天津市",
				Region:   "河西区",
			},
		},
		{
			name:     "紧凑格式地址",
			input:    "马云13593464918陕西省西安市雁塔区丈八沟街道高新四路南江国际",
			withUser: true,
			want: &AddressInfo{
				Name:     "马云",
				Mobile:   "13593464918",
				Province: "陕西省",
				City:     "西安市",
				Region:   "雁塔区",
				Street:   "丈八沟街道高新四路南江国际",
			},
		},
		{
			name:     "带座机号格式",
			input:    "姓名：马云\n联系电话：800-8585222\n所在地区：河北省石家庄市新华区\n详细地址:中华北大街68号鹿城商务中心6号楼1413室",
			withUser: true,
			want: &AddressInfo{
				Name:     "马云",
				Mobile:   "800-8585222",
				Province: "河北省",
				City:     "石家庄市",
				Region:   "新华区",
				Street:   "中华北大街68号鹿城商务中心6号楼1413室",
			},
		},
		{
			name:     "北京市重复格式",
			input:    "北京市北京市市辖区东城区",
			withUser: false,
			want: &AddressInfo{
				Province: "北京",
				City:     "北京市",
				Region:   "东城区",
				Street:   "",
			},
		},
		{
			name:     "河北省新乐市地址",
			input:    "河北省石家庄市新乐市经济开发区兴工街10号来优品仓库",
			withUser: false,
			want: &AddressInfo{
				Province: "河北省",
				City:     "石家庄市",
				Region:   "新乐市",
				Street:   "经济开发区兴工街10号来优品仓库",
			},
		},
		{
			name:     "江苏仪征市地址",
			input:    "江苏省扬州市仪征市真州镇解放东路99号",
			withUser: false,
			want: &AddressInfo{
				Province: "江苏省",
				City:     "扬州市",
				Region:   "仪征市",
				Street:   "真州镇解放东路99号",
			},
		},
		{
			name:     "新疆石河子市地址",
			input:    "新疆石河子市北三路25小区",
			withUser: false,
			want: &AddressInfo{
				Province: "新疆维吾尔自治区",
				City:     "自治区直辖县级市",
				Region:   "石河子市",
			},
		},
		{
			name:     "新疆石河子市-简化格式省+县级市",
			input:    "新疆维吾尔自治区石河子市",
			withUser: false,
			want: &AddressInfo{
				Province: "新疆维吾尔自治区",
				City:     "自治区直辖县级市",
				Region:   "石河子市",
				Street:   "",
			},
		},
		{
			name:     "新疆石河子市-完整行政区划表述",
			input:    "新疆维吾尔自治区自治区直辖县级市石河子市",
			withUser: false,
			want: &AddressInfo{
				Province: "新疆维吾尔自治区",
				City:     "自治区直辖县级市",
				Region:   "石河子市",
				Street:   "",
			},
		},
		{
			name:     "浙江杭州西湖区重复地址",
			input:    "浙江省杭州市西湖区杭州市西湖区人民政府109号",
			withUser: false,
			want: &AddressInfo{
				Province: "浙江省",
				City:     "杭州市",
				Region:   "西湖区",
				Street:   "人民政府109号",
			},
		},
		{
			name:     "湖南长沙市重复地址",
			input:    "湖南省长沙市岳麓区银盆岭街道长沙市人民政府长沙市政府大楼",
			withUser: false,
			want: &AddressInfo{
				Province: "湖南省",
				City:     "长沙市",
				Region:   "岳麓区",
				Street:   "银盆岭街道人民政府政府大楼",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseCNAddress(tt.input, tt.withUser)

			// 打印结果便于调试
			jsonData, _ := json.MarshalIndent(got, "", "  ")
			t.Logf("Result: %s", jsonData)

			// 验证主要字段
			if tt.want.Name != "" && got.Name != tt.want.Name {
				t.Errorf("Name = %v, want %v", got.Name, tt.want.Name)
			}
			if tt.want.Mobile != "" && got.Mobile != tt.want.Mobile {
				t.Errorf("Mobile = %v, want %v", got.Mobile, tt.want.Mobile)
			}
			if tt.want.Province != "" && got.Province != tt.want.Province {
				t.Errorf("Province = %v, want %v", got.Province, tt.want.Province)
			}
			if tt.want.City != "" && got.City != tt.want.City {
				t.Errorf("City = %v, want %v", got.City, tt.want.City)
			}
			if tt.want.Region != "" && got.Region != tt.want.Region {
				t.Errorf("Region = %v, want %v", got.Region, tt.want.Region)
			}
		})
	}
}

func TestParsePersonInfo(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		verify func(*testing.T, *AddressInfo)
	}{
		{
			name:  "提取姓名和手机号",
			input: "张三 13800138000 北京市朝阳区",
			verify: func(t *testing.T, got *AddressInfo) {
				if got.Name != "张三" {
					t.Errorf("Name = %v, want 张三", got.Name)
				}
				if got.Mobile != "13800138000" {
					t.Errorf("Mobile = %v, want 13800138000", got.Mobile)
				}
			},
		},
		{
			name:  "提取身份证号",
			input: "李四 110101199001011234 上海市",
			verify: func(t *testing.T, got *AddressInfo) {
				if got.Name != "李四" {
					t.Errorf("Name = %v, want 李四", got.Name)
				}
				if got.IDN != "110101199001011234" {
					t.Errorf("IDN = %v, want 110101199001011234", got.IDN)
				}
			},
		},
		{
			name:  "提取邮编",
			input: "王五 100000 天津市",
			verify: func(t *testing.T, got *AddressInfo) {
				if got.Name != "王五" {
					t.Errorf("Name = %v, want 王五", got.Name)
				}
				if got.Postcode != "100000" {
					t.Errorf("Postcode = %v, want 100000", got.Postcode)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParsePersonInfo(tt.input)
			jsonData, _ := json.MarshalIndent(got, "", "  ")
			t.Logf("Result: %s", jsonData)
			tt.verify(t, got)
		})
	}
}

func TestFuzz(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *fuzzyResult
	}{
		{
			name:  "包含区",
			input: "北京市朝阳区建国路1号",
			want: &fuzzyResult{
				A2:     "北京市",
				A3:     "朝阳区",
				Street: "建国路1号",
			},
		},
		{
			name:  "包含县",
			input: "河北省石家庄市正定县",
			want: &fuzzyResult{
				A2: "石家庄市",
				A3: "正定县",
			},
		},
		{
			name:  "复杂街道地址",
			input: "浙江省杭州市拱墅区武林街道杭州锦麟宾馆中河片区",
			want: &fuzzyResult{
				A2:     "杭州市",
				A3:     "拱墅区",
				Street: "武林街道杭州锦麟宾馆中河片区",
			},
		},
		{
			name:  "北京市重复格式",
			input: "北京市北京市市辖区东城区",
			want: &fuzzyResult{
				A2:     "北京市",
				A3:     "东城区",
				Street: "",
			},
		},
		{
			name:  "详细地址包含市字",
			input: "北京市朝阳区建外大街1号国贸商城",
			want: &fuzzyResult{
				A2:     "北京市",
				A3:     "朝阳区",
				Street: "建外大街1号国贸商城",
			},
		},
		{
			name:  "详细地址真的包含市字",
			input: "北京市朝阳区农贸市场路1号",
			want: &fuzzyResult{
				A2:     "北京市",
				A3:     "朝阳区",
				Street: "农贸市场路1号",
			},
		},
		{
			name:  "河北省新乐市地址",
			input: "河北省石家庄市新乐市经济开发区兴工街10号来优品仓库",
			want: &fuzzyResult{
				A2:     "石家庄市",
				A3:     "新乐市",
				Street: "经济开发区兴工街10号来优品仓库",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fuzz(tt.input)
			jsonData, _ := json.MarshalIndent(got, "", "  ")
			t.Logf("Result: %s", jsonData)

			if got.A2 != tt.want.A2 {
				t.Errorf("A2 = %v, want %v", got.A2, tt.want.A2)
			}
			if got.A3 != tt.want.A3 {
				t.Errorf("A3 = %v, want %v", got.A3, tt.want.A3)
			}
			if tt.want.Street != "" && got.Street != tt.want.Street {
				t.Errorf("Street = %v, want %v", got.Street, tt.want.Street)
			}
		})
	}
}

func ExampleParseCNAddress() {
	// 解析包含用户信息的完整地址
	result := ParseCNAddress("张三 13800138000 北京市朝阳区建国路1号", true)
	jsonData, _ := json.MarshalIndent(result, "", "  ")
	println(string(jsonData))
}

func ExampleParsePersonInfo() {
	// 分离用户信息
	result := ParsePersonInfo("收货人：李四 电话：18612345678 地址：上海市浦东新区世纪大道100号")
	jsonData, _ := json.MarshalIndent(result, "", "  ")
	println(string(jsonData))
}
