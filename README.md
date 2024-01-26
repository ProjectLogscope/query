## Filter Endpoint:
The Filter endpoint is not your everyday search—this is a precision-engineered, high-performance log hunter. Imagine it as a sleek, ninja cat gracefully sifting through vast log data with the agility of a seasoned developer dodging code bugs.

### Purpose:
The primary objective of the Filter endpoint is to empower users to perform a surgical search across the log universe. You're not just looking for any logs; you're on a quest for specific ones, armed with a variety of search parameters.

### Parameters:
- __Level__: Specify the log level you're interested in, whether it's a calm "info," a warning "warning," or the ever-dreaded "error." It's like choosing your log adventure difficulty level.

- __Message__: Enter keywords to uncover logs with messages like hidden treasure chests. "Failed to connect," "Database connection lost," or "Unexpected EOF while parsing JSON" are your secret passcodes.

- __Resource ID__: Unearth logs specific to a resource. Your servers have unique IDs, and this parameter is your map to the exact spot in the log jungle.

- __Timestamp Range__: Define a time window for your search expedition. Need logs from a quiet Sunday morning or the chaotic Friday rush? We've got you covered.

- __Trace and Span IDs__: Trace your logs' lineage with trace and span IDs. It's like a genealogy chart for your logs. Find logs related to a specific operation and understand their ancestry.

- __Commit__: Ever wondered what was happening in your logs during a specific commit? The commit parameter reveals the log secrets tied to a particular commit hash.

- __Metadata Parent Resource ID__: Dive deeper into the metadata rabbit hole. Sometimes, logs have parents too, and this parameter lets you trace logs born from a specific parent.

- __Pagination__: Nobody likes an overwhelming scroll. Paginate through your results with paginationPage and paginationCount to keep things manageable.

### Advanced Features (Bonus):
- __Date Range Filtering__: Feeling nostalgic for logs from a bygone era? Use this feature to filter logs between specific date ranges. It's like time travel for your logs.

- __Regular Expressions__: For the regex virtuosos, unleash your pattern-matching sorcery. Craft intricate regex queries and watch logs surrender to your mighty expressions.

- __Combining Multiple Filters__: Why settle for one filter when you can have them all? Combine multiple filters to create powerful, nuanced queries. It's like crafting a potion with just the right ingredients.

- __Real-time Log Magic__: Experience logs unfolding in real-time. Witness the magic of logs popping into existence and disappearing into the void as your system updates in real-time.

- __Role-Based Access__: Keep the log kingdom secure with role-based access. Ensure that only the chosen ones have the privilege to summon and command the logs.

### Response:
The Filter endpoint doesn't just find logs; it returns a treasure trove—a JSON array of CommonResponse objects. Each object is a log gem waiting to be explored, revealing details about its level, message, resource ID, timestamp, trace ID, span ID, commit, and metadata.

### Success:
When your search is successful, the HTTP status code 200 acknowledges your log-hunting prowess. Your results are ready to be analyzed, and your developer instincts validated.


## Rank Endpoint:
Now, let's turn our attention to the Rank endpoint—an endpoint that doesn't just search logs; it ranks them, giving you a leaderboard of logs. It's like turning your logs into a competitive sport.

### Purpose:
The Rank endpoint isn't for the faint of heart. It's for those who want more than just logs; they want insights, patterns, and a touch of drama. With a multiline query string, it's like writing the script for your log opera.

### Parameters:
- __Multiline Query__: This is where the magic happens. Craft a multiline query string that tells a story. Whether it's an epic saga or a short story, this parameter transforms your logs into a narrative.

- __Timestamp Range__: Set the stage for your log drama with a specific time range. Logs from a weekend escapade or a Monday morning crisis—choose your log setting.

- __Pagination__: Sometimes, you want to savor your log insights in bite-sized pieces. paginationPage and paginationCount are your companions on this journey.

## Advanced Features (Bonus):
Real-time Drama: Live log drama unfolds before your eyes. As logs stream in, the ranking adapts, giving you a front-row seat to the evolving log storyline.

### Role-Based Access:
Not everyone gets a backstage pass to the log opera. With role-based access, ensure that only the privileged few can witness the log drama unfold.

### Response:
The Rank endpoint doesn't just return logs; it presents a spectacle—a JSON array of CommonResponse objects, each log vying for attention on the leaderboard. It's not just about finding logs; it's about discovering the star logs that steal the show.

### Success:
When your ranked search is successful, the HTTP status code 200 is the standing ovation. Your log leaderboard is ready, and you're the director of this log performance, shaping the narrative with each query.
