package routes

import "github.com/gin-gonic/gin"

func (ar *AppRouter) GetUsersList(c *gin.Context) {
	ar.personsRepo.LoadPersons()
}
