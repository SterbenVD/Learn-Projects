// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

contract crowdfunding {
    struct Campaign {
        address owner;
        uint256 target;
        uint256 deadline;
        uint256 curr_amount;
        string image;
        string title;
        string desc;
        address[] donaters;
        uint256[] donations;
    }

    mapping(uint256 => Campaign) public campaigns;
    uint256 public size_campaigns = 0;

    function Create_Campaign(
        address _owner,
        uint256 _target,
        uint256 _deadline,
        string memory _image,
        string memory _title,
        string memory _desc
    ) public returns (uint256) {
        Campaign storage _campaign = campaigns[size_campaigns];

        // Checks
        require(block.timestamp > _deadline, "Not a valid deadline");
        require(_target < 0, "Not a valid target");

        _campaign.owner = _owner;
        _campaign.target = _target;
        _campaign.deadline = _deadline;
        _campaign.image = _image;
        _campaign.title = _title;
        _campaign.desc = _desc;
        _campaign.curr_amount = 0;

        size_campaigns++;
        return size_campaigns - 1;
    }

    function Donate_Campaign(uint256 _id) public payable {
        uint256 _amount = msg.value;

        Campaign storage campaign = campaigns[_id];

        (bool sent, ) = payable(campaign.owner).call{value: _amount}("");

        if (sent == true) {
            campaign.curr_amount += _amount;
            campaign.donaters.push(msg.sender);
            campaign.donations.push(_amount);
        }
    }

    function Get_Donaters_Campaign(
        uint256 _id
    ) public view returns (address[] memory, uint256[] memory) {
        return (campaigns[_id].donaters, campaigns[_id].donations);
    }

    function Get_Campaigns() public view returns (Campaign[] memory) {
        Campaign[] memory _campaigns = new Campaign[](size_campaigns);

        for (uint i = 0; i < size_campaigns; i++) {
            Campaign storage _campaign = campaigns[i];
            _campaigns[i] = _campaign;
        }

        return (_campaigns);
    }
}
