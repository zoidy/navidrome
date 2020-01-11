package ledis

import "github.com/google/wire"

var Set = wire.NewSet(
	NewCheckSumRepository,
	NewMediaFolderRepository,
	NewNowPlayingRepository,
	NewPlaylistRepository,
)