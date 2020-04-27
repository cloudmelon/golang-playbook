package main

type HTTPRequest struct {
	Method string
}

func main() {
	r := HTTPRequest{Method: "HEAD"}

	switch r.Method {
	case "GET":
		println("Get method")
	//	fallthrough
	case "POST":
		println("Post method")
	//	fallthrough
	case "PUT":
		println("Put method")
	//	fallthrough
	case "DELETE":
		println("Delete method")
	//	fallthrough
	default:
		println("Unhandled method")
	}

}
