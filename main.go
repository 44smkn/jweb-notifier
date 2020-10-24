package main

import (
	"jweb-notifier/presentation"
)

func main() {
	// TODO: 実行時のflagをparseする
	presentation.ServerRun(8080)
	// TODO: SignalをHookしてGracefulに処理を終了させる
}
