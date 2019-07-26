package service

import (
	"CephMonitorAPI/goceph/cephfs"
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strings"
)

const (
	commandPrefix = "getfattr -d -m"
	entries       = "ceph.dir.entries"
	filesommond   = "ceph.dir.files"
	rebates       = "ceph.dir.rbytes" //目录暂用的字节数
	rctime        = "ceph.dir.rctime"
	rentries      = "ceph.dir.rentries"
	rfiles        = "ceph.dir.rfiles"   //目录下的文件数量
	rsubdirs      = "ceph.dir.rsubdirs" //子目录个数
	subdirs       = "ceph.dir.subdirs"
)

type CephfsService struct {
	Dir string `url:"dir"`
}

type CephfsDirInfo struct {
	Dir       string `json:"dir"`
	UsedBytes string `json:"usedBytes"`
}

type CephfsDirBatchService struct {
	CephfsDirs []CephfsDirInfo `json:"cephfsDirs"`
}

func (cephfsService *CephfsService) CreateDir() error {
	mount, err := cephfs.CreateMount()
	err = mount.ReadDefaultConfigFile()
	err = mount.Mount()
	if err != nil {
		log.Println("create mount error")
		return err
	}
	currentDir := mount.CurrentDir()
	log.Println(currentDir)
	err = mount.MakeDir("/k8s/wangcheng", 0755)
	if err != nil {
		log.Printf("Make dir:%s error", cephfsService.Dir)
		return err
	}
	return nil
}

func (cephfsService *CephfsService) DeleteDir() error {
	mount, err := cephfs.CreateMount()
	err = mount.ReadDefaultConfigFile()
	err = mount.Mount()
	if err != nil {
		log.Println("create mount error")
		return err
	}
	err = mount.RemoveDir(cephfsService.Dir)
	if err != nil {
		log.Printf("删除cephfs dir失败：%s", err.Error())
		return err
	}
	return nil
}

func (cephfsService *CephfsService) GetDirUsage() (CephfsDirInfo, error) {
	log.Printf("执行统计命令dir:%s", cephfsService.Dir)
	cmd := exec.Command("/bin/bash", "-c", commandPrefix, rebates, "/mnt" + cephfsService.Dir)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return CephfsDirInfo{}, errors.New("云盘统计命令执行失败")
	}
	resultStringArray := strings.Split("\"", out.String())
	usedStr := resultStringArray[1]
	return CephfsDirInfo{
		Dir:       cephfsService.Dir,
		UsedBytes: usedStr,
	}, nil
}

func (cephfsDirBatchService *CephfsDirBatchService) GetCephDirsUsage() ([]CephfsDirInfo, error) {
	var dirArray = cephfsDirBatchService.CephfsDirs
	dirNum := len(dirArray)
	if dirNum == 0 {
		return nil, errors.New("文件目录不能为空")
	}
	infoChan := make(chan CephfsDirInfo, 3)
	boundary := dirNum % 5
	for index := 0; index <= dirNum; index += 5 {
		if (dirNum > boundary && dirNum-index < 5) || dirNum <= boundary {
			go calculateCephfsUsed(dirArray[index:], infoChan)
			break
		}
		go calculateCephfsUsed(dirArray[index:index+5], infoChan)
	}
	var resultDirInfos []CephfsDirInfo

	for {
		dirInfo := <-infoChan
		resultDirInfos = append(resultDirInfos, dirInfo)
		if len(resultDirInfos) == dirNum {
			close(infoChan)
			return resultDirInfos, nil
		}
	}
}

func calculateCephfsUsed(dirs []CephfsDirInfo, c chan CephfsDirInfo) {
	for _, dir := range dirs {
		cmd := exec.Command("/bin/bash", "-c", commandPrefix, rebates, "/mnt" + dir.Dir)
		var out bytes.Buffer
		cmd.Stdout = &out
		if err := cmd.Run(); err != nil {
		}
		resultStringArray := strings.Split("\"", out.String())
		usedStr := resultStringArray[1]
		c <- CephfsDirInfo{
			Dir:       dir.Dir,
			UsedBytes: usedStr,
		}
	}
}