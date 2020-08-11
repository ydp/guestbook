package http

type Guestbook struct {
	SignatureCount int
	Signatures     []string
}

func createGuestbook(signatures []Signature) Guestbook {
	var result []string
	for _, s := range signatures {
		result = append(result, s.Signature)
	}
	return Guestbook{ SignatureCount: len(result), Signatures: result}
}