// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package protocol

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"numerator\",\"type\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\"},{\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"isRoundingError\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5][]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6][]\"},{\"name\":\"fillTakerTokenAmount\",\"type\":\"uint256\"},{\"name\":\"shouldThrowOnInsufficientBalanceOrAllowance\",\"type\":\"bool\"},{\"name\":\"v\",\"type\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"fillOrdersUpTo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6]\"},{\"name\":\"cancelTakerTokenAmount\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZRX_TOKEN_CONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5][]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6][]\"},{\"name\":\"fillTakerTokenAmounts\",\"type\":\"uint256[]\"},{\"name\":\"v\",\"type\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"batchFillOrKillOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6]\"},{\"name\":\"fillTakerTokenAmount\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getUnavailableTakerTokenAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"numerator\",\"type\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\"},{\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"getPartialAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_TRANSFER_PROXY_CONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5][]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6][]\"},{\"name\":\"fillTakerTokenAmounts\",\"type\":\"uint256[]\"},{\"name\":\"shouldThrowOnInsufficientBalanceOrAllowance\",\"type\":\"bool\"},{\"name\":\"v\",\"type\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"batchFillOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5][]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6][]\"},{\"name\":\"cancelTakerTokenAmounts\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelOrders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6]\"},{\"name\":\"fillTakerTokenAmount\",\"type\":\"uint256\"},{\"name\":\"shouldThrowOnInsufficientBalanceOrAllowance\",\"type\":\"bool\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"fillOrder\",\"outputs\":[{\"name\":\"filledTakerTokenAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6]\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXTERNAL_QUERY_GAS_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_zrxToken\",\"type\":\"address\"},{\"name\":\"_tokenTransferProxy\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"makerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"takerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"filledMakerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"filledTakerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"paidMakerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"paidTakerFee\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"tokens\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"LogFill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"makerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"takerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cancelledMakerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cancelledTakerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"tokens\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"LogCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"errorId\",\"type\":\"uint8\"},{\"indexed\":true,\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"LogError\",\"type\":\"event\"}]"

