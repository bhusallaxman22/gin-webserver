package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vsouza/go-gin-boilerplate/config"
	"github.com/vsouza/go-gin-boilerplate/controllers"
)

func Test(t *testing.T) {
	TestingT(t)
}

var _ = Suite(&UserSuite{})

type UserSuite struct {
	config *viper.Viper
	router *gin.Engine
}

func (s *UserSuite) SetUpTest(c *C) {
	config.Init("test")
	s.config = config.GetConfig()
	s.router = SetupRouter()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	health := new(controllers.HealthController)
	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
		}
	}
	return router
}

func TestMain(m *testing.M) {
	SetupRouter()
}

func (s *UserSuite) TestRetrieve(c *C) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/user/1", nil)
	s.router.ServeHTTP(w, req)
	c.Assert(w.Code, Equals, 200)
}
