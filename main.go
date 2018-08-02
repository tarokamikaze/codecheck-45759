package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/pkg/errors"
	"strconv"
	"net/url"
	"net/http"
	"encoding/json"
)

type (
	args struct {
		seed string
		n    int
	}
)

func main() {
	args, err := parseArgs(os.Args)

	if err != nil {
		// returning when an error occurred.
		printError(err)
		return
	}

	ans, err := solve(args)
	if err != nil {
		printError(err)
	}
	fmt.Println(ans)
}
func printError(err error) {
	msg := fmt.Sprintf("error! %s", err.Error())
	fmt.Println(msg)
	os.Stderr.WriteString(msg)
}
func parseArgs(args []string) (*args, error) {
	// checking arguments length
	if len(args) < 2 {
		return nil, errors.Errorf("invalid arguments' %s", strings.Join(os.Args, ","))
	}
	// checking n as an integer
	n, err := strconv.Atoi(args[1])
	if err != nil {
		return nil, errors.Wrapf(err, "The second argument(n) must be an integer, %s given", args[1])
	}

	// seed の文字列長チェックを入れたい気がするが、テスト仕様上要求されていないので行わない

	// valid pattern
	return &args{
		n:    n,
		seed: args[0],
	}, nil
}
func solve(args *args) (int, error) {
	switch args.n {
	case 0:
		//f(0) = 1
		return 1, nil
	case 2:
		//f(2) = 2
		return 2, nil
	}

	// 奇数ならサーバーを叩いた結果を返す
	if args.n%2 != 0 {
		return askServer(args)
	}
	// 偶数ならf(n - 1..4)の合計を返す
	var ans int
	for i := 1; i <= 4; i++ {
		tmpArgs := &args{
			n:    args.n - i,
			seed: args.seed,
		}
		res, err := solve(tmpArgs)
		if err != nil {
			return 0, err
		}
		ans += res
	}
	return ans, nil
}
func askServer(args *args) (int, error) {
	q := url.Values{}
	q.Add("seed", args.seed)
	q.Add("n", strconv.Itoa(args.n))
	u := "http://challenge-server.code-check.io/api/recursive/ask?" + q.Encode()

	resp, err := http.Get(u)
	if err != nil {
		return -1, errors.Wrapf(err, "cannot get the url: %s", u)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return -1, errors.Errorf("invalid response code %d", resp.StatusCode)
	}

	bd := map[string]interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&bd); err != nil {
		return -1, errors.Wrap(err, "cannot parse the response")
	}

	rsl, ok := bd["result"]
	if !ok {
		return -1, errors.Errorf("cannot find the key 'result' in the response: %v", bd)
	}
	nRsl, err := strconv.Atoi(rsl.(string))
	if err != nil {
		return -1, errors.Wrapf(err, "the 'result' response is not an integer: %s given", rsl)
	}

	return nRsl, nil
}
