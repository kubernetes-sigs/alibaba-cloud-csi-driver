package dara

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
	"fmt"
)

// 定义 Event 结构体
type SSEEvent struct {
	ID    *string
	Event *string
	Data  *string
	Retry *int
}

// 解析单个事件
func parseEvent(lines []string) *SSEEvent {
	event := &SSEEvent{}
	for _, line := range lines {
		if strings.HasPrefix(line, "data: ") {
			data := strings.TrimPrefix(line, "data: ") + "\n"
			if event.Data == nil {
				event.Data = new(string)
			}
			*event.Data += data
		} else if strings.HasPrefix(line, "event: ") {
			eventName := strings.TrimPrefix(line, "event: ")
			event.Event = &eventName
		} else if strings.HasPrefix(line, "id: ") {
			id := strings.TrimPrefix(line, "id: ")
			event.ID = &id
		} else if strings.HasPrefix(line, "retry: ") {
			var retry int
			fmt.Sscanf(strings.TrimPrefix(line, "retry: "), "%d", &retry)
			event.Retry = &retry
		}
	}
	// Remove last newline from data
	if event.Data != nil {
		data := strings.TrimRight(*event.Data, "\n")
		event.Data = &data
	}
	return event
}

func ReadAsBytes(body io.Reader) ([]byte, error) {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	r, ok := body.(io.ReadCloser)
	if ok {
		r.Close()
	}
	return byt, nil
}

func ReadAsJSON(body io.Reader) (result interface{}, err error) {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	if string(byt) == "" {
		return
	}
	r, ok := body.(io.ReadCloser)
	if ok {
		r.Close()
	}
	d := json.NewDecoder(bytes.NewReader(byt))
	d.UseNumber()
	err = d.Decode(&result)
	return
}

func ReadAsString(body io.Reader) (string, error) {
	byt, err := ioutil.ReadAll(body)
	if err != nil {
		return "", err
	}
	r, ok := body.(io.ReadCloser)
	if ok {
		r.Close()
	}
	return string(byt), nil
}

func ReadAsSSE(body io.ReadCloser, eventChannel chan *SSEEvent, errorChannel chan error) {

	go func() {
		defer func() {
			body.Close()
			close(eventChannel)
		}()

		reader := bufio.NewReader(body)
		var eventLines []string

		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					// Handle the end of the stream and possibly pending event
					if len(eventLines) > 0 {
						event := parseEvent(eventLines)
						eventChannel <- event
					}
					errorChannel <- nil
					return
				}
				errorChannel <- err
				return
			}

			line = strings.TrimRight(line, "\n")

			if line == "" {
				// End of an SSE event
				if len(eventLines) > 0 {
					event := parseEvent(eventLines)
					eventChannel <- event
					eventLines = []string{} // Reset for the next event
				}
				continue
			}

			eventLines = append(eventLines, line)
		}
	}()

	return
}
