package handler

import (
	"log/slog"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gofiber/fiber/v2"
	"github.com/hardeepnarang10/query/pkg/authorization"
	"github.com/hardeepnarang10/query/service/common/header"
	"github.com/hardeepnarang10/query/service/db/elasticsearch/store"
	"github.com/hardeepnarang10/query/service/db/elasticsearch/store/search"
	"github.com/hardeepnarang10/query/service/handler/v1/internal/jsonmap"
	"github.com/hardeepnarang10/query/service/handler/v1/internal/pagination"
	"github.com/hardeepnarang10/query/service/handler/v1/message"
)

func (h *handler) Rank(c *fiber.Ctx) error {
	msgRank := message.Rank{}
	if err := msgRank.Unmarshal(c.Queries()); err != nil {
		slog.DebugContext(c.Context(),
			"unable to parse incoming request body to rank message type",
			slog.Any("error", err),
		)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	jsonMap := jsonmap.Parse(msgRank)
	queries := make([]types.Query, 0)
	if msgRank.Query != nil {
		termQuery := search.Terms(search.TermFieldsNonTimestamp, *msgRank.Query)
		queries = append(queries, termQuery...)
	}
	rangeQuery := search.Range(jsonMap)
	queries = append(queries, rangeQuery...)
	if len(queries) == 0 {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	authorization := c.Get(header.AuthorizationHeaderKey, authorization.UserAccessDefault)
	if err := h.service.AuthorizationClient.Validate(authorization); err != nil {
		slog.DebugContext(c.Context(),
			"authorization validation failed",
			slog.String("authorization_token", authorization),
			slog.Any("error", err),
		)
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	fields := h.service.AuthorizationClient.Fields(authorization)

	pn := pagination.Parse(c.Queries())
	searchParams := store.SearchParams{
		Fields:  fields,
		Queries: queries,
		Page:    pn.Page,
		Size:    pn.Count,
	}
	jsonRawMessageSlice, err := h.ess.SearchExecutor(c.Context(), searchParams)
	if err != nil {
		slog.ErrorContext(c.Context(),
			"unable to perform rank search on datastore",
			slog.Any("search_params", searchParams),
			slog.Any("error", err),
		)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	documentSlice, err := jsonRawMessageSlice.Unmarshal()
	if err != nil {
		slog.ErrorContext(c.Context(),
			"unable to unmarshal json raw message slice",
			slog.Any("json_raw_message_slice", jsonRawMessageSlice),
			slog.Any("error", err),
		)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(documentSlice, fiber.MIMEApplicationJSON)
}
