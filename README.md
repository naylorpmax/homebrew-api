# homebrew-api

API to discover Homebrew D&D content for Patreon users of Power Attack Publishing.

## Local Setup

**Prerequisites**


**Start DB**

```bash
export POSTGRES_HOST=postgres
export POSTGRES_DB=postgres
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=<postgres-password>
make db-init

export DATA_PATH=data/spells.csv
export TABLE_NAME=spells
make db-write
```

**Start API**

```bash
make api
```

**Make requests**

- Navigate to `http://localhost:8080/` in browser
- Navigate to `http://localhost:8080/login` in browser
- Login to Patreon
- Make scripted requests for spells to the API

```
./scripts/sample-requests.sh
```

**Locally explore data**

```bash
make db-conn
```
