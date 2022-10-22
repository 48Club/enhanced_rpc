// SPDX-License-Identifier: MIT

pragma solidity ^0.8.17;

import "@openzeppelin/contracts/access/Ownable.sol";

contract PuissantIndicator is Ownable {
    address[] public puissants;
    mapping(address => bool) public isPuissant;

    event PuissantDeployed(address coinbase);
    event PuissantDisabled(address coinbase);

    function addValidator(address coinbase) external onlyOwner {
        require(!isPuissant[coinbase], "already exist");

        isPuissant[coinbase] = true;
        puissants.push(coinbase);

        emit PuissantDeployed(coinbase);
    }

    function removeValidator(address coinbase) external onlyOwner {
        require(isPuissant[coinbase], "!exist");

        isPuissant[coinbase] = false;
        for (uint256 index = 0; index < puissants.length; index++) {
            if (puissants[index] == coinbase) {
                puissants[index] = puissants[puissants.length - 1];
                puissants.pop();
                break;
            }
        }

        emit PuissantDisabled(coinbase);
    }

    function puissantsLength() external view returns (uint256) {
        return puissants.length;
    }
}
