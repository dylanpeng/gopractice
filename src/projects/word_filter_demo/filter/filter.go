package filter

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type word struct {
	end   bool
	child map[rune]*word
}

// 过滤服务
type FilterManager struct {
	ctx   context.Context
	repl  rune             // 替换字符串
	regex []*regexp.Regexp // 正则表达式
	words *word            // 敏感词
}

func NewFilterManager(repl rune) *FilterManager {
	return &FilterManager{
		repl:  repl,
		regex: make([]*regexp.Regexp, 0),
		words: newWord(),
	}
}

// 正则表达式过滤
func (f *FilterManager) RegexFilter(text string) string {
	str := text
	for _, r := range f.regex {
		str = r.ReplaceAllStringFunc(str, f.toChange)
	}
	return str
}

// 敏感词过滤
func (f *FilterManager) WordsFilter(text string) string {
	uchars := []rune(text)
	idexs := f.doIndexes(uchars)
	for i := 0; i < len(idexs); i++ {
		uchars[idexs[i]] = rune(f.repl)
	}
	return string(uchars)
}

// 批量设置正则表达式
func (f *FilterManager) BatchSetRegex(rgs []string) {
	for _, v := range rgs {
		re, err := regexp.Compile(v)
		if err != nil {
			fmt.Println("batch set regex failed", err)
			continue
		}
		f.regex = append(f.regex, re)
	}
}

// 批量设置敏感词
func (f *FilterManager) BatchSetWords(ws []string) {
	for _, w := range ws {
		f.addWord(w)
	}
}

// 设置正则表达式
func (f *FilterManager) SetRegex(r string) {
	re, err := regexp.Compile(r)
	if err != nil {
		fmt.Println("parses a regular expression failed", err)
		return
	}
	f.regex = append(f.regex, re)
}

// 设置敏感词
func (f *FilterManager) SetWord(w string) {
	f.addWord(w)
}

func newWord() *word {
	return &word{
		child: make(map[rune]*word),
	}
}

func (f *FilterManager) addWord(word string) {
	word = strings.ToLower(word)
	w := f.words
	uchars := []rune(word)
	for i, l := 0, len(uchars); i < l; i++ {
		if unicode.IsSpace(uchars[i]) {
			continue
		}
		if _, ok := w.child[uchars[i]]; !ok {
			w.child[uchars[i]] = newWord()
		}
		w = w.child[uchars[i]]
	}
	w.end = true
}

func (f *FilterManager) toChange(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		b.WriteByte(byte(f.repl))
	}
	return b.String()
}

func (f *FilterManager) doIndexes(uchars []rune) (idexs []int) {
	var (
		tIdexs []int
		ul     = len(uchars)
		n      = f.words
	)
	for i := 0; i < ul; i++ {
		charLower := unicode.ToLower(uchars[i]) // 转换为小写
		if _, ok := n.child[charLower]; !ok {
			continue
		}
		n = n.child[charLower]
		tIdexs = append(tIdexs, i)
		if n.end {
			idexs = f.appendTo(idexs, tIdexs)
			tIdexs = nil
		}
		for j := i + 1; j < ul; j++ {
			charLower = unicode.ToLower(uchars[j]) // 转换为小写
			if _, ok := n.child[charLower]; !ok {
				break
			}
			n = n.child[charLower]
			tIdexs = append(tIdexs, j)
			if n.end {
				idexs = f.appendTo(idexs, tIdexs)
			}
		}
		if tIdexs != nil {
			tIdexs = nil
		}
		n = f.words
	}
	return
}

func (f *FilterManager) appendTo(dst, src []int) []int {
	var t []int
	for i, il := 0, len(src); i < il; i++ {
		var exist bool
		for j, jl := 0, len(dst); j < jl; j++ {
			if src[i] == dst[j] {
				exist = true
				break
			}
		}
		if !exist {
			t = append(t, src[i])
		}
	}
	return append(dst, t...)
}
