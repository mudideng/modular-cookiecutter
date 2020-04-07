package web

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	. "github.com/smartystreets/goconvey/convey"
)

func CreateServer(i *helloWorld) *httptest.Server {
	router := chi.NewRouter()
	router.Mount("/", (i).Routes())
	server := httptest.NewServer(router)
	return server
}

func TestMessageReturned(t *testing.T) {
	Convey("Given a running webserver", t, func() {
		server := CreateServer(&helloWorld{
			Message: "Test",
		})
		defer server.Close()

		Convey("When making a request to the root of the server", func() {
			req, err := http.NewRequest("GET", server.URL, nil)
			So(err, ShouldBeNil)
			res, err := http.DefaultClient.Do(req)
			So(err, ShouldBeNil)
			Convey("Then response status should be Ok and response the specified message", func() {
				So(res.StatusCode, ShouldEqual, http.StatusOK)
				responseBuffer, err := ioutil.ReadAll(res.Body)
				So(err, ShouldBeNil)
				So(string(responseBuffer), ShouldResemble, "Test")
			})
		})
	})
}
