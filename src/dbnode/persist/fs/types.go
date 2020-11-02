// FileSetFileIdentifier contains all the information required to identify a FileSetFile.
// DataWriterOpenOptions is the options struct for the Open method on the DataFileSetWriter.
// that contains information specific to writing snapshot files.
// DataFileSetWriter provides an unsynchronized writer for a TSDB file set.
// DataFileSetReaderStatus describes the status of a file set reader.
// DataFileSetReader provides an unsynchronized reader for a TSDB file set.
	// StreamingRead returns the next unpooled id, encodedTags, data, checksum
	// values ordered by id, or error, will return io.EOF at end of volume.
	// Validate validates both the metadata and data and returns an error if either is corrupted.
	// ValidateMetadata validates the data and returns an error if the data is corrupted.
	// ValidateData validates the data and returns an error if the data is corrupted.
	// Range returns the time range associated with data in the volume.
	// Entries returns the count of entries in the volume.
	// EntriesRead returns the position read into the volume.
	// MetadataRead returns the position of metadata read into the volume.
	// StreamingEnabled returns true if the reader is opened in streaming mode.
// DataFileSetSeeker provides an out of order reader for a TSDB file set.
	// Open opens the files for the given shard and version for reading.
	// SeekWideEntry seeks in a manner similar to SeekIndexEntry, but
	// instead yields a wide entry checksum of the series.
	SeekWideEntry(id ident.ID, resources ReusableSeekerResources) (xio.WideEntry, error)
	// SeekWideEntry is the same as in DataFileSetSeeker.
	SeekWideEntry(id ident.ID, resources ReusableSeekerResources) (xio.WideEntry, error)
// DataBlockRetriever provides a block retriever for TSDB file sets.
	// Open the block retriever to retrieve from a namespace.
// RetrievableDataBlockSegmentReader is a retrievable block reader.
	// rate to use for the index bloom filter size and k hashes estimation.
	// SetInfoReaderBufferSize sets the buffer size for reading TSDB info,
	//  digest and checkpoint files.
	// InfoReaderBufferSize returns the buffer size for reading TSDB info,
	//  digest and checkpoint files.
// BlockRetrieverOptions represents the options for block retrieval.
	// MergeAndCleanup merges the specified fileset file with a merge target and
	// removes the previous version of the fileset. This should only be called
	// within the bootstrapper. Any other file deletions outside of the
	// bootstrapper should be handled by the CleanupManager.
// CrossBlockReader allows reading data (encoded bytes) from multiple
// DataFileSetReaders of the same shard, ordered lexographically by series ID,
// then by block time.
	// Next advances to the next data record, returning true if it exists.
	// Current returns distinct series id and encodedTags, plus a slice with data
	// and checksums from all blocks corresponding to that series (in temporal order).
	// ID, encodedTags, records, and underlying data are invalidated on each call to Next().