## Reviews

The following project consists of api for the following.

- [ ] Get app reviews
- [ ] Search apps
- [ ] Get app details


> At present the app supports, AppStore, the following will be added soon.
> - [X] AppStore
> - [ ] PlayStore

> Note: cd in git clone directory

### Search apps

```bash
go run main.go server -s -n Uber

go run main.go server -s -n "Google home"

# default -n whatsapp
go run main.go server -s 
```

### Get app details

```bash
go run main.go server -a -i 368677368

go run main.go server -a -i 310633997

# default -i 368677368 
# default id: UBER
go run main.go server -a
```

### Get app reviews

```bash
# id: UBER , country: US
go run main.go server -r -i 368677368 -c 143441
# id: whatsapp, country: US
go run main.go server -r -i 310633997 -c 143441
# id: whatsapp, country: US Page: 1
go run main.go server -r -i 310633997 -c 143441 -p 1
# default -i 368677368 -c 143441 -p 0
# default id: UBER , country: US, page: 0
go run main.go server -r
```