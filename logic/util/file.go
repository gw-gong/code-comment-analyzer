package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"code-comment-analyzer/config"
	"code-comment-analyzer/protocol"
)

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		if strings.HasPrefix(f.Name, "__MACOSX/") {
			continue // 跳过 __MACOSX 目录和其中的文件 (针对mac系统)
		}

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
			continue
		}

		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}

		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateUUIDProjectName() string {
	return config.Cfg.UuidProjectPrefix + GenerateUUIDName()
}

func GenerateUUIDAvatarName() string {
	return config.Cfg.UuidAvatarPrefix + GenerateUUIDName()
}

func ReadFileContentByPath(path string) (fileContent string, err error) {
	file, err := os.Open(path)
	if err != nil {
		err = fmt.Errorf("打开文件失败")
		return
	}
	defer file.Close()

	fileContentBytes, err := io.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("读取文件失败")
		return
	}

	return string(fileContentBytes), nil
}

func BuildDirectoryTree(currentPath, rootPath, projectStorageName string) protocol.FileNode {
	node := protocol.FileNode{
		Label: filepath.Base(currentPath),
	}

	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return node
	}

	for _, entry := range entries {
		fullPath := filepath.Join(currentPath, entry.Name())
		relPath, _ := filepath.Rel(rootPath, fullPath)
		relPath = filepath.ToSlash(relPath) // 统一使用斜杠

		if entry.IsDir() {
			child := BuildDirectoryTree(fullPath, rootPath, projectStorageName)
			node.Children = append(node.Children, child)
		} else {
			node.Children = append(node.Children, protocol.FileNode{
				Label: entry.Name(),
				Value: filepath.Join(projectStorageName, relPath),
			})
		}
	}

	return node
}

func TransformProfilePictureUrlToResourceUrl(notUseDefaultAvatar bool, profilePictureUrl string) *string {
	profilePictureFileName := config.Cfg.DefaultAvatar
	if notUseDefaultAvatar {
		profilePictureFileName = filepath.Base(profilePictureUrl)
	}
	avatarsStorageRootPath := config.Cfg.FileStoragePath.Avatars
	avatarUrl := filepath.Join("/", avatarsStorageRootPath, profilePictureFileName)
	return &avatarUrl
}

func GetAvatarStoragePath(userID uint64, avatarFileName string) string {
	if avatarFileName == config.Cfg.DefaultAvatar {
		return filepath.Join(config.Cfg.FileStoragePath.Avatars, avatarFileName)
	}
	return filepath.Join(config.Cfg.FileStoragePath.Avatars, fmt.Sprintf("%v", userID), avatarFileName)
}
