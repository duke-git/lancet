// Package formatter implements some functions to format string, struct.
package formatter

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// AddressInfo represents the parsed address information including user details and location.
// AddressInfo 表示解析后的地址信息，包括用户详细信息和位置信息
type AddressInfo struct {
	Name     string `json:"name"`     // Name of the recipient / 姓名
	Mobile   string `json:"mobile"`   // Mobile phone number or landline / 手机号或座机
	IDN      string `json:"idn"`      // ID card number / 身份证号
	Postcode string `json:"postcode"` // Postal code / 邮编
	Province string `json:"province"` // Province / 省
	City     string `json:"city"`     // City / 市
	Region   string `json:"region"`   // District or county / 区/县
	Street   string `json:"street"`   // Street address / 街道详细地址
	Addr     string `json:"addr"`     // Original address string / 原始地址字符串
}

// fuzzyResult represents the result of fuzzy address parsing.
// fuzzyResult 表示模糊地址解析的结果
type fuzzyResult struct {
	A1     string // Province level / 省级
	A2     string // City level / 市级
	A3     string // District/County level / 区/县级
	Street string // Street address / 街道地址
}

// ParseCNAddress parses a Chinese address string intelligently and extracts structured information.
// It can parse addresses with or without user information (name, phone, ID card, etc.).
// When withUser is true, it extracts user information from the address string.
// When withUser is false, it only parses the location information.
// The function handles various address formats including:
// - Standard format: "Province City District Street"
// - Compact format: "Name Phone Province City District Street"
// - With keywords: "Name: xxx Phone: xxx Address: xxx"
// - County-level cities: "Province City CountyCity District" (e.g., "河北省石家庄市新乐市")
// ParseCNAddress 智能解析中国地址字符串并提取结构化信息。
// 可以解析带或不带用户信息（姓名、电话、身份证等）的地址。
// 当 withUser 为 true 时，从地址字符串中提取用户信息。
// 当 withUser 为 false 时，仅解析位置信息。
// 该函数处理多种地址格式，包括：
// - 标准格式："省 市 区 街道"
// - 紧凑格式："姓名 电话 省 市 区 街道"
// - 带关键词："姓名:xxx 电话:xxx 地址:xxx"
// - 县级市："省 市 县级市 区"（如"河北省石家庄市新乐市"）
func ParseCNAddress(str string, withUser bool) *AddressInfo {
	result := &AddressInfo{}

	if withUser {
		ParsePersonInfo := ParsePersonInfo(str)
		result = ParsePersonInfo
	} else {
		result.Addr = str
	}

	fuzz := fuzz(result.Addr)
	parse := parse(fuzz.A1, fuzz.A2, fuzz.A3)

	result.Province = parse.Province
	result.City = parse.City
	result.Region = parse.Region

	// 提取街道地址：从原始地址中找到区/县的位置，提取后面的内容
	if result.Region != "" && result.Addr != "" {
		// 在原始地址中查找区/县的位置（转换为rune数组以正确处理中文）
		addrRunes := []rune(result.Addr)
		regionRunes := []rune(result.Region)
		regionPos := mbStrpos(result.Addr, result.Region)

		if regionPos != -1 {
			// 提取区/县后面的内容作为街道地址
			streetStart := regionPos + len(regionRunes)
			if streetStart < len(addrRunes) {
				result.Street = string(addrRunes[streetStart:])
			}
		} else if fuzz.Street != "" {
			// 如果没找到区/县，使用fuzz返回的街道
			result.Street = fuzz.Street
		}
	} else if fuzz.Street != "" {
		result.Street = fuzz.Street
	}

	// 清理街道地址中的重复省市区信息（可能存在部分匹配的残留）
	result.Street = strings.ReplaceAll(result.Street, result.Region, "")
	result.Street = strings.ReplaceAll(result.Street, result.City, "")
	result.Street = strings.ReplaceAll(result.Street, result.Province, "")
	// 清理街道地址中的残留片段（如"自治区直辖县级市"被替换后的残留）
	result.Street = strings.ReplaceAll(result.Street, "自治区直辖县级市", "")
	result.Street = strings.ReplaceAll(result.Street, "直辖县级市", "")
	result.Street = strings.TrimSpace(result.Street)

	return result
}

