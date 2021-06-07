package server

import (
	"apihut-server/model"
	"errors"

	"github.com/go-ego/gse"
)

var (
	ErrCutType = errors.New("类型错误")
)

const (
	CutTypeBase   = "base"
	CutTypeSearch = "search"
	CutTypeAll    = "all"
)

// GetWords 获取分词
func GetWords(w *model.Words) (words []string, err error) {
	var seg gse.Segmenter
	err = seg.LoadDict("./data/dict/dictionary.txt")
	if err != nil {
		return nil, err
	}
	//seg.LoadDict("./data/dict/dictionary.txt")
	err = seg.LoadStop()
	if err != nil {
		return nil, err
	}

	if w.Type == CutTypeBase {
		words = seg.Cut(w.Word, w.Hmm)
	} else if w.Type == CutTypeSearch {
		words = seg.CutSearch(w.Word, w.Hmm)
	} else if w.Type == CutTypeAll {
		words = seg.CutAll(w.Word)
	} else {
		return nil, ErrCutType
	}

	return
}
