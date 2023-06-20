package main

type user struct {
    ID int64 `json:"id"`
    Email string `json:"email"`
    Username string `json:"username"`
}

type updateUserDAO struct {
    Email string `json:"email"`
    Username string `json:"username"`
}

func toUser(id int64, userDAO updateUserDAO) user {
    var newUser user = user {
        ID: id,
        Email: userDAO.Email,
        Username: userDAO.Username,
    }

    return newUser
}