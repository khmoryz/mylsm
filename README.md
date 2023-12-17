# mylsm
Simple implementation of LSM-Tree.

## Usage

### Put

```sh
$ go run ./...
>put key=foo
ok
```

### Get

```sh
$ go run ./...
>put key=foo
ok
>get key
foo true
```

### Delete

```sh
$ go run ./...    
>put key=foo
ok
>get key
foo true
>del key
ok
>get key
 false
```

## Test

```sh
$ go test ./... -v
```
