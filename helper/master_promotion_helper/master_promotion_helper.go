package master_promotion_helper

import "github.com/labstack/echo"

//parse filter statement
func ParseFilterStatement(filters map[string]interface{}) (res string){
	if val,ok := filters["name"];ok {
		res +=` and lower(name) like '%`+val.(string)+`%'`
	}

	if val,ok := filters["updated_at"];ok {
		res +=` and lower(cast(updated_at as varchar)) like '%`+val.(string)+`%'`
	}

	return res
}

//set filter params
func SetFilterParams(ctx echo.Context) map[string]interface{}{
	res := make(map[string]interface{})
	if ctx.QueryParam("name") != ""{
		res["name"] = ctx.QueryParam("name")
	}
	if ctx.QueryParam("updated_at") != ""{
		res["updated_at"] = ctx.QueryParam("updated_at")
	}

	return res
}
