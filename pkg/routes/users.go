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

func (ar *AppRouter) UpdateUser(c *gin.Context) {
	var u UserUpdateRequest
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"ok": false, "error": "Invalid request"})
		return
	}
	ok, err := ar.personsRepo.UpdateUser(u.UserID, u.FirstName, u.LastName, u.Email, u.Active)
	if err != nil {
		logger.Error("error update single user", err)
		c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "server error"})
		return
	}
	if !ok {
		c.JSON(http.StatusOK, gin.H{"ok": true, "msg": "can't update this user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true, "msg": "user updated"})
}
