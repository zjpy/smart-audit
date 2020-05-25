pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./interface/IService.sol";


// 规则抽象合约，封装规则注册以及验证相关的逻辑
contract Rules {
    // 人脸识别服务合约地址
    address internal faceContract;
    // 物体识别服务合约地址
    address internal identifyContract;
    // 时间服务合约地址
    address internal timeContract;
    // 定位服务合约地址
    address internal locationContract;

    /// @dev 规则合约的构造方法.
    /// @param face 人脸识别服务合约地址.
    /// @param identify 物体识别服务合约地址.
    /// @param time 时间服务合约地址.
    /// @param location 定位服务合约地址.
    constructor(
        address face,
        address identify,
        address time,
        address location
    ) internal {
        faceContract = face;
        identifyContract = identify;
        timeContract = time;
        locationContract = location;
    }

    // 规则类型枚举
    enum RuleType {None, Time, Location, FaceRecognize, ObjectRecognize}
    enum LogicOperator {NONE, NOT, AND, OR}

    // 用于定义单个验证规则的结构
    struct ValidationExpression {
        // 规则类型
        RuleType Type;
        // 具体验证规则
        string Expression;
    }

    // 用于定义一个验证组合中的各个验证规则的关系
    struct ValidationRelationship {
        // 逻辑操作符
        LogicOperator Operator;
        // 用于记录一组规则，Key值对应一个规则类型，Value值为注册规则表达式时预言机服务返回的相应的ID值
        mapping(uint32 => uint32) Rules;
        // 规则唯一标识
        uint32 ID;
    }

    // 用于定义单个验证值
    struct ValidationValue {
        // 规则类型
        RuleType Type;
        // 验证值
        string ActualValues;
        // 规则表达式在预言机中对应的ID
        uint32 ID;
    }

    /// @dev 注册一个规则项.
    /// @param t 规则类型，这里会根据规则类型匹配到对应的预言机服务.
    /// @param expression 需要注册到预言机上的规则表达式.
    /// @return ruleID 返回在预言机服务中注册后对应的规则ID
    /// @return errorCode 错误码，如果为0则表示没有错误，否则发生注册错误.
    /// @return message 返回结果信息.
    function registerRule(RuleType t, string expression)
        internal
        returns (
            uint32 ruleID,
            uint32 errorCode,
            string memory message
        )
    {
        IService service;
        if (t == RuleType.FaceRecognize) {
            service = IService(faceContract);
        } else if (t == RuleType.ObjectRecognize) {
            service = IService(identifyContract);
        } else if (t == RuleType.Time) {
            service = IService(timeContract);
        } else if (t == RuleType.Location) {
            service = IService(locationContract);
        } else {
            return (0, 1, "编码类型尚未支持");
        }

        string[] args;
        args.push(expression);
        return service.register(args);
    }

    /// @dev 验证一个规则项.
    /// @param t 规则类型，这里会根据规则类型匹配到对应的预言机服务.
    /// @param ruleID 返回在预言机服务中注册后对应的规则ID
    /// @param value 需要验证的规则值.
    /// @return errorCode 错误码，如果为0则表示没有错误，否则发生注册错误.
    /// @return message 返回结果信息.
    function validateRule(
        RuleType t,
        uint32 ruleID,
        string values
    ) internal returns (uint32 errorCode, string memory message) {
        IService service;
        if (t == RuleType.FaceRecognize) {
            service = IService(faceContract);
        } else if (t == RuleType.ObjectRecognize) {
            service = IService(identifyContract);
        } else if (t == RuleType.Time) {
            service = IService(timeContract);
        } else if (t == RuleType.Location) {
            service = IService(locationContract);
        } else {
            return (1, "编码类型尚未支持");
        }

        string[] args;
        args.push(values);
        return service.validate(ruleID, args);
    }
}
