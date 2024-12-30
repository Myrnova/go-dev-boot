package CH8

import (
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	newMessages := make([]sms, 0, len(messages))
	for _, m := range messages {
		m.tags = tagger(m)
		newMessages = append(newMessages, m)
	}

	return newMessages
}

func tagger(msg sms) []string {
	tags := []string{}
	msg.content = strings.ToLower(msg.content)
	if strings.Contains(msg.content, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(msg.content, "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}
