package check

import (
	"errors"
	"net/http"
	"strconv"
)

type checker interface {
	Check() (int, error)
}

var (
	ErrInvalid    = errors.New("The provided link is Invalid")
	ErrOverMaxInt = errors.New("The number of chuncks surpass Max")
	ErrUnxcpected = errors.New("Unxcpected error happened")
)

// getStatusCode will try and make an HTTP Get request to a giving URL
// and return it's HTTP Status Code
func getStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return -1, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func mockGetStatusCode(url string) (int, error) {
	var m = make(map[string]int)
	for i := 0; i < 15; i++ {
		link := "https://xxxxx.cloudfront.net/gebrish_user_XXXXXX_XXXXXXX/chunked/" + strconv.Itoa(i) + ".ts"
		m[link] = http.StatusOK
	}

	return m[url], nil
}
