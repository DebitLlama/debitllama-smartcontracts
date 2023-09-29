// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//   o          o    o__ __o__/_   o              o        o__ __o__/_  \o       o/  ____o__ __o____   o__ __o                o               o__ __o    ____o__ __o____          o           o__ __o     o            o__ __o__/_       o__ __o__/_   o__ __o__/_   o__ __o__/_ 
//  <|\        /|>  <|    v       <|>            <|>      <|    v        v\     /v    /   \   /   \   <|     v\              <|>             /v     v\    /   \   /   \          <|>         <|     v\   <|>          <|    v           <|    v       <|    v       <|    v      
//  / \\o    o// \  < >           < >            < >      < >             <\   />          \o/        / \     <\             / \            />       <\        \o/               / \         / \     <\  / \          < >               < >           < >           < >          
//  \o/ v\  /v \o/   |             \o            o/        |                \o/             |         \o/     o/           o/   \o        o/                    |              o/   \o       \o/     o/  \o/           |                 |             |             |           
//   |   <\/>   |    o__/_          v\          /v         o__/_             |             < >         |__  _<|           <|__ __|>      <|                    < >            <|__ __|>       |__  _<|    |            o__/_             o__/_         o__/_         o__/_       
//  / \        / \   |               <\        />          |                / \             |          |       \          /       \       \\                    |             /       \       |       \  / \           |                 |             |             |           
//  \o/        \o/  <o>                \o    o/           <o>             o/   \o           o         <o>       \o      o/         \o       \         /         o           o/         \o    <o>      /  \o/          <o>               <o>           <o>           <o>          
//   |          |    |                  v\  /v             |             /v     v\         <|          |         v\    /v           v\       o       o         <|          /v           v\    |      o    |            |                 |             |             |           
//  / \        / \  / \  _\o__/_         <\/>             / \  _\o__/_  />       <\        / \        / \         <\  />             <\      <\__ __/>         / \        />             <\  / \  __/>   / \ _\o__/_  / \  _\o__/_      / \           / \  _\o__/_  / \  _\o__/_ 
                                                                                                                                                                                                                                                                              
                                                                                                                                                                                                                                                                              
                                                                                                                                                                                                                                                                              


/**
This contract specifies an MEV extractable fee, that allows searchers to solve payment intents for a percentage of the profit
The searchers must apply to be whitelisted and aquire a signature from an approved signer that can authorize searchers
After aquireing the signature, they can whitelist themselves and extract MEV by solving Payment Intents
 */

abstract contract ExtractableFee {
    // 1% fee by default goes to the solver, which can be updated later
    // Fee Dividers are used to divide the amount to calculate the fee (payment /100) is 1% of payment
    uint256 public solverFeeDivider = 100;

    /**
     Throws when the signer to whitelist a searcher is invalid
     */
    error InvalidSigner();

    /**
      Throws when the searcher is not whitelisted
     */
    error InvalidSearcher();

    /**
    The mapping stores the whitelisted searchers allowed to solve a payment intent
    */
    mapping(bytes32 => mapping(address => bool)) public whitelistedSearchers;

    /**
     Returns the recipient of the fee, reverts if the searcher is not whitelisted
   */
    function _getSolver(
        bytes32 paymentIntent
    ) internal view returns (address) {
        if (!whitelistedSearchers[paymentIntent][msg.sender])
            revert InvalidSearcher();
        return msg.sender;
    }

    /**
     Updates the solver fee
   */
    function _updateSolverFee(uint256 newSolverFeeDivider) internal {
        solverFeeDivider = newSolverFeeDivider;
    }

    /**
      Whitelist a searcher
      A searcher can request to be whitelisted to solve a payment intent
      The backend specifically selects the intents based on conditions before the searcher is allowed to apply for it.
      multiple searchers can be whitelisted to handle the same intent
      Whitelisting is needed due to the nature of the recurring payments, as they can be repeated only approved addresses should solve them
     */

    function whitelistSearcher(
        uint8 v,
        bytes32 r,
        bytes32 s,
        bytes32 paymentIntent,
        address solver
    ) external {
        address signer = verifySignature(v, r, s, paymentIntent, solver);
        if (signer != _getApprovedSigner()) revert InvalidSigner();
        whitelistedSearchers[paymentIntent][msg.sender] = true;
    }

    /**
     Get the approved signer that can approve searchers so they can whitelist themselves
    */
    function _getApprovedSigner() internal virtual returns (address);

    function verifySignature(
        uint8 v,
        bytes32 r,
        bytes32 s,
        bytes32 paymentIntent,
        address solver
    ) public view returns (address) {
        uint256 chainId;
        assembly {
            chainId := chainid()
        }
        bytes32 eip712DomainHash = keccak256(
            abi.encode(
                keccak256(
                    "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
                ),
                keccak256(bytes("DirectDebit")),
                keccak256(bytes("1")),
                chainId,
                address(this)
            )
        );
        bytes32 hashStruct = keccak256(
            abi.encode(
                keccak256(bytes("doc(bytes32 paymentIntent,address solver)")),
                keccak256(abi.encode(paymentIntent)),
                keccak256(abi.encode(solver))
            )
        );
        bytes32 hash = keccak256(
            abi.encodePacked("\x19\x01", eip712DomainHash, hashStruct)
        );
        return ecrecover(hash, v, r, s);
    }
}
