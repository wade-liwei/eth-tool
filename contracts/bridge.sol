pragma solidity ^0.4.17;

//import "../token/ERC20/StandardToken.sol";
//pragma experimental ABIEncoderV2;
/// general helpers.
library Helpers {
    /// returns whether `array` contains `value`.
    function addressArrayContains(address[] array, address value) internal pure returns (bool) {
        for (uint256 i = 0; i < array.length; i++) {
            if (array[i] == value) {
                return true;
            }
        }
        return false;
    }
}

//contract ForeignBridge is StandardToken {
contract ForeignBridge  {
  uint256 public requiredSignatures;
  address[] public authorities;
  string public name;
  string public symbol;
  uint8 public decimals;
  uint256  public  batchSize;
  uint256  public  blockRange;
  uint256 public   minBalance;


  struct MessageCollection {
      uint  blockNum;
      bytes32 msgHash;
      bytes  message;
      bytes[] signatures;
  }
  mapping (bytes32 => MessageCollection[])  messagesByEthTx;
  /// Pending deposits and authorities who confirmed them
  mapping (bytes32 => address[]) deposits;
  //mapping (bytes32 => uint256) depositsCount;
  //限制地址重复
  mapping (bytes => uint) relayBlockNum;
  //限制指定eth块高度范围里交易数量
  mapping (uint256 => uint256) oneBlockTxNum;
  //限制指定eth块高度范围里deposit交易数量
  mapping (uint256 =>uint256) oneBlockDepositNum;


    /// triggered when an authority confirms a deposit
    event DepositConfirmation(address recipient, uint256 value, bytes32 transactionHash);
    /// triggered when enough authorities have confirmed a deposit
    event Deposit(address recipient, uint256 value, bytes32 transactionHash);
    /// Event created on money withdraw.
    event Withdraw(bytes recipient, uint256 value);
    event WithdrawWithEth(address contractAddr,address from, bytes recipient, uint256 value);

    constructor(
      uint256 _requiredSignatures,
      address[] _authorities,
      string _name,
      string _symbol,
      uint8 _decimals,
      uint256 _batchSize,
      uint256 _blockRange,
      uint256 _minBalance)public {
            require(_requiredSignatures != 0);
            require(_requiredSignatures <= _authorities.length);
            requiredSignatures = _requiredSignatures;
            authorities = _authorities;
            name = _name;
            symbol = _symbol;
            decimals = _decimals;
            batchSize = _batchSize;
            blockRange = _blockRange;
            minBalance = _minBalance;
      }

    /// require that sender is an authority
    modifier onlyAuthority() {
        require(Helpers.addressArrayContains(authorities, msg.sender));
        _;
    }

    /// Used to deposit money to the contract.
    /// deposit recipient (bytes20)
    /// deposit value (uint256)
    /// mainnet transaction hash (bytes32) // to avoid transaction duplication
    /// utox tx vout idx to avoid tx vout duplication
    function deposit(address recipient, uint256 value, bytes32 transactionHash,uint64 idx) onlyAuthority()  public {
        require( !(minBalance   > value ));



        // Protection from misbehaving authority
        bytes32 newhash;
        newhash = keccak256(recipient, value, transactionHash,idx);
        //newhash = keccak256(recipient, value, transactionHash,idx);
        // don't allow authority to confirm deposit twice
        require(!Helpers.addressArrayContains(deposits[newhash], msg.sender));
        uint256 total;
        for (uint256 i = block.number -  blockRange; i <=block.number; i++){
            total +=oneBlockDepositNum[i];
        }
        require(total  < batchSize);

        deposits[newhash].push(msg.sender);
        //  depositsCount[newhash] +=1;
        // TODO: this may cause troubles if requiredSignatures len is changed
        if (deposits[newhash].length != requiredSignatures) {
            DepositConfirmation(recipient, value, transactionHash);
            return ;
        }

        oneBlockDepositNum[block.number] += 1;
        recipient.transfer(value);

        Deposit(recipient, value, transactionHash);
    }


  function GetAllowSubmitBatchSize()public view  returns(uint256){
     uint256 total;
      for (uint256 i = block.number -  blockRange; i <=block.number; i++){
          total +=oneBlockDepositNum[i];
      }
      return batchSize - total;
 }


 function checkDepositBatchSizeWithBlockRangeTotal()public view returns (bool){
      uint256 total;
      for (uint256 i = block.number -  blockRange; i <=block.number; i++){
          total +=oneBlockDepositNum[i];
      }

      if (total  < batchSize){
        return true;
      }

      return  false;
 }
//mapping (bytes32 => address[]) deposits;
  function  checkMyDepositAlreadySubmit(address recipient, uint256 value, bytes32 transactionHash,uint64 idx)public view returns(address[]){
      bytes32 newhash;
      newhash = keccak256(recipient, value, transactionHash,idx);
      return deposits[newhash];
  }
/*
    function getAlreadySubmitedCount(address recipient, uint256 value, bytes32 transactionHash,uint64 idx)public view returns (uint256){
      bytes32 newhash;
      newhash = keccak256(recipient, value, transactionHash,idx);
      return depositsCount[newhash];
    }

    function getAlreadySubmitedItem(address recipient, uint256 value, bytes32 transactionHash,uint64 idx, uint256 addrIdx)public view returns(address){
      bytes32 newhash;
      newhash = keccak256(recipient, value, transactionHash,idx);
      return deposits[newhash][addrIdx];
    } */



    /// Should be used to withdraw eth.
  function () public payable {
    require( !(minBalance > msg.value));
    //Withdraw(msg.data, msg.value  );
    //WithdrawWithEth(this,msg.sender,msg.data,msg.value / );
    Withdraw(msg.data, msg.value);
    //WithdrawWithEth(this,msg.sender,msg.data,msg.value);
  }


  function SubmitMessage(bytes32 ethTxHash, bytes message)public   onlyAuthority() {
    bytes32  newHash;
    newHash = keccak256(message);
    MessageCollection   memory tmp;
    tmp.message = message;
    tmp.msgHash = newHash;
    tmp.blockNum = block.number;
    for(uint256 i= 0; i < messagesByEthTx[ethTxHash].length; i++){
      //按刘岩的说法有可能重复；
      if (messagesByEthTx[ethTxHash][i].msgHash == newHash){
        return;
      }
    }
    messagesByEthTx[ethTxHash].push(tmp);
  }
  function getMessage(bytes32 ethTxHash,uint256 idx)public view returns (bytes){
    return messagesByEthTx[ethTxHash][idx].message;
  }
  function getMessageLenth(bytes32 ethTxHash)public view returns (uint){
    return messagesByEthTx[ethTxHash].length;
  }
  function getMessageBlockNumber(bytes32 ethTxHash,uint256 idx)public view returns (uint){
    return messagesByEthTx[ethTxHash][idx].blockNum;
  }
  function SubmitMessageSignature(bytes32 ethTxHash, bytes message, bytes signatureBytes)public   onlyAuthority() {
     bytes32  newHash;
     newHash = keccak256(message);
     for (uint256 i = 0; i < messagesByEthTx[ethTxHash].length; i++){
       if (messagesByEthTx[ethTxHash][i].msgHash == newHash){
          messagesByEthTx[ethTxHash][i].signatures.push(signatureBytes);
          return;
       }
     }
     require(false);
  }
  function getMessageSignature(bytes32 ethTxHash, bytes message ,uint256 sigIdx)public view returns (bytes){
    bytes32  newHash;
    newHash = keccak256(message);
      for (uint256 i = 0; i < messagesByEthTx[ethTxHash].length; i++){
        if (messagesByEthTx[ethTxHash][i].msgHash == newHash){
          return messagesByEthTx[ethTxHash][i].signatures[sigIdx];
        }
      }
   }

   function getMessageSignatureNum(bytes32 ethTxHash, bytes message)public view returns (uint){
     bytes32 newHash;
     newHash = keccak256(message);
     for (uint256 i = 0; i < messagesByEthTx[ethTxHash].length; i++){
        if (messagesByEthTx[ethTxHash][i].msgHash == newHash){
         return  messagesByEthTx[ethTxHash][i].signatures.length;
          }
       }
    }

}
