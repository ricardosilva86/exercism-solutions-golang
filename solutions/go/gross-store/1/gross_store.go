package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := make(map[string]int)
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    for _ = range units {
        if _, ok := units[unit]; !ok {
            return false
        }
    }
    if _, ok := bill[item]; !ok {
        bill[item] = units[unit]
        return true
    }
    bill[item] += units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    for _ = range units {
        if _, ok := units[unit]; !ok {
            return false
        }
    }
    
	if _, ok := bill[item]; ok {
        if bill[item] - units[unit] == 0 {
            delete(bill, item)
            return true
        }
        if bill[item] - units[unit] < 0 { 
            return false
        } else {
            bill[item] -= units[unit]
            return true
        }
        delete(bill, item)
        return true
    }
    
    return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    if value, ok := bill[item]; ok {
        return value, true
    }
    
    return 0, false
	
}