// ExchangeBin is the compiled bytecode used for deploying new contracts.
const ExchangeBin = `0x608060405234801561001057600080fd5b50604051604080611d8f83398101604052805160209091015160008054600160a060020a03938416600160a060020a03199182161790915560018054939092169216919091179055611d28806100676000396000f3006080604052600436106100fb5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166314df96ee8114610100578063288cdc91146101325780632ac126221461015c578063363349be14610174578063394c21e7146103285780633b30ba59146103935780634f150787146103c4578063741bcc93146105ab5780637e9abb50146106285780638163681e1461064057806398024a8b14610670578063add1cbc51461068e578063b7b2c7d6146106a3578063baa0181d14610891578063bc61394a146109cb578063cfc4d0ec14610a50578063f06bbf7514610ab9578063ffa1ad7414610ae5575b600080fd5b34801561010c57600080fd5b5061011e600435602435604435610b6f565b604080519115158252519081900360200190f35b34801561013e57600080fd5b5061014a600435610bc5565b60408051918252519081900360200190f35b34801561016857600080fd5b5061014a600435610bd7565b34801561018057600080fd5b5060408051600480358082013560208181028501810190955280845261014a943694602493909290840191819060009085015b828210156101ef576040805160a081810190925290808402870190600590839083908082843750505091835250506001909101906020016101b3565b50506040805186358801803560208181028401810190945280835296999897830196919550820193509150819060009085015b8282101561025e576040805160c08181019092529080840287019060069083908390808284375050509183525050600190910190602001610222565b505060408051602087830135890180358281028085018401909552808452979a89359a838b013515159a919990985060609091019650929450810192829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610be99650505050505050565b34801561033457600080fd5b506040805160a081810190925261014a91369160049160a491908390600590839083908082843750506040805160c0818101909252949796958181019594509250600691508390839080828437509396505092359350610d0892505050565b34801561039f57600080fd5b506103a8610f86565b60408051600160a060020a039092168252519081900360200190f35b3480156103d057600080fd5b506040805160048035808201356020818102850181019095528084526105a9943694602493909290840191819060009085015b8282101561043f576040805160a08181019092529080840287019060059083908390808284375050509183525050600190910190602001610403565b50506040805186358801803560208181028401810190945280835296999897830196919550820193509150819060009085015b828210156104ae576040805160c08181019092529080840287019060069083908390808284375050509183525050600190910190602001610472565b50505050509192919290803590602001908201803590602001908080602002602001604051908101604052809392919081815260200183836020028082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610f959650505050505050565b005b3480156105b757600080fd5b506040805160a08181019092526105a991369160049160a491908390600590839083908082843750506040805160c0818101909252949796958181019594509250600691508390839080828437509396505083359450505050602081013560ff16906040810135906060013561104a565b34801561063457600080fd5b5061014a60043561106d565b34801561064c57600080fd5b5061011e600160a060020a036004351660243560ff60443516606435608435611097565b34801561067c57600080fd5b5061014a60043560243560443561114e565b34801561069a57600080fd5b506103a861116b565b3480156106af57600080fd5b506040805160048035808201356020818102850181019095528084526105a9943694602493909290840191819060009085015b8282101561071e576040805160a081810190925290808402870190600590839083908082843750505091835250506001909101906020016106e2565b50506040805186358801803560208181028401810190945280835296999897830196919550820193509150819060009085015b8282101561078d576040805160c08181019092529080840287019060069083908390808284375050509183525050600190910190602001610751565b5050505050919291929080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284375050604080516020808901358a01803580830284810184018652818552999c8b3515159c909b909a950198509296508101945090925082919085019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a99890198929750908201955093508392508501908490808284375094975061117a9650505050505050565b34801561089d57600080fd5b506040805160048035808201356020818102850181019095528084526105a9943694602493909290840191819060009085015b8282101561090c576040805160a081810190925290808402870190600590839083908082843750505091835250506001909101906020016108d0565b50506040805186358801803560208181028401810190945280835296999897830196919550820193509150819060009085015b8282101561097b576040805160c0818101909252908084028701906006908390839080828437505050918352505060019091019060200161093f565b505050505091929192908035906020019082018035906020019080806020026020016040519081016040528093929190818152602001838360200280828437509497506112239650505050505050565b3480156109d757600080fd5b506040805160a081810190925261014a91369160049160a491908390600590839083908082843750506040805160c0818101909252949796958181019594509250600691508390839080828437509396505083359450505050602081013515159060ff604082013516906060810135906080013561128e565b348015610a5c57600080fd5b506040805160a081810190925261014a91369160049160a491908390600590839083908082843750506040805160c0818101909252949796958181019594509250600691508390839080828437509396506116f795505050505050565b348015610ac557600080fd5b50610ace6117ba565b6040805161ffff9092168252519081900360200190f35b348015610af157600080fd5b50610afa6117c0565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610b34578181015183820152602001610b1c565b50505050905090810190601f168015610b615780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600080600084801515610b7e57fe5b8685099150811515610b935760009250610bbc565b610bb2610ba383620f42406117f7565b610bad88876117f7565b611822565b90506103e8811192505b50509392505050565b60026020526000908152604090205481565b60036020526000908152604090205481565b600080805b8951811015610cfb57896000815181101515610c0657fe5b6020908102909101015160036020020151600160a060020a03168a82815181101515610c2e57fe5b602090810291909101015160600151600160a060020a031614610c5057600080fd5b610ce482610cdf8c84815181101515610c6557fe5b906020019060200201518c85815181101515610c7d57fe5b90602001906020020151610c918d88611839565b8c8c88815181101515610ca057fe5b906020019060200201518c89815181101515610cb857fe5b906020019060200201518c8a815181101515610cd057fe5b9060200190602002015161128e565b61184b565b915087821415610cf357610cfb565b600101610bee565b5098975050505050505050565b6000610d12611ca0565b60408051610160810182528651600160a060020a039081168252602080890151821681840152888401518216838501526060808a01518316818501526080808b015190931683850152885160a08501529088015160c08401529287015160e0830152918601516101008201529085015161012082015260009081906101408101610d9c89896116f7565b9052805190935033600160a060020a03908116911614610dbb57600080fd5b60008360a00151118015610dd3575060008360c00151115b8015610ddf5750600085115b1515610dea57600080fd5b6101208301514210610e375761014083015160005b60ff167f36d86c59e00bd73dc19ba3adfe068e4b64ac7e92be35546adeddf1b956a87e9060405160405180910390a360009350610f7c565b610e528360c00151610e4d85610140015161106d565b611839565b9150610e5e858361185a565b9050801515610e74576101408301516001610dff565b610140830151600090815260036020526040902054610e93908261184b565b6101408401516000908152600360205260409081902091909155808401805160608601805184516c01000000000000000000000000600160a060020a03948516810282529184169091026014820152935193849003602801909320608087015187519351945160c089015160a08a0151939692851695909416937f67d66f160bc93d925d05dae1794c90d2d6d6688b29b84ff069398a9b0458713193610f3a91899161114e565b6101408a015160408051600160a060020a0395861681529390941660208401528284019190915260608201889052608082015290519081900360a00190a48093505b5050509392505050565b600054600160a060020a031681565b60005b8651811015611041576110398782815181101515610fb257fe5b906020019060200201518783815181101515610fca57fe5b906020019060200201518784815181101515610fe257fe5b906020019060200201518785815181101515610ffa57fe5b90602001906020020151878681518110151561101257fe5b90602001906020020151878781518110151561102a57fe5b9060200190602002015161104a565b600101610f98565b50505050505050565b8361105b878787600088888861128e565b1461106557600080fd5b505050505050565b6000818152600260209081526040808320546003909252822054611091919061184b565b92915050565b604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101869052815190819003603c018120600080835260208381018086529290925260ff8716838501526060830186905260808301859052925160019260a08082019392601f198101928190039091019086865af1158015611125573d6000803e3d6000fd5b50505060206040510351600160a060020a031686600160a060020a031614905095945050505050565b600061116361115d85846117f7565b84611822565b949350505050565b600154600160a060020a031681565b60005b875181101561121957611210888281518110151561119757fe5b9060200190602002015188838151811015156111af57fe5b9060200190602002015188848151811015156111c757fe5b906020019060200201518888868151811015156111e057fe5b9060200190602002015188878151811015156111f857fe5b906020019060200201518888815181101515610cd057fe5b5060010161117d565b5050505050505050565b60005b83518110156112885761127f848281518110151561124057fe5b90602001906020020151848381518110151561125857fe5b90602001906020020151848481518110151561127057fe5b90602001906020020151610d08565b50600101611226565b50505050565b6000611298611ca0565b600080600080610160604051908101604052808e60006005811015156112ba57fe5b60209081029190910151600160a060020a03168252018e600160209081029190910151600160a060020a03168252018e600260209081029190910151600160a060020a03168252018e600360209081029190910151600160a060020a03168252018e600460209081029190910151600160a060020a03168252018d6000602090810291909101518252018d6001602090810291909101518252018d6002602090810291909101518252018d6003602090810291909101518252018d60046020020151815260200161138b8f8f6116f7565b90526020810151909550600160a060020a031615806113bf575033600160a060020a03168560200151600160a060020a0316145b15156113ca57600080fd5b60008560a001511180156113e2575060008560c00151115b80156113ee575060008b115b15156113f957600080fd5b61140f85600001518661014001518b8b8b611097565b151561141a57600080fd5b61012085015142106114675761014085015160005b60ff167f36d86c59e00bd73dc19ba3adfe068e4b64ac7e92be35546adeddf1b956a87e9060405160405180910390a3600095506116e7565b61147d8560c00151610e4d87610140015161106d565b93506114898b8561185a565b955085151561149f57610140850151600161142f565b6114b2868660c001518760a00151610b6f565b156114c457610140850151600261142f565b891580156114d957506114d78587611870565b155b156114eb57610140850151600361142f565b6114fe868660c001518760a0015161114e565b610140860151600090815260026020526040902054909350611520908761184b565b6101408601516000908152600260205260409081902091909155850151855161154b91903386611ac3565b151561155657600080fd5b61156a856060015133876000015189611ac3565b151561157557600080fd5b6080850151600160a060020a0316156116265760008560e0015111156115d6576115a8868660c001518760e0015161114e565b600054865160808801519294506115cb92600160a060020a039092169185611ac3565b15156115d657600080fd5b60008561010001511115611626576115f8868660c0015187610100015161114e565b600054608087015191925061161b91600160a060020a0390911690339084611ac3565b151561162657600080fd5b604080860180516060808901805185516c01000000000000000000000000600160a060020a0395861681028252918516909102601482015285519081900360280181206080808d01518d51975194516101408f0151338916865295881660208601528716848a01529483018b905282018d905260a0820189905260c0820188905260e08201929092529451909491831693909216917f0d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb3918190036101000190a45b5050505050979650505050505050565b815160208084015160408086015160608088015160809889015188519689015189860151938a01519a8a015160a0909a01518651600160a060020a033081166c0100000000000000000000000090810283529b81168c0260148301529889168b0260288201529588168a02603c87015292871689026050860152951690960260648301526078820194909452609881019290925260b882019290925260d881019490945260f8840192909252610118830152516101389181900391909101902090565b61138781565b60408051808201909152600581527f312e302e30000000000000000000000000000000000000000000000000000000602082015281565b6000828202831580611813575082848281151561181057fe5b04145b151561181b57fe5b9392505050565b600080828481151561183057fe5b04949350505050565b60008282111561184557fe5b50900390565b60008282018381101561181b57fe5b6000818310611869578161181b565b5090919050565b60008060008060008060008060003397506118948a8c60c001518d60a0015161114e565b60808c0151909750600160a060020a031615611a455760005460408c015160608d015160c08e015160e08f0151600160a060020a0394851693851684149a509390911690911496506118e8918c919061114e565b93506118fe8a8c60c001518d610100015161114e565b92508561190b5783611915565b611915878561184b565b915084611922578261192c565b61192c8a8461184b565b6000548c51919250839161194991600160a060020a031690611b7a565b108061196e57506000548b51839161196c91600160a060020a0390911690611c2b565b105b8061198e5750600054819061198c90600160a060020a03168a611b7a565b105b806119ae575060005481906119ac90600160a060020a03168a611c2b565b105b156119bc5760009850611ab5565b851580156119f45750866119d88c604001518d60000151611b7a565b10806119f45750866119f28c604001518d60000151611c2b565b105b15611a025760009850611ab5565b84158015611a32575089611a1a8c606001518a611b7a565b1080611a32575089611a308c606001518a611c2b565b105b15611a405760009850611ab5565b611ab0565b86611a588c604001518d60000151611b7a565b1080611a74575086611a728c604001518d60000151611c2b565b105b80611a8b575089611a898c606001518a611b7a565b105b80611aa2575089611aa08c606001518a611c2b565b105b15611ab05760009850611ab5565b600198505b505050505050505092915050565b600154604080517f15dacbea000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301528681166024830152858116604483015260648201859052915160009392909216916315dacbea9160848082019260209290919082900301818787803b158015611b4557600080fd5b505af1158015611b59573d6000803e3d6000fd5b505050506040513d6020811015611b6f57600080fd5b505195945050505050565b600082600160a060020a03166370a0823161138761ffff16846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018082600160a060020a0316600160a060020a03168152602001915050602060405180830381600088803b158015611bf757600080fd5b5087f1158015611c0b573d6000803e3d6000fd5b50505050506040513d6020811015611c2257600080fd5b50519392505050565b600154604080517fdd62ed3e000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529283166024820152905160009285169163dd62ed3e916113879160448082019260209290919082900301818888803b158015611bf757600080fd5b6040805161016081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810191909152905600a165627a7a7230582045c14714bac80c2f2f08840c759663da43446e76b90f792633ee2130ef7d7efd0029`

