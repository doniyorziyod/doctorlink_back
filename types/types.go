package types

import "github.com/google/uuid"

type User struct {
    Id uuid.UUID `json:"id" form:"id" query:"id"`
    Phone string `json:phone"" form:"phone" query:"phone"`
    Name string `json:name"" form:"name" query:"name"`
    Surname string `json:surname"" form:"surname" query:"surname"`
    Gender bool `json:gender"" form:"gender" query:"gender"`
    Birthday string `json:birthday"" form:"birthday" query:"birthday"`
    Subregion string `json:subregion"" form:"subregion" query:"subregion"`
    Region string `json:region"" form:"region" query:"region"`
}

type InputUser struct {
    Phone string `json:phone"" form:"phone" query:"phone"`
    Name string `json:name"" form:"name" query:"name"`
    Surname string `json:surname"" form:"surname" query:"surname"`
    Gender bool `json:gender"" form:"gender" query:"gender"`
    Birthday string `json:birthday"" form:"birthday" query:"birthday"`
    Subregion string `json:subregion"" form:"subregion" query:"subregion"`
    Region string `json:region"" form:"region" query:"region"`
}

type Sms struct {
    Phone string `json:phone"" form:"phone" query:"phone"`
    Sms int `json:sms"" form:"sms" query:"sms"`
}
