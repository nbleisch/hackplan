package handler

import (
	"fmt"
	"github.com/hackdaysspring2017/hackplan/model"
	"github.com/kataras/iris"
)

var userrepo map[int]model.User

func init() {
	userrepo = map[int]model.User{}
	userrepo[1] = model.User{Id: 1, SureName: "Power", FirstName: "Max", IsAdmin: true, Email: "admin@hackplan.de", Password: "hackplan"}
}

func LoginHandler(ctx *iris.Context) {

	auth := &model.Auth{}

	if err := ctx.ReadJSON(&auth); err != nil {
		ctx.JSON(500, err.Error())
	} else {
		for _, value := range userrepo {
			if value.Email == auth.Email && value.Password == auth.Password {
				token := &model.Token{Token: "2138123ASDYXCASD"}
				ctx.JSON(200, token)
				return
			} else {
				ctx.JSON(403, "no found")
				return
			}
		}
	}
	ctx.JSON(403, "no found")

}

func UserGetHandler(ctx *iris.Context) {

	userID, err := ctx.ParamInt("id")
	if err == nil && userID != 0 {
		currentuser := userrepo[userID]
		ctx.JSON(iris.StatusOK, currentuser)
	} else {
		users := make([]model.User, 0)

		for _, value := range userrepo {
			users = append(users, value)
		}

		ctx.JSON(iris.StatusOK, map[string]interface{}{
			"users": &users,
		})
	}

}

func UserPostHandler(ctx *iris.Context) {

	user := &model.User{}
	if err := ctx.ReadJSON(&user); err != nil {
		ctx.JSON(500, err.Error())
	} else {
		userrepo[user.Id] = *user
	}

	ctx.JSON(200, userrepo)
}

func UserDeleteHandler(ctx *iris.Context) {

	userID, _ := ctx.ParamInt("id")
	delete(userrepo, userID)
	msg := fmt.Sprintf("deleted user with id %v", userID)
	ctx.JSON(200, msg)
}
