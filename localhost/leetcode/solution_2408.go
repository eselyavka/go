package solutions

type SQL struct {
	storage map[string][][]string
}

func Constructor(names []string, columns []int) SQL {
	sql := SQL{storage: make(map[string][][]string)}
	return sql
}

func (this *SQL) InsertRow(name string, row []string) {
	this.storage[name] = append(this.storage[name], row)
}

func (this *SQL) DeleteRow(name string, rowId int) {
	this.storage[name][rowId-1] = nil
}

func (this *SQL) SelectCell(name string, rowId int, columnId int) string {
	if this.storage[name][rowId-1] != nil {
		return this.storage[name][rowId-1][columnId-1]
	}
	return ""
}
