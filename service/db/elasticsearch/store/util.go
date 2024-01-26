package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/hardeepnarang10/query/service/common/document"
)

func elasticInt(i int) *int {
	return &i
}

func elasticString(s string) *string {
	return &s
}

type JSONRawMessageSlice []json.RawMessage

func (j *JSONRawMessageSlice) LogValue() slog.Value {
	if j == nil {
		return slog.Value{}
	}

	attrs := []slog.Attr{}
	for i, message := range *j {
		attrs = append(attrs, slog.String(fmt.Sprintf("message_%02d", i), string(message)))
	}
	return slog.GroupValue(attrs...)
}

func (j *JSONRawMessageSlice) Unmarshal() ([]document.Plain, error) {
	if j == nil {
		return nil, errors.New("unmarshal attempted on empty JSONRawMessageSlice type")
	}

	documentPlainSlice := make([]document.Plain, len(*j))
	for i, message := range *j {
		var documentExtend document.Extended
		if err := json.Unmarshal(message, &documentExtend); err != nil {
			return nil, fmt.Errorf("unable to unmarshal query result to log document extended container %q: %w", string(message), err)
		}
		documentPlainSlice[i] = *documentExtend.Reduce()
	}

	return documentPlainSlice, nil
}
