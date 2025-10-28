# library-auth-service

## Setting up the Library Auth Service

Either connect a local PostgreSQL database and setup according to .env configs.
Or just run a docker compose:

```bash
docker-compose up -d
```

Install the migration tool if needed:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Run the migrations:
```bash
make migrate-up
```

NB! If you use Windows and do not have `make` here is an option:

Run PowerShell as administrator, set execution policy to unrestricted:
```powershell
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser
```
restart PowerShell and install Scoop:
```powershell
irm get.scoop.sh | iex
```
install `make`:
```powershell
scoop install make
```
then run in project:
```powershell
make migrate-up
```

## Check the Swagger UI
Open in your browser:
```web
http://localhost:8284/swagger/index.html
```