package types

import (
	"sort"
	"time"
)

// VouchersList is a list of vouchers
type VouchersList []CompleteVoucher

// Revert filter a list of ticket vouchers
func (list VouchersList) Revert() []CompleteVoucher {
	var ticketsAvailable []CompleteVoucher
	for _, voucher := range list {
		ticketsAvailable = append(ticketsAvailable, voucher)
	}
	return ticketsAvailable
}

// Before filter a list of ticket vouchers
func (list VouchersList) Before(day time.Time) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.CreationDate.Time().Before(day) {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// FilterByDate filter a list of ticket vouchers
func (list VouchersList) FilterByDate(day time.Time) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Day() == day {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// FilterDrinks filter a list of ticket vouchers
func (list VouchersList) FilterDrinks() VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Product.Type != "TICKET" {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// Size filter a list of ticket vouchers
func (list VouchersList) Size() int {
	return len(list)
}

// FilterByCategory filter a list of ticket vouchers
func (list VouchersList) FilterByCategory(tipo string) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.PartyProduct != nil {
			if voucher.PartyProduct.Category != nil {
				if *voucher.PartyProduct.Category == tipo {
					ticketsAvailable = append(ticketsAvailable, voucher)
				}
			}
		}
	}
	return ticketsAvailable
}

// ExcludeErrors filter a list of ticket vouchers
func (list VouchersList) ExcludeErrors() VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Status != "ERROR" {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// ExcludeCanceled filter a list of ticket vouchers
func (list VouchersList) ExcludeCanceled() VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Status != "CANCELED" {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// ExcludePending filter a list of ticket vouchers
func (list VouchersList) ExcludePending() VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Status != "PENDING" {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// ExcludeCanceled filter a list of ticket vouchers
func (list VouchersList) ExcludeProcessing() VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Status != "PROCESSING" {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// ExcludeWeirdos TODO: NEEDS COMMENT INFO
func (list VouchersList) ExcludeWeirdos() VouchersList {
	return list.ExcludeCanceled().ExcludeErrors().ExcludeProcessing().ExcludePending()
}

// FilterByStatus filter a list of ticket vouchers
func (list VouchersList) FilterByStatus(status string) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Status == status {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// FilterByStatuses filter a list of ticket vouchers
func (list VouchersList) FilterByStatuses(status ...string) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		for _, statu := range status {
			if voucher.Status == statu {
				ticketsAvailable = append(ticketsAvailable, voucher)
			}
		}
	}
	return ticketsAvailable
}

// FilterByNotStatuses filter a list of ticket vouchers
func (list VouchersList) FilterByNotStatuses(status ...string) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		for _, statu := range status {
			if voucher.Status != statu {
				ticketsAvailable = append(ticketsAvailable, voucher)
			}
		}
	}
	return ticketsAvailable
}

// FilterByNotStatus filter a list of ticket vouchers
func (list VouchersList) FilterByNotStatus(status string) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Status != status {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// FilterByName filter a list of ticket vouchers
func (list VouchersList) FilterByName(name string) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Product.Name == name {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// Accountable filter a list of ticket vouchers
func (list VouchersList) Accountable() VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Type == "NORMAL" || voucher.Type == "PROMOTION" || voucher.Type == "TRANSFERED" {
			if voucher.Status != "CANCELED" && voucher.Status != "ERROR" && voucher.Status != "TRANSFERED" && voucher.Status != "PROCESSING" {
				ticketsAvailable = append(ticketsAvailable, voucher)
			}
		}
	}
	return ticketsAvailable
}

// FilterByType filter a list of ticket vouchers
func (list VouchersList) FilterByType(tipo string) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Type == tipo {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// FilterByType filter a list of ticket vouchers
func (list VouchersList) FilterByTypes(tipos ...string) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		for _, typ := range tipos {
			if voucher.Type == typ {
				ticketsAvailable = append(ticketsAvailable, voucher)
			}
		}
	}
	return ticketsAvailable
}

// FilterTickets filter a list of ticket vouchers
func (list VouchersList) FilterTickets() VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Product.Type == "TICKET" {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// TimeSlice is a sslice of time
type TimeSlice []time.Time

// Less TODO: NEEDS COMMENT INFO
func (s TimeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }

// Swap TODO: NEEDS COMMENT INFO
func (s TimeSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Len TODO: NEEDS COMMENT INFO
func (s TimeSlice) Len() int { return len(s) }

// ConsumedDates is awesome
func (list VouchersList) ConsumedDates() TimeSlice {
	var ticketsAvailable []time.Time
	for _, voucher := range list {
		if voucher.Status != "CANCELED" {
			ticketsAvailable = sameDate(ticketsAvailable, voucher.Day())
		}
	}
	hey := TimeSlice(ticketsAvailable)
	sort.Sort(hey)
	return hey
}

// LiquidSumTickets soma todos os vouchers de uma lista
func (list VouchersList) LiquidSumTickets(percent float64, assume bool) float64 {
	list = list.FilterTickets()
	var ticketsAvailable float64
	for _, voucher := range list {
		ticketsAvailable = ticketsAvailable + float64(voucher.Price.Value)
	}
	var ratio float64
	ratio = ticketsAvailable * (percent / 100)
	if !assume {
		ratio = ticketsAvailable / (1 + (1 - (percent / 100)))
	}
	return ratio
}

// LiquidSumDrinks soma todos os vouchers de uma lista
func (list VouchersList) LiquidSumDrinks(percent float64) float64 {
	list = list.FilterDrinks()
	var ticketsAvailable float64
	for _, voucher := range list {
		ticketsAvailable = ticketsAvailable + float64(voucher.Price.Value)
	}
	return ticketsAvailable * (percent / 100)
}

// Sum soma todos os vouchers de uma lista
func (list VouchersList) Sum() float64 {
	var ticketsAvailable float64
	for _, voucher := range list {
		ticketsAvailable = ticketsAvailable + float64(voucher.Price.Value)
	}
	return ticketsAvailable
}

// FilterByCustomerMail filter a list of ticket vouchers
func (list VouchersList) FilterByCustomerMail(mail string) VouchersList {
	var ticketsAvailable VouchersList
	for _, voucher := range list {
		if voucher.Customer.Mail == mail {
			ticketsAvailable = append(ticketsAvailable, voucher)
		}
	}
	return ticketsAvailable
}

// AvailableCustomers sdkfjn
func (list VouchersList) AvailableCustomers() []string {
	var ticketsAvailable []string
	for _, name := range list {
		if name.Customer != nil {
			ticketsAvailable = appendIfMissing(ticketsAvailable, name.Customer.Mail)
		}
	}
	return ticketsAvailable
}

// AvailableTickets sdkfjn
func (list VouchersList) AvailableTickets() []string {
	var ticketsAvailable []string
	for _, name := range list {
		if name.Product.Type == "TICKET" {
			ticketsAvailable = appendIfMissing(ticketsAvailable, name.Product.Name)
		}
	}
	return ticketsAvailable
}

// AvailableDrinks sdkfjn
func (list VouchersList) AvailableDrinks() []string {
	var ticketsAvailable []string
	for _, name := range list {
		if name.Product.Type != "TICKET" {
			ticketsAvailable = appendIfMissing(ticketsAvailable, name.Product.Name)
		}
	}
	return ticketsAvailable
}

// AvailableFromCategory sdfkjn
func (list VouchersList) AvailableFromCategory(cat string) []string {
	var ConsumedBeers []string
	for _, name := range list {
		if name.PartyProduct != nil {
			if name.PartyProduct.Category != nil {
				if *name.PartyProduct.Category == cat {
					ConsumedBeers = appendIfMissing(ConsumedBeers, name.Product.Name)
				}
			}
		}
	}
	return ConsumedBeers
}

var appendIfMissing = func(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

var sameDate = func(slice []time.Time, i time.Time) []time.Time {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

// amount : 0
// amountTotal : 0
// countFree : 0
// countPromotion : 0
// countSold : 0
// countTotal : 0
// countUsed : 0
// currency : "BRL"
// productName : "Cortesia Feminino (válido até 1h)"
type DashSoftVoucherSummary struct {
	Amount         float64 `json:"amount" bson:"amount"`
	AmountTotal    float64 `json:"amountTotal" bson:"amountTotal"`
	CountFree      int     `json:"countFree" bson:"countFree"`
	CountPromotion int     `json:"countPromotion" bson:"countPromotion"`
	CountSold      int     `json:"countSold" bson:"countSold"`
	CountTotal     int     `json:"countTotal" bson:"countTotal"`
	CountUsed      int     `json:"countUsed" bson:"countUsed"`
	Currency       string  `json:"currency" bson:"currency"`
	ProductName    string  `json:"productName" bson:"productName"`
}

// DrinksSummary filter a list of ticket vouchers
func (list VouchersList) DrinksSummary() []DashSoftVoucherSummary {
	ticketsAvailable := []DashSoftVoucherSummary{}
	for _, name := range list.AvailableDrinks() {
		kind := list.FilterByName(name).ExcludeWeirdos()
		instance := DashSoftVoucherSummary{
			Amount:         kind.Sum(),
			AmountTotal:    kind.Sum(),
			CountFree:      kind.FilterByTypes("FREE", "ANNIVERSARY").Size(),
			CountPromotion: kind.FilterByType("PROMOTION").Size(),
			CountSold:      kind.FilterByType("NORMAL").Size(),
			CountTotal:     kind.Size(),
			CountUsed:      kind.FilterByStatus("USED").Size(),
			Currency:       "BRL",
			ProductName:    name,
		}
		ticketsAvailable = append(ticketsAvailable, instance)
	}
	return ticketsAvailable
}

// TicketsSummary filter a list of ticket vouchers
func (list VouchersList) TicketsSummary() []DashSoftVoucherSummary {
	ticketsAvailable := []DashSoftVoucherSummary{}
	for _, name := range list.AvailableTickets() {
		kind := list.FilterByName(name).ExcludeWeirdos()
		instance := DashSoftVoucherSummary{
			Amount:         kind.Sum(),
			AmountTotal:    kind.Sum(),
			CountFree:      kind.FilterByTypes("FREE", "ANNIVERSARY").Size(),
			CountPromotion: kind.FilterByType("PROMOTION").Size(),
			CountSold:      kind.FilterByType("NORMAL").Size(),
			CountTotal:     kind.Size(),
			CountUsed:      kind.FilterByStatus("USED").Size(),
			Currency:       "BRL",
			ProductName:    name,
		}
		ticketsAvailable = append(ticketsAvailable, instance)
	}
	return ticketsAvailable
}
