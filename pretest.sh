cd sql/schema
goose postgres "postgres://daviddo:@localhost:5432/gator" down
goose postgres "postgres://daviddo:@localhost:5432/gator" up