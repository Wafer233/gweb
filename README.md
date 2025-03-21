# gweb

## Preface
This project is a simple Go web framework based on net/http, implementing some basic functionalities similar to those in `gin`.

## Why gweb
As we all know, `net/http` provides fundamental web functionalities, such as listening on ports, mapping static routes, and parsing HTTP messages. However, directly using net/http in actual development can be quite cumbersome. 

Therefore, I have implemented several essential features, including creating a gweb instance, adding routes, and finally starting the web service. Through this process, I have also gained a deeper understanding of `gin` and other web frameworks.


## Run gweb
Put the following code inside of an empty `.go` file.
```
package main

import (
	"github.com/Wafer233/gweb"
	"net/http"
)

func main() {
	route := gweb.Default()
	route.POST("/", func(ctx *gweb.Context) {
		ctx.SendTextResponse(http.StatusOK, "Hello Wafer")
	})
	route.Run("localhost:8080")
}
```
Then visit `localhost:8080/` in your browser to see the response!

## Feature


### Static Route
#### Get Start
```
func main() {
	r := gweb.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.Run("localhost:8080")
}
```
#### Explanation
1. Use `New()` to create an instance of gweb.
2. Add a route using the `GET()` method.
3. Start the web service with `Run()`.
4. Use `POST()` is optional.


#### Postman
![](./img/1.png)


### Contex
#### Get Start
```
func main() {
	r := gweb.New()

	r.GET("/", func(c *Context) {
		c.SendHTMLResponse(http.StatusOK, "<h1> Hello Gweb </h1>")
	})

	r.Run("localhost:8080")

}
```



#### Explanation
1. Use `New()` to create an instance of gweb.
2. Add a route using the `GET()` method.
3. Send a response using `SendHTMLResponse()`.
4. Start the web service with `Run()`.
5. `GetFormValue()` `GetQueryParam()` `SendTextResponse()` `SendJSONResponse()` are optional

#### Postman
![](./img/2.png)

### Dynamic Route

#### Get start
#### Explanation
#### Postman

### Group
#### Get start
#### Explanation
#### Postman
### Middleware
#### Get start
#### Explanation
#### Postman

### Recovery
#### Get start
#### Explanation
#### Postman













