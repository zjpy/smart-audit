pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./rules.sol";

// Smart Audit 合约
contract Audit is Rules {

    /// @dev 规则合约的构造方法，这里实现Rules.
    /// @param face 人脸识别服务合约地址.
    /// @param identify 物体识别服务合约地址.
    /// @param time 时间服务合约地址.
    /// @param location 定位服务合约地址.
    constructor(
        string[] names,
        address face,
        address identify,
        address time,
        address location
    ) Rules (face, identify, time, location) {
        numMaintainers = uint32(names.length);
        for (uint i = 0; i < numMaintainers; i++) {
            maintainers[uint32(i)] = names[i];
        }
    }

    /// @dev 注册一个验证人的事件.
    /// @param AuditeeID Auditee ID.
    event registerAuditeeEvent(uint32 AuditeeID);

    /// @dev 注册一个项目的事件.
    /// @param ProjectID Project ID.
    event registerProjectEvent(uint32 ProjectID);

    // 审计当事人计数器.
    uint32 numAuditees;
    // 用于存储所有审计当事人.
    mapping (uint32 => string) auditees;

    // 合约运维人员计数器
    uint32 numMaintainers;
    // 用于存储所有合约运维人员
    mapping (uint32 => string) maintainers; 

    // 用于定义单个项目.
    struct Project {
        string name;
        string detail;
        uint32 auditeeID;
        uint32 relationID;
    }
    // 项目计数器.
    uint32 numProjects;
    // 用于存储所有的项目实例.
    mapping (uint32 => Project) projects;

    // 用于存储所有的审计事件.
    mapping (uint32 =>  mapping (uint32 => string[])) events;

    /// @dev 注册一个 auditee.
    /// @param name auditee 名字.
    function registerAuditee(string name) public {
        numAuditees++;
        auditees[numAuditees] = name;

        emit registerAuditeeEvent(numAuditees);
    }

    /// @dev 注册一个项目.
    /// @param detail 项目信息.
    /// @param auditeeID auditee ID.
    /// @param relationID 相关规则 ID.
    function registerProject(string name, string detail, uint32 auditeeID, uint32 relationID) public {
        string storage auditee = auditees[auditeeID];
 //       require(auditee, "auditee 不存在");

        ValidationRelationship storage relation = relationMap[relationID];
 //       require(relation, "relation 不存在");

        numProjects++;
        projects[numProjects] = Project(name, detail, auditeeID, relationID);

        emit registerProjectEvent(numProjects);
    }

    /// @dev 添加一个审计事件.
    /// @param auditeeID 被审计人ID.
    /// @param projectID 项目ID.
    /// @param value 具体验证内容表达式.
    function addEvent(uint32 auditeeID, uint32 projectID, string[] value) public {
        string storage auditee = auditees[auditeeID];
 //       require(auditee, "auditee 不存在");

        Project storage project = projects[projectID];
 //       require(project, "项目不存在");

        validateRules(project.relationID, value);

        events[auditeeID][projectID] = value;
    }

    /// @dev 获取一个 auditee.
    /// @param auditeeID auditee ID.
    /// @return string auditee 名字.
    function getAuditee(uint32 auditeeID) public view returns(string) {
        return auditees[auditeeID];
    }

    /// @dev 获取一个项目.
    /// @param projectID 项目ID.
    /// @return Project 项目内容.
    function getProject(uint32 projectID) public view returns(Project) {
        return projects[projectID];
    }

    /// @dev 获取一个维护者列表.
    /// @return uint32[] 维护者列表.
    function getMaintainers() public view returns(string[] memory) {
        string[] memory rtn = new string[](numMaintainers);
        for (uint32 i = 0; i < numMaintainers; i++) {
            rtn[i] = maintainers[i];
        }
        return rtn;
    }

    /// @dev 获取一个审计事件.
    /// @param auditeeID 被审计人ID.
    /// @param projectID 项目ID.
    /// @return string[] 审计内容.
    function queryEvents(uint32 auditeeID, uint32 projectID) public view returns(string[]) {
       return events[auditeeID][projectID];
    }
}