pragma solidity 0.4.24;
pragma experimental ABIEncoderV2;

contract Face {
      function init() public {
           // 初始化人脸识别预言机服务相关信息……
      }

      function invoke(uint _data) public pure returns (bool) {
            return true;
      }
}

contract Identify {
      function init() public {
           // 初始化身份识别预言机服务相关信息……
      }

      function invoke(uint _data) public pure returns (bool) {
            return true;
      }
}

contract Location {
      function init() public {
           // 初始化位置识别预言机服务相关信息……
      }

      function invoke(uint _data) public pure returns (bool) {
            return true;
      }
}

contract Time {
      function init() public {
           // 初始化时间服务预言机服务相关信息……
      }

      function invoke(uint _data) public pure returns (bool) {
            return true;
      }
}

contract Audit {
    uint256 numAuditees;
    mapping (uint256 => string) auditees;

    struct Rule {
        string time;
        string location;
        string faceRecognize;
        string objectRecognize;
    }
    uint256 numRules;
    mapping (uint256 => Rule) rules;
    mapping (uint256 => string) ruleNames;

    struct Project {
        string detail;
        uint256 auditeeId;
        uint256 ruleId;
    }
    uint256 numProjects;
    mapping (uint256 => Project) projects;

    mapping (uint256 => mapping (uint256 => mapping (uint256 => Rule))) events;

    function registerAuditee(string name) public {
        numAuditees++;
        auditees[numAuditees] = name;
    }

    function registerRule(string name, string time, string location, string faceRec, string objectRec) public {
        numRules++;
        rules[numRules] = Rule(time, location, faceRec, objectRec);
        ruleNames[numRules] = name;
    }

    function registerProject(string detail, uint256 auditeeId, uint256 ruleId) public {
        numProjects++;
        projects[numProjects] = Project(detail, auditeeId, ruleId);
    }

    function addEvent(uint256 auditeeId, uint256 projectId, uint256 ruleId,
        string time, string location, string faceRec, string objectRec) public {
        events[auditeeId][projectId][ruleId] = Rule(time, location, faceRec, objectRec);
    }

    function getAuditee(uint256 auditeeId) public view returns(string) {
        return auditees[auditeeId];
    }

    function getRule(uint256 ruleId)
        public
        view
        returns(Rule)
    {
        return rules[ruleId];
    }

    function getProject(uint256 projectId) public view returns(Project) {
        return projects[projectId];
    }

    function getMaintainers() public {

    }

    function queryEvents(uint256 auditeeId, uint256 projectId, uint256 ruleId) public view returns(Rule) {
       return events[auditeeId][projectId][ruleId];
    }
}