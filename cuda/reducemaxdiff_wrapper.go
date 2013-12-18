package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"sync"
	"unsafe"
)

var reducemaxdiff_code cu.Function

type reducemaxdiff_args_t struct {
	arg_src1    unsafe.Pointer
	arg_src2    unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	argptr      [5]unsafe.Pointer
	sync.Mutex
}

var reducemaxdiff_args reducemaxdiff_args_t

func init() {
	reducemaxdiff_args.argptr[0] = unsafe.Pointer(&reducemaxdiff_args.arg_src1)
	reducemaxdiff_args.argptr[1] = unsafe.Pointer(&reducemaxdiff_args.arg_src2)
	reducemaxdiff_args.argptr[2] = unsafe.Pointer(&reducemaxdiff_args.arg_dst)
	reducemaxdiff_args.argptr[3] = unsafe.Pointer(&reducemaxdiff_args.arg_initVal)
	reducemaxdiff_args.argptr[4] = unsafe.Pointer(&reducemaxdiff_args.arg_n)

}

// Wrapper for reducemaxdiff CUDA kernel, asynchronous.
func k_reducemaxdiff_async(src1 unsafe.Pointer, src2 unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	reducemaxdiff_args.Lock()
	defer reducemaxdiff_args.Unlock()

	if reducemaxdiff_code == 0 {
		reducemaxdiff_code = fatbinLoad(reducemaxdiff_map, "reducemaxdiff")
	}

	reducemaxdiff_args.arg_src1 = src1
	reducemaxdiff_args.arg_src2 = src2
	reducemaxdiff_args.arg_dst = dst
	reducemaxdiff_args.arg_initVal = initVal
	reducemaxdiff_args.arg_n = n

	args := reducemaxdiff_args.argptr[:]
	cu.LaunchKernel(reducemaxdiff_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

var reducemaxdiff_map = map[int]string{0: "",
	20: reducemaxdiff_ptx_20,
	30: reducemaxdiff_ptx_30,
	35: reducemaxdiff_ptx_35}

const (
	reducemaxdiff_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64

.global .align 1 .b8 $str[11] = {95, 95, 67, 85, 68, 65, 95, 70, 84, 90, 0};

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<17>;
	.reg .f32 	%f<33>;
	.reg .s64 	%rd<16>;
	// demoted variable
	.shared .align 4 .b8 reducemaxdiff$__cuda_local_var_33858_35_non_const_sdata[2048];

	ld.param.u64 	%rd5, [reducemaxdiff_param_0];
	ld.param.u64 	%rd6, [reducemaxdiff_param_1];
	ld.param.u64 	%rd7, [reducemaxdiff_param_2];
	ld.param.f32 	%f32, [reducemaxdiff_param_3];
	ld.param.u32 	%r9, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd7;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 1 8 1
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r15, %r16, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r16, %r11;
	.loc 1 8 1
	setp.ge.s32	%p1, %r15, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	.loc 1 8 1
	mul.wide.s32 	%rd8, %r15, 4;
	add.s64 	%rd9, %rd3, %rd8;
	add.s64 	%rd10, %rd2, %rd8;
	ld.global.f32 	%f5, [%rd10];
	ld.global.f32 	%f6, [%rd9];
	sub.f32 	%f7, %f6, %f5;
	.loc 2 2750 10
	abs.f32 	%f8, %f7;
	.loc 2 2770 10
	max.f32 	%f32, %f32, %f8;
	.loc 1 8 1
	add.s32 	%r15, %r15, %r4;
	.loc 1 8 1
	setp.lt.s32	%p2, %r15, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	.loc 1 8 1
	mul.wide.s32 	%rd11, %r2, 4;
	mov.u64 	%rd12, reducemaxdiff$__cuda_local_var_33858_35_non_const_sdata;
	add.s64 	%rd4, %rd12, %rd11;
	st.shared.f32 	[%rd4], %f32;
	bar.sync 	0;
	.loc 1 8 1
	setp.lt.u32	%p3, %r16, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	.loc 1 8 1
	mov.u32 	%r7, %r16;
	shr.u32 	%r16, %r7, 1;
	.loc 1 8 1
	setp.ge.u32	%p4, %r2, %r16;
	@%p4 bra 	BB0_5;

	.loc 1 8 1
	ld.shared.f32 	%f9, [%rd4];
	add.s32 	%r12, %r16, %r2;
	mul.wide.u32 	%rd13, %r12, 4;
	add.s64 	%rd15, %rd12, %rd13;
	ld.shared.f32 	%f10, [%rd15];
	.loc 2 2770 10
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%rd4], %f11;

BB0_5:
	.loc 1 8 1
	bar.sync 	0;
	.loc 1 8 1
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	.loc 1 8 1
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	.loc 1 8 1
	ld.volatile.shared.f32 	%f12, [%rd4];
	ld.volatile.shared.f32 	%f13, [%rd4+128];
	.loc 2 2770 10
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%rd4], %f14;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f15, [%rd4+64];
	ld.volatile.shared.f32 	%f16, [%rd4];
	.loc 2 2770 10
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%rd4], %f17;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f18, [%rd4+32];
	ld.volatile.shared.f32 	%f19, [%rd4];
	.loc 2 2770 10
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%rd4], %f20;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f21, [%rd4+16];
	ld.volatile.shared.f32 	%f22, [%rd4];
	.loc 2 2770 10
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%rd4], %f23;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f24, [%rd4+8];
	ld.volatile.shared.f32 	%f25, [%rd4];
	.loc 2 2770 10
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%rd4], %f26;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f27, [%rd4+4];
	ld.volatile.shared.f32 	%f28, [%rd4];
	.loc 2 2770 10
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%rd4], %f29;

