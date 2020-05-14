package dummy

import (
	"audit/contract"
	"crypto/rand"
	"errors"
	"face/service"
	mrand "math/rand"
)

type FaceValidation struct {
}

func (f *FaceValidation) Validate(id contract.ServiceRuleID, args []string) error {
	if len(args) == 0 {
		return errors.New("不允许没有参数的验证")
	}

	// 云从科技的人脸对比服务支持base64编码的图片以及由图片生成的特征码，为了减轻人脸识别存储端的压力，
	//	且考虑到对个人肖像隐私保护，我们在预言机服务中使用的是特征码形式存储。一个典型的特征码如下所示（
	//	注意特征码内容量较大，这里只列出了收尾部分，中间以...形式省略）：
	// Q+tGPz8InT9O7A0+8bwUwJy4oL+55MS+ADWFv0Bze75mG8o/...//bLSfOk3W7zoAAIA/U3xdwg==

	feature, err := f.getFaceFeature(args[0])
	if err != nil {
		return err
	}

	return f.faceCompare(feature)
}

// 调用人脸预言机服务中的人脸特征提取接口用以返回图片中响应人脸的特征值。
// 这里我们返回一个随机值以模拟该特征值生成过程
func (f *FaceValidation) getFaceFeature(faceRaw string) (rtn []byte, err error) {
	rtn = make([]byte, mrand.Int31n(1000))
	if _, err = rand.Read(rtn); err != nil {
		return
	}
	return
}

// 调用人脸预言机服务中的人脸比对接口，并返回包含FaceCompareResult中内容的比对结果。
// 这里我们返回一个人脸评分在0.93~1之间的数值，以模拟在以0.95为界有>80%通过率的情况
func (f *FaceValidation) getFaceCompareResult(rtn []byte) service.FaceCompareResult {
	return service.FaceCompareResult{
		Result:   0,
		Score:    0.93 + float32(mrand.Int31n(7))*0.01,
		Sequence: "XXXXX",
	}
}

// 这里对人脸比对评分进行评价，评分超过或等于0.95则验证成功，否则视为非本人的情况
func (f *FaceValidation) faceCompare(rtn []byte) error {
	result := f.getFaceCompareResult(rtn)
	if result.Score < 0.95 {
		return errors.New("非本人操作")
	}
	return nil
}
