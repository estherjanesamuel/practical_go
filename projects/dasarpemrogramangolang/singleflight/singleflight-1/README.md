# Single Flight API

```bash
curl -s -w "elapsed: %{time_total}s\n" -i -X GET "http://localhost:8000/api/report/download?reportid=sales-2021-10"
```

```sh
curl -G --data-urlencode 'comment=this cookbook is awesome' https://catonmat.net
```