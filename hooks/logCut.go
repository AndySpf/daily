package hooks

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"daily/g"
	"daily/utils"

	log "github.com/sirupsen/logrus"
)

var (
	mux sync.Mutex
)

type MyHookCut struct{}

func (m *MyHookCut) Levels() []log.Level {
	return log.AllLevels
}

func (*MyHookCut) Fire(entry *log.Entry) error {
	// 如果out不是os.File类型则可能是os.Stdout等非文件类型，无需处理,跳过该hook即可
	if _, ok := entry.Logger.Out.(*os.File); !ok {
		return nil
	}

	v := entry.Logger.Out.(*os.File)
	fi, err := entry.Logger.Out.(*os.File).Stat()
	if err != nil {
		return err
	}
	fmt.Println("当前大小为:", fi.Size())

	// 文件大小没有超限，则直接返回
	if fi.Size() < int64(getFileSize()) {
		return nil
	}

	// 超限的话由于要改动logger的Outer，以及删除压缩文件等 上锁
	mux.Lock()
	defer mux.Unlock()
	entry.Logger.Out = nil
	v.Close() // 关掉写满了的文件的句柄

	dir := g.Cfg.LogConfig.LogDir
	// 获取最旧的日志文件和压缩文件序号
	oldestLogFileIndex, err := getLastedLogFileIndex(dir)
	if err != nil {
		return err
	}
	oldestCompressIndex, err := getLastedCompressFileIndex(dir)
	if err != nil {
		return err
	}

	// 在文件大小超限的前提下，日志文件个数达到最大值，则开始压缩逻辑，没有达到则滚动日志文件即可
	if oldestLogFileIndex == g.Cfg.LogConfig.LogFileCount {
		// 滚动压缩文件逻辑
		err = scrollCompress(dir, oldestCompressIndex)
		if err != nil {
			return err
		}

		// 压缩逻辑
		err = compressLogFile(dir, oldestCompressIndex)
		if err != nil {
			panic("压缩文件失败" + err.Error())
		}

		// 压缩后新建log.1 将logger的out指向新文件
		file, err := os.OpenFile(path.Join(dir, "log.1"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			return errors.New("创建新的日志文件失败:" + err.Error())
		}
		entry.Logger.Out = file
		return nil
	} else {
		// 没有达到最大日志文件个数，则滚动日志文件序号，新建log.1, 并重定向输出
		err = scrollLogFile(dir)
		if err != nil {
			return err
		}

		file, err := os.OpenFile(path.Join(dir, "log.1"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			return errors.New("创建新的日志文件失败:" + err.Error())
		}
		entry.Logger.Out = file
	}
	return nil
}

func scrollLogFile(dir string) error {
	res, err := utils.ExecuteCMD(fmt.Sprintf("ls -lrt %s | grep log | grep -v tgz | awk '{print $9}'", dir))
	if err != nil {
		return err
	}
	data := strings.Trim(res, "\n")
	if data == "" {
		return nil
	}
	nameList := strings.Split(data, "\n")
	for _, name := range nameList {
		index := getLogFileIndex(name)
		format := "%s/log.%d"
		err := os.Rename(fmt.Sprintf(format, dir, index), fmt.Sprintf(format, dir, index+1))
		if err != nil {
			return err
		}
	}
	return nil
}

func getLastedLogFileIndex(dir string) (int, error) {
	res, err := utils.ExecuteCMD(fmt.Sprintf("ls -lrt %s | grep log | grep -v tgz | awk '{print $9}'", dir))
	if err != nil {
		panic("执行命令失败")
	}
	data := strings.Trim(res, "\n")
	if data == "" {
		data += "log.1"
	}
	resList := strings.Split(data, "\n")
	lastedLogFile := resList[0]

	index := getLogFileIndex(lastedLogFile)
	return index, nil
}

func getLastedCompressFileIndex(dir string) (int, error) {
	res, err := utils.ExecuteCMD(fmt.Sprintf("ls -lrt %s | grep tgz | awk '{print $9}'", dir))
	if err != nil {
		return 0, err
	}
	data := strings.Trim(res, "\n")
	if data == "" {
		return 0, nil
	}
	resList := strings.Split(data, "\n")
	lastedCompressFile := resList[0]

	index := getCompressIndex(lastedCompressFile)
	return index, nil
}

func getCompressIndex(name string) int {
	nameList := strings.Split(name, ".")
	indexString := nameList[len(nameList)-2]
	index, err := strconv.Atoi(indexString)
	if err != nil {
		return 0
	}
	return index
}

func compressLogFile(dir string, compressIndex int) error {

	// 产生新的压缩文件logs.1.tgz, 删除日志文件log.*
	var files string
	for i := 1; i <= g.Cfg.LogConfig.LogFileCount; i++ {
		files += fmt.Sprintf(" log.%d", i)
	}

	cmdFormat := "tar -zcvf %s/logs.1.tgz -C %s %s;rm -rf %s/log.*"
	_, err := utils.ExecuteCMD(fmt.Sprintf(cmdFormat, dir, dir, files, dir))
	if err != nil {
		return err
	}
	return nil
}

func getLogFileIndex(name string) int {
	nameList := strings.Split(name, ".")
	indexString := nameList[len(nameList)-1]
	index, err := strconv.Atoi(indexString)
	if err != nil {
		return 0
	}
	return index
}

func getFileSize() int {
	switch g.Cfg.LogConfig.LogFileUnit {
	case "KB":
		return 1024 * g.Cfg.LogConfig.LogFileSize
	case "MB":
		return 1024 * 1024 * g.Cfg.LogConfig.LogFileSize
	case "GB":
		return 1024 * 1024 * 1024 * g.Cfg.LogConfig.LogFileSize
	default:
		return 1024 * g.Cfg.LogConfig.LogFileSize
	}
}

func scrollCompress(dir string, oldestCompressIndex int) error {
	if oldestCompressIndex == 0 { // 第一次产生压缩文件前，oldestCompressIndex为0
		return nil
	}

	// 删除最老的文件， 即标号最大的文件， 然后继续压缩逻辑，压缩步骤会将其他所有压缩文件重命名
	if oldestCompressIndex == g.Cfg.LogConfig.CompressFileCount {
		err := os.Remove(path.Join(dir, fmt.Sprintf("logs.%d.tgz", g.Cfg.LogConfig.CompressFileCount)))
		if err != nil {
			return err
		}
	}

	// 每次产生新的压缩文件，都要将其他压缩文件重命名，保证序号为1的文件是最新的
	res, err := utils.ExecuteCMD(fmt.Sprintf("ls -lrt %s | grep tgz | awk '{print $9}'", dir))
	if err != nil {
		return err
	}
	data := strings.Trim(res, "\n")
	if data == "" {
		return nil
	}
	nameList := strings.Split(data, "\n")

	// 遍历所有压缩文件并重命名，序号加一
	for _, name := range nameList {
		index := getCompressIndex(name)
		format := "%s/logs.%d.tgz"
		err := os.Rename(fmt.Sprintf(format, dir, index), fmt.Sprintf(format, dir, index+1))
		if err != nil {
			return err
		}
	}
	return nil
}
