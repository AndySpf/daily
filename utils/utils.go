package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strings"
	"time"
	"unsafe"

	_ "github.com/go-sql-driver/mysql"
)

func ExecuteCMD(cmd string) (string, error) {
	fmt.Println(cmd)
	c := exec.Command("sh", "-c", cmd)
	bs, err := c.Output()
	if err != nil {
		log.Println(err)
		return "nil", err
	}
	return string(bs), nil
}

// 管道形式
//func ExecuteCMD1(cmd string) {
//	c := exec.Command("ls")
//	c1 := exec.Command("wc", "-l")
//
//	c1.Stdin, _ = c.StdoutPipe()
//	c1.Stdout = os.Stdout
//	c1.Start()
//	c.Run()
//	c.Wait()
//}

// Index2Col 将excel的列序号转换为列名
func Index2Col(index int) string {
	if index < 1 {
		return "A"
	}

	start := int('A' - 1)
	suffix := 0
	prefix := ""

	if i := index % 26; i == 0 {
		suffix = start + 26
	} else {
		suffix = start + i
	}

	prefixCount := index / 27
	prefix = strings.Repeat("A", prefixCount)

	return prefix + string(byte(suffix))
}

func T() error {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", "root", "root", "tcp", "192.168.29.154", 3306, "alarms")
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return err
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)                                   //设置闲置连接数
	rows, _ := DB.Query("SELECT * FROM event_cases LIMIT 1") // Note: Ignoring errors for brevity
	cols, _ := rows.Columns()

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		fmt.Print(m)
	}
	return nil

}

type tf struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func format() {
	t := &tf{}
	t.Name = "tom"
	s := `{"age": 7}`
	json.Unmarshal([]byte(s), t)
	fmt.Println(t)
}

func RemoveItemFromSlice(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func ReverseSlice(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}

func SliceDeduplication(in []int) []int {
	sort.Ints(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++

		in[j] = in[i]
	}
	result := in[:j+1]
	return result
}

// time.Time序列化时会变成2006-01-02T15:04:05+8:00类型，使用重写后的Time可以正常json序列化
type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