// DeployExchange deploys a new Ethereum contract, binding an instance of Exchange to it.
func DeployExchange(auth *bind.TransactOpts, backend bind.ContractBackend, _zrxToken common.Address, _tokenTransferProxy common.Address) (common.Address, *types.Transaction, *Exchange, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExchangeBin), backend, _zrxToken, _tokenTransferProxy)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// EXTERNALQUERYGASLIMIT is a free data retrieval call binding the contract method 0xf06bbf75.
//
// Solidity: function EXTERNAL_QUERY_GAS_LIMIT() constant returns(uint16)
func (_Exchange *ExchangeCaller) EXTERNALQUERYGASLIMIT(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "EXTERNAL_QUERY_GAS_LIMIT")
	return *ret0, err
}

// EXTERNALQUERYGASLIMIT is a free data retrieval call binding the contract method 0xf06bbf75.
//
// Solidity: function EXTERNAL_QUERY_GAS_LIMIT() constant returns(uint16)
func (_Exchange *ExchangeSession) EXTERNALQUERYGASLIMIT() (uint16, error) {
	return _Exchange.Contract.EXTERNALQUERYGASLIMIT(&_Exchange.CallOpts)
}

// EXTERNALQUERYGASLIMIT is a free data retrieval call binding the contract method 0xf06bbf75.
//
// Solidity: function EXTERNAL_QUERY_GAS_LIMIT() constant returns(uint16)
func (_Exchange *ExchangeCallerSession) EXTERNALQUERYGASLIMIT() (uint16, error) {
	return _Exchange.Contract.EXTERNALQUERYGASLIMIT(&_Exchange.CallOpts)
}

// TOKENTRANSFERPROXYCONTRACT is a free data retrieval call binding the contract method 0xadd1cbc5.
//
// Solidity: function TOKEN_TRANSFER_PROXY_CONTRACT() constant returns(address)
func (_Exchange *ExchangeCaller) TOKENTRANSFERPROXYCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "TOKEN_TRANSFER_PROXY_CONTRACT")
	return *ret0, err
}

// TOKENTRANSFERPROXYCONTRACT is a free data retrieval call binding the contract method 0xadd1cbc5.
//
// Solidity: function TOKEN_TRANSFER_PROXY_CONTRACT() constant returns(address)
func (_Exchange *ExchangeSession) TOKENTRANSFERPROXYCONTRACT() (common.Address, error) {
	return _Exchange.Contract.TOKENTRANSFERPROXYCONTRACT(&_Exchange.CallOpts)
}

// TOKENTRANSFERPROXYCONTRACT is a free data retrieval call binding the contract method 0xadd1cbc5.
//
// Solidity: function TOKEN_TRANSFER_PROXY_CONTRACT() constant returns(address)
func (_Exchange *ExchangeCallerSession) TOKENTRANSFERPROXYCONTRACT() (common.Address, error) {
	return _Exchange.Contract.TOKENTRANSFERPROXYCONTRACT(&_Exchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Exchange *ExchangeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Exchange *ExchangeSession) VERSION() (string, error) {
	return _Exchange.Contract.VERSION(&_Exchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Exchange *ExchangeCallerSession) VERSION() (string, error) {
	return _Exchange.Contract.VERSION(&_Exchange.CallOpts)
}

// ZRXTOKENCONTRACT is a free data retrieval call binding the contract method 0x3b30ba59.
//
// Solidity: function ZRX_TOKEN_CONTRACT() constant returns(address)
func (_Exchange *ExchangeCaller) ZRXTOKENCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "ZRX_TOKEN_CONTRACT")
	return *ret0, err
}

// ZRXTOKENCONTRACT is a free data retrieval call binding the contract method 0x3b30ba59.
//
// Solidity: function ZRX_TOKEN_CONTRACT() constant returns(address)
func (_Exchange *ExchangeSession) ZRXTOKENCONTRACT() (common.Address, error) {
	return _Exchange.Contract.ZRXTOKENCONTRACT(&_Exchange.CallOpts)
}

