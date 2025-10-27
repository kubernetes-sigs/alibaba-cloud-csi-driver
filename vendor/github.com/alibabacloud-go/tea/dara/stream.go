package dara

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type SSEEvent struct {
	Id    *string
	Event *string
	Data  *string
	Retry *int
}

func parseEvent(lines []string) *SSEEvent {
	event := &SSEEvent{}
	for _, line := range lines {
		if strings.HasPrefix(line, "data:") {
			var data string
			if strings.HasPrefix(line, "data: ") {
				data = strings.TrimPrefix(line, "data: ") + "\n"
			} else {
				data = strings.TrimPrefix(line, "data:") + "\n"
			}
			if event.Data == nil {
				event.Data = new(string)
			}
			*event.Data += data
		} else if strings.HasPrefix(line, "event:") {
			var eventName string
			if strings.HasPrefix(line, "event: ") {
				eventName = strings.TrimPrefix(line, "event: ")
			} else {
				eventName = strings.TrimPrefix(line, "event:")
			}
			event.Event = &eventName
		} else if strings.HasPrefix(line, "id:") {
			var id string
			if strings.HasPrefix(line, "id: ") {
				id = strings.TrimPrefix(line, "id: ")
			} else {
				id = strings.TrimPrefix(line, "id:")
			}
			event.Id = &id
		} else if strings.HasPrefix(line, "retry:") {
			var retryStr string
			if strings.HasPrefix(line, "retry: ") {
				retryStr = strings.TrimPrefix(line, "retry: ")
			} else {
				retryStr = strings.TrimPrefix(line, "retry:")
			}
			var retry int
			fmt.Sscanf(retryStr, "%d", &retry)
			event.Retry = &retry
		}
	}
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
}

func ReadAsSSEWithContext(ctx context.Context, body io.ReadCloser, eventChannel chan *SSEEvent, errorChannel chan error) {

	go func() {
		defer func() {
			body.Close()
			close(eventChannel)
		}()

		reader := bufio.NewReader(body)
		var eventLines []string

		for {
			select {
			case <-ctx.Done():
				errorChannel <- ctx.Err()
				return
			default:
			}

			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					// Handle the end of the stream and possibly pending event
					if len(eventLines) > 0 {
						event := parseEvent(eventLines)
						select {
						case eventChannel <- event:
						case <-ctx.Done():
							errorChannel <- ctx.Err()
							return
						}
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
					select {
					case eventChannel <- event:
					case <-ctx.Done():
						errorChannel <- ctx.Err()
						return
					}
					eventLines = []string{} // Reset for the next event
				}
				continue
			}

			eventLines = append(eventLines, line)
		}
	}()
}
