package radiot

/*
Метод для интерфейса сортировкиsort.Sort
*/
func (p Podcast) Len() int {
	return len(p.Episodes)
}

/*
Метод для интерфейса сортировкиsort.Sort
*/
func (p Podcast) Less(i, j int) bool {
	return p.Episodes[i].Time.Before(p.Episodes[j].Time)
}

/*
Метод для интерфейса сортировкиsort.Sort
*/
func (p Podcast) Swap(i, j int) {
	p.Episodes[i], p.Episodes[j] = p.Episodes[j], p.Episodes[i]
}