// ZRXTOKENCONTRACT is a free data retrieval call binding the contract method 0x3b30ba59.
//
// Solidity: function ZRX_TOKEN_CONTRACT() constant returns(address)
func (_Exchange *ExchangeCallerSession) ZRXTOKENCONTRACT() (common.Address, error) {
	return _Exchange.Contract.ZRXTOKENCONTRACT(&_Exchange.CallOpts)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeCaller) Cancelled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "cancelled", arg0)
	return *ret0, err
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeSession) Cancelled(arg0 [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Cancelled(&_Exchange.CallOpts, arg0)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) Cancelled(arg0 [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Cancelled(&_Exchange.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "filled", arg0)
	return *ret0, err
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Filled(&_Exchange.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Filled(&_Exchange.CallOpts, arg0)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xcfc4d0ec.
//
// Solidity: function getOrderHash(orderAddresses address[5], orderValues uint256[6]) constant returns(bytes32)
func (_Exchange *ExchangeCaller) GetOrderHash(opts *bind.CallOpts, orderAddresses [5]common.Address, orderValues [6]*big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getOrderHash", orderAddresses, orderValues)
	return *ret0, err
}

// GetOrderHash is a free data retrieval call binding the contract method 0xcfc4d0ec.
//
// Solidity: function getOrderHash(orderAddresses address[5], orderValues uint256[6]) constant returns(bytes32)
func (_Exchange *ExchangeSession) GetOrderHash(orderAddresses [5]common.Address, orderValues [6]*big.Int) ([32]byte, error) {
	return _Exchange.Contract.GetOrderHash(&_Exchange.CallOpts, orderAddresses, orderValues)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xcfc4d0ec.
//
// Solidity: function getOrderHash(orderAddresses address[5], orderValues uint256[6]) constant returns(bytes32)
func (_Exchange *ExchangeCallerSession) GetOrderHash(orderAddresses [5]common.Address, orderValues [6]*big.Int) ([32]byte, error) {
	return _Exchange.Contract.GetOrderHash(&_Exchange.CallOpts, orderAddresses, orderValues)
}

// GetPartialAmount is a free data retrieval call binding the contract method 0x98024a8b.
//
// Solidity: function getPartialAmount(numerator uint256, denominator uint256, target uint256) constant returns(uint256)
func (_Exchange *ExchangeCaller) GetPartialAmount(opts *bind.CallOpts, numerator *big.Int, denominator *big.Int, target *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getPartialAmount", numerator, denominator, target)
	return *ret0, err
}

// GetPartialAmount is a free data retrieval call binding the contract method 0x98024a8b.
//
// Solidity: function getPartialAmount(numerator uint256, denominator uint256, target uint256) constant returns(uint256)
func (_Exchange *ExchangeSession) GetPartialAmount(numerator *big.Int, denominator *big.Int, target *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetPartialAmount(&_Exchange.CallOpts, numerator, denominator, target)
}

// GetPartialAmount is a free data retrieval call binding the contract method 0x98024a8b.
//
// Solidity: function getPartialAmount(numerator uint256, denominator uint256, target uint256) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) GetPartialAmount(numerator *big.Int, denominator *big.Int, target *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetPartialAmount(&_Exchange.CallOpts, numerator, denominator, target)
}

// GetUnavailableTakerTokenAmount is a free data retrieval call binding the contract method 0x7e9abb50.
//
// Solidity: function getUnavailableTakerTokenAmount(orderHash bytes32) constant returns(uint256)
func (_Exchange *ExchangeCaller) GetUnavailableTakerTokenAmount(opts *bind.CallOpts, orderHash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getUnavailableTakerTokenAmount", orderHash)
	return *ret0, err
}

// GetUnavailableTakerTokenAmount is a free data retrieval call binding the contract method 0x7e9abb50.
//
// Solidity: function getUnavailableTakerTokenAmount(orderHash bytes32) constant returns(uint256)
func (_Exchange *ExchangeSession) GetUnavailableTakerTokenAmount(orderHash [32]byte) (*big.Int, error) {
	return _Exchange.Contract.GetUnavailableTakerTokenAmount(&_Exchange.CallOpts, orderHash)
}

// GetUnavailableTakerTokenAmount is a free data retrieval call binding the contract method 0x7e9abb50.
//
// Solidity: function getUnavailableTakerTokenAmount(orderHash bytes32) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) GetUnavailableTakerTokenAmount(orderHash [32]byte) (*big.Int, error) {
	return _Exchange.Contract.GetUnavailableTakerTokenAmount(&_Exchange.CallOpts, orderHash)
}

// IsRoundingError is a free data retrieval call binding the contract method 0x14df96ee.
//
// Solidity: function isRoundingError(numerator uint256, denominator uint256, target uint256) constant returns(bool)
func (_Exchange *ExchangeCaller) IsRoundingError(opts *bind.CallOpts, numerator *big.Int, denominator *big.Int, target *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "isRoundingError", numerator, denominator, target)
	return *ret0, err
}

// IsRoundingError is a free data retrieval call binding the contract method 0x14df96ee.
//
// Solidity: function isRoundingError(numerator uint256, denominator uint256, target uint256) constant returns(bool)
func (_Exchange *ExchangeSession) IsRoundingError(numerator *big.Int, denominator *big.Int, target *big.Int) (bool, error) {
	return _Exchange.Contract.IsRoundingError(&_Exchange.CallOpts, numerator, denominator, target)
}

// IsRoundingError is a free data retrieval call binding the contract method 0x14df96ee.
//
// Solidity: function isRoundingError(numerator uint256, denominator uint256, target uint256) constant returns(bool)
func (_Exchange *ExchangeCallerSession) IsRoundingError(numerator *big.Int, denominator *big.Int, target *big.Int) (bool, error) {
	return _Exchange.Contract.IsRoundingError(&_Exchange.CallOpts, numerator, denominator, target)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x8163681e.
//
// Solidity: function isValidSignature(signer address, hash bytes32, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_Exchange *ExchangeCaller) IsValidSignature(opts *bind.CallOpts, signer common.Address, hash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "isValidSignature", signer, hash, v, r, s)
	return *ret0, err
}

// IsValidSignature is a free data retrieval call binding the contract method 0x8163681e.
//
// Solidity: function isValidSignature(signer address, hash bytes32, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_Exchange *ExchangeSession) IsValidSignature(signer common.Address, hash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Exchange.Contract.IsValidSignature(&_Exchange.CallOpts, signer, hash, v, r, s)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x8163681e.
//
// Solidity: function isValidSignature(signer address, hash bytes32, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_Exchange *ExchangeCallerSession) IsValidSignature(signer common.Address, hash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Exchange.Contract.IsValidSignature(&_Exchange.CallOpts, signer, hash, v, r, s)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xbaa0181d.
//
// Solidity: function batchCancelOrders(orderAddresses address[5][], orderValues uint256[6][], cancelTakerTokenAmounts uint256[]) returns()
func (_Exchange *ExchangeTransactor) BatchCancelOrders(opts *bind.TransactOpts, orderAddresses [][5]common.Address, orderValues [][6]*big.Int, cancelTakerTokenAmounts []*big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "batchCancelOrders", orderAddresses, orderValues, cancelTakerTokenAmounts)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xbaa0181d.
//
// Solidity: function batchCancelOrders(orderAddresses address[5][], orderValues uint256[6][], cancelTakerTokenAmounts uint256[]) returns()
func (_Exchange *ExchangeSession) BatchCancelOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, cancelTakerTokenAmounts []*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.BatchCancelOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, cancelTakerTokenAmounts)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xbaa0181d.
//
// Solidity: function batchCancelOrders(orderAddresses address[5][], orderValues uint256[6][], cancelTakerTokenAmounts uint256[]) returns()
func (_Exchange *ExchangeTransactorSession) BatchCancelOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, cancelTakerTokenAmounts []*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.BatchCancelOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, cancelTakerTokenAmounts)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0x4f150787.
//
// Solidity: function batchFillOrKillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeTransactor) BatchFillOrKillOrders(opts *bind.TransactOpts, orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "batchFillOrKillOrders", orderAddresses, orderValues, fillTakerTokenAmounts, v, r, s)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0x4f150787.
//
// Solidity: function batchFillOrKillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeSession) BatchFillOrKillOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.BatchFillOrKillOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmounts, v, r, s)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0x4f150787.
//
// Solidity: function batchFillOrKillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeTransactorSession) BatchFillOrKillOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.BatchFillOrKillOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmounts, v, r, s)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0xb7b2c7d6.
//
// Solidity: function batchFillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeTransactor) BatchFillOrders(opts *bind.TransactOpts, orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "batchFillOrders", orderAddresses, orderValues, fillTakerTokenAmounts, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0xb7b2c7d6.
//
// Solidity: function batchFillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeSession) BatchFillOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.BatchFillOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmounts, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0xb7b2c7d6.
//
// Solidity: function batchFillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeTransactorSession) BatchFillOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.BatchFillOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmounts, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x394c21e7.
//
// Solidity: function cancelOrder(orderAddresses address[5], orderValues uint256[6], cancelTakerTokenAmount uint256) returns(uint256)
func (_Exchange *ExchangeTransactor) CancelOrder(opts *bind.TransactOpts, orderAddresses [5]common.Address, orderValues [6]*big.Int, cancelTakerTokenAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancelOrder", orderAddresses, orderValues, cancelTakerTokenAmount)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x394c21e7.
//
// Solidity: function cancelOrder(orderAddresses address[5], orderValues uint256[6], cancelTakerTokenAmount uint256) returns(uint256)
func (_Exchange *ExchangeSession) CancelOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, cancelTakerTokenAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, cancelTakerTokenAmount)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x394c21e7.
//
// Solidity: function cancelOrder(orderAddresses address[5], orderValues uint256[6], cancelTakerTokenAmount uint256) returns(uint256)
func (_Exchange *ExchangeTransactorSession) CancelOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, cancelTakerTokenAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, cancelTakerTokenAmount)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0x741bcc93.
//
// Solidity: function fillOrKillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, v uint8, r bytes32, s bytes32) returns()
func (_Exchange *ExchangeTransactor) FillOrKillOrder(opts *bind.TransactOpts, orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "fillOrKillOrder", orderAddresses, orderValues, fillTakerTokenAmount, v, r, s)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0x741bcc93.
//
// Solidity: function fillOrKillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, v uint8, r bytes32, s bytes32) returns()
func (_Exchange *ExchangeSession) FillOrKillOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrKillOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, v, r, s)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0x741bcc93.
//
// Solidity: function fillOrKillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, v uint8, r bytes32, s bytes32) returns()
func (_Exchange *ExchangeTransactorSession) FillOrKillOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrKillOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, v, r, s)
}

