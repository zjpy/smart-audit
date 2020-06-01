pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./interface/IService.sol";
import "./IdentifyResult.sol";


// 该智能合约用于模拟物体识别的预言机服务调用
contract DummyIdentifyService is IService {
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
        require(args.length > 0, "缺少需要识别物体的数据");

        // 与人脸识别类似，云从科技的物体识别服务支持base64编码的图片以及由图片生成的特征码。
        //	一个典型的特征码如下所示（注意特征码内容量较大，这里只列出了首尾部分，中间以...形式省略）：
        // g1T4vTuetL2grRVANLFpvwiGEz7YZkbAcWcwv/.../6ulAOO9aYQToAAIA/yZduwg==

        bytes memory feature;
        feature = getEntityFeature(args[0]);

        return entityCompare(args[0], feature);
    }

    /// @dev 调用物体识别预言机服务中的特征提取接口用以返回图片中相应物体的特征值。
    /// @param faceRaw 返回在预言机服务中注册后对应的规则ID
    /// @return rtn 物体图像的特征值.
    function getEntityFeature(string memory faceRaw)
        private
        returns (bytes memory rtn)
    {
        // 这里我们简单返回一个单值模拟该特征值生成过程
        rtn = new bytes(1);
        rtn[0] = bytes1(1);
        return rtn;
    }

    /// @dev 这里对物体识别评分进行评价，评分超过或等于90则验证成功，否则视为其他物体
    /// @param ruleID 返回在预言机服务中注册后对应的规则ID
    /// @param feature 物体图像的特征值
    function entityCompare(string ruleID, bytes memory feature) private {
        IdentifyResult.IdentifyCompare memory result = getEntityCompareResult(
            ruleID,
            feature
        );
        require(result.Score >= 90, "未识别到指定物体");
    }

    /// @dev 调用物体识别预言机服务中的物体比对接口，并返回包含EntityIdentifyResult中内容的比对结果。
    ///        这里我们返回一个物体评分在88~100之间的数值，以模拟在以90为界有>80%通过率的情况
    function getEntityCompareResult(string ruleID, bytes memory feature)
        private
        returns (IdentifyResult.IdentifyCompare memory rtn)
    {
        rtn.Result = 0;
        rtn.Score = 88 + random(12);
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
