// Code generated DO NOT EDIT

package cmds

import "testing"

func model0(s Builder) {
	s.AiModeldel().Key("1").Build()
	s.AiModelget().Key("1").Meta().Blob().Build()
	s.AiModelget().Key("1").Meta().Blob().Cache()
	s.AiModelget().Key("1").Meta().Build()
	s.AiModelget().Key("1").Meta().Cache()
	s.AiModelget().Key("1").Blob().Build()
	s.AiModelget().Key("1").Blob().Cache()
	s.AiModelget().Key("1").Build()
	s.AiModelget().Key("1").Cache()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Minbatchtimeout(1).Inputs(1).Input("1").Input("1").Outputs(1).Output("1").Output("1").Blob("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Minbatchtimeout(1).Inputs(1).Input("1").Input("1").Outputs(1).Output("1").Output("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Minbatchtimeout(1).Inputs(1).Input("1").Input("1").Blob("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Minbatchtimeout(1).Inputs(1).Input("1").Input("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Minbatchtimeout(1).Outputs(1).Output("1").Output("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Minbatchtimeout(1).Blob("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Minbatchtimeout(1).Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Inputs(1).Input("1").Input("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Outputs(1).Output("1").Output("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Blob("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchsize(1).Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Minbatchtimeout(1).Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Inputs(1).Input("1").Input("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Outputs(1).Output("1").Output("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Blob("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Batchsize(1).Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Minbatchsize(1).Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Minbatchtimeout(1).Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Inputs(1).Input("1").Input("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Outputs(1).Output("1").Output("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Blob("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Tag("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Batchsize(1).Build()
	s.AiModelstore().Key("1").Tf().Cpu().Minbatchsize(1).Build()
	s.AiModelstore().Key("1").Tf().Cpu().Minbatchtimeout(1).Build()
	s.AiModelstore().Key("1").Tf().Cpu().Inputs(1).Input("1").Input("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Outputs(1).Output("1").Output("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Blob("1").Build()
	s.AiModelstore().Key("1").Tf().Cpu().Build()
	s.AiModelstore().Key("1").Tf().Gpu().Tag("1").Build()
	s.AiModelstore().Key("1").Tf().Gpu().Batchsize(1).Build()
	s.AiModelstore().Key("1").Tf().Gpu().Minbatchsize(1).Build()
	s.AiModelstore().Key("1").Tf().Gpu().Minbatchtimeout(1).Build()
	s.AiModelstore().Key("1").Tf().Gpu().Inputs(1).Input("1").Input("1").Build()
	s.AiModelstore().Key("1").Tf().Gpu().Outputs(1).Output("1").Output("1").Build()
	s.AiModelstore().Key("1").Tf().Gpu().Blob("1").Build()
	s.AiModelstore().Key("1").Tf().Gpu().Build()
	s.AiModelstore().Key("1").Torch().Cpu().Build()
	s.AiModelstore().Key("1").Torch().Gpu().Build()
	s.AiModelstore().Key("1").Onnx().Cpu().Build()
	s.AiModelstore().Key("1").Onnx().Gpu().Build()
}

func TestCommand_InitSlot_model(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { model0(s) })
}

func TestCommand_NoSlot_model(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { model0(s) })
}
