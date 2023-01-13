package store_test

import (
	"int_test/store"
	"sort"
	"testing"
)

//length: 100
var UnsortedData = []store.HashData{
	{"", 7781}, {"", 3102}, {"", 7055}, {"", 4022}, {"", 3136}, {"", 1086}, {"", 2966},
	{"", 2932}, {"", 2921}, {"", 9014}, {"", 1110}, {"", 2542}, {"", 6338}, {"", 8507},
	{"", 9373}, {"", 6631}, {"", 7679}, {"", 2895}, {"", 1626}, {"", 2338}, {"", 7914},
	{"", 8106}, {"", 1289}, {"", 5319}, {"", 6211}, {"", 7558}, {"", 2534}, {"", 9726},
	{"", 6336}, {"", 7009}, {"", 4035}, {"", 1443}, {"", 3527}, {"", 6920}, {"", 5132},
	{"", 6091}, {"", 1220}, {"", 7401}, {"", 6623}, {"", 9454}, {"", 5145}, {"", 9173},
	{"", 3936}, {"", 9655}, {"", 4817}, {"", 5662}, {"", 5435}, {"", 8546}, {"", 6127},
	{"", 3550}, {"", 8055}, {"", 4498}, {"", 1400}, {"", 7482}, {"", 6997}, {"", 9973},
	{"", 2080}, {"", 6186}, {"", 8948}, {"", 8438}, {"", 0244}, {"", 8929}, {"", 3532},
	{"", 8555}, {"", 6112}, {"", 9027}, {"", 1664}, {"", 9178}, {"", 4970}, {"", 4666},
	{"", 0763}, {"", 1240}, {"", 4608}, {"", 6017}, {"", 2244}, {"", 6277}, {"", 8219},
	{"", 8365}, {"", 1700}, {"", 3898}, {"", 3044}, {"", 6283}, {"", 9256}, {"", 3543},
	{"", 6157}, {"", 3335}, {"", 3885}, {"", 6653}, {"", 4462}, {"", 2988}, {"", 9532},
	{"", 8259}, {"", 8774}, {"", 7266}, {"", 3271}, {"", 9345}, {"", 7209}, {"", 8461},
	{"", 7941}, {"", 8585},
}

func TestInsertSort(t *testing.T) {
	t.Parallel()
	hashStore := store.NewHashStore()
	updateChan := make(chan store.HashData, 10)
	go hashStore.ReceiveUpdate(updateChan)

	for _, d := range UnsortedData {
		updateChan <- d
	}

	for len(hashStore.Hashes) != 100 {
		continue
	}

	newSlc := make([]int, 0)
	for _, i := range hashStore.Hashes {
		newSlc = append(newSlc, int(i.Sum))
	}

	if !sort.IntsAreSorted(newSlc) {
		t.Errorf("hashed data are not sorted")
	}
}