// FillOrder is a paid mutator transaction binding the contract method 0xbc61394a.
//
// Solidity: function fillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r bytes32, s bytes32) returns(filledTakerTokenAmount uint256)
func (_Exchange *ExchangeTransactor) FillOrder(opts *bind.TransactOpts, orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "fillOrder", orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// FillOrder is a paid mutator transaction binding the contract method 0xbc61394a.
//
// Solidity: function fillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r bytes32, s bytes32) returns(filledTakerTokenAmount uint256)
func (_Exchange *ExchangeSession) FillOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// FillOrder is a paid mutator transaction binding the contract method 0xbc61394a.
//
// Solidity: function fillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r bytes32, s bytes32) returns(filledTakerTokenAmount uint256)
func (_Exchange *ExchangeTransactorSession) FillOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// FillOrdersUpTo is a paid mutator transaction binding the contract method 0x363349be.
//
// Solidity: function fillOrdersUpTo(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns(uint256)
func (_Exchange *ExchangeTransactor) FillOrdersUpTo(opts *bind.TransactOpts, orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "fillOrdersUpTo", orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// FillOrdersUpTo is a paid mutator transaction binding the contract method 0x363349be.
//
// Solidity: function fillOrdersUpTo(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns(uint256)
func (_Exchange *ExchangeSession) FillOrdersUpTo(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrdersUpTo(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// FillOrdersUpTo is a paid mutator transaction binding the contract method 0x363349be.
//
// Solidity: function fillOrdersUpTo(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns(uint256)
func (_Exchange *ExchangeTransactorSession) FillOrdersUpTo(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrdersUpTo(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// ExchangeLogCancelIterator is returned from FilterLogCancel and is used to iterate over the raw logs and unpacked data for LogCancel events raised by the Exchange contract.
type ExchangeLogCancelIterator struct {
	Event *ExchangeLogCancel // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeLogCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogCancel)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeLogCancel)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeLogCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogCancel represents a LogCancel event raised by the Exchange contract.
type ExchangeLogCancel struct {
	Maker                     common.Address
	FeeRecipient              common.Address
	MakerToken                common.Address
	TakerToken                common.Address
	CancelledMakerTokenAmount *big.Int
	CancelledTakerTokenAmount *big.Int
	Tokens                    [32]byte
	OrderHash                 [32]byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterLogCancel is a free log retrieval operation binding the contract event 0x67d66f160bc93d925d05dae1794c90d2d6d6688b29b84ff069398a9b04587131.
//
// Solidity: event LogCancel(maker indexed address, feeRecipient indexed address, makerToken address, takerToken address, cancelledMakerTokenAmount uint256, cancelledTakerTokenAmount uint256, tokens indexed bytes32, orderHash bytes32)
func (_Exchange *ExchangeFilterer) FilterLogCancel(opts *bind.FilterOpts, maker []common.Address, feeRecipient []common.Address, tokens [][32]byte) (*ExchangeLogCancelIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogCancel", makerRule, feeRecipientRule, tokensRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeLogCancelIterator{contract: _Exchange.contract, event: "LogCancel", logs: logs, sub: sub}, nil
}

// WatchLogCancel is a free log subscription operation binding the contract event 0x67d66f160bc93d925d05dae1794c90d2d6d6688b29b84ff069398a9b04587131.
//
// Solidity: event LogCancel(maker indexed address, feeRecipient indexed address, makerToken address, takerToken address, cancelledMakerTokenAmount uint256, cancelledTakerTokenAmount uint256, tokens indexed bytes32, orderHash bytes32)
func (_Exchange *ExchangeFilterer) WatchLogCancel(opts *bind.WatchOpts, sink chan<- *ExchangeLogCancel, maker []common.Address, feeRecipient []common.Address, tokens [][32]byte) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogCancel", makerRule, feeRecipientRule, tokensRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogCancel)
				if err := _Exchange.contract.UnpackLog(event, "LogCancel", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeLogErrorIterator is returned from FilterLogError and is used to iterate over the raw logs and unpacked data for LogError events raised by the Exchange contract.
type ExchangeLogErrorIterator struct {
	Event *ExchangeLogError // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeLogErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogError)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeLogError)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeLogErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogError represents a LogError event raised by the Exchange contract.
type ExchangeLogError struct {
	ErrorId   uint8
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogError is a free log retrieval operation binding the contract event 0x36d86c59e00bd73dc19ba3adfe068e4b64ac7e92be35546adeddf1b956a87e90.
//
// Solidity: event LogError(errorId indexed uint8, orderHash indexed bytes32)
func (_Exchange *ExchangeFilterer) FilterLogError(opts *bind.FilterOpts, errorId []uint8, orderHash [][32]byte) (*ExchangeLogErrorIterator, error) {

	var errorIdRule []interface{}
	for _, errorIdItem := range errorId {
		errorIdRule = append(errorIdRule, errorIdItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogError", errorIdRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeLogErrorIterator{contract: _Exchange.contract, event: "LogError", logs: logs, sub: sub}, nil
}

// WatchLogError is a free log subscription operation binding the contract event 0x36d86c59e00bd73dc19ba3adfe068e4b64ac7e92be35546adeddf1b956a87e90.
//
// Solidity: event LogError(errorId indexed uint8, orderHash indexed bytes32)
func (_Exchange *ExchangeFilterer) WatchLogError(opts *bind.WatchOpts, sink chan<- *ExchangeLogError, errorId []uint8, orderHash [][32]byte) (event.Subscription, error) {

	var errorIdRule []interface{}
	for _, errorIdItem := range errorId {
		errorIdRule = append(errorIdRule, errorIdItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogError", errorIdRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogError)
				if err := _Exchange.contract.UnpackLog(event, "LogError", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ExchangeLogFillIterator is returned from FilterLogFill and is used to iterate over the raw logs and unpacked data for LogFill events raised by the Exchange contract.
type ExchangeLogFillIterator struct {
	Event *ExchangeLogFill // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ExchangeLogFillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogFill)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ExchangeLogFill)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ExchangeLogFillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogFill represents a LogFill event raised by the Exchange contract.
type ExchangeLogFill struct {
	Maker                  common.Address
	Taker                  common.Address
	FeeRecipient           common.Address
	MakerToken             common.Address
	TakerToken             common.Address
	FilledMakerTokenAmount *big.Int
	FilledTakerTokenAmount *big.Int
	PaidMakerFee           *big.Int
	PaidTakerFee           *big.Int
	Tokens                 [32]byte
	OrderHash              [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterLogFill is a free log retrieval operation binding the contract event 0x0d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb3.
//
// Solidity: event LogFill(maker indexed address, taker address, feeRecipient indexed address, makerToken address, takerToken address, filledMakerTokenAmount uint256, filledTakerTokenAmount uint256, paidMakerFee uint256, paidTakerFee uint256, tokens indexed bytes32, orderHash bytes32)
func (_Exchange *ExchangeFilterer) FilterLogFill(opts *bind.FilterOpts, maker []common.Address, feeRecipient []common.Address, tokens [][32]byte) (*ExchangeLogFillIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogFill", makerRule, feeRecipientRule, tokensRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeLogFillIterator{contract: _Exchange.contract, event: "LogFill", logs: logs, sub: sub}, nil
}

// WatchLogFill is a free log subscription operation binding the contract event 0x0d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb3.
//
// Solidity: event LogFill(maker indexed address, taker address, feeRecipient indexed address, makerToken address, takerToken address, filledMakerTokenAmount uint256, filledTakerTokenAmount uint256, paidMakerFee uint256, paidTakerFee uint256, tokens indexed bytes32, orderHash bytes32)
func (_Exchange *ExchangeFilterer) WatchLogFill(opts *bind.WatchOpts, sink chan<- *ExchangeLogFill, maker []common.Address, feeRecipient []common.Address, tokens [][32]byte) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogFill", makerRule, feeRecipientRule, tokensRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogFill)
				if err := _Exchange.contract.UnpackLog(event, "LogFill", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a033316600160a060020a03199091161790556101838061003d6000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461008e575b600080fd5b34801561005c57600080fd5b506100656100be565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009a57600080fd5b506100bc73ffffffffffffffffffffffffffffffffffffffff600435166100da565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff90811691161461010257600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615610154576000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b505600a165627a7a72305820ae465235618fb4f72de129241c31dd67397ea660b294075f43d9c0132c58c57b0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a723058204657073b0d876a80c5a1526774f9a6a2a87eeacad36c7686fa8914081fe6ade30029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x608060405234801561001057600080fd5b506101a0806100206000396000f30060806040526004361061006c5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461007157806323b872dd146100b657806370a08231146100ed578063a9059cbb14610071578063dd62ed3e1461012d575b600080fd5b34801561007d57600080fd5b506100a273ffffffffffffffffffffffffffffffffffffffff6004351660243561015d565b604080519115158252519081900360200190f35b3480156100c257600080fd5b506100a273ffffffffffffffffffffffffffffffffffffffff60043581169060243516604435610165565b3480156100f957600080fd5b5061011b73ffffffffffffffffffffffffffffffffffffffff6004351661016e565b60408051918252519081900360200190f35b34801561013957600080fd5b5061011b73ffffffffffffffffffffffffffffffffffffffff600435811690602435165b600092915050565b60009392505050565b506000905600a165627a7a723058208f4ac20e0cb525d708d87e1d8816b60a4de9b91769e05743ad3b2e4863c0d2eb0029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenTransferProxyABI is the input ABI used to generate the binding from.
const TokenTransferProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"addAuthorizedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"authorities\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"target\",\"type\":\"address\"}],\"name\":\"removeAuthorizedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAuthorizedAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"LogAuthorizedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"LogAuthorizedAddressRemoved\",\"type\":\"event\"}]"

// TokenTransferProxyBin is the compiled bytecode used for deploying new contracts.
const TokenTransferProxyBin = `0x608060405260008054600160a060020a033316600160a060020a03199091161790556106b2806100306000396000f30060806040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166315dacbea811461009257806342f1181e146100d6578063494503d4146100f9578063707129391461012d5780638da5cb5b1461014e578063b918161114610163578063d39de6e914610184578063f2fde38b146101e9575b600080fd5b34801561009e57600080fd5b506100c2600160a060020a036004358116906024358116906044351660643561020a565b604080519115158252519081900360200190f35b3480156100e257600080fd5b506100f7600160a060020a03600435166102da565b005b34801561010557600080fd5b506101116004356103bd565b60408051600160a060020a039092168252519081900360200190f35b34801561013957600080fd5b506100f7600160a060020a03600435166103e5565b34801561015a57600080fd5b50610111610562565b34801561016f57600080fd5b506100c2600160a060020a0360043516610571565b34801561019057600080fd5b50610199610586565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101d55781810151838201526020016101bd565b505050509050019250505060405180910390f35b3480156101f557600080fd5b506100f7600160a060020a03600435166105e9565b600160a060020a03331660009081526001602052604081205460ff16151561023157600080fd5b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301528581166024830152604482018590529151918716916323b872dd916064808201926020929091908290030181600087803b1580156102a557600080fd5b505af11580156102b9573d6000803e3d6000fd5b505050506040513d60208110156102cf57600080fd5b505195945050505050565b60005433600160a060020a039081169116146102f557600080fd5b600160a060020a038116600090815260016020526040902054819060ff161561031d57600080fd5b600160a060020a038083166000818152600160208190526040808320805460ff19168317905560028054928301815583527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace909101805473ffffffffffffffffffffffffffffffffffffffff1916841790555133909316927f94bb87f4c15c4587ff559a7584006fa01ddf9299359be6b512b94527aa961aca9190a35050565b60028054829081106103cb57fe5b600091825260209091200154600160a060020a0316905081565b6000805433600160a060020a0390811691161461040157600080fd5b600160a060020a038216600090815260016020526040902054829060ff16151561042a57600080fd5b600160a060020a0383166000908152600160205260408120805460ff1916905591505b60025482101561051d5782600160a060020a031660028381548110151561047057fe5b600091825260209091200154600160a060020a031614156105125760028054600019810190811061049d57fe5b60009182526020909120015460028054600160a060020a0390921691849081106104c357fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560028054600019019061050c908261063f565b5061051d565b60019091019061044d565b33600160a060020a031683600160a060020a03167ff5b347a1e40749dd050f5f07fbdbeb7e3efa9756903044dd29401fd1d4bb4a1c60405160405180910390a3505050565b600054600160a060020a031681565b60016020526000908152604090205460ff1681565b606060028054806020026020016040519081016040528092919081815260200182805480156105de57602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116105c0575b505050505090505b90565b60005433600160a060020a0390811691161461060457600080fd5b600160a060020a0381161561063c576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b81548183558181111561066357600083815260209020610663918101908301610668565b505050565b6105e691905b80821115610682576000815560010161066e565b50905600a165627a7a723058203835161910301010d7b796f9ac6ce1b793363687e2ad40197aa5941e0c9f3eb80029`

// DeployTokenTransferProxy deploys a new Ethereum contract, binding an instance of TokenTransferProxy to it.
func DeployTokenTransferProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenTransferProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenTransferProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenTransferProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenTransferProxy{TokenTransferProxyCaller: TokenTransferProxyCaller{contract: contract}, TokenTransferProxyTransactor: TokenTransferProxyTransactor{contract: contract}, TokenTransferProxyFilterer: TokenTransferProxyFilterer{contract: contract}}, nil
}

// TokenTransferProxy is an auto generated Go binding around an Ethereum contract.
type TokenTransferProxy struct {
	TokenTransferProxyCaller     // Read-only binding to the contract
	TokenTransferProxyTransactor // Write-only binding to the contract
	TokenTransferProxyFilterer   // Log filterer for contract events
}

// TokenTransferProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenTransferProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransferProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransferProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransferProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenTransferProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransferProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenTransferProxySession struct {
	Contract     *TokenTransferProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TokenTransferProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenTransferProxyCallerSession struct {
	Contract *TokenTransferProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TokenTransferProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransferProxyTransactorSession struct {
	Contract     *TokenTransferProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TokenTransferProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenTransferProxyRaw struct {
	Contract *TokenTransferProxy // Generic contract binding to access the raw methods on
}

// TokenTransferProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenTransferProxyCallerRaw struct {
	Contract *TokenTransferProxyCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransferProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransferProxyTransactorRaw struct {
	Contract *TokenTransferProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenTransferProxy creates a new instance of TokenTransferProxy, bound to a specific deployed contract.
func NewTokenTransferProxy(address common.Address, backend bind.ContractBackend) (*TokenTransferProxy, error) {
	contract, err := bindTokenTransferProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenTransferProxy{TokenTransferProxyCaller: TokenTransferProxyCaller{contract: contract}, TokenTransferProxyTransactor: TokenTransferProxyTransactor{contract: contract}, TokenTransferProxyFilterer: TokenTransferProxyFilterer{contract: contract}}, nil
}

// NewTokenTransferProxyCaller creates a new read-only instance of TokenTransferProxy, bound to a specific deployed contract.
func NewTokenTransferProxyCaller(address common.Address, caller bind.ContractCaller) (*TokenTransferProxyCaller, error) {
	contract, err := bindTokenTransferProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransferProxyCaller{contract: contract}, nil
}

// NewTokenTransferProxyTransactor creates a new write-only instance of TokenTransferProxy, bound to a specific deployed contract.
func NewTokenTransferProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransferProxyTransactor, error) {
	contract, err := bindTokenTransferProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransferProxyTransactor{contract: contract}, nil
}

// NewTokenTransferProxyFilterer creates a new log filterer instance of TokenTransferProxy, bound to a specific deployed contract.
func NewTokenTransferProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenTransferProxyFilterer, error) {
	contract, err := bindTokenTransferProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenTransferProxyFilterer{contract: contract}, nil
}

// bindTokenTransferProxy binds a generic wrapper to an already deployed contract.
func bindTokenTransferProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenTransferProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenTransferProxy *TokenTransferProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenTransferProxy.Contract.TokenTransferProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenTransferProxy *TokenTransferProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.TokenTransferProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenTransferProxy *TokenTransferProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.TokenTransferProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenTransferProxy *TokenTransferProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenTransferProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenTransferProxy *TokenTransferProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenTransferProxy *TokenTransferProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.contract.Transact(opts, method, params...)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_TokenTransferProxy *TokenTransferProxyCaller) Authorities(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenTransferProxy.contract.Call(opts, out, "authorities", arg0)
	return *ret0, err
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_TokenTransferProxy *TokenTransferProxySession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _TokenTransferProxy.Contract.Authorities(&_TokenTransferProxy.CallOpts, arg0)
}

// Authorities is a free data retrieval call binding the contract method 0x494503d4.
//
// Solidity: function authorities( uint256) constant returns(address)
func (_TokenTransferProxy *TokenTransferProxyCallerSession) Authorities(arg0 *big.Int) (common.Address, error) {
	return _TokenTransferProxy.Contract.Authorities(&_TokenTransferProxy.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized( address) constant returns(bool)
func (_TokenTransferProxy *TokenTransferProxyCaller) Authorized(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenTransferProxy.contract.Call(opts, out, "authorized", arg0)
	return *ret0, err
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized( address) constant returns(bool)
func (_TokenTransferProxy *TokenTransferProxySession) Authorized(arg0 common.Address) (bool, error) {
	return _TokenTransferProxy.Contract.Authorized(&_TokenTransferProxy.CallOpts, arg0)
}

// Authorized is a free data retrieval call binding the contract method 0xb9181611.
//
// Solidity: function authorized( address) constant returns(bool)
func (_TokenTransferProxy *TokenTransferProxyCallerSession) Authorized(arg0 common.Address) (bool, error) {
	return _TokenTransferProxy.Contract.Authorized(&_TokenTransferProxy.CallOpts, arg0)
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_TokenTransferProxy *TokenTransferProxyCaller) GetAuthorizedAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _TokenTransferProxy.contract.Call(opts, out, "getAuthorizedAddresses")
	return *ret0, err
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_TokenTransferProxy *TokenTransferProxySession) GetAuthorizedAddresses() ([]common.Address, error) {
	return _TokenTransferProxy.Contract.GetAuthorizedAddresses(&_TokenTransferProxy.CallOpts)
}

// GetAuthorizedAddresses is a free data retrieval call binding the contract method 0xd39de6e9.
//
// Solidity: function getAuthorizedAddresses() constant returns(address[])
func (_TokenTransferProxy *TokenTransferProxyCallerSession) GetAuthorizedAddresses() ([]common.Address, error) {
	return _TokenTransferProxy.Contract.GetAuthorizedAddresses(&_TokenTransferProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenTransferProxy *TokenTransferProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenTransferProxy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenTransferProxy *TokenTransferProxySession) Owner() (common.Address, error) {
	return _TokenTransferProxy.Contract.Owner(&_TokenTransferProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenTransferProxy *TokenTransferProxyCallerSession) Owner() (common.Address, error) {
	return _TokenTransferProxy.Contract.Owner(&_TokenTransferProxy.CallOpts)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(target address) returns()
func (_TokenTransferProxy *TokenTransferProxyTransactor) AddAuthorizedAddress(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _TokenTransferProxy.contract.Transact(opts, "addAuthorizedAddress", target)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(target address) returns()
func (_TokenTransferProxy *TokenTransferProxySession) AddAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.AddAuthorizedAddress(&_TokenTransferProxy.TransactOpts, target)
}

// AddAuthorizedAddress is a paid mutator transaction binding the contract method 0x42f1181e.
//
// Solidity: function addAuthorizedAddress(target address) returns()
func (_TokenTransferProxy *TokenTransferProxyTransactorSession) AddAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.AddAuthorizedAddress(&_TokenTransferProxy.TransactOpts, target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(target address) returns()
func (_TokenTransferProxy *TokenTransferProxyTransactor) RemoveAuthorizedAddress(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _TokenTransferProxy.contract.Transact(opts, "removeAuthorizedAddress", target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(target address) returns()
func (_TokenTransferProxy *TokenTransferProxySession) RemoveAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.RemoveAuthorizedAddress(&_TokenTransferProxy.TransactOpts, target)
}

// RemoveAuthorizedAddress is a paid mutator transaction binding the contract method 0x70712939.
//
// Solidity: function removeAuthorizedAddress(target address) returns()
func (_TokenTransferProxy *TokenTransferProxyTransactorSession) RemoveAuthorizedAddress(target common.Address) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.RemoveAuthorizedAddress(&_TokenTransferProxy.TransactOpts, target)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(token address, from address, to address, value uint256) returns(bool)
func (_TokenTransferProxy *TokenTransferProxyTransactor) TransferFrom(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenTransferProxy.contract.Transact(opts, "transferFrom", token, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(token address, from address, to address, value uint256) returns(bool)
func (_TokenTransferProxy *TokenTransferProxySession) TransferFrom(token common.Address, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.TransferFrom(&_TokenTransferProxy.TransactOpts, token, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(token address, from address, to address, value uint256) returns(bool)
func (_TokenTransferProxy *TokenTransferProxyTransactorSession) TransferFrom(token common.Address, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.TransferFrom(&_TokenTransferProxy.TransactOpts, token, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenTransferProxy *TokenTransferProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenTransferProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenTransferProxy *TokenTransferProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.TransferOwnership(&_TokenTransferProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenTransferProxy *TokenTransferProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenTransferProxy.Contract.TransferOwnership(&_TokenTransferProxy.TransactOpts, newOwner)
}

// TokenTransferProxyLogAuthorizedAddressAddedIterator is returned from FilterLogAuthorizedAddressAdded and is used to iterate over the raw logs and unpacked data for LogAuthorizedAddressAdded events raised by the TokenTransferProxy contract.
type TokenTransferProxyLogAuthorizedAddressAddedIterator struct {
	Event *TokenTransferProxyLogAuthorizedAddressAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenTransferProxyLogAuthorizedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransferProxyLogAuthorizedAddressAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenTransferProxyLogAuthorizedAddressAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenTransferProxyLogAuthorizedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferProxyLogAuthorizedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransferProxyLogAuthorizedAddressAdded represents a LogAuthorizedAddressAdded event raised by the TokenTransferProxy contract.
type TokenTransferProxyLogAuthorizedAddressAdded struct {
	Target common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogAuthorizedAddressAdded is a free log retrieval operation binding the contract event 0x94bb87f4c15c4587ff559a7584006fa01ddf9299359be6b512b94527aa961aca.
//
// Solidity: event LogAuthorizedAddressAdded(target indexed address, caller indexed address)
func (_TokenTransferProxy *TokenTransferProxyFilterer) FilterLogAuthorizedAddressAdded(opts *bind.FilterOpts, target []common.Address, caller []common.Address) (*TokenTransferProxyLogAuthorizedAddressAddedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenTransferProxy.contract.FilterLogs(opts, "LogAuthorizedAddressAdded", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferProxyLogAuthorizedAddressAddedIterator{contract: _TokenTransferProxy.contract, event: "LogAuthorizedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchLogAuthorizedAddressAdded is a free log subscription operation binding the contract event 0x94bb87f4c15c4587ff559a7584006fa01ddf9299359be6b512b94527aa961aca.
//
// Solidity: event LogAuthorizedAddressAdded(target indexed address, caller indexed address)
func (_TokenTransferProxy *TokenTransferProxyFilterer) WatchLogAuthorizedAddressAdded(opts *bind.WatchOpts, sink chan<- *TokenTransferProxyLogAuthorizedAddressAdded, target []common.Address, caller []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenTransferProxy.contract.WatchLogs(opts, "LogAuthorizedAddressAdded", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransferProxyLogAuthorizedAddressAdded)
				if err := _TokenTransferProxy.contract.UnpackLog(event, "LogAuthorizedAddressAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// TokenTransferProxyLogAuthorizedAddressRemovedIterator is returned from FilterLogAuthorizedAddressRemoved and is used to iterate over the raw logs and unpacked data for LogAuthorizedAddressRemoved events raised by the TokenTransferProxy contract.
type TokenTransferProxyLogAuthorizedAddressRemovedIterator struct {
	Event *TokenTransferProxyLogAuthorizedAddressRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TokenTransferProxyLogAuthorizedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransferProxyLogAuthorizedAddressRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TokenTransferProxyLogAuthorizedAddressRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TokenTransferProxyLogAuthorizedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferProxyLogAuthorizedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransferProxyLogAuthorizedAddressRemoved represents a LogAuthorizedAddressRemoved event raised by the TokenTransferProxy contract.
type TokenTransferProxyLogAuthorizedAddressRemoved struct {
	Target common.Address
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogAuthorizedAddressRemoved is a free log retrieval operation binding the contract event 0xf5b347a1e40749dd050f5f07fbdbeb7e3efa9756903044dd29401fd1d4bb4a1c.
//
// Solidity: event LogAuthorizedAddressRemoved(target indexed address, caller indexed address)
func (_TokenTransferProxy *TokenTransferProxyFilterer) FilterLogAuthorizedAddressRemoved(opts *bind.FilterOpts, target []common.Address, caller []common.Address) (*TokenTransferProxyLogAuthorizedAddressRemovedIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenTransferProxy.contract.FilterLogs(opts, "LogAuthorizedAddressRemoved", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferProxyLogAuthorizedAddressRemovedIterator{contract: _TokenTransferProxy.contract, event: "LogAuthorizedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchLogAuthorizedAddressRemoved is a free log subscription operation binding the contract event 0xf5b347a1e40749dd050f5f07fbdbeb7e3efa9756903044dd29401fd1d4bb4a1c.
//
// Solidity: event LogAuthorizedAddressRemoved(target indexed address, caller indexed address)
func (_TokenTransferProxy *TokenTransferProxyFilterer) WatchLogAuthorizedAddressRemoved(opts *bind.WatchOpts, sink chan<- *TokenTransferProxyLogAuthorizedAddressRemoved, target []common.Address, caller []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenTransferProxy.contract.WatchLogs(opts, "LogAuthorizedAddressRemoved", targetRule, callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransferProxyLogAuthorizedAddressRemoved)
				if err := _TokenTransferProxy.contract.UnpackLog(event, "LogAuthorizedAddressRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
