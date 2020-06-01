pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./interface/IService.sol";
import "./FaceResult.sol";


// 该智能合约用于模拟人脸识别的预言机服务调用
contract DummyFaceService is IService {
    uint32 nonce;

    /// @dev 构造方法，目前仅用于初始化Nonce值.
    constructor() public {
        // 这里初始化nonce值
        nonce = 10;
    }

    /// @dev 注册一个规则项.
    /// @param args 注册规则参数，以数组形式表示.
    /// @return ruleID 返回在预言机服务中注册后对应的规则ID
    function register(string[] args) public returns (uint32 ruleID) {
        // 由于假设是固定的验证规则，所以这里不需要额外工作
        return 0;
    }

    /// @dev 验证一个规则项.
    /// @param ruleID 返回在预言机服务中注册后对应的规则ID
    /// @param args 验证规则所需的值，以数组形式表示.
    function validate(uint32 ruleID, string[] args) public {
        require(args.length > 0, "缺少人脸数据");

        // 云从科技的人脸对比服务支持base64编码的图片以及由图片生成的特征码，为了减轻人脸识别存储端的压力，
        //	且考虑到对个人肖像隐私保护，我们在预言机服务中使用的是特征码形式存储。一个典型的特征码如下所示（
        //	注意特征码内容量较大，这里只列出了首尾部分，中间以...形式省略）：
        // Q+tGPz8InT9O7A0+8bwUwJy4oL+55MS+ADWFv0Bze75mG8o/.../bLSfOk3W7zoAAIA/U3xdwg==

        bytes memory feature;
        feature = getFaceFeature(args[0]);

        return faceCompare(args[0], feature);
    }

    /// @dev 调用人脸预言机服务中的人脸特征提取接口用以返回图片中相应人脸的特征值。
    /// @param faceRaw 返回在预言机服务中注册后对应的规则ID
    /// @return rtn 人脸图像的特征值.
    function getFaceFeature(string memory faceRaw)
        private
        returns (bytes memory rtn)
    {
        // 这里我们简单返回一个单值模拟该特征值生成过程
        rtn = new bytes(1);
        rtn[0] = bytes1(1);
        return rtn;
    }

    /// @dev 这里对人脸比对评分进行评价，评分超过或等于95则验证成功，否则视为非本人的情况
    /// @param ruleID 返回在预言机服务中注册后对应的规则ID
    /// @param feature 人脸图像的特征值
    /// @return errorCode 错误码，如果为0则表示没有错误，否则发生注册错误.
    /// @return message 返回结果信息.
    function faceCompare(string ruleID, bytes memory feature) private {
        FaceResult.FaceCompare memory result = getFaceCompareResult(
            ruleID,
            feature
        );
        require(result.Score >= 95, "非本人操作");
    }

    /// @dev 调用人脸预言机服务中的人脸比对接口，并返回包含FaceCompareResult中内容的比对结果。
    ///        这里我们返回一个人脸评分在93~100之间的数值，以模拟在以95为界有>80%通过率的情况
    function getFaceCompareResult(string ruleID, bytes memory feature)
        private
        returns (FaceResult.FaceCompare memory rtn)
    {
        rtn.Result = 0;
        rtn.Score = 93 + random(7);
        rtn.Sequence = "XXXXX";
        return (rtn);
    }

    /// @dev 模拟生成一个随机值。
    /// @param n 给定随机值的最大值。
    function random(uint32 n) private returns (uint32) {
        uint32 random = uint32(keccak256(now, msg.sender, nonce)) % n;
        nonce++;
        return random;
    }
}
