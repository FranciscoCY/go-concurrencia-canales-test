# go-concurrencia-canales-test
Proyecto donde se utiliza los conceptos concurrencia, canales y testing

## Tech Stack  
**Client:** Go, Golang  

**Server:** Golang  
 
## Run Locally  
Unit Test

~~~bash  
  go test
~~~

Coverage Unit test

~~~bash  
  go test -cover
~~~

Report coverage
~~~bash
 go test -coverprofile=coverage.out
~~~

Report coverage
~~~bash
 go tool cover -func=coverage.out
~~~

Report coverage html
~~~bash
 go tool cover -html=coverage.out
~~~

Report coverage - time
~~~bash
 go test -cpuprofile=cou.out
~~~

~~~bash
 go tool pprof cou.out
~~~