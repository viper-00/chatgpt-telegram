package sse

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	URL          string
	EventChannel chan string
	Headers      map[string]string
}

func Init(url string) Client {
	return Client{
		URL:          url,
		EventChannel: make(chan string),
	}
}

type CHATGPT struct {
	Model    string           `json:"model"`
	Messages []CHATGPTMessage `json:"messages"`
}

type CHATGPTMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CHATGPTResponse struct {
	Id      string                  `json:"id"`
	Object  string                  `json:"object"`
	Created float64                 `json:"created"`
	Model   string                  `json:"model"`
	Usage   CHATGPTResponseUsage    `json:"usage"`
	Choices []CHATGPTResponseChoice `json:"choices"`
}

type CHATGPTResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type CHATGPTResponseChoice struct {
	Message      CHATGPTResponseChoiceMessage `json:"message"`
	FinishReason string                       `json:"finish_reason"`
	Index        int                          `json:"index"`
}

type CHATGPTResponseChoiceMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (c *Client) Connect(message string, conversationId string, parentMessageId string) error {
	// messages, err := json.Marshal([]string{message})
	// if err != nil {
	// 	return errors.New(fmt.Sprintf("failed to encode message: %v", err))
	// }

	// if parentMessageId == "" {
	// 	parentMessageId = uuid.NewString()
	// }

	// var conversationIdString string
	// if conversationId != "" {
	// 	conversationIdString = fmt.Sprintf(`, "conversation_id": "%s"`, conversationId)
	// }

	// if conversation id is empty, don't send it
	// body := fmt.Sprintf(`{
	//     "action": "next",
	//     "messages": [
	//         {
	//             "id": "%s",
	//             "role": "user",
	//             "content": {
	//                 "content_type": "text",
	//                 "parts": %s
	//             }
	//         }
	//     ],
	//     "model": "gpt-3.5-turbo",
	// 	"parent_message_id": "%s"%s
	// }`, uuid.NewString(), string(messages), parentMessageId, conversationIdString)

	var gpt CHATGPT
	gpt.Model = "gpt-3.5-turbo"
	var cMsg CHATGPTMessage
	cMsg.Role = "user"
	cMsg.Content = message
	gpt.Messages = append(gpt.Messages, cMsg)

	body, err := json.Marshal(gpt)
	if err != nil {
		return fmt.Errorf("failed to encode source: %v", err)
	}
	// gpt.Messages

	// body := fmt.Sprintf(`{
	// 	"model": "gpt-3.5-turbo",
	// 	"messages": [{"role": "user", "content": "%s"}]
	// }`, string(messages))

	req, err := http.NewRequest("POST", c.URL, strings.NewReader(string(body)))
	if err != nil {
		return errors.New(fmt.Sprintf("failed to create request: %v", err))
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	http := &http.Client{}
	resp, err := http.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to connect to SSE: %v", err))
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("failed to connect to SSE: %v", resp.Status))
	}

	// decoder := eventsource.NewDecoder(resp.Body)

	go func() {
		defer resp.Body.Close()
		defer close(c.EventChannel)

		for {
			// event, err := decoder.Decode()
			// if err != nil {
			// 	log.Println(errors.New(fmt.Sprintf("failed to decode event: %v", err)))
			// 	break
			// }
			// if event.Data() == "[DONE]" || event.Data() == "" {
			// 	break
			// }

			var dest CHATGPTResponse
			err = json.NewDecoder(resp.Body).Decode(&dest)
			if err != nil {
				log.Println(errors.New(fmt.Sprintf("failed to decode: %v", err)))
				break
			}

			c.EventChannel <- dest.Choices[0].Message.Content
		}
	}()

	return nil
}
