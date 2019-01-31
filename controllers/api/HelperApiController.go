package api

import (
	"MPMS/services/db"
	"MPMS/structure"
)

type HelperApiController struct {
	Controller
}

func (h *HelperApiController) RefreshDBConPools() {
	if err := db.CheckAndRefreshCon(); err != nil {
		h.ApiReturn(structure.Response{Error: 1, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}
	h.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{}})
}
