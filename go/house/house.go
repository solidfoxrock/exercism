package house

func Embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

func Verse(subject string, relPhrases []string, nounPhrase string) string {
	ret := subject + " "
	for _, r := range relPhrases {
		ret += r + " "
	}
	ret += nounPhrase
	return ret
}

func Song() string {

	song := ""
	s := "This is"
	rels := []string{
		"the horse and the hound and the horn\nthat belonged to",
		"the farmer sowing his corn\nthat kept",
		"the rooster that crowed in the morn\nthat woke",
		"the priest all shaven and shorn\nthat married",
		"the man all tattered and torn\nthat kissed",
		"the maiden all forlorn\nthat milked",
		"the cow with the crumpled horn\nthat tossed",
		"the dog\nthat worried",
		"the cat\nthat killed",
		"the rat\nthat ate",
		"the malt\nthat lay in",
	}
	n := "the house that Jack built."

	song += Verse(s, []string{}, n)
	for i := len(rels) - 1; i >= 0; i-- {
		song += "\n\n" + Verse(s, rels[i:], n)
	}
	return song
}
