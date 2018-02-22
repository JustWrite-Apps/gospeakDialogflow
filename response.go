package gospeakDialogflow

import (
	"encoding/json"
	"fmt"

	"github.com/blforce/gospeakCommon"
)

type eventInput struct {
	Name         string                 `json:"name,omitempty"`
	Parameters   map[string]interface{} `json:"parameters,omitempty"`
	LanguageCode string                 `json:"languageCode,omitempty"`
}

type Response struct {
	FulfillmentText     string      `json:"fulfillmentText,omitempty"`
	FulfillmentMessages []message   `json:"fulfillmentMessages,omitempty"`
	Source              string      `json:"source,omitempty"`
	Payload             interface{} `json:"payload,omitempty"`
	OutputContexts      []context   `json:"outputContexts,omitempty"`
	FollowupEventInput  interface{} `json:"followupEventInput,omitempty"`
}

func (r Response) SetText(value string) gospeakCommon.Response {
	r.FulfillmentText = value
	return r
}

func (r Response) SetImageCard(title, imageURL, text string) gospeakCommon.Response {
	cardMessage := message{
		Card: &card{
			Title:    title,
			ImageURI: imageURL,
			Subtitle: text,
		},
	}

	r.FulfillmentMessages = append(r.FulfillmentMessages, cardMessage)

	return r
}

func (r Response) GetBytes() []byte {
	result, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return result
}
