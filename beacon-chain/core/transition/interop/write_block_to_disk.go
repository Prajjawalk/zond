package interop

import (
	"fmt"
	"os"
	"path"

	"github.com/theQRL/zond/config/features"
	"github.com/theQRL/zond/consensus-types/interfaces"
	"github.com/theQRL/zond/io/file"
)

// WriteBlockToDisk as a block ssz. Writes to temp directory. Debug!
func WriteBlockToDisk(block interfaces.SignedBeaconBlock, failed bool) {
	if !features.Get().WriteSSZStateTransitions {
		return
	}

	filename := fmt.Sprintf("beacon_block_%d.ssz", block.Block().Slot())
	if failed {
		filename = "failed_" + filename
	}
	fp := path.Join(os.TempDir(), filename)
	log.Warnf("Writing block to disk at %s", fp)
	enc, err := block.MarshalSSZ()
	if err != nil {
		log.WithError(err).Error("Failed to ssz encode block")
		return
	}
	if err := file.WriteFile(fp, enc); err != nil {
		log.WithError(err).Error("Failed to write to disk")
	}
}
