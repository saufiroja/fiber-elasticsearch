## Searching with Elasticsearch

### Searching match

```
  {
		"query": {
			"match": {
				"first_name": "john"
			}
		}
	}
```

### Searching firtname OR lastname

```
  {
		"query": {
			"bool": {
				"should": [
					{
						"match": {
							"first_name": "hallo"
						}
					},
					{
						"match": {
							"last_name": "world"
						}
					}
				]
			}
		}
	}
```

### Searching firstname AND age

```
  {
		"query": {
			"bool": {
				"must": [
					{
						"match": {
							"first_name": "halo"
						}
					},
					{
						"match": {
							"age": 10
						}
					}
				]
			}
		}
	}
```

## Searching age range

```
  {
		"query": {
			"range": {
				"age": {
					"gte": 10,
					"lte": 20
				}
			}
		}
	}
```

### Searching match pharse query

```
  {
		"query": {
			"match_phrase": {
				"hobbies": "basket"
			}
		}
	}

  {
		"query": {
			"match_phrase": {
				"hobbies": {
					"query": "basket",
					"analyzer": "standard"
				}
			}
		}
	}
```
