# fake-data

Git clone & change directory to `fake-data`

```bash
git clone https://github.com/reishou/fake-data.git && cd fake-data
```

Run generate example csv

With *nix
```bash
./fake-data -template=default -count=1000
```

With windows
```bash
fake-data.exe -template=default -count=1000
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

With *nix
```bash
./fake-data -template=user -count=1000000
```

With windows
```bash
fake-data.exe -template=user -count=1000000
```

List all tags supported

```yaml
---
schema:
  - cc_number
  - cc_type
  - email
  - domain_name
  - ipv4
  - ipv6
  - password
  - jwt
  - phone_number
  - mac_address
  - url
  - username
  - toll_free_number
  - e_164_phone_number
  - title_male
  - title_female
  - first_name
  - first_name_male
  - first_name_female
  - last_name
  - name
  - gender
  - date
  - time
  - month_name
  - year
  - day_of_week
  - day_of_month
  - timestamp
  - century
  - timezone
  - time_period
  - word
  - sentence
  - paragraph
  - currency
  - amount_with_currency
  - uuid_digit
  - uuid_hyphenated
  - lat
  - long
  - unix_time
```