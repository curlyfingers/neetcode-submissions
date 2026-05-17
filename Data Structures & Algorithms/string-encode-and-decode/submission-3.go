import stringies "strings"

type Solution struct{
	delim string
	emptySeq string
}

func (s *Solution) Encode(strs []string) string {
	s.emptySeq = "&^(#^&"
	if s.delim == "" {
		s.delim = "#**#"
	}
	if len(strs) == 0 {
		return s.emptySeq
	}
	return stringies.Join(strs, s.delim)
}

func (s *Solution) Decode(encoded string) []string {
	s.emptySeq = "&^(#^&"
	if encoded == s.emptySeq {
		return []string{}
	}
	if s.delim == "" {
		s.delim = "#**#"
	}
	
	return stringies.Split(encoded, s.delim)
}
