package types

import "strconv"

// ReturnCode enum
const (
	Success     int32 = 0
	Fail        int32 = -1
	ServerError int32 = 500
)

// Error message
const (
	// Dir
	ListDirError            = "Error while listing dir"
	DirNotFoundConst        = "Directory not found: "
	FileMoveErrorConst      = "Failed to move file: "
	CreateDirErrorConst     = "Failed to create dir: "
	ReadDirErrorConst       = "Failed to read directory: "
	DeleteDirErrorConst     = "Failed to delete directory: "
	RenameDirErrorConst     = "Failed to rename directory: "
	DirAlreadyExistsConst   = "Directory already exists: "
	GetFileNumberErrorConst = "Failed to get file number for directory: "

	// File
	CloseFileErrorConst    = "Failed to close file: "
	WriteFileErrorConst    = "Failed to write file: "
	OpenFileErrorConst     = "Failed to open file: "
	DeleteFileErrorConst   = "Failed to delete file: "
	CopyFileErrorConst     = "Failed to copy file: "
	CreateFileErrorConst   = "Failed to create file: "
	FileNotFoundErrorConst = "File not found: "
	RenameFileErrorConst   = "Failed to rename file: "

	// User
	SignInErrorConst       = "Failed to sign in: "
	UserAlreadyExistsConst = "User already exists: "

	// DB
	DBConnectError = "Failed to connect to database"
)

// Success message
const (
	// Dir
	ListDirSuccess            = "List dir created"
	DirMoveSuccessConst       = "Directory moved: "
	DeleteDirSuccessConst     = "Directory deleted: "
	RenameDirSuccessConst     = "Directory renamed: "
	CreateDirSuccessConst     = "Directory created: "
	GetRootPathSuccessConst   = "Get root path: "
	GetFileNumberSuccessConst = "Number of file in directory: "
	SearchSuccessConst = "Search success: "

	// File
	WriteFileChunkSuccessConst = "File Chunk written: "
	MoveFileSuccessConst       = "File moved: "
	DeleteFileSuccessConst     = "File deleted: "
	RenameFileSuccessConst     = "File renamed: "
	FileSizeSuccessConst       = "File size retrieved for: "

	// User
	SignInSuccess    = "Sign In successful"
	UserCreatedConst = "User created: "
)

func DirNotFound(dirName string) string {
	return DirNotFoundConst + dirName
}

func DirMoveSuccess(dirName string) string {
	return DirMoveSuccessConst + dirName
}

func FileMoveError(dirName string) string {
	return FileMoveErrorConst + dirName
}

func CreateDirError(dirName string) string {
	return CreateDirErrorConst + dirName
}

func ReadDirError(dirName string) string {
	return ReadDirErrorConst + dirName
}

func DeleteDirSuccess(dirName string) string {
	return DeleteDirSuccessConst + dirName
}

func DeleteDirError(dirName string) string {
	return DeleteDirErrorConst + dirName
}

func RenameDirSuccess(oldName, newName string) string {
	return RenameDirSuccessConst + oldName + " to " + newName
}

func RenameDirError(dirName string) string {
	return RenameDirErrorConst + dirName
}

func DirAlreadyExists(dirName string) string {
	return DirAlreadyExistsConst + dirName
}

func CreateDirSuccess(dirName string) string {
	return CreateDirSuccessConst + dirName
}

func WriteFileChunkSuccess(fileName string) string {
	return WriteFileChunkSuccessConst + fileName
}

func CloseFileError(fileName string) string {
	return CloseFileErrorConst + fileName
}

func WriteFileError(fileName string) string {
	return WriteFileErrorConst + fileName
}

func OpenFileError(fileName string) string {
	return OpenFileErrorConst + fileName
}

func MoveFileSuccess(fileName string) string {
	return MoveFileSuccessConst + fileName
}

func DeleteFileError(fileName string) string {
	return DeleteFileErrorConst + fileName
}

func CopyFileError(fileName string) string {
	return CopyFileErrorConst + fileName
}

func CreateFileError(fileName string) string {
	return CreateFileErrorConst + fileName
}

func FileNotFound(fileName string) string {
	return FileNotFoundErrorConst + fileName
}

func DeleteFileSuccess(fileName string) string {
	return DeleteFileSuccessConst + fileName
}

func RenameFileSuccess(oldName, newName string) string {
	return RenameFileSuccessConst + oldName + " to " + newName
}

func RenameFileError(oldName, newName string) string {
	return RenameFileErrorConst + oldName + " to " + newName
}

func SignInError(userName string) string {
	return SignInErrorConst + userName
}

func UserAlreadyExists(userName string) string {
	return UserAlreadyExistsConst + userName
}

func UserCreated(userName string) string {
	return UserCreatedConst + userName
}

func GetRootPathSuccess(dirName string) string {
	return GetRootPathSuccessConst + dirName
}

func FileSizeSuccess(dirName string) string {
	return FileSizeSuccessConst + dirName
}

func GetFileNumberSuccess(dirName string, number int64) string {
	return GetFileNumberSuccessConst + dirName + " (" + strconv.FormatInt(number, 10) + ")"
}

func GetFileNumberError(dirName string) string {
	return GetFileNumberErrorConst + dirName
}

func SearchSuccess(search string) string {
	return SearchSuccessConst + search
}