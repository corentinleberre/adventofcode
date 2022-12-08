package day7

type Directory struct {
	parentDirectory *Directory
	name            string
	files           []File
	directories     []*Directory
	totalSize       int64
}

type File struct {
	parentDirectory *Directory
	name            string
	size            int64
}
