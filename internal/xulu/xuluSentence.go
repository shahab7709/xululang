package xulu

type ixuluSentence interface {
	compute() (int, error)
}

type xuluSentence struct {
	verb      xuluVerb
	sentences []ixuluSentence
}

func NewXuluSentence(verb xuluVerb) xuluSentence {
	return xuluSentence{verb: verb, sentences: make([]ixuluSentence, 0)}
}

func (xs *xuluSentence) compute() (int, error) {
	result := 0
	var err error
	switch xs.verb {
	case "abcd":
		result, err = xs.sum()
	case "bcde":
		result, err = xs.sub()
	case "dede":
		result, err = xs.mul()
	}
	return result, err
}

func (xs *xuluSentence) sum() (int, error) {
	result := 0
	for _, item := range xs.sentences {
		value, err := item.compute()
		if err != nil {
			return 0, err
		}
		result += value
	}
	return result, nil
}

func (xs *xuluSentence) sub() (int, error) {
	result := 0
	for _, item := range xs.sentences {
		value, err := item.compute()
		if err != nil {
			return 0, err
		}
		if result == 0 {
			result = value
		} else {
			result -= value
		}

	}
	return result, nil
}

func (xs *xuluSentence) mul() (int, error) {
	result := 1
	for _, item := range xs.sentences {
		value, err := item.compute()
		if err != nil {
			return 0, err
		}
		result *= value
	}
	return result, nil
}
