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
	// args はコマンド引数を格納します。
	args struct {
		seed string
		n    int
	}
	// ApiBody は、API結果を格納します。
	ApiBody struct {
		Result int
	}
	// askRepository は、APIへのリクエストを実行します。
	askRepository interface {
		Ask(int) (int, error)
	}
	// askRepositoryImpl は、askRepositoryの実装です。
	askRepositoryImpl struct {
		seed   string
		result map[int]int
	}
	// solver は、チャレンジのロジック部を司ります。
	solver struct {
		repo  askRepository
		cache map[int]int
	}
)

func main() {
	args, err := parseArgs(os.Args)

	if err != nil {
		// returning when an error occurred.
		printError(err)
		return
	}
	s := newSolver(args.seed)
	ans, err := s.solve(args.n)

	if err != nil {
		printError(err)
		return
	}
	fmt.Println(ans)
	os.Exit(0)
}

// printErrorは、エラー結果をprintします
func printError(err error) {
	msg := fmt.Sprintf("error! %s", err.Error())
	fmt.Println(msg)
	os.Stderr.WriteString(msg)
	os.Exit(1)
}

// parseArgs は引数のパースとバリデーションを実施します。
func parseArgs(a []string) (*args, error) {
	// checking arguments length
	if len(a) < 3 {
		return nil, errors.Errorf("invalid arguments' %s", strings.Join(os.Args, ","))
	}
	// checking n as an integer
	n, err := strconv.Atoi(a[2])
	if err != nil {
		return nil, errors.Wrapf(err, "The second argument(n) must be an integer, %s given", a[1])
	}

	// seed の文字列長チェックを入れたい気がするが、テスト仕様上要求されていないので行わない

	// valid pattern
	return &args{
		n:    n,
		seed: a[1],
	}, nil
}

// newSolver は、solverインスタンスを返します。
func newSolver(seed string) *solver {
	return &solver{
		repo:  newAskRepositoryImpl(seed),
		cache: map[int]int{},
	}
}

// solve は、課題のロジックを実行します。
func (s *solver) solve(n int) (int, error) {
	switch n {
	case 0:
		//f(0) = 1
		return 1, nil
	case 2:
		//f(2) = 2
		return 2, nil
	}

	// 奇数ならサーバーを叩いた結果を返す
	if n%2 != 0 {
		return s.repo.Ask(n)
	}
	// 偶数ならf(n - 1..4)の合計を返す
	var ans int
	for i := 1; i <= 4; i++ {
		tgt := n - i
		// 同じ数で再帰的に処理が走ってしまって重いので、solver側の偶数処理結果もキャッシュを活用する。
		// api側はrepository内部キャッシュで済んでいるので、solverでは気にしないこととする。
		if res, ok := s.cache[tgt]; ok {
			ans += res
			continue
		}
		res, err := s.solve(tgt)
		if err != nil {
			return 0, err
		}
		s.cache[tgt] = res
		ans += res
	}
	return ans, nil
}

// newAskRepositoryImpl は、askRepositoryImplインスタンスを生成します。
func newAskRepositoryImpl(seed string) *askRepositoryImpl {
	return &askRepositoryImpl{
		seed:   seed,
		result: map[int]int{},
	}
}

// Ask はAPIの結果を返します。キャッシュがあればキャッシュを返します。
func (r *askRepositoryImpl) Ask(n int) (int, error) {
	if res, ok := r.result[n]; ok {
		return res, nil
	}
	res, err := r.askServer(n)
	if err != nil {
		return -1, err
	}
	r.result[n] = res
	return res, nil
}

func (r *askRepositoryImpl) askServer(n int) (int, error) {
	q := url.Values{}
	q.Add("seed", r.seed)
	q.Add("n", strconv.Itoa(n))
	u := "http://challenge-server.code-check.io/api/recursive/ask?" + q.Encode()

	resp, err := http.Get(u)
	if err != nil {
		return -1, errors.Wrapf(err, "cannot get the url: %s", u)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return -1, errors.Errorf("invalid response code %d", resp.StatusCode)
	}

	bd := &ApiBody{}
	if err := json.NewDecoder(resp.Body).Decode(bd); err != nil {
		return -1, errors.Wrap(err, "cannot parse the response")
	}

	return bd.Result, nil
}
