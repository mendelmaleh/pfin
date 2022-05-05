# pfin: personal finance tools

Collection of tools for parsing bank statements and csvs.

[![Go Documentation](https://godocs.io/git.sr.ht/~mendelmaleh/pfin?status.svg)](https://godocs.io/git.sr.ht/~mendelmaleh/pfin)

# structure

- `pfin`: base library, defines transaction and parser interfaces, parser loading utility like database/sql
- `pfin/util`: Utils for directory parsing, filter flags, and formatting functions
- `pfin/parser`: parser implementations
- `pfin/parser/util`: common code for parsers
- `pfin/parser/all`: metapackage to load all parser implementations

- `cmd/main`: prints out all transactions, can filter by account/user and sums by user
- `cmd/unpaid`: get a list of unpaid transactions for a user
- `cmd/status`: wip, get an overview of the statements tree
