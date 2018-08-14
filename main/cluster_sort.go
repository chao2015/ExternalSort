package main

import (
	"externalsort/pipeline"
)

// STEP 2
// 对 small.in 或 large.in 文件里的数据进行排序
func main() {
	// small 512
	// large 80000000
	const filename_prefix = "small"  // small/large

	// 网络版
	p := pipeline.CreateNetworkPipeline(filename_prefix+".in", 512, 4)

	pipeline.WriteToFile(p, filename_prefix+".out")
	pipeline.PrintFile(filename_prefix + ".out")
}