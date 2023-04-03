#./abigen.exe --abi scroll_abi.json --pkg scroll_abi --type scrollABI --out scroll_abi.go
./main.exe --abi scroll_abi.json --pkg abi --type scrollABI --out scroll_abi.go
./main.exe --abi L2_ABI.json --pkg abi --type L2ABI --out l2_abi.go
./main --abi usdc/usdc.abi --pkg usdc --type USDCABI --out usdc/usdc_abi.go
./main --abi usdc/test_usdc_abi.json --pkg usdc --type TestUsdcABI --out usdc/test_usdc_abi.go


./main --abi swap/swap.abi.json --pkg swap --type SwapABI --out swap/swap_abi.go
./main --abi weth/weth.abi.json --pkg weth --type WEthABI --out weth/weth_abi.go
#./abigen.exe --abi simple_abi.json --pkg scroll_abi --type simpleABI --out simple_abi.go