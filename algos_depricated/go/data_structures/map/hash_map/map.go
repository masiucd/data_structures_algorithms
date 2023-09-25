package hash_map

type Pair struct {
	key string
	val string
}

type HashMap struct {
	size     int
	capacity int
	box      []*Pair
}

func newHashMap(capacity int) *HashMap {
	return &HashMap{
		size:     0,
		capacity: capacity,
		box:      make([]*Pair, capacity),
	}
}

func (hm *HashMap) hash(key string) int {
	var index = 0
	for i := 0; i < len(key); i++ {
		index += int(key[i])
	}
	return index % hm.capacity
}

func (hm *HashMap) get(key string) string {
	var index = hm.hash(key)
	for hm.box[index] != nil {
		if hm.box[index].key == key {
			return hm.box[index].val
		}
		index += 1
		index = index % hm.capacity
	}
	return ""
}

func (hm *HashMap) put(key string, val string) {
	index := hm.hash(key)

	for {
		if hm.box[index] == nil {
			hm.box[index] = &Pair{key, val}
			hm.size += 1
			if hm.size >= hm.capacity/2 {
				hm.rehash()
			}
			return
		} else if hm.box[index].key == key {
			hm.box[index].val = val
			return
		}
		index += 1
		index = index % hm.capacity
	}
}

func (hm *HashMap) rehash() {
	hm.capacity = 2 * hm.capacity
	newBox := make([]*Pair, hm.capacity)
	oldMap := hm.box
	hm.box = newBox
	hm.size = 0

	for _, p := range oldMap {
		if p != nil {
			hm.put(p.key, p.val)
		}
	}
}

func (hm *HashMap) remove(key string) {
	if hm.get(key) == "" {
		return
	}
	index := hm.hash(key)
	for {
		if hm.box[index].key == key {
			// Removing an element using open-addressing actually causes a bug,
			// because we may create a hole in the list, and our get() may
			// stop searching early when it reaches hm hole.
			hm.box[index] = nil
			hm.size -= 1
			return
		}
		index += 1
		index = index % hm.capacity
	}
}
