package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if score, ok := units[unit]; ok {
		bill[item] += score
		return true
	} else {
		return false
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	}

	quantity, ok := bill[item]
	if !ok {
		return false
	}
	switch {
	case quantity < units[unit]:
		return false
	case quantity == 0 || quantity == units[unit]:
		delete(bill, item)
		return true
	default:
		bill[item] -= units[unit]
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if quantity, ok := bill[item]; !ok {
		return 0, false
	} else {
		return quantity, true
	}
}
