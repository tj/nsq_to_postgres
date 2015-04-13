
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

 A configuration file _must_ be specified via `--config`, I prefer making the path explicit so no one is left guessing of its whereabouts (unlike most programs, grr!).

### Postgres

 Two sections are available for tweaking, first the `postgres` section which defines the connection information, the target table name, target column name, and verbosity. For example:

```yml
postgres:
  connection: user=tj password=sloth host=localhost port=5432 sslmode=disable
  table: logs
  column: log
  max_open_connections: 10
  verbose: no
```

  When nsq_to_postgres first establishes a connect the table will be automatically created for you, if you have not already done so.

### NSQ

 The next section available is `nsq` which defines the topic to consume from, the nsqd or nsqlookupd addresses, max number of retry attempts and so on.

 For more nsq configuration options visit [segmentio/go-queue](https://github.com/segmentio/go-queue).

```yml
nsq:
  topic: logs
  nsqd: localhost:4150
  max_attempts: 5
  msg_timeout: 15s
  max_in_flight: 300
  concurrency: 50
```

## Development

Boot postgres, nsqd, and nsqlookupd (requires docker and docker-compose):

```
$ make
```

Connect via `psql`:

```
$ make client
```

# License

MIT