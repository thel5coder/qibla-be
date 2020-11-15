package odoohelper

import (
	"qibla-backend/usecase/viewmodel"
	"sort"
)

type ByPrice []viewmodel.OdooMasterPackageRoomRate

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Less(i, j int) bool { return a[i].Price < a[j].Price }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func SortByPrice(tourPackagePrices []viewmodel.OdooMasterPackageRoomRate) (res []viewmodel.OdooMasterPackageRoomRate) {
	sort.Sort(ByPrice(tourPackagePrices))
	res = tourPackagePrices

	return res
}
