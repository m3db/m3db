package placements

import (
	"testing"
)

func makeStub() (PlacementArgs, PlacementFlags) {
	placementArgs := PlacementArgs{}
	placementFlagSets := SetupFlags(&placementArgs)

	placementFlagSets.PlacementDoer = func(*PlacementArgs, string) { return }
	placementFlagSets.AddDoer = func(*PlacementArgs, string) { return }
	placementFlagSets.DeleteDoer = func(*PlacementArgs, string) { return }
	placementFlagSets.InitDoer = func(*PlacementArgs, string) { return }
	placementFlagSets.ReplaceDoer = func(*PlacementArgs, string) { return }

	return placementArgs, placementFlagSets
}
func TestBasic(t *testing.T) {

	// default is to list placements
	// so no args is OK
	args, flags := makeStub()
	if err := parseAndDo([]string{""}, &args, &flags, ""); err != nil {
		t.Error("It should not return error no args")
	}

	args, flags = makeStub()
	if err := parseAndDo([]string{"junk"}, &args, &flags, ""); err == nil {
		t.Error("It should return an error on junk")
	}

	args, flags = makeStub()
	if err := parseAndDo([]string{"delete", "-all"}, &args, &flags, ""); err != nil {
		t.Error("It should not return error no sane args")
	}

	args, flags = makeStub()
	if err := parseAndDo([]string{"delete", "-node", "nodeName"}, &args, &flags, ""); err != nil {
		t.Error("It should not return error no sane args")
	}

	args, flags = makeStub()
	if err := parseAndDo([]string{"add", "-f", "somefile"}, &args, &flags, ""); err != nil {
		t.Error("It should not return error no sane args")
	}

	args, flags = makeStub()
	if err := parseAndDo([]string{"init", "-f", "somefile"}, &args, &flags, ""); err != nil {
		t.Error("It should not return error no sane args")
	}

	args, flags = makeStub()
	if err := parseAndDo([]string{"replace", "-f", "somefile"}, &args, &flags, ""); err != nil {
		t.Error("It should not return error no sane args")
	}

}
