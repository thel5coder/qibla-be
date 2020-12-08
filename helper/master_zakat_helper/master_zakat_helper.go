package master_zakat_helper

import "github.com/labstack/echo"

func SetFilterParams(ctx echo.Context) map[string]interface{}{
	res := make(map[string]interface{})
	if ctx.QueryParam("type_zakat")!= ""{
		res["type_zakat"] = ctx.QueryParam("type_zakat")
	}

	if ctx.QueryParam("name")!= ""{
		res["name"] = ctx.QueryParam("name")
	}

	if ctx.QueryParam("description")!= ""{
		res["description"] = ctx.QueryParam("description")
	}

	if ctx.QueryParam("updated_at")!= ""{
		res["updated_at"] = ctx.QueryParam("updated_at")
	}

	return res
}
