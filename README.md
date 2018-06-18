## Reviews

The following project consists of api for the following.

- [ ] Get app reviews
- [ ] Search apps
- [ ] Get app details


> At present the app supports, AppStore, Playstore support will be added soon.
> - [X] AppStore
> - [ ] PlayStore


### Dependencies
1. dep (go dependency manager)

> Note: cd into git clone directory

### Search apps

```go
go run main.go server -s -n Uber

go run main.go server -s -n "Google home"

# default -n whatsapp
go run main.go server -s 
```

### Get app details

```go
go run main.go server -a -i 368677368

go run main.go server -a -i 310633997

# default -i 368677368 
# default id: UBER
go run main.go server -a
```

### Get app reviews

```go
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

> The repo is not an official guide to get reviews, I've predominantly used some of the suggested methods in stackoverflow and web.
> If you have any questions mail me at nareshkumarganesan at **g**oogle e**mail** dot com