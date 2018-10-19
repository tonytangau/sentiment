package sentiment

import "strings"

// SentimentAnalysis takes in a (possibly 'dirty')
// sentence (or any block of text,) cleans the
// text, finds the sentiment of each word in the
// text, finds the sentiment of the sentence as
// a whole, adn returns an Analysis struct
func (m Models) SentimentAnalysis(sentence string, lang Language) *Analysis {
	if _, ok := m[lang]; !ok {
		lang = English
	}

	analysis := &Analysis{
		Language: lang,
		Words:    []Score{},
		Score:    uint8(0),
	}

	sentences := strings.FieldsFunc(sentence, SplitSentences)
	if len(sentences) > 1 {
		analysis.Sentences = []SentenceScore{}

		for _, s := range sentences {
			analysis.Sentences = append(analysis.Sentences, SentenceScore{
				Sentence: s,
				Score:    m[lang].Predict(s),
			})
		}
	}

	w := strings.Split(sentence, " ")
	sum := uint8(0)

	for _, word := range w {
		p := m[lang].Predict(word)

		analysis.Words = append(analysis.Words, Score{
			Word:  word,
			Score: p,
		})

		sum = sum + p
	}

	analysis.Score = m[lang].Predict(sentence)
	analysis.Average = float32(sum) / float32(len(w))

	return analysis
}
