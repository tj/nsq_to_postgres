
## nsq_to_postgres

 Pull messages from NSQ and write to a Postgres JSONB column.

## Example

Run:

```
$ nsq_to_postgres --config config.yml
```

Query:

```sql
select
  log->'program' as program,
  log->'level' as level,
  count(*)
from logs
  group by program, level
  order by count desc;
```

Result:

```
 program |  level  | count
---------+---------+--------
 "site"  | "info"  | 142276
 "api"   | "info"  |   8176
 "api"   | "error" |   1638
```

## Configuration

 TODO:

## Development

Boot postgres, nsqd, and nsqlookupd (requires docker and docker-compose):

```
$ make
```

Connect via `psql`:

```
$ make client
```

