package routes

import (
	"grinder/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ar *AppRouter) GetUsersList(c *gin.Context) {
	persons, personsRights, err := ar.personsRepo.LoadPersons(0)
	if err != nil {
		logger.Error("error load persons", err)
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "persons": persons, "person_rights": personsRights})
}

func (ar *AppRouter) GetUsersRoles(c *gin.Context) {
	rightsMap := ar.personsRepo.GetRightsMap()
	c.JSON(http.StatusOK, gin.H{"ok": true, "rights": rightsMap})
}