// ParsePersonInfo extracts user information (name, phone, ID card, postal code) from an address string.
// It separates personal information from the address, supporting various formats:
// - Labeled format: "Name: xxx Phone: xxx Address: xxx"
// - Compact format: "Name Phone Address" (e.g., "张三13800138000北京市朝阳区")
// - With separators: using colons, commas, newlines as delimiters
// Returns an AddressInfo with extracted user information and cleaned address string in Addr field.
// ParsePersonInfo 从地址字符串中提取用户信息（姓名、电话、身份证、邮编）。
// 将个人信息与地址分离，支持多种格式：
// - 带标签格式："姓名:xxx 电话:xxx 地址:xxx"
// - 紧凑格式："姓名 电话 地址"（如"张三13800138000北京市朝阳区"）
// - 带分隔符：使用冒号、逗号、换行符作为分隔符
// 返回包含提取的用户信息和清理后地址字符串（在 Addr 字段中）的 AddressInfo。
func ParsePersonInfo(str string) *AddressInfo {
	compose := &AddressInfo{}

	// 先尝试提取带标签的信息
	// 提取姓名 (支持: 姓名:xxx, 收货人:xxx, 收件人:xxx)
	nameRe := regexp.MustCompile(`(?:姓名|收货人|收件人)[：:]\s*([^\s\d\n]+)`)
	if match := nameRe.FindStringSubmatch(str); len(match) > 1 {
		compose.Name = strings.TrimSpace(match[1])
		str = nameRe.ReplaceAllString(str, " ")
	}

	// 提取手机号或座机号 (支持: 电话:xxx, 手机:xxx, 联系电话:xxx)
	phoneRe := regexp.MustCompile(`(?:电话|手机号码|手机|联系电话)[：:]\s*([\d\-]+)`)
	if match := phoneRe.FindStringSubmatch(str); len(match) > 1 {
		compose.Mobile = strings.TrimSpace(match[1])
		str = phoneRe.ReplaceAllString(str, " ")
	}

	// 提取所在地区 (支持: 所在地区:xxx)
	regionRe := regexp.MustCompile(`所在地区[：:]\s*([^\n]+)`)
	if match := regionRe.FindStringSubmatch(str); len(match) > 1 {
		// 将所在地区保留在字符串中,不删除
		// str 保持不变,让后续的地址解析处理
	}

	// 提取详细地址 (支持: 详细地址:xxx, 收货地址:xxx, 地址:xxx)
	addrRe := regexp.MustCompile(`(?:详细地址|收货地址|地址)[：:]\s*([^\n]+)`)
	if match := addrRe.FindStringSubmatch(str); len(match) > 1 {
		// 保留详细地址在字符串中
		str = addrRe.ReplaceAllString(str, " "+match[1])
	}

	// 如果还没有提取到姓名和手机号,尝试识别紧凑格式 (如: 马云13593464918陕西省...)
	if compose.Name == "" && compose.Mobile == "" {
		// 匹配: 2-4个汉字 + 7-12位数字 + 剩余内容
		compactRe := regexp.MustCompile(`^([\x{4e00}-\x{9fa5}]{2,4})(\d{7,12})(.*)$`)
		if match := compactRe.FindStringSubmatch(str); len(match) > 3 {
			compose.Name = match[1]
			compose.Mobile = match[2]
			str = match[3] // 保留剩余的地址部分
		}
	}

	// 替换常见的地址关键词为空格
	replacements := map[string]string{
		"收货地址": " ", "详细地址": " ", "地址": " ", "收货人": " ",
		"收件人": " ", "收货": " ", "所在地区": " ", "邮编": " ",
		"电话": " ", "手机号码": " ", "身份证号码": " ", "身份证号": " ",
		"身份证": " ", "姓名": " ", "联系电话": " ", "手机": " ",
		"：": " ", ":": " ", "；": " ", ";": " ",
		"，": " ", ",": " ", "。": " ", "\n": " ", "\r": " ",
	}

	for old, new := range replacements {
		str = strings.ReplaceAll(str, old, new)
	}

	// 将多个空格合并为一个
	spaceRe := regexp.MustCompile(`\s{1,}`)
	str = spaceRe.ReplaceAllString(str, " ")

	// 处理座机号格式 (如: 800-8585222)
	telRe := regexp.MustCompile(`(\d{3,4})-(\d{6,8})`)
	str = telRe.ReplaceAllString(str, "$1$2")

	// 提取身份证号 (18位或17位+X)
	idnRe := regexp.MustCompile(`\d{18}|\d{17}[Xx]`)
	if match := idnRe.FindString(str); match != "" {
		compose.IDN = strings.ToUpper(match)
		str = strings.ReplaceAll(str, match, "")
	}

	// 如果之前没有提取到手机号,现在提取
	if compose.Mobile == "" {
		mobileRe := regexp.MustCompile(`\d{7,12}`)
		if match := mobileRe.FindString(str); match != "" {
			compose.Mobile = match
			str = strings.ReplaceAll(str, match, "")
		}
	} else {
		// 已经提取过手机号,从字符串中删除
		str = strings.ReplaceAll(str, compose.Mobile, "")
	}

	// 提取邮编
	postcodeRe := regexp.MustCompile(`\d{6}`)
	if match := postcodeRe.FindString(str); match != "" {
		compose.Postcode = match
		str = strings.ReplaceAll(str, match, "")
	}

	// 清理多余空格
	str = strings.TrimSpace(spaceRe.ReplaceAllString(str, " "))

	// 如果之前没有提取到姓名,现在提取
	if compose.Name == "" {
		// 提取姓名（取最短的词作为姓名,排除空字符串）
		splitArr := strings.Split(str, " ")
		if len(splitArr) > 0 {
			for _, value := range splitArr {
				value = strings.TrimSpace(value)
				if value == "" {
					continue
				}
				if compose.Name == "" {
					compose.Name = value
				} else if utf8.RuneCountInString(value) < utf8.RuneCountInString(compose.Name) && utf8.RuneCountInString(value) >= 2 {
					compose.Name = value
				}
			}
			if compose.Name != "" {
				str = strings.TrimSpace(strings.ReplaceAll(str, compose.Name, ""))
			}
		}
	} else {
		// 已经提取过姓名,从字符串中删除
		str = strings.TrimSpace(strings.ReplaceAll(str, compose.Name, ""))
	}

	compose.Addr = str
	return compose
}

