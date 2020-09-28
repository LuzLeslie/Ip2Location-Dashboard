

## Backend Ip2location Dashboard 

---

Analyze and obtain statistics, request only the data you require using **graphql**.

### Technologies

---

* [Graphql](https://graphql.org/)

### Requirements

---

- [Go](https://golang.org)
- [IP2Location and IP2Proxy Database](https://ip2location.com/)

### Setup

----

```bash
git clone --depth=1 -b backend https://github.com/LuzLeslie/Ip2Location-Dashboard.git
cd Ip2Location-Dashboard
go run ./cmd/ipdebugger/ -l /path/ip2location.bin -p /path/ip2proxy.bin
```

### Build for production

```bash
go build -v -o ip2 ./cmd/ipdebugger/
```

> run with `./ip2 -l /path/ip2location.bin -p /path/ip2proxy.bin` 
>
> graphql api server is started at: http://localhost:4000/graphql
>
> subscriptions api server is started at: http://localhost:4000/subscriptions

## How to use?

---

  ![Analyze ip](https://i.ibb.co/M565631/graphql-Post.png)

```bash
curl --request POST \
  --url http://localhost:4000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"mutation{\n  analyze(ip:\"103.243.82.198\"){\n    country_code\n    country_name\n  }\n}"}'
```



## License

---

This project is licensed under the terms of the **MIT** license.