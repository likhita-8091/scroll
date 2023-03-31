#./abigen.exe --abi scroll_abi.json --pkg scroll_abi --type scrollABI --out scroll_abi.go
./main.exe --abi scroll_abi.json --pkg abi --type scrollABI --out scroll_abi.go
./main.exe --abi L2_ABI.json --pkg abi --type L2ABI --out l2_abi.go

#./abigen.exe --abi simple_abi.json --pkg scroll_abi --type simpleABI --out simple_abi.go