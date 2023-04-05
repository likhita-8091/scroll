#./abigen.exe --abi scroll_abi.json --pkg scroll_abi --type scrollABI --out scroll_abi.go
./main.exe --abi scroll_abi.json --pkg abi --type scrollABI --out scroll_abi.go
./main.exe --abi L2_ABI.json --pkg abi --type L2ABI --out l2_abi.go
./main --abi usdc/usdc.abi --pkg usdc --type USDCABI --out usdc/usdc_abi.go
./main --abi usdc/test_usdc_abi.json --pkg usdc --type TestUsdcABI --out usdc/test_usdc_abi.go

./main --bin=deploy/storage.bin --abi deploy/storage.abi --pkg deploy --type StorageABI --out deploy/storage.go
./main --bin=deploy/lock.bin --abi deploy/lock.abi --pkg deploy --type LockABI --out deploy/lock.go

# 0xbd1A5920303F45d628630E88aFbAF012bA078F37
./main.exe --abi lp/lp.abi.json --pkg lp --type LPABI --out lp/lp_abi.go
./main --abi swap/swap.abi.json --pkg swap --type SwapABI --out swap/swap_abi.go
./main --abi weth/weth.abi.json --pkg weth --type WEthABI --out weth/weth_abi.go
#./abigen.exe --abi simple_abi.json --pkg scroll_abi --type simpleABI --out simple_abi.go