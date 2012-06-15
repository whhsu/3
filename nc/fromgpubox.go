package nc

// Copies from gpu.

import (
	"github.com/barnex/cuda4/cu"
	"unsafe"
)

type FromGpuBox struct {
	Input  <-chan GpuFloats
	Output []chan<- []float32
	stream cu.Stream
}

func NewFromGpuBox() *FromGpuBox {
	box := new(FromGpuBox)
	box.stream = cu.StreamCreate()
	Register(box)
	return box
}

func (box *FromGpuBox) Run() {
	for {
		in := RecvGpu(box.Input)
		sendToHost(in, box.Output, box.stream)
	}
}

func sendToHost(in GpuFloats, out []chan<- []float32, stream cu.Stream) {
	buffer := Buffer()
	SetCudaCtx()
	cu.MemcpyDtoHAsync(unsafe.Pointer(&buffer[0]), cu.DevicePtr(in),
		cu.SIZEOF_FLOAT32*int64(WarpLen()), stream)
	stream.Synchronize()
	RecycleGpu(in)
	Send(out, buffer)
}
