
package files

// FTP server management
type FTPManager struct {
    accounts map[string]*FTPAccount
}

type FTPAccount struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    HomeDir  string `json:"home_dir"`
    Quota    int64  `json:"quota"`
}

func NewFTPManager() *FTPManager {
    return &FTPManager{
        accounts: make(map[string]*FTPAccount),
    }
}
