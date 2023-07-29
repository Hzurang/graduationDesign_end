package utils

import (
	"fmt"
	"ginStudy/global"
	jsonIter "github.com/json-iterator/go"
	"go.uber.org/zap"
	"time"
)

func SetWithMarshal(redisKey string, data interface{}, expiration time.Duration) bool {
	if data == nil || len(redisKey) <= 0 {
		return false
	}
	//处理同步
	dataMarshal, err := jsonIter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		zap.L().Error("SetWithMarshal marshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
		return false
	}

	err = global.RD.Set(redisKey, dataMarshal, expiration).Err()
	if err != nil {
		zap.L().Error("SetWithMarshal redis set err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
	}

	return err == nil
}

func GetWithUnMarshal(redisKey string, outputData interface{}) (bool, error) {
	if outputData == nil || len(redisKey) <= 0 {
		return false, nil
	}

	redisData, err := global.RD.Get(redisKey).Result()

	if err != nil {
		return false, err
	}

	err = jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(redisData), outputData)
	if err != nil {
		zap.L().Error("GetWithUnMarshal Json Unmarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("redisData", redisData),
		)
		return false, err
	}

	return true, nil
}

func HSetWithMarshal(redisKey string, fieldName string, data interface{}) (bool, error) {
	if data == nil || len(redisKey) <= 0 || len(fieldName) <= 0 {
		return false, nil
	}

	dataMarshal, err := jsonIter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		zap.L().Error("HSetWithMarshal err",
			zap.String("Param", fmt.Sprintf("(%v,%v)", redisKey, fieldName)),
			zap.Error(err),
			zap.ByteString("dataMarshal", dataMarshal),
		)
		return false, err
	}

	err = global.RD.HSet(redisKey, fieldName, dataMarshal).Err()
	if err != nil {
		zap.L().Error("HSetWithMarshal set err",
			zap.String("Param", fmt.Sprintf("(%v,%v)", redisKey, fieldName)),
			zap.ByteString("dataMarshal", dataMarshal))
		return false, err
	}

	return true, nil
}

func HGetWithUnMarshal(redisKey string, fieldName string, outputData interface{}) (bool, error) {
	if outputData == nil || len(redisKey) <= 0 || len(fieldName) <= 0 {
		return false, nil
	}

	redisData, err := global.RD.HGet(redisKey, fieldName).Result()
	if err != nil {
		zap.L().Error("LoadRedis", zap.Error(err))
		return false, err
	}

	err = jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(redisData), outputData)
	if err != nil {
		zap.L().Error("HGetWithUnMarshal Json Unmarshal err",
			zap.String("param", fmt.Sprintf("(%v,%v)", redisKey, fieldName)),
			zap.Error(err),
			zap.String("redisData", redisData),
		)
		return false, err
	}

	return true, nil
}

func RPushWithMarshal(redisKey string, data interface{}) (bool, error) {
	if data == nil || len(redisKey) <= 0 {
		return false, nil
	}

	//处理同步
	dataMarshal, err := jsonIter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		zap.L().Error("RPushWithMarshal marshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
		return false, err
	}

	err = global.RD.RPush(redisKey, dataMarshal).Err()
	if err != nil {
		zap.L().Error("RPushWithMarshal redis set err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
		return false, err
	}

	return true, nil
}

func LPushWithMarshal(redisKey string, data interface{}) (bool, error) {
	if data == nil || len(redisKey) <= 0 {
		return false, nil
	}

	//处理同步
	dataMarshal, err := jsonIter.ConfigCompatibleWithStandardLibrary.Marshal(data)
	if err != nil {
		zap.L().Error("LPushWithMarshal marshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
		return false, err
	}

	err = global.RD.LPush(redisKey, dataMarshal).Err()
	if err != nil {
		zap.L().Error("LPushWithMarshal redis set err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.ByteString("userData", dataMarshal),
		)
		return false, err
	}

	return true, nil
}

func LPopWithUnMarshal(redisKey string, outputData interface{}) (bool, error) {
	if outputData == nil || len(redisKey) <= 0 {
		return false, nil
	}

	redisData, err := global.RD.LPop(redisKey).Result()
	if err == nil {
		err = jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(redisData), outputData)
		if err != nil {
			zap.L().Error("LPopWithUnMarshal Json Unmarshal err",
				zap.String("param", fmt.Sprintf("(%v)", redisKey)),
				zap.Error(err),
				zap.String("redisData", redisData),
			)
			return false, err
		}
	} else {
		zap.L().Error("LPopWithUnMarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("data", redisData),
		)
		return false, err
	}

	return true, nil
}

func RPopWithUnMarshal(redisKey string, outputData interface{}) (bool, error) {
	if outputData == nil || len(redisKey) <= 0 {
		return false, nil
	}

	redisData, err := global.RD.RPop(redisKey).Result()
	if err == nil {
		err = jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(redisData), outputData)
		if err != nil {
			zap.L().Error("RPopWithUnMarshal Json Unmarshal err",
				zap.String("param", fmt.Sprintf("(%v)", redisKey)),
				zap.Error(err),
				zap.String("redisData", redisData),
			)
			return false, err
		}
	} else {
		zap.L().Error("RPopWithUnMarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("data", redisData),
		)
		return false, err
	}

	return true, nil
}

func LTrimWithUnMarshal(redisKey string, start, stop int64) (bool, error) {
	if len(redisKey) <= 0 {
		return false, nil
	}

	redisData, err := global.RD.LTrim(redisKey, start, stop).Result()
	if err != nil {
		zap.L().Error("LTrimWithUnMarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("data", redisData))
		return false, err
	}

	return true, nil
}

func LIndexWithUnMarshal(redisKey string, index int64, outputData interface{}) (bool, error) {
	if outputData == nil || len(redisKey) <= 0 {
		return false, nil
	}

	redisData, err := global.RD.LIndex(redisKey, index).Result()
	if err != nil {
		zap.L().Error("LIndexWithUnMarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("data", redisData))
		return false, err
	}

	err = jsonIter.ConfigCompatibleWithStandardLibrary.Unmarshal([]byte(redisData), outputData)
	if err != nil {
		zap.L().Error("LIndexWithUnMarshal Json Unmarshal err",
			zap.String("param", fmt.Sprintf("(%v)", redisKey)),
			zap.Error(err),
			zap.String("redisData", redisData))
		return false, err
	}

	return true, nil
}
