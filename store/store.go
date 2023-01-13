package store

type HashData struct {
	Hash string
	Sum  int64
}

type HashStore struct {
	Hashes []HashData
}

func NewHashStore() HashStore {
	return HashStore{
		Hashes: make([]HashData, 0),
	}
}

func (h *HashStore) ReceiveUpdate(hashDataChan chan HashData) {
	for hashData := range hashDataChan {
		h.Hashes = insertSort(h.Hashes, hashData)
	}
}

func insertSort(hs []HashData, data HashData) []HashData {
	size := len(hs)

	if size == 0 {
		return append(hs, data)
	}

	for i := 0; i < size; i++ {
		if data.Sum < hs[i].Sum {
			if i == 0 {
				return append([]HashData{data}, hs...)
			} else {
				prev := make([]HashData, i)
				after := make([]HashData, size-i)
				copy(prev, hs[:i])
				copy(after, hs[i:])
				prev = append(prev, data)

				return append(prev, after...)
			}
		}
	}

	return append(hs, data)
}
