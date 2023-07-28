package netutil

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/fileutil"
)

// GetInternalIp return internal ipv4.
// Play: https://go.dev/play/p/5mbu-gFp7ei
func GetInternalIp() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		panic(err.Error())
	}
	for _, a := range addr {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}

	return ""
}

// GetRequestPublicIp return the requested public ip.
// Play: https://go.dev/play/p/kxU-YDc_eBo
func GetRequestPublicIp(req *http.Request) string {
	var ip string
	for _, ip = range strings.Split(req.Header.Get("X-Forwarded-For"), ",") {
		if ip = strings.TrimSpace(ip); ip != "" && !IsInternalIP(net.ParseIP(ip)) {
			return ip
		}
	}

	if ip = strings.TrimSpace(req.Header.Get("X-Real-Ip")); ip != "" && !IsInternalIP(net.ParseIP(ip)) {
		return ip
	}

	if ip, _, _ = net.SplitHostPort(req.RemoteAddr); !IsInternalIP(net.ParseIP(ip)) {
		return ip
	}

	return ip
}

// GetPublicIpInfo return public ip information
// return the PublicIpInfo struct.
// Play: https://go.dev/play/p/YDxIfozsRHR
func GetPublicIpInfo() (*PublicIpInfo, error) {
	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ip PublicIpInfo
	err = json.Unmarshal(body, &ip)
	if err != nil {
		return nil, err
	}

	return &ip, nil
}

// GetIps return all ipv4 of system.
// Play: https://go.dev/play/p/NUFfcEmukx1
func GetIps() []string {
	var ips []string

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}

	for _, addr := range addrs {
		ipNet, isValid := addr.(*net.IPNet)
		if isValid && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

// GetMacAddrs get mac address.
// Play: https://go.dev/play/p/Rq9UUBS_Xp1
func GetMacAddrs() []string {
	var macAddrs []string

	nets, err := net.Interfaces()
	if err != nil {
		return macAddrs
	}

	for _, net := range nets {
		macAddr := net.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}

	return macAddrs
}

// PublicIpInfo public ip info: country, region, isp, city, lat, lon, ip
type PublicIpInfo struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Ip          string  `json:"query"`
}

// IsPublicIP verify a ip is public or not.
// Play: https://go.dev/play/p/nmktSQpJZnn
func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// IsInternalIP verify an ip is intranet or not.
// Play: https://go.dev/play/p/sYGhXbgO4Cb
func IsInternalIP(IP net.IP) bool {
	if IP.IsLoopback() {
		return true
	}
	if ip4 := IP.To4(); ip4 != nil {
		return ip4[0] == 10 ||
			(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) ||
			(ip4[0] == 169 && ip4[1] == 254) ||
			(ip4[0] == 192 && ip4[1] == 168)
	}
	return false
}

// EncodeUrl encode url.
// Play: https://go.dev/play/p/bsZ6BRC4uKI
func EncodeUrl(urlStr string) (string, error) {
	URL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	URL.RawQuery = URL.Query().Encode()

	return URL.String(), nil
}

// DownloadFile will upload the file to a server.
func UploadFile(filepath string, server string) (bool, error) {
	if !fileutil.IsExist(filepath) {
		return false, errors.New("file not exist")
	}

	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}

	bodyBuffer := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuffer)

	formFile, err := writer.CreateFormFile("uploadfile", fileInfo.Name())
	if err != nil {
		return false, err
	}

	srcFile, err := os.Open(filepath)
	if err != nil {
		return false, err
	}
	defer srcFile.Close()

	_, err = io.Copy(formFile, srcFile)
	if err != nil {
		return false, err
	}

	contentType := writer.FormDataContentType()
	writer.Close()

	_, err = http.Post(server, contentType, bodyBuffer)
	if err != nil {
		return false, err
	}

	return true, nil
}

// DownloadFile will download the file exist in url to a local file.
// Play: todo
func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}

// IsPingConnected checks if can ping specified host or not.
// Play: https://go.dev/play/p/q8OzTijsA87
func IsPingConnected(host string) bool {
	cmd := exec.Command("ping", host, "-c", "4", "-W", "6")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("ping", host, "-n", "4", "-w", "6")
	}

	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

// IsTelnetConnected checks if can telnet specified host or not.
// Play: https://go.dev/play/p/yiLCGtQv_ZG
func IsTelnetConnected(host string, port string) bool {
	adder := host + ":" + port
	conn, err := net.DialTimeout("tcp", adder, 5*time.Second)

	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}