// fuzz 根据统计规律分析出二三级地址
func fuzz(addr string) *fuzzyResult {
	addrOrigin := addr
	addr = strings.ReplaceAll(addr, " ", "")
	addr = strings.ReplaceAll(addr, ",", "")
	// 先替换"自治区直辖县级市"为"市"，避免后续"自治区"替换时产生问题
	addr = strings.ReplaceAll(addr, "自治区直辖县级市", "市")
	addr = strings.ReplaceAll(addr, "自治区", "省")
	addr = strings.ReplaceAll(addr, "自治州", "州")
	addr = strings.ReplaceAll(addr, "小区", "")
	addr = strings.ReplaceAll(addr, "校区", "")
	// 过滤"市辖区" - 这是一个行政术语占位符,不是真正的区名
	addr = strings.ReplaceAll(addr, "市辖区", "")

	a1 := ""
	a2 := ""
	a3 := ""
	street := ""

	deep3KeywordPos := -1

	// 判断是否包含县/区/旗
	countyPos := mbStrpos(addr, "县")
	districtPos := mbStrpos(addr, "区")
	bannerPos := mbStrpos(addr, "旗")

	// 只要存在这些关键词就处理，不再限制位置
	hasEarlyCounty := countyPos != -1
	hasEarlyDistrict := districtPos != -1
	hasEarlyBanner := bannerPos != -1

	if hasEarlyCounty || hasEarlyDistrict || hasEarlyBanner {
		// 优先检查是否存在县级市（如"新乐市"）
		// 如果同时存在"XX市"和"XX区"/"XX县"，优先处理"市"
		hasCountyLevelCity := false
		if mbStrstr(addr, "市") {
			// 查找所有"市"的位置
			cityCount := mbSubstrCount(addr, "市")
			if cityCount >= 2 {
				// 找到第二个"市"的位置（可能是县级市）
				firstCityPos := mbStrpos(addr, "市")
				// 从第一个"市"之后继续查找
				addrAfterFirstCity := mbSubstr(addr, firstCityPos+1, utf8.RuneCountInString(addr)-firstCityPos-1)
				secondCityPos := mbStrpos(addrAfterFirstCity, "市")
				if secondCityPos != -1 {
					secondCityAbsPos := firstCityPos + 1 + secondCityPos
					// 检查第二个"市"后面是否存在"区"或"县"
					addrAfterSecondCity := mbSubstr(addr, secondCityAbsPos+1, utf8.RuneCountInString(addr)-secondCityAbsPos-1)
					if mbStrstr(addrAfterSecondCity, "区") || mbStrstr(addrAfterSecondCity, "县") {
						// 提取两个"市"之间的内容
						betweenCities := mbSubstr(addr, firstCityPos+1, secondCityAbsPos-firstCityPos)
						// 检查是否是重复的地名（如"北京市北京市"或"杭州市西湖区杭州市"）
						// 如果两个"市"之间包含"区"或"县"，说明不是县级市，而是重复地名
						if !mbStrstr(betweenCities, "区") && !mbStrstr(betweenCities, "县") {
							// 第一个"市"及之前的内容
							firstCityFull := mbSubstr(addr, 0, firstCityPos+1)
							if betweenCities != firstCityFull {
								// 不是重复地名，这是县级市
								a3 = betweenCities
								deep3KeywordPos = secondCityAbsPos
								hasCountyLevelCity = true
							}
						}
					}
				}
			}
		}

		if !hasCountyLevelCity {
			// 处理旗
			if mbStrstr(addr, "旗") {
				deep3KeywordPos = mbStrpos(addr, "旗")
				a3 = mbSubstr(addr, deep3KeywordPos-1, 2)
			}

			// 处理区
			if mbStrstr(addr, "区") {
				// 使用第一个"区"（避免重复地名干扰，如"西湖区杭州市西湖区"）
				deep3KeywordPos = mbStrpos(addr, "区")

				if mbStrstr(addr, "市") {
					// 策略:找到"区"之前的最后一个"市"
					// 这样可以避免详细地址中的"市"字干扰(如"农贸市场")
					zonePos := deep3KeywordPos
					// 从开头到"区"的子串中,查找最后一个"市"
					addrBeforeZone := mbSubstr(addr, 0, zonePos)
					cityPos := mbStrripos(addrBeforeZone, "市")
					if cityPos != -1 {
						a3 = mbSubstr(addr, cityPos+1, zonePos-cityPos)
					} else {
						// 没有找到"市",使用默认逻辑
						a3 = mbSubstr(addr, deep3KeywordPos-2, 3)
					}
				} else {
					a3 = mbSubstr(addr, deep3KeywordPos-2, 3)
				}
			}

			// 处理县
			if mbStrstr(addr, "县") {
				// 使用第一个"县"（避免重复地名干扰）
				deep3KeywordPos = mbStrpos(addr, "县")
				if mbStrstr(addr, "市") {
					// 从开头到"县"的子串中,查找最后一个"市"
					addrBeforeCounty := mbSubstr(addr, 0, deep3KeywordPos)
					cityPos := mbStrripos(addrBeforeCounty, "市")
					if cityPos != -1 {
						a3 = mbSubstr(addr, cityPos+1, deep3KeywordPos-cityPos)
					} else {
						if mbStrstr(addr, "自治县") {
							a3 = mbSubstr(addr, deep3KeywordPos-6, 7)
							firstChar := mbSubstr(a3, 0, 1)
							if firstChar == "省" || firstChar == "市" || firstChar == "州" {
								a3 = mbSubstr(a3, 1, utf8.RuneCountInString(a3)-1)
							}
						} else {
							a3 = mbSubstr(addr, deep3KeywordPos-2, 3)
						}
					}
				} else {
					if mbStrstr(addr, "自治县") {
						a3 = mbSubstr(addr, deep3KeywordPos-6, 7)
						firstChar := mbSubstr(a3, 0, 1)
						if firstChar == "省" || firstChar == "市" || firstChar == "州" {
							a3 = mbSubstr(a3, 1, utf8.RuneCountInString(a3)-1)
						}
					} else {
						a3 = mbSubstr(addr, deep3KeywordPos-2, 3)
					}
				}
			}
		}

		if deep3KeywordPos != -1 {
			street = mbSubstr(addrOrigin, deep3KeywordPos+1, utf8.RuneCountInString(addrOrigin)-deep3KeywordPos-1)
		}
	} else {
		// 处理市
		if mbStrripos(addr, "市") != -1 {
			cityCount := mbSubstrCount(addr, "市")
			if cityCount == 1 {
				deep3KeywordPos = mbStrripos(addr, "市")
				a3 = mbSubstr(addr, deep3KeywordPos-2, 3)
				street = mbSubstr(addrOrigin, deep3KeywordPos+1, utf8.RuneCountInString(addrOrigin)-deep3KeywordPos-1)
			} else if cityCount >= 2 {
				deep3KeywordPos = mbStrripos(addr, "市")
				a3 = mbSubstr(addr, deep3KeywordPos-2, 3)
				street = mbSubstr(addrOrigin, deep3KeywordPos+1, utf8.RuneCountInString(addrOrigin)-deep3KeywordPos-1)
			}
		} else {
			a3 = ""
			street = addr
		}
	}

	// 提取市级地址
	if mbStrpos(addr, "市") != -1 || mbStrstr(addr, "盟") || mbStrstr(addr, "州") {
		tmpPos := -1
		if tmpPos = mbStrpos(addr, "市"); tmpPos != -1 {
			// 使用第一个"市"（避免重复地名干扰，如"杭州市西湖区杭州市"）
			// 向前查找省的位置，如果有省就从省后开始，否则从开头开始
			addrBeforeCity := mbSubstr(addr, 0, tmpPos)
			provincePos := mbStrripos(addrBeforeCity, "省")
			startPos := 0
			if provincePos != -1 {
				startPos = provincePos + 1
			}
			a2 = mbSubstr(addr, startPos, tmpPos-startPos+1)
		} else if tmpPos = mbStrpos(addr, "盟"); tmpPos != -1 {
			a2 = mbSubstr(addr, tmpPos-2, 3)
		} else if mbStrpos(addr, "州") != -1 {
			if tmpPos = mbStrpos(addr, "自治州"); tmpPos != -1 {
				a2 = mbSubstr(addr, tmpPos-4, 5)
			} else {
				tmpPos = mbStrpos(addr, "州")
				a2 = mbSubstr(addr, tmpPos-2, 3)
			}
		}
	}

	return &fuzzyResult{
		A1:     a1,
		A2:     a2,
		A3:     a3,
		Street: street,
	}
}

