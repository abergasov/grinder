package repository

import (
	"grinder/pkg/config"
	"grinder/pkg/logger"
	"grinder/pkg/storage"
	"grinder/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type rights struct {
	UserID  int64 `db:"user_id"`
	RightID int64 `db:"right_id"`
}

type RightsManager struct {
	db *storage.DBConnector
}

func InitRightManager(cnf *config.AppConfig, db *storage.DBConnector) *RightsManager {
	return &RightsManager{
		db: db,
	}
}

func (r *RightsManager) CheckRight(rights []int64, ver func(*gin.Context) (int64, int64, bool)) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _, valid := ver(c)
		if !valid {
			c.JSON(http.StatusUnauthorized, gin.H{"ok": false, "error": "access denied"})
			return
		}

		var userRights []int64
		err := r.db.Client.Select(&userRights, "SELECT right_id FROM users_rights WHERE user_id = ?", userID)
		if err != nil {
			logger.Error("error get user rights", err)
			c.JSON(http.StatusInternalServerError, gin.H{"ok": false, "error": "server error"})
			return
		}

		if !utils.SliceHasIntersections(userRights, rights) {
			c.JSON(http.StatusForbidden, gin.H{"ok": false, "error": "access denied"})
			return
		}
		c.Next()
	}
}
