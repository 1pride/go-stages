Query parameters are the `?key=value` pairs in the URL. For example:

`q` and `page` are query parameters

```sql
GET
/search?q=queryName&page=2
```

### Diff between expressions

| Expression      | Contains...                     | Populated by...        |
|-----------------|---------------------------------|------------------------|
| `r.URL.Query()` | Only query params               | Always available       |
| `r.Form`        | Both query and POST form values | After `r.ParseForm()`  |
| `r.PostForm`    | Only POST form values           | After `r.ParseForm()`  |
| `r.FormValue()` | First value from `r.Form`       | Internally parses form |
