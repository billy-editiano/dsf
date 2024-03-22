# DSF
> Datasaur Fetcher

## `cmd/fetcher`
Run a cron to fetch Datasaur instance healthcheck (currently only support every minute).

It will store the result in `$HOME/dsfetcher/result.json`.

### Building

Pre-requisites
- go 1.21+
- make
```shell
make build -s app=fetcher
```

The binary file will be available under `bin` folder.

### Configuration
Create a file under `$HOME/.config/dsfetcher/config.json`, here is the format example:
```json
{
  "mydsinstance": "https://your-datasaur-installation.com/api/health"
}
```