BB0_8:
	.loc 1 8 1
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	.loc 1 8 1
	ld.shared.f32 	%f30, [reducemaxdiff$__cuda_local_var_33858_35_non_const_sdata];
	.loc 2 2750 10
	abs.f32 	%f31, %f30;
	.loc 1 8 37
	mov.b32 	 %r13, %f31;
	.loc 2 3781 3
	atom.global.max.s32 	%r14, [%rd1], %r13;

BB0_10:
	.loc 1 9 2
	ret;
}


`
	reducemaxdiff_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64

.global .align 1 .b8 $str[11] = {95, 95, 67, 85, 68, 65, 95, 70, 84, 90, 0};

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<17>;
	.reg .f32 	%f<33>;
	.reg .s64 	%rd<16>;
	// demoted variable
	.shared .align 4 .b8 reducemaxdiff$__cuda_local_var_33931_35_non_const_sdata[2048];

	ld.param.u64 	%rd5, [reducemaxdiff_param_0];
	ld.param.u64 	%rd6, [reducemaxdiff_param_1];
	ld.param.u64 	%rd7, [reducemaxdiff_param_2];
	ld.param.f32 	%f32, [reducemaxdiff_param_3];
	ld.param.u32 	%r9, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd7;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 1 8 1
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r15, %r16, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r16, %r11;
	.loc 1 8 1
	setp.ge.s32	%p1, %r15, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	.loc 1 8 1
	mul.wide.s32 	%rd8, %r15, 4;
	add.s64 	%rd9, %rd3, %rd8;
	add.s64 	%rd10, %rd2, %rd8;
	ld.global.f32 	%f5, [%rd10];
	ld.global.f32 	%f6, [%rd9];
	sub.f32 	%f7, %f6, %f5;
	.loc 2 2750 10
	abs.f32 	%f8, %f7;
	.loc 2 2770 10
	max.f32 	%f32, %f32, %f8;
	.loc 1 8 1
	add.s32 	%r15, %r15, %r4;
	.loc 1 8 1
	setp.lt.s32	%p2, %r15, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	.loc 1 8 1
	mul.wide.s32 	%rd11, %r2, 4;
	mov.u64 	%rd12, reducemaxdiff$__cuda_local_var_33931_35_non_const_sdata;
	add.s64 	%rd4, %rd12, %rd11;
	st.shared.f32 	[%rd4], %f32;
	bar.sync 	0;
	.loc 1 8 1
	setp.lt.u32	%p3, %r16, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	.loc 1 8 1
	mov.u32 	%r7, %r16;
	shr.u32 	%r16, %r7, 1;
	.loc 1 8 1
	setp.ge.u32	%p4, %r2, %r16;
	@%p4 bra 	BB0_5;

	.loc 1 8 1
	ld.shared.f32 	%f9, [%rd4];
	add.s32 	%r12, %r16, %r2;
	mul.wide.u32 	%rd13, %r12, 4;
	add.s64 	%rd15, %rd12, %rd13;
	ld.shared.f32 	%f10, [%rd15];
	.loc 2 2770 10
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%rd4], %f11;

BB0_5:
	.loc 1 8 1
	bar.sync 	0;
	.loc 1 8 1
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	.loc 1 8 1
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	.loc 1 8 1
	ld.volatile.shared.f32 	%f12, [%rd4];
	ld.volatile.shared.f32 	%f13, [%rd4+128];
	.loc 2 2770 10
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%rd4], %f14;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f15, [%rd4+64];
	ld.volatile.shared.f32 	%f16, [%rd4];
	.loc 2 2770 10
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%rd4], %f17;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f18, [%rd4+32];
	ld.volatile.shared.f32 	%f19, [%rd4];
	.loc 2 2770 10
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%rd4], %f20;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f21, [%rd4+16];
	ld.volatile.shared.f32 	%f22, [%rd4];
	.loc 2 2770 10
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%rd4], %f23;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f24, [%rd4+8];
	ld.volatile.shared.f32 	%f25, [%rd4];
	.loc 2 2770 10
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%rd4], %f26;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f27, [%rd4+4];
	ld.volatile.shared.f32 	%f28, [%rd4];
	.loc 2 2770 10
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%rd4], %f29;

BB0_8:
	.loc 1 8 1
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	.loc 1 8 1
	ld.shared.f32 	%f30, [reducemaxdiff$__cuda_local_var_33931_35_non_const_sdata];
	.loc 2 2750 10
	abs.f32 	%f31, %f30;
	.loc 1 8 37
	mov.b32 	 %r13, %f31;
	.loc 2 3781 3
	atom.global.max.s32 	%r14, [%rd1], %r13;

BB0_10:
	.loc 1 9 2
	ret;
}


`
	reducemaxdiff_ptx_35 = `
.version 3.2
.target sm_35
.address_size 64

.global .align 1 .b8 $str[11] = {95, 95, 67, 85, 68, 65, 95, 70, 84, 90, 0};

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<17>;
	.reg .f32 	%f<33>;
	.reg .s64 	%rd<16>;
	// demoted variable
	.shared .align 4 .b8 reducemaxdiff$__cuda_local_var_34094_35_non_const_sdata[2048];

	ld.param.u64 	%rd5, [reducemaxdiff_param_0];
	ld.param.u64 	%rd6, [reducemaxdiff_param_1];
	ld.param.u64 	%rd7, [reducemaxdiff_param_2];
	ld.param.f32 	%f32, [reducemaxdiff_param_3];
	ld.param.u32 	%r9, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd7;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 1 8 1
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r15, %r16, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r16, %r11;
	.loc 1 8 1
	setp.ge.s32	%p1, %r15, %r9;
	@%p1 bra 	BB2_2;

BB2_1:
	.loc 1 8 1
	mul.wide.s32 	%rd8, %r15, 4;
	add.s64 	%rd9, %rd3, %rd8;
	add.s64 	%rd10, %rd2, %rd8;
	ld.global.f32 	%f5, [%rd10];
	ld.global.f32 	%f6, [%rd9];
	sub.f32 	%f7, %f6, %f5;
	.loc 3 2750 10
	abs.f32 	%f8, %f7;
	.loc 3 2770 10
	max.f32 	%f32, %f32, %f8;
	.loc 1 8 1
	add.s32 	%r15, %r15, %r4;
	.loc 1 8 1
	setp.lt.s32	%p2, %r15, %r9;
	@%p2 bra 	BB2_1;

BB2_2:
	.loc 1 8 1
	mul.wide.s32 	%rd11, %r2, 4;
	mov.u64 	%rd12, reducemaxdiff$__cuda_local_var_34094_35_non_const_sdata;
	add.s64 	%rd4, %rd12, %rd11;
	st.shared.f32 	[%rd4], %f32;
	bar.sync 	0;
	.loc 1 8 1
	setp.lt.u32	%p3, %r16, 66;
	@%p3 bra 	BB2_6;

BB2_3:
	.loc 1 8 1
	mov.u32 	%r7, %r16;
	shr.u32 	%r16, %r7, 1;
	.loc 1 8 1
	setp.ge.u32	%p4, %r2, %r16;
	@%p4 bra 	BB2_5;

	.loc 1 8 1
	ld.shared.f32 	%f9, [%rd4];
	add.s32 	%r12, %r16, %r2;
	mul.wide.u32 	%rd13, %r12, 4;
	add.s64 	%rd15, %rd12, %rd13;
	ld.shared.f32 	%f10, [%rd15];
	.loc 3 2770 10
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%rd4], %f11;

BB2_5:
	.loc 1 8 1
	bar.sync 	0;
	.loc 1 8 1
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB2_3;

BB2_6:
	.loc 1 8 1
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB2_8;

	.loc 1 8 1
	ld.volatile.shared.f32 	%f12, [%rd4];
	ld.volatile.shared.f32 	%f13, [%rd4+128];
	.loc 3 2770 10
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%rd4], %f14;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f15, [%rd4+64];
	ld.volatile.shared.f32 	%f16, [%rd4];
	.loc 3 2770 10
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%rd4], %f17;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f18, [%rd4+32];
	ld.volatile.shared.f32 	%f19, [%rd4];
	.loc 3 2770 10
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%rd4], %f20;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f21, [%rd4+16];
	ld.volatile.shared.f32 	%f22, [%rd4];
	.loc 3 2770 10
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%rd4], %f23;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f24, [%rd4+8];
	ld.volatile.shared.f32 	%f25, [%rd4];
	.loc 3 2770 10
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%rd4], %f26;
	.loc 1 8 1
	ld.volatile.shared.f32 	%f27, [%rd4+4];
	ld.volatile.shared.f32 	%f28, [%rd4];
	.loc 3 2770 10
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%rd4], %f29;

BB2_8:
	.loc 1 8 1
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB2_10;

	.loc 1 8 1
	ld.shared.f32 	%f30, [reducemaxdiff$__cuda_local_var_34094_35_non_const_sdata];
	.loc 3 2750 10
	abs.f32 	%f31, %f30;
	.loc 1 8 37
	mov.b32 	 %r13, %f31;
	.loc 3 3781 3
	atom.global.max.s32 	%r14, [%rd1], %r13;

BB2_10:
	.loc 1 9 2
	ret;
}


`
)
