## Ip2Location Dashboard

Tool to analyze IP addresses and visualize the results graphically with statistics. To make decisions, useful for public servers with multiple connections.

![](https://img.shields.io/github/issues/LuzLeslie/Ip2Location-Dashboard)
![](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)
![](https://img.shields.io/badge/Go-%5E1.13-blue)
![](https://img.shields.io/badge/Vue-2.x-blue)
![](https://img.shields.io/badge/NodeJS-12.x-blue)

---

![Fronted Analyze ip](https://i.ibb.co/N9vvPps/dashboard.png)

![Backend Analyze ip](https://i.ibb.co/M565631/graphql-Post.png)

### Features

---

* Request only the data you require using **graphql**
* Real-time analysis activities
* Map of ip origins
* Proxy graph
* Number of IPs analyzed

### Tech/framework used

- [Graphql](https://graphql.org)
- [Vue.js](https://vuejs.org)

### Prerequisites

---
- [Go](https://golang.org)
- [IP2Location - IP2Proxy(optional) - Database .bin](https://ip2location.com/)
- [NodeJS](https://nodejs.org)

### Installation
---
:rocket: Get the latest functions directly from the specified branch.

##### Backend 

```bash
git clone --depth=1 -b backend https://github.com/LuzLeslie/Ip2Location-Dashboard.git ip2Dashboard-backend
cd ip2Dashboard-backend
go run ./cmd/ipdebugger -l /path/ip2location.bin -p /path/ip2proxy.bin
```
> More details in  [![](https://img.shields.io/badge/branch-backend-orange?style=flat)](https://github.com/LuzLeslie/Ip2Location-Dashboard/tree/backend)

##### Fronted
```bash
git clone --depth=1 -b backend https://github.com/LuzLeslie/Ip2Location-Dashboard.git ip2Dashboard-fronted
cd ip2Dashboard-fronted
yarn install
yarn serve
```
> More details in  [![](https://img.shields.io/badge/branch-fronted-orange?style=flat)](https://github.com/LuzLeslie/Ip2Location-Dashboard/tree/fronted)


### Usage
* Send the ip to analyze to the backend directly using curl or through its own tool.

```bash
curl --request POST \
  --url http://localhost:4000/graphql \
  --header 'content-type: application/json' \
  --data '{"query":"mutation{\n  analyze(ip:\"103.243.82.198\"){\n    country_code\n    country_name\n  }\n}"}'
```

* Once the web app is running, you can also analyze your ip and view its results graphically in real-time.
[http://localhost:3000](http://localhost:3000)

### Used tools
* https://lite.ip2location.com
* https://altair.sirmuel.design
* https://github.com/StartBootstrap/startbootstrap-sb-admin-2
* https://bootstrap-vue.org

## License

---

This project is licensed under the terms of the ![](https://badgen.net/github/license/LuzLeslie/Ip2Location-Dashboard)
