package awesomeProject1

import "testing"

const goodURL = "http://www.baidu.com"
const badURL = "http://aa4ee2d7-d961-48da-abf1-4d3f83ccb726.com"

func TestIPv4Fetcher(t *testing.T) {
  IPv4Fetcher(goodURL)
  IPv4Fetcher(badURL)
}
