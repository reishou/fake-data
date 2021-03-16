# fake-data

Git clone & change directory to `fake-data`

```bash
git clone https://github.com/reishou/fake-data.git && cd fake-data
```

Run generate example csv

```bash
./fake-data -template=default -count=1000
```

Create a custom template yml `template/user.yml`

```yaml
---
schema:
  - uuid_digit
  - username
  - first_name
  - last_name
  - gender
  - date
  - password
  - sentence
  - lat
  - long
```

Generate file `csv/user.csv`

```bash
./fake-data -template=user -count=1000000
```