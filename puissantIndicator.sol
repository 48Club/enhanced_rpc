// SPDX-License-Identifier: MIT

pragma solidity ^0.8.1;

import "@openzeppelin/contracts/access/Ownable.sol";

contract Indicator is Ownable {
    address[] public _validators;
    mapping(address => bool) public validators;

    function addValidator(address validator) public onlyOwner {
        validators[validator] = true;
        for (uint256 index = 0; index < _validators.length; index++) {
            if (_validators[index] == validator) return;
        }
        _validators.push(validator);
    }

    function removeValidator(address validator) public onlyOwner {
        validators[validator] = false;
        for (uint256 index = 0; index < _validators.length; index++) {
            if (_validators[index] == validator) {
                _validators[index] = _validators[_validators.length - 1];
                _validators.pop();
                return;
            }
        }
    }

    function getAllValidators() public view returns (address[] memory) {
        return _validators;
    }
}
