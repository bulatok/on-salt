package salt

import (
	"fmt"
	"net/http"
)

type bot struct {
	Token string
	Dst   []Dst
}

func (b *bot) send(what string) error {
	for _, v := range b.Dst {
		dstID := v.ID()
		link := createLink(b.Token, dstID, what)
		resp, err := http.Post(link, "", nil)
		if err != nil {
			return err
		}

		switch resp.StatusCode {
		case 200:
			continue
		case 401:
			return ErrInvalidToken
		case 400:
			return ErrChatNotFound
		}
	}
	return nil
}

func createLink(token, dstID, text string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		token, dstID, text)
}
