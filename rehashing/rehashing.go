package rehashing

type Rehash map[uint32]map[uint32]uint32

func (this Rehash) Increase(oldBucket uint32, newBucket uint32) {
	if _, ok := this[oldBucket]; !ok {
		this[oldBucket] = make(map[uint32]uint32)
	}
	this[oldBucket][newBucket]++
}

var NumOfBuckets uint32 = 11

type Total struct {
	TotalMoved   uint64
	TotalStayied uint64
	Total        uint64
}
