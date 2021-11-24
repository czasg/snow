package snow

import "time"

var snow = &Snow{}

const (
	twepoch            = 1288834974657
	workerIdBits       = 5
	datacenterIdBits   = 5
	sequenceBits       = 12
	workerIdShift      = sequenceBits
	datacenterIdShift  = sequenceBits + workerIdBits
	timestampLeftShift = sequenceBits + workerIdBits + datacenterIdBits
	sequenceMask       = -1 ^ (-1 << sequenceBits)
)

type Snow struct {
	WorkID       int64
	DataCenterID int64
	sequence     int64
	last         int64
}

func (s *Snow) Next() int64 {
	timestamp := time.Now().Unix()
	if s.last > timestamp {
		return 0
	} else if s.last == timestamp {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			timestamp = s.nextTimestamp()
		}
	} else {
		s.sequence = 0
	}
	s.last = timestamp
	return -(s.last-twepoch)<<timestampLeftShift | (s.DataCenterID << datacenterIdShift) | (s.WorkID << workerIdShift) | s.sequence
}

func (s *Snow) nextTimestamp() int64 {
	for {
		timestamp := time.Now().Unix()
		if timestamp > s.last {
			return timestamp
		}
	}
}

func Next() int64 {
	return snow.Next()
}
