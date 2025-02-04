// Code generated DO NOT EDIT

package cmds

import "testing"

func cf0(s Builder) {
	s.CfAdd().Key("1").Item("1").Build()
	s.CfAddnx().Key("1").Item("1").Build()
	s.CfCount().Key("1").Item("1").Build()
	s.CfCount().Key("1").Item("1").Cache()
	s.CfDel().Key("1").Item("1").Build()
	s.CfExists().Key("1").Item("1").Build()
	s.CfExists().Key("1").Item("1").Cache()
	s.CfInfo().Key("1").Build()
	s.CfInfo().Key("1").Cache()
	s.CfInsert().Key("1").Capacity(1).Nocreate().Items().Item("1").Item("1").Build()
	s.CfInsert().Key("1").Capacity(1).Items().Item("1").Item("1").Build()
	s.CfInsert().Key("1").Nocreate().Items().Item("1").Item("1").Build()
	s.CfInsert().Key("1").Items().Item("1").Item("1").Build()
	s.CfInsertnx().Key("1").Capacity(1).Nocreate().Items().Item("1").Item("1").Build()
	s.CfInsertnx().Key("1").Capacity(1).Items().Item("1").Item("1").Build()
	s.CfInsertnx().Key("1").Nocreate().Items().Item("1").Item("1").Build()
	s.CfInsertnx().Key("1").Items().Item("1").Item("1").Build()
	s.CfLoadchunk().Key("1").Iterator(1).Data("1").Build()
	s.CfMexists().Key("1").Item("1").Item("1").Build()
	s.CfReserve().Key("1").Capacity(1).Bucketsize(1).Maxiterations(1).Expansion(1).Build()
	s.CfReserve().Key("1").Capacity(1).Bucketsize(1).Maxiterations(1).Build()
	s.CfReserve().Key("1").Capacity(1).Bucketsize(1).Expansion(1).Build()
	s.CfReserve().Key("1").Capacity(1).Bucketsize(1).Build()
	s.CfReserve().Key("1").Capacity(1).Maxiterations(1).Build()
	s.CfReserve().Key("1").Capacity(1).Expansion(1).Build()
	s.CfReserve().Key("1").Capacity(1).Build()
	s.CfScandump().Key("1").Iterator(1).Build()
}

func TestCommand_InitSlot_cf(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { cf0(s) })
}

func TestCommand_NoSlot_cf(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { cf0(s) })
}
