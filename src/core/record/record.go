package record

import (
	"core/contract"
	"strconv"
)

// 用于抽象记录的存储操作
type Record interface {
	CountKey() string

	// 获取记录数量
	GetCount() uint32
}

func StoreCount(record Record, context contract.Context) error {
	return context.PutState(record.CountKey(), []byte(strconv.FormatUint(uint64(
		record.GetCount()), 10)))
}

func GetRecordCount(key string, context contract.Context) (uint32, error) {
	countBuf, err := context.GetState(key)
	if err != nil {
		return 0, err
	}

	// 未找到则从0开始
	if countBuf == nil {
		return 0, nil
	}

	count, err := strconv.ParseUint(string(countBuf), 10, 32)
	return uint32(count), err
}
