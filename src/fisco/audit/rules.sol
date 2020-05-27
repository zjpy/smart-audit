pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./interface/IService.sol";
import "./Utils.sol";


// 规则抽象合约，封装规则注册以及验证相关的逻辑
contract Rules {
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
    }

    // 用于定义单个验证值
    struct ValidationValue {
        // 规则类型
        RuleType Type;
        // 验证值
        string ActualValues;
    }

    // 人脸识别服务合约地址
    address internal faceContract;
    // 物体识别服务合约地址
    address internal identifyContract;
    // 时间服务合约地址
    address internal timeContract;
    // 定位服务合约地址
    address internal locationContract;
    // 记录已存储的规则关系数
    uint32 relationsCount;
    // 用于维护一组验证规则关系，Key值对应一个规则关系ID，Value为具体的规则关系
    mapping(uint32 => ValidationRelationship) relationMap;

    /// @dev 规则合约的构造方法，通过将访问值设为internal使得Rules成为抽象合约.
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
        relationsCount = 0;
    }

    /// @dev 注册一个规则项的事件.
    /// @param ruleType 规则类型，以uint32形式表示.
    /// @param expression 需要注册到预言机上的规则表达式.
    event registerRuleEvent(uint32 ruleType, string expression);

    /// @dev 验证一个规则项的事件.
    /// @param ruleType 规则类型，以uint32形式表示.
    /// @param ruleID 返回在预言机服务中注册后对应的规则ID
    /// @param values 需要验证的规则值.
    event validateRuleEvent(uint32 ruleType, uint32 ruleID, string values);

    /// @dev 注册一条完整规则参数值中的所有规则项.
    /// @param args 所有表达式列表.
    /// @return relationID 返回在智能合约中注册后对应的规则关系ID
    function registerRules(string[] args) public returns (uint32 relationID) {
        LogicOperator op;
        ValidationExpression[] memory expressions;
        (op, expressions) = parseRuleExpressions(args);

        ValidationRelationship storage relation;
        relation.Operator = op;
        for (uint32 i = 0; i < expressions.length; i++) {
            emit registerRuleEvent(
                uint32(expressions[i].Type),
                expressions[i].Expression
            );

            uint32 id = registerRule(
                expressions[i].Type,
                expressions[i].Expression
            );
            relation.Rules[uint32(expressions[i].Type)] = id;
        }

        relationID = relationsCount;
        relationMap[relationID] = relation;
        relationsCount++;

        return relationID;
    }

    /// @dev 验证一条完整规则参数值中的所有规则项.
    /// @param relationID 在智能合约中注册后对应的规则关系ID
    /// @param expressions 所有表达式列表.
    function validateRules(uint32 relationID, string[] memory expressions)
        public
    {
        ValidationValue[] memory valueList;
        valueList = parseRuleValues(expressions);

        ValidationRelationship storage relation = relationMap[relationID];
        for (uint32 i = 0; i < valueList.length; i++) {
            RuleType t = valueList[i].Type;
            uint32 ruleID = relation.Rules[uint32(t)];
            require(ruleID != 0, "规则类型不在需要验证的列表中");

            emit validateRuleEvent(
                uint32(t),
                ruleID,
                valueList[i].ActualValues
            );

            validateRule(t, ruleID, valueList[i].ActualValues);
        }
    }

    /// @dev 解析参数值中的逻辑操作符及所有规则项.
    /// @param args 所有表达式列表
    /// @return op 返回逻辑操作符
    /// @return expressions 返回所有规则项
    function parseRuleExpressions(string[] args)
        internal
        returns (LogicOperator op, ValidationExpression[] memory expressions)
    {
        require(args.length > 0, "规则解析参数不足");

        op = getLogicOperator(args[0]);

        expressions = new ValidationExpression[](args.length / 2);
        uint32 index = 0;
        for (uint32 i = 1; i + 1 < args.length; i += 2) {
            RuleType t = getRuleType(args[i]);
            expressions[index] = ValidationExpression({
                Type: t,
                Expression: args[i + 1]
            });
            index++;
        }
        return (op, expressions);
    }

    /// @dev 注册一个规则项.
    /// @param t 规则类型，这里会根据规则类型匹配到对应的预言机服务.
    /// @param expression 需要注册到预言机上的规则表达式.
    /// @return ruleID 返回在预言机服务中注册后对应的规则ID
    /// @return errorCode 错误码，如果为0则表示没有错误，否则发生注册错误.
    /// @return message 返回结果信息.
    function registerRule(RuleType t, string expression)
        internal
        returns (uint32 ruleID)
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
            require(false, "编码类型尚未支持");
        }

        string[] memory args = new string[](1);
        args[0] = expression;
        return service.register(args);
    }

    /// @dev 根据表达式列表解析出所有的规则规则验证值.
    /// @param expressions 所有表达式列表.
    /// @return valueList 规则验证值列表.
    function parseRuleValues(string[] expressions)
        internal
        returns (ValidationValue[] storage valueList)
    {
        for (uint32 i = 0; i + 1 < expressions.length; i += 2) {
            RuleType t = getRuleType(expressions[i]);
            require(t != RuleType.None, "未找到规则");

            valueList.push(
                ValidationValue({Type: t, ActualValues: expressions[i + 1]})
            );
        }
        return valueList;
    }

    /// @dev 根据类型对应的字符串返回相应的LogicOperator类型.
    /// @param word 类型对应的字符串表示.
    /// @return op 一个LogicOperator类型.
    function getLogicOperator(string memory word)
        internal
        pure
        returns (LogicOperator op)
    {
        if (Utils.compareStrings(word, "AND")) {
            return LogicOperator.AND;
        } else if (Utils.compareStrings(word, "OR")) {
            return LogicOperator.OR;
        } else if (Utils.compareStrings(word, "NOT")) {
            return LogicOperator.NOT;
        } else {
            return LogicOperator.NONE;
        }
    }

    /// @dev 根据类型对应的字符串返回相应的RuleType类型.
    /// @param word 类型对应的字符串表示.
    /// @return t 一个RuleType类型.
    function getRuleType(string memory word)
        internal
        pure
        returns (RuleType t)
    {
        if (Utils.compareStrings(word, "Time")) {
            return RuleType.Time;
        } else if (Utils.compareStrings(word, "Location")) {
            return RuleType.Location;
        } else if (Utils.compareStrings(word, "FaceRecognize")) {
            return RuleType.FaceRecognize;
        } else if (Utils.compareStrings(word, "ObjectRecognize")) {
            return RuleType.ObjectRecognize;
        } else {
            return RuleType.None;
        }
    }

    /// @dev 返回一个RuleType类型对应的字符串表示.
    /// @param t 一个RuleType类型.
    /// @return word 类型对应的字符串表示.
    function ruleTypeToString(RuleType t)
        internal
        pure
        returns (string memory word)
    {
        if (t == RuleType.Time) {
            return "Time";
        } else if (t == RuleType.Location) {
            return "Location";
        } else if (t == RuleType.FaceRecognize) {
            return "FaceRecognize";
        } else if (t == RuleType.ObjectRecognize) {
            return "ObjectRecognize";
        } else {
            return "Unknown";
        }
    }

    /// @dev 验证一个规则项.
    /// @param t 规则类型，这里会根据规则类型匹配到对应的预言机服务.
    /// @param ruleID 返回在预言机服务中注册后对应的规则ID
    /// @param values 需要验证的规则值.
    function validateRule(
        RuleType t,
        uint32 ruleID,
        string values
    ) internal {
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
            require(false, "编码类型尚未支持");
        }

        string[] memory args = new string[](1);
        args[0] = values;
        return service.validate(ruleID, args);
    }
}
