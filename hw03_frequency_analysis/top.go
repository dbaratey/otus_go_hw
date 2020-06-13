package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

var tpl = regexp.MustCompile("[.,?!:()\"]|(\\.\\.\\.)")

type Word struct {
	word  string
	count int
}

type Words []Word

func (kvs Words) Len() int           { return len(kvs) }
func (kvs Words) Less(i, j int) bool { return kvs[i].count > kvs[j].count }
func (kvs Words) Swap(i, j int)      { kvs[i], kvs[j] = kvs[j], kvs[i] }

func sortMap(m map[string]int) Words {
	res := make(Words, 0, 100)
	for k, v := range m {
		res = append(res, Word{
			word:  k,
			count: v,
		})
	}
	sort.Sort(res)
	return res
}

func normalize(s string) string {
	tmp := tpl.ReplaceAllString(s, "")
	tmp = replaceMinus(tmp)
	res := strings.ToLower(tmp)
	return res
}

func replaceMinus(s string) string {
	return strings.Trim(s, "-")
}

func Top10(s string, asterisk bool) []string {
	// Place your code here
	res := make([]string, 0, 10)
	if s == "" {
		return res
	}
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		tmp := v
		if asterisk {
			tmp = normalize(v)
		}
		if tmp != "" {
			m[tmp]++
		}
	}
	tmpArr := sortMap(m)
	if len(tmpArr) >= 10 {
		tmpArr = tmpArr[:10]
	}
	for _, w := range tmpArr {
		res = append(res, w.word)
	}

	return res
}
