	"github.com/m3db/m3/src/dbnode/persist/fs/wide"
// FileSetFileIdentifier contains all the information required to identify a FileSetFile
// DataWriterOpenOptions is the options struct for the Open method on the DataFileSetWriter
// that contains information specific to writing snapshot files
// DataFileSetWriter provides an unsynchronized writer for a TSDB file set
// DataFileSetReaderStatus describes the status of a file set reader
// StreamedChecksum yields a schema.IndexChecksum value asynchronously,
// and any errors encountered during execution.
type StreamedChecksum interface {
	// RetrieveIndexChecksum retrieves the index checksum.
	RetrieveIndexChecksum() (xio.IndexChecksum, error)
}

type emptyStreamedChecksum struct{}

func (emptyStreamedChecksum) RetrieveIndexChecksum() (xio.IndexChecksum, error) {
	return xio.IndexChecksum{}, nil
}

// EmptyStreamedChecksum is an empty streamed checksum.
var EmptyStreamedChecksum StreamedChecksum = emptyStreamedChecksum{}

// DataFileSetReader provides an unsynchronized reader for a TSDB file set
	// StreamingRead returns the next unpooled id, encodedTags, data, checksum values ordered by id,
	// or error, will return io.EOF at end of volume.
	// Validate validates both the metadata and data and returns an error if either is corrupted
	// ValidateMetadata validates the data and returns an error if the data is corrupted
	// ValidateData validates the data and returns an error if the data is corrupted
	// Range returns the time range associated with data in the volume
	// Entries returns the count of entries in the volume
	// EntriesRead returns the position read into the volume
	// MetadataRead returns the position of metadata read into the volume
	// StreamingEnabled returns true if the reader is opened in streaming mode
// DataFileSetSeeker provides an out of order reader for a TSDB file set
	// Open opens the files for the given shard and version for reading
	// SeekReadMismatchesByIndexChecksum seeks in a manner similar to
	// SeekIndexByEntry, checking against a set of streamed index checksums.
	SeekReadMismatchesByIndexChecksum(
		checksum xio.IndexChecksum,
		mismatchChecker wide.EntryChecksumMismatchChecker,
		resources ReusableSeekerResources,
	) (wide.ReadMismatch, error)

	// SeekIndexEntryToIndexChecksum seeks in a manner similar to SeekIndexEntry, but
	// instead yields a minimal structure describing a checksum of the series.
	SeekIndexEntryToIndexChecksum(id ident.ID, resources ReusableSeekerResources) (xio.IndexChecksum, error)
	// SeekReadMismatchesByIndexChecksum is the same as in DataFileSetSeeker.
	SeekReadMismatchesByIndexChecksum(
		checksum xio.IndexChecksum,
		mismatchChecker wide.EntryChecksumMismatchChecker,
		resources ReusableSeekerResources,
	) (wide.ReadMismatch, error)

	// SeekIndexEntryToIndexChecksum is the same as in DataFileSetSeeker.
	SeekIndexEntryToIndexChecksum(id ident.ID, resources ReusableSeekerResources) (xio.IndexChecksum, error)
// DataBlockRetriever provides a block retriever for TSDB file sets
	// Open the block retriever to retrieve from a namespace
// RetrievableDataBlockSegmentReader is a retrievable block reader
	// rate to use for the index bloom filter size and k hashes estimation
	// SetInfoReaderBufferSize sets the buffer size for reading TSDB info, digest and checkpoint files.
	// InfoReaderBufferSize returns the buffer size for reading TSDB info, digest and checkpoint files.
// BlockRetrieverOptions represents the options for block retrieval
	// MergeAndCleanup merges the specified fileset file with a merge target and removes the previous version of the
	// fileset. This should only be called within the bootstrapper. Any other file deletions outside of the bootstrapper
	// should be handled by the CleanupManager.
// CrossBlockReader allows reading data (encoded bytes) from multiple DataFileSetReaders of the same shard,
// ordered by series id first, and block start time next.
	// Next advances to the next data record and returns true, or returns false if no more data exists.
	// Current returns distinct series id and encodedTags, plus a slice with data and checksums from all
	// blocks corresponding to that series (in temporal order).
	// id, encodedTags, records slice and underlying data are being invalidated on each call to Next().

// CrossBlockIterator iterates across BlockRecords.
type CrossBlockIterator interface {
	encoding.Iterator

	// Reset resets the iterator to the given block records.
	Reset(records []BlockRecord)
}