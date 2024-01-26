package definition

import (
	"github.com/gofiber/fiber/v2"
)

// @Summary 			Perform filtered search
// @Description		Perform a filtered search across multiple fields.
// @Tags 					Search
// @Param 				level 										query	string	false	"Log level"
// @Param 				message 									query	string	false	"Log message"
// @Param 				resourceId 								query	string	false	"Log resourceId"
// @Param 				timestampStart 						query	string	false	"Timerange start value"
// @Param 				timestampEnd 							query	string	false	"Timerange end value"
// @Param 				traceId 									query	string	false	"Log traceId"
// @Param 				spanId 										query	string	false	"Log spanId"
// @Param 				commit 										query	string	false	"Log commit"
// @Param 				metadateParentResourceId 	query	string	false	"Log metadata parentResourceId"
// @Param 				paginationPage 						query	integer	false	"Pagination page value"
// @Param 				paginationCount 					query	integer	false	"Pagination count value"
// @Produce 			json
// @Success 			200							{object}	[]CommonResponse
// @Router 				/search/filter	[get]
func Filter(c *fiber.Ctx) error

// @Summary 			Perform ranked search
// @Description 	Perform a ranked search across multiple fields
// @Tags 					Search
// @Param 				query											query	string	false	"Log multiline query string"
// @Param 				timestampStart						query	string	false	"Timerange start value"
// @Param 				timestampEnd							query	string	false	"Timerange end value"
// @Param 				paginationPage						query	integer	false	"Pagination page value"
// @Param 				paginationCount						query	integer	false	"Pagination count value"
// @Produce 			json
// @Success 			200							{object}	[]CommonResponse
// @Router 				/search/rank		[get]
func Rank(c *fiber.Ctx) error