// parse 智能解析出省市区
func parse(a1, a2, a3 string) *AddressInfo {
	r := &AddressInfo{}

	if a3 == "" {
		return r
	}

	// 在三级地址数据中查找匹配
	area3Matches := make(map[int]*Region)
	for id, v := range A3Data {
		if mbStrpos(v.Name, a3) != -1 {
			area3Matches[id] = v
		}
	}

	// 多个匹配项，需要通过二级地址筛选
	if len(area3Matches) > 1 {
		if a2 != "" {
			area2Matches := make(map[int]*Region)
			for id, v := range A2Data {
				if mbStrpos(v.Name, a2) != -1 {
					area2Matches[id] = v
				}
			}

			if len(area2Matches) > 0 {
				for _, v := range area3Matches {
					if city, ok := area2Matches[v.PID]; ok {
						r.City = city.Name
						r.Region = v.Name
						if province, ok := A1Data[city.PID]; ok {
							r.Province = province.Name
						}
					}
				}
			}
		} else {
			r.Province = ""
			r.City = ""
			r.Region = a3
		}
	} else if len(area3Matches) == 1 {
		// 唯一匹配
		for _, v := range area3Matches {
			r.Region = v.Name
			if city, ok := A2Data[v.PID]; ok {
				r.City = city.Name
				if province, ok := A1Data[city.PID]; ok {
					r.Province = province.Name
				}
			}
		}
	} else if len(area3Matches) == 0 && a2 == a3 {
		// 没有匹配到三级地址，但二级地址等于三级地址，可能是直辖市
		shengID := 0
		for _, v := range A2Data {
			if mbStrpos(v.Name, a2) != -1 {
				r.City = v.Name
				shengID = v.PID
				break
			}
		}
		if province, ok := A1Data[shengID]; ok {
			r.Province = province.Name
		}
		r.Region = ""
	}

	return r
}
